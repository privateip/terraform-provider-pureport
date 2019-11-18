// Code generated by generators/listtags/main.go; DO NOT EDIT.

package keyvaluetags

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

// AcmListTags lists acm service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AcmListTags(conn *acm.ACM, identifier string) (KeyValueTags, error) {
	input := &acm.ListTagsForCertificateInput{
		CertificateArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForCertificate(input)

	if err != nil {
		return New(nil), err
	}

	return AcmKeyValueTags(output.Tags), nil
}

// AcmpcaListTags lists acmpca service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AcmpcaListTags(conn *acmpca.ACMPCA, identifier string) (KeyValueTags, error) {
	input := &acmpca.ListTagsInput{
		CertificateAuthorityArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return AcmpcaKeyValueTags(output.Tags), nil
}

// AmplifyListTags lists amplify service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AmplifyListTags(conn *amplify.Amplify, identifier string) (KeyValueTags, error) {
	input := &amplify.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return AmplifyKeyValueTags(output.Tags), nil
}

// AppmeshListTags lists appmesh service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AppmeshListTags(conn *appmesh.AppMesh, identifier string) (KeyValueTags, error) {
	input := &appmesh.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return AppmeshKeyValueTags(output.Tags), nil
}

// AppstreamListTags lists appstream service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AppstreamListTags(conn *appstream.AppStream, identifier string) (KeyValueTags, error) {
	input := &appstream.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return AppstreamKeyValueTags(output.Tags), nil
}

// AppsyncListTags lists appsync service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AppsyncListTags(conn *appsync.AppSync, identifier string) (KeyValueTags, error) {
	input := &appsync.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return AppsyncKeyValueTags(output.Tags), nil
}

// AthenaListTags lists athena service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func AthenaListTags(conn *athena.Athena, identifier string) (KeyValueTags, error) {
	input := &athena.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return AthenaKeyValueTags(output.Tags), nil
}

// BackupListTags lists backup service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func BackupListTags(conn *backup.Backup, identifier string) (KeyValueTags, error) {
	input := &backup.ListTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return BackupKeyValueTags(output.Tags), nil
}

// Cloudhsmv2ListTags lists cloudhsmv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func Cloudhsmv2ListTags(conn *cloudhsmv2.CloudHSMV2, identifier string) (KeyValueTags, error) {
	input := &cloudhsmv2.ListTagsInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return Cloudhsmv2KeyValueTags(output.TagList), nil
}

// CloudwatchListTags lists cloudwatch service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CloudwatchListTags(conn *cloudwatch.CloudWatch, identifier string) (KeyValueTags, error) {
	input := &cloudwatch.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CloudwatchKeyValueTags(output.Tags), nil
}

// CloudwatcheventsListTags lists cloudwatchevents service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CloudwatcheventsListTags(conn *cloudwatchevents.CloudWatchEvents, identifier string) (KeyValueTags, error) {
	input := &cloudwatchevents.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CloudwatcheventsKeyValueTags(output.Tags), nil
}

// CloudwatchlogsListTags lists cloudwatchlogs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CloudwatchlogsListTags(conn *cloudwatchlogs.CloudWatchLogs, identifier string) (KeyValueTags, error) {
	input := &cloudwatchlogs.ListTagsLogGroupInput{
		LogGroupName: aws.String(identifier),
	}

	output, err := conn.ListTagsLogGroup(input)

	if err != nil {
		return New(nil), err
	}

	return CloudwatchlogsKeyValueTags(output.Tags), nil
}

// CodecommitListTags lists codecommit service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CodecommitListTags(conn *codecommit.CodeCommit, identifier string) (KeyValueTags, error) {
	input := &codecommit.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CodecommitKeyValueTags(output.Tags), nil
}

// CodedeployListTags lists codedeploy service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CodedeployListTags(conn *codedeploy.CodeDeploy, identifier string) (KeyValueTags, error) {
	input := &codedeploy.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CodedeployKeyValueTags(output.Tags), nil
}

