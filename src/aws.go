package pike

import (
	"encoding/json"
	"log"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result ResourceV2) []string {

	var Permissions []string
	if result.TypeName == "resource" {
		Permissions = GetAWSResourcePermissions(result)
	} else {
		Permissions = GetAWSDataPermissions(result)
	}

	return Permissions
}

// GetAWSResourcePermissions looks up permissions required for resources
func GetAWSResourcePermissions(result ResourceV2) []string {

	TFLookup := map[string]interface{}{
		"aws_s3_bucket":            awsS3Bucket,
		"aws_s3_bucket_acl":        awsS3BucketACL,
		"aws_s3_bucket_versioning": awsS3BucketVersioning,
		"aws_s3_bucket_server_side_encryption_configuration": awsS3BucketServerSideEncryptionConfiguration,
		"aws_s3_bucket_public_access_block":                  awsS3BucketPublicAccessBlock,
		"aws_instance":                                       awsInstance,
		"aws_security_group":                                 awsSecurityGroup,
		"aws_security_group_rule":                            awsSecurityGroupRule,
		"aws_lambda_function":                                awsLambdaFunction,
		"aws_vpc":                                            awsVpc,
		"aws_subnet":                                         awsSubnet,
		"aws_network_acl":                                    awsNetworkACL,
		"aws_kms_key":                                        awsKmsKey,
		"aws_iam_role":                                       awsIamRole,
		"aws_iam_role_policy":                                awsIamRolePolicy,
		"aws_iam_role_policy_attachment":                     awsIamRolePolicyAttachment,
		"aws_iam_policy":                                     awsIamPolicy,
		"aws_iam_instance_profile":                           awsIamInstanceProfile,
		"aws_iam_access_key":                                 awsIamAccessKey,
		"aws_iam_group":                                      awsIamGroup,
		"aws_iam_group_membership":                           awsIamGroupMembership,
		"aws_iam_group_policy":                               awsIamGroupPolicy,
		"aws_iam_group_policy_attachment":                    awsIamGroupPolicyAttachment,
		"aws_iam_policy_attachment":                          awsIamPolicyAttachment,
		"aws_iam_service_linked_role":                        awsIamServiceLinkedRole,
		"aws_iam_user":                                       awsIamUser,
		"aws_iam_user_login_profile":                         awsIamUserLoginProfile,
		"aws_iam_user_policy":                                awsIamUserPolicy,
		"aws_iam_user_policy_attachment":                     awsIamUserPolicyAttachment,
		"aws_mq_broker":                                      awsMqBroker,
		"aws_mq_configuration":                               awsMqConfiguration,
		"aws_cloudwatch_log_group":                           awsCloudwatchLogGroup,
		"aws_cloudwatch_event_rule":                          awsCloudwatchEventRule,
		"aws_cloudwatch_event_target":                        awsCloudwatchEventTarget,
		"aws_cloudwatch_log_metric_filter":                   awsCloudwatchLogMetricFilter,
		"aws_cloudwatch_log_resource_policy":                 awsCloudwatchLogResourcePolicy,
		"aws_cloudwatch_log_subscription_filter":             awsCloudwatchLogSubscriptionFilter,
		"aws_cloudwatch_metric_alarm":                        awsCloudwatchMetricAlarm,
		"aws_route53_record":                                 awsRoute53Record,
		"aws_sns_topic":                                      awsSnsTopic,
		"aws_key_pair":                                       awsKeyPair,
		"aws_db_instance":                                    awsDbInstance,
		"aws_dynamodb_table":                                 awsDynamodbTable,
		"aws_ssm_parameter":                                  awsSsmParameter,
		"aws_route":                                          awsRoute,
	}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetPermissionMap Anonymous parsing
func GetPermissionMap(raw []byte, attributes []string) []string {
	var mappings []interface{}
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		log.Print(err)
	}
	temp := mappings[0].(map[string]interface{})
	myAttributes := temp["attributes"].(map[string]interface{})

	var found []string

	for _, attribute := range attributes {
		if myAttributes[attribute] != nil {
			entries := myAttributes[attribute].([]interface{})
			for _, entry := range entries {
				found = append(found, entry.(string))
			}
		}
	}

	actions := []string{"apply", "plan", "modify", "destroy"}

	for _, action := range actions {
		if temp[action] != nil {
			myentries := temp[action].([]interface{})
			for _, entry := range myentries {
				found = append(found, entry.(string))
			}
		}
	}

	return found
}
