package pureport

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pureport/pureport-sdk-go/pureport/client"
	"github.com/pureport/pureport-sdk-go/pureport/session"
)

var (
	DeletableState = map[string]bool{
		"FAILED_TO_PROVISION": true,
		"ACTIVE":              true,
		"DOWN":                true,
		"FAILED_TO_UPDATE":    true,
		"FAILED_TO_DELETE":    true,
		"DELETED":             true,
	}
)

func getBaseConnectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"customer_asn": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"customer_networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"address": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.CIDRNetwork(16, 32),
					},
				},
			},
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"high_availability": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"location_href": {
			Type:     schema.TypeString,
			Required: true,
		},
		"billing_term": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HOURLY",
		},
		"nat_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},
					"mappings": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"native_cidr": {
									Type:     schema.TypeString,
									Required: true,
								},
								"nat_cidr": {
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"blocks": {
						Type:     schema.TypeList,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"pnat_cidr": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"href": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"speed": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntInSlice([]int{50, 100, 200, 300, 400, 500, 1000, 10000}),
		},
	}
}

func flattenConnection(connection client.Connection) map[string]interface{} {
	return map[string]interface{}{
		"customer_asn":      connection.CustomerASN,
		"customer_networks": flattenCustomerNetworks(connection.CustomerNetworks),
		"description":       connection.Description,
		"high_availability": connection.HighAvailability,
		"location":          flattenLink(connection.Location),
		"name":              connection.Name,
		"network":           flattenLink(connection.Network),
		"speed":             connection.Speed,
	}
}

func flattenLink(link *client.Link) map[string]interface{} {
	return map[string]interface{}{
		"id":   link.Id,
		"href": link.Href,
	}
}

func flattenCustomerNetworks(networks []client.CustomerNetwork) (out []map[string]interface{}) {

	for _, network := range networks {

		n := map[string]interface{}{
			"name":    network.Name,
			"address": network.Address,
		}

		out = append(out, n)
	}

	return
}

func flattenNatConfig(config client.NatConfig) map[string]interface{} {
	return map[string]interface{}{
		"blocks":    config.Blocks,
		"enabled":   config.Enabled,
		"pnat_cidr": config.PnatCidr,
		"mappings":  flattenMappings(config.Mappings),
	}
}

func flattenMappings(mappings []client.NatMapping) (out []map[string]interface{}) {

	for _, mapping := range mappings {

		m := map[string]interface{}{
			"nat_cidr":    mapping.NatCidr,
			"native_cidr": mapping.NativeCidr,
		}

		out = append(out, m)
	}

	return
}

func DeleteConnection(name string, d *schema.ResourceData, m interface{}) error {

	sess := m.(*session.Session)
	ctx := sess.GetSessionContext()
	connectionId := d.Id()

	// Wait until we are in a state that we can trigger a delete from
	log.Printf("[Info] Waiting to trigger a delete.")

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Minute

	wait_to_delete := func() error {

		c, resp, err := sess.Client.ConnectionsApi.GetConnection(ctx, connectionId)
		if err != nil {
			return backoff.Permanent(
				fmt.Errorf("Error deleting data for %s: %s", name, err),
			)
		}

		if resp.StatusCode >= 300 {
			return backoff.Permanent(
				fmt.Errorf("Error Response while attempting to delete %s: code=%v", name, resp.StatusCode),
			)
		}

		conn := reflect.ValueOf(c)
		if DeletableState[conn.FieldByName("State").String()] {
			return nil
		} else {
			return fmt.Errorf("Not Completed ...")
		}
	}

	if err := backoff.Retry(wait_to_delete, b); err != nil {
		return err
	}

	// Delete
	_, resp, err := sess.Client.ConnectionsApi.DeleteConnection(ctx, connectionId)
	if err != nil {
		return fmt.Errorf("Error deleting data for %s: %s", name, err)
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("Error Response while deleting %s: code=%v", name, resp.StatusCode)
	}

	log.Printf("[Info] Waiting for connection to be deleted")
	wait_for_delete := func() error {

		log.Printf("Retrying ...%+v", b.GetElapsedTime())
		_, resp, _ := sess.Client.ConnectionsApi.GetConnection(ctx, connectionId)

		if resp.StatusCode == 404 {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("Not Completed ...")
		}
	}

	return backoff.Retry(wait_for_delete, b)
}

// AddCustomerNetworks to decode the customer network information
func AddCustomerNetworks(d *schema.ResourceData) []client.CustomerNetwork {

	if data, ok := d.GetOk("customer_networks"); ok {

		customerNetworks := []client.CustomerNetwork{}

		for _, cn := range data.([]map[string]string) {

			new := client.CustomerNetwork{
				Name:    cn["name"],
				Address: cn["Address"],
			}

			customerNetworks = append(customerNetworks, new)
		}

		return customerNetworks
	}

	return nil
}

func AddNATConfiguration(d *schema.ResourceData) *client.NatConfig {

	if data, ok := d.GetOk("nat_config"); ok {

		natConfig := &client.NatConfig{}

		config := data.(map[string]interface{})
		natConfig.Enabled = config["enabled"].(bool)

		for _, m := range config["mappings"].([]map[string]string) {

			new := client.NatMapping{
				NativeCidr: m["native_cidr"],
			}

			natConfig.Mappings = append(natConfig.Mappings, new)
		}
		return natConfig
	}

	return nil
}

func AddCloudServices(d *schema.ResourceData) []client.Link {

	if data, ok := d.GetOk("cloud_services"); ok {

		cloudServices := []client.Link{}
		for _, cs := range data.([]map[string]string) {
			new := client.Link{
				Href: cs["href"],
			}
			cloudServices = append(cloudServices, new)
		}

		return cloudServices
	}

	return nil
}

func AddPeeringType(d *schema.ResourceData) *client.PeeringConfiguration {

	peeringConfig := &client.PeeringConfiguration{}

	if data, ok := d.GetOk("peering"); ok {
		peeringConfig.Type_ = data.(string)
	} else {
		peeringConfig.Type_ = "Private"
	}

	return peeringConfig
}