// CodepipelineListTags lists codepipeline service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CodepipelineListTags(conn *codepipeline.CodePipeline, identifier string) (KeyValueTags, error) {
	input := &codepipeline.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CodepipelineKeyValueTags(output.Tags), nil
}

// CognitoidentityListTags lists cognitoidentity service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CognitoidentityListTags(conn *cognitoidentity.CognitoIdentity, identifier string) (KeyValueTags, error) {
	input := &cognitoidentity.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CognitoidentityKeyValueTags(output.Tags), nil
}

// CognitoidentityproviderListTags lists cognitoidentityprovider service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func CognitoidentityproviderListTags(conn *cognitoidentityprovider.CognitoIdentityProvider, identifier string) (KeyValueTags, error) {
	input := &cognitoidentityprovider.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return CognitoidentityproviderKeyValueTags(output.Tags), nil
}

// ConfigserviceListTags lists configservice service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ConfigserviceListTags(conn *configservice.ConfigService, identifier string) (KeyValueTags, error) {
	input := &configservice.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return ConfigserviceKeyValueTags(output.Tags), nil
}

// DatabasemigrationserviceListTags lists databasemigrationservice service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DatabasemigrationserviceListTags(conn *databasemigrationservice.DatabaseMigrationService, identifier string) (KeyValueTags, error) {
	input := &databasemigrationservice.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return DatabasemigrationserviceKeyValueTags(output.TagList), nil
}

// DatasyncListTags lists datasync service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DatasyncListTags(conn *datasync.DataSync, identifier string) (KeyValueTags, error) {
	input := &datasync.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return DatasyncKeyValueTags(output.Tags), nil
}

// DaxListTags lists dax service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DaxListTags(conn *dax.DAX, identifier string) (KeyValueTags, error) {
	input := &dax.ListTagsInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return DaxKeyValueTags(output.Tags), nil
}

// DevicefarmListTags lists devicefarm service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DevicefarmListTags(conn *devicefarm.DeviceFarm, identifier string) (KeyValueTags, error) {
	input := &devicefarm.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return DevicefarmKeyValueTags(output.Tags), nil
}

// DirectoryserviceListTags lists directoryservice service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DirectoryserviceListTags(conn *directoryservice.DirectoryService, identifier string) (KeyValueTags, error) {
	input := &directoryservice.ListTagsForResourceInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return DirectoryserviceKeyValueTags(output.Tags), nil
}

// DocdbListTags lists docdb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DocdbListTags(conn *docdb.DocDB, identifier string) (KeyValueTags, error) {
	input := &docdb.ListTagsForResourceInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return DocdbKeyValueTags(output.TagList), nil
}

// DynamodbListTags lists dynamodb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func DynamodbListTags(conn *dynamodb.DynamoDB, identifier string) (KeyValueTags, error) {
	input := &dynamodb.ListTagsOfResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsOfResource(input)

	if err != nil {
		return New(nil), err
	}

	return DynamodbKeyValueTags(output.Tags), nil
}

// EcrListTags lists ecr service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func EcrListTags(conn *ecr.ECR, identifier string) (KeyValueTags, error) {
	input := &ecr.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return EcrKeyValueTags(output.Tags), nil
}

// EcsListTags lists ecs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func EcsListTags(conn *ecs.ECS, identifier string) (KeyValueTags, error) {
	input := &ecs.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return EcsKeyValueTags(output.Tags), nil
}

// EfsListTags lists efs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func EfsListTags(conn *efs.EFS, identifier string) (KeyValueTags, error) {
	input := &efs.DescribeTagsInput{
		FileSystemId: aws.String(identifier),
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return New(nil), err
	}

	return EfsKeyValueTags(output.Tags), nil
}

// EksListTags lists eks service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func EksListTags(conn *eks.EKS, identifier string) (KeyValueTags, error) {
	input := &eks.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return EksKeyValueTags(output.Tags), nil
}

// ElasticacheListTags lists elasticache service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ElasticacheListTags(conn *elasticache.ElastiCache, identifier string) (KeyValueTags, error) {
	input := &elasticache.ListTagsForResourceInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return ElasticacheKeyValueTags(output.TagList), nil
}

// ElasticbeanstalkListTags lists elasticbeanstalk service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ElasticbeanstalkListTags(conn *elasticbeanstalk.ElasticBeanstalk, identifier string) (KeyValueTags, error) {
	input := &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return ElasticbeanstalkKeyValueTags(output.ResourceTags), nil
}

// ElasticsearchserviceListTags lists elasticsearchservice service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ElasticsearchserviceListTags(conn *elasticsearchservice.ElasticsearchService, identifier string) (KeyValueTags, error) {
	input := &elasticsearchservice.ListTagsInput{
		ARN: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return ElasticsearchserviceKeyValueTags(output.TagList), nil
}

// Elbv2ListTags lists elbv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func Elbv2ListTags(conn *elbv2.ELBV2, identifier string) (KeyValueTags, error) {
	input := &elbv2.DescribeTagsInput{
		ResourceArns: aws.StringSlice([]string{identifier}),
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return New(nil), err
	}

	return Elbv2KeyValueTags(output.TagDescriptions[0].Tags), nil
}

// FirehoseListTags lists firehose service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func FirehoseListTags(conn *firehose.Firehose, identifier string) (KeyValueTags, error) {
	input := &firehose.ListTagsForDeliveryStreamInput{
		DeliveryStreamName: aws.String(identifier),
	}

	output, err := conn.ListTagsForDeliveryStream(input)

	if err != nil {
		return New(nil), err
	}

	return FirehoseKeyValueTags(output.Tags), nil
}

// FsxListTags lists fsx service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func FsxListTags(conn *fsx.FSx, identifier string) (KeyValueTags, error) {
	input := &fsx.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return FsxKeyValueTags(output.Tags), nil
}

// GlueListTags lists glue service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GlueListTags(conn *glue.Glue, identifier string) (KeyValueTags, error) {
	input := &glue.GetTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.GetTags(input)

	if err != nil {
		return New(nil), err
	}

	return GlueKeyValueTags(output.Tags), nil
}

// GuarddutyListTags lists guardduty service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GuarddutyListTags(conn *guardduty.GuardDuty, identifier string) (KeyValueTags, error) {
	input := &guardduty.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return GuarddutyKeyValueTags(output.Tags), nil
}

// InspectorListTags lists inspector service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func InspectorListTags(conn *inspector.Inspector, identifier string) (KeyValueTags, error) {
	input := &inspector.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return InspectorKeyValueTags(output.Tags), nil
}

// IotListTags lists iot service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func IotListTags(conn *iot.IoT, identifier string) (KeyValueTags, error) {
	input := &iot.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return IotKeyValueTags(output.Tags), nil
}

// IotanalyticsListTags lists iotanalytics service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func IotanalyticsListTags(conn *iotanalytics.IoTAnalytics, identifier string) (KeyValueTags, error) {
	input := &iotanalytics.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return IotanalyticsKeyValueTags(output.Tags), nil
}

// IoteventsListTags lists iotevents service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func IoteventsListTags(conn *iotevents.IoTEvents, identifier string) (KeyValueTags, error) {
	input := &iotevents.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return IoteventsKeyValueTags(output.Tags), nil
}

// KafkaListTags lists kafka service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func KafkaListTags(conn *kafka.Kafka, identifier string) (KeyValueTags, error) {
	input := &kafka.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return KafkaKeyValueTags(output.Tags), nil
}

// KinesisanalyticsListTags lists kinesisanalytics service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func KinesisanalyticsListTags(conn *kinesisanalytics.KinesisAnalytics, identifier string) (KeyValueTags, error) {
	input := &kinesisanalytics.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return KinesisanalyticsKeyValueTags(output.Tags), nil
}

// Kinesisanalyticsv2ListTags lists kinesisanalyticsv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func Kinesisanalyticsv2ListTags(conn *kinesisanalyticsv2.KinesisAnalyticsV2, identifier string) (KeyValueTags, error) {
	input := &kinesisanalyticsv2.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return Kinesisanalyticsv2KeyValueTags(output.Tags), nil
}

// KmsListTags lists kms service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func KmsListTags(conn *kms.KMS, identifier string) (KeyValueTags, error) {
	input := &kms.ListResourceTagsInput{
		KeyId: aws.String(identifier),
	}

	output, err := conn.ListResourceTags(input)

	if err != nil {
		return New(nil), err
	}

	return KmsKeyValueTags(output.Tags), nil
}

// LambdaListTags lists lambda service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func LambdaListTags(conn *lambda.Lambda, identifier string) (KeyValueTags, error) {
	input := &lambda.ListTagsInput{
		Resource: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return LambdaKeyValueTags(output.Tags), nil
}

// LicensemanagerListTags lists licensemanager service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func LicensemanagerListTags(conn *licensemanager.LicenseManager, identifier string) (KeyValueTags, error) {
	input := &licensemanager.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return LicensemanagerKeyValueTags(output.Tags), nil
}

// MediaconnectListTags lists mediaconnect service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func MediaconnectListTags(conn *mediaconnect.MediaConnect, identifier string) (KeyValueTags, error) {
	input := &mediaconnect.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return MediaconnectKeyValueTags(output.Tags), nil
}

// MedialiveListTags lists medialive service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func MedialiveListTags(conn *medialive.MediaLive, identifier string) (KeyValueTags, error) {
	input := &medialive.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return MedialiveKeyValueTags(output.Tags), nil
}

// MediapackageListTags lists mediapackage service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func MediapackageListTags(conn *mediapackage.MediaPackage, identifier string) (KeyValueTags, error) {
	input := &mediapackage.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return MediapackageKeyValueTags(output.Tags), nil
}

// MediastoreListTags lists mediastore service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func MediastoreListTags(conn *mediastore.MediaStore, identifier string) (KeyValueTags, error) {
	input := &mediastore.ListTagsForResourceInput{
		Resource: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return MediastoreKeyValueTags(output.Tags), nil
}

// MqListTags lists mq service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func MqListTags(conn *mq.MQ, identifier string) (KeyValueTags, error) {
	input := &mq.ListTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return MqKeyValueTags(output.Tags), nil
}

// NeptuneListTags lists neptune service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func NeptuneListTags(conn *neptune.Neptune, identifier string) (KeyValueTags, error) {
	input := &neptune.ListTagsForResourceInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return NeptuneKeyValueTags(output.TagList), nil
}

// OpsworksListTags lists opsworks service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func OpsworksListTags(conn *opsworks.OpsWorks, identifier string) (KeyValueTags, error) {
	input := &opsworks.ListTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return OpsworksKeyValueTags(output.Tags), nil
}

// OrganizationsListTags lists organizations service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func OrganizationsListTags(conn *organizations.Organizations, identifier string) (KeyValueTags, error) {
	input := &organizations.ListTagsForResourceInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return OrganizationsKeyValueTags(output.Tags), nil
}

// QldbListTags lists qldb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func QldbListTags(conn *qldb.QLDB, identifier string) (KeyValueTags, error) {
	input := &qldb.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return QldbKeyValueTags(output.Tags), nil
}

// RdsListTags lists rds service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func RdsListTags(conn *rds.RDS, identifier string) (KeyValueTags, error) {
	input := &rds.ListTagsForResourceInput{
		ResourceName: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return RdsKeyValueTags(output.TagList), nil
}

// ResourcegroupsListTags lists resourcegroups service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ResourcegroupsListTags(conn *resourcegroups.ResourceGroups, identifier string) (KeyValueTags, error) {
	input := &resourcegroups.GetTagsInput{
		Arn: aws.String(identifier),
	}

	output, err := conn.GetTags(input)

	if err != nil {
		return New(nil), err
	}

	return ResourcegroupsKeyValueTags(output.Tags), nil
}

// Route53resolverListTags lists route53resolver service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func Route53resolverListTags(conn *route53resolver.Route53Resolver, identifier string) (KeyValueTags, error) {
	input := &route53resolver.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return Route53resolverKeyValueTags(output.Tags), nil
}

// SagemakerListTags lists sagemaker service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SagemakerListTags(conn *sagemaker.SageMaker, identifier string) (KeyValueTags, error) {
	input := &sagemaker.ListTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return New(nil), err
	}

	return SagemakerKeyValueTags(output.Tags), nil
}

// SecurityhubListTags lists securityhub service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SecurityhubListTags(conn *securityhub.SecurityHub, identifier string) (KeyValueTags, error) {
	input := &securityhub.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return SecurityhubKeyValueTags(output.Tags), nil
}

// SfnListTags lists sfn service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SfnListTags(conn *sfn.SFN, identifier string) (KeyValueTags, error) {
	input := &sfn.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return SfnKeyValueTags(output.Tags), nil
}

// SnsListTags lists sns service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SnsListTags(conn *sns.SNS, identifier string) (KeyValueTags, error) {
	input := &sns.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return SnsKeyValueTags(output.Tags), nil
}

// SqsListTags lists sqs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SqsListTags(conn *sqs.SQS, identifier string) (KeyValueTags, error) {
	input := &sqs.ListQueueTagsInput{
		QueueUrl: aws.String(identifier),
	}

	output, err := conn.ListQueueTags(input)

	if err != nil {
		return New(nil), err
	}

	return SqsKeyValueTags(output.Tags), nil
}

// SsmListTags lists ssm service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SsmListTags(conn *ssm.SSM, identifier string, resourceType string) (KeyValueTags, error) {
	input := &ssm.ListTagsForResourceInput{
		ResourceId:   aws.String(identifier),
		ResourceType: aws.String(resourceType),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return SsmKeyValueTags(output.TagList), nil
}

// StoragegatewayListTags lists storagegateway service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func StoragegatewayListTags(conn *storagegateway.StorageGateway, identifier string) (KeyValueTags, error) {
	input := &storagegateway.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return StoragegatewayKeyValueTags(output.Tags), nil
}

// SwfListTags lists swf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func SwfListTags(conn *swf.SWF, identifier string) (KeyValueTags, error) {
	input := &swf.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return SwfKeyValueTags(output.Tags), nil
}

// TransferListTags lists transfer service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func TransferListTags(conn *transfer.Transfer, identifier string) (KeyValueTags, error) {
	input := &transfer.ListTagsForResourceInput{
		Arn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return TransferKeyValueTags(output.Tags), nil
}

// WafListTags lists waf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func WafListTags(conn *waf.WAF, identifier string) (KeyValueTags, error) {
	input := &waf.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return WafKeyValueTags(output.TagInfoForResource.TagList), nil
}

// WafregionalListTags lists wafregional service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func WafregionalListTags(conn *wafregional.WAFRegional, identifier string) (KeyValueTags, error) {
	input := &waf.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return New(nil), err
	}

	return WafregionalKeyValueTags(output.TagInfoForResource.TagList), nil
}

// WorkspacesListTags lists workspaces service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func WorkspacesListTags(conn *workspaces.WorkSpaces, identifier string) (KeyValueTags, error) {
	input := &workspaces.DescribeTagsInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return New(nil), err
	}

	return WorkspacesKeyValueTags(output.TagList), nil
}
