package pike

import "fmt"

// GetGCPPermissions for GCP resources.
func GetGCPPermissions(result ResourceV2) ([]string, error) {
	var (
		err         error
		Permissions []string
	)

	if result.TypeName == "resource" {
		Permissions, err = GetGCPResourcePermissions(result)
		if err != nil {
			return Permissions, err
		}
	} else {
		Permissions, err = GetGCPDataPermissions(result)
		if err != nil {
			return Permissions, err
		}
	}

	return Permissions, err
}

// GetGCPResourcePermissions looks up permissions required for resources.
func GetGCPResourcePermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := GCPLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("%s not implemented", result.Name)
	}

	return Permissions, err
}

func GCPLookup(result string) interface{} {
	TFLookup := map[string]interface{}{
		"google_artifact_registry_repository":                     googleArtifactRegistryRepository,
		"google_artifact_registry_repository_iam_binding":         googleArtifactRegistryRepositoryIamBinding,
		"google_artifact_registry_repository_iam_member":          googleArtifactRegistryRepositoryIamMember,
		"google_artifact_registry_repository_iam_policy":          googleArtifactRegistryRepositoryIamPolicy,
		"google_bigquery_dataset":                                 googleBigqueryDataset,
		"google_bigquery_job":                                     googleBigqueryJob,
		"google_bigquery_table":                                   placeholder,
		"google_bigtable_instance":                                googleBigtableInstance,
		"google_bigtable_instance_iam_binding":                    googleBigTableInstanceIam,
		"google_bigtable_instance_iam_member":                     googleBigTableInstanceIam,
		"google_bigtable_instance_iam_policy":                     googleBigTableInstanceIam,
		"google_bigtable_table":                                   googleBigtableTable,
		"google_bigtable_table_iam_binding":                       googleBigTableTableIam,
		"google_bigtable_table_iam_member":                        googleBigTableTableIam,
		"google_bigtable_table_iam_policy":                        googleBigTableTableIam,
		"google_cloud_run_v2_job":                                 googleCloudRunV2Job,
		"google_cloud_run_v2_service":                             googleCloudRunV2Job,
		"google_cloud_scheduler_job":                              googleCloudSchedulerJob,
		"google_cloudfunctions_function":                          googleCloudfunctionsFunction,
		"google_cloudfunctions_function_iam_member":               googleCloudfunctionsFunctionIamPolicy,
		"google_cloudfunctions_function_iam_policy":               googleCloudfunctionsFunctionIamPolicy,
		"google_compute_address":                                  googleComputeAddress,
		"google_compute_firewall":                                 googleComputeFirewall,
		"google_compute_global_address":                           googleComputeGlobalAddress,
		"google_compute_instance":                                 googleComputeInstance,
		"google_compute_instance_template":                        googleComputeInstanceTemplate,
		"google_compute_network":                                  googleComputeNetwork,
		"google_compute_project_metadata_item":                    googleComputeProjectMetadataItem,
		"google_compute_region_ssl_certificate":                   googleComputeRegionSslCertificate,
		"google_compute_security_policy":                          googleComputeSecurityPolicy,
		"google_compute_subnetwork":                               googleComputeSubnetwork,
		"google_container_cluster":                                googleContainerCluster,
		"google_container_node_pool":                              googleContainerNodePool,
		"google_dns_managed_zone":                                 googleDnsmanagedZone,
		"google_dns_policy":                                       googleDNSPolicy,
		"google_dns_record_set":                                   googleDNSRecordSet,
		"google_kms_crypto_key":                                   googleKmsCryptoKey,
		"google_kms_crypto_key_iam_binding":                       googlekmsCryptoKeyIamBinding,
		"google_kms_crypto_key_iam_member":                        googlekmsCryptoKeyIamMember,
		"google_kms_crypto_key_iam_policy":                        googlekmsCryptoKeyIamPolicy,
		"google_kms_key_ring":                                     googleKmsKeyRing,
		"google_project_iam_binding":                              googleProjectIamBinding,
		"google_project_iam_custom_role":                          googleProjectIamCustomRole,
		"google_project_iam_member":                               googleProjectIamBinding,
		"google_project_service":                                  googleProjectService,
		"google_project_service_identity":                         placeholder,
		"google_pubsub_lite_reservation":                          googlePubsubLiteReservation,
		"google_pubsub_lite_subscription":                         googlePubsubLiteSubscription,
		"google_pubsub_lite_topic":                                googlePubsubLiteTopic,
		"google_pubsub_schema":                                    googlePubsubSchema,
		"google_pubsub_subscription":                              googlePubsubSubscription,
		"google_pubsub_topic":                                     googlePubsubTopic,
		"google_pubsub_topic_iam_binding":                         googlePubsubTopicIam,
		"google_pubsub_topic_iam_member":                          googlePubsubTopicIam,
		"google_pubsub_topic_iam_policy":                          googlePubsubTopicIam,
		"google_redis_instance":                                   googleRedisInstance,
		"google_secret_manager_secret":                            googleSecretManagerSecret,
		"google_secret_manager_secret_iam_binding":                googleSecretManagerSecretIam,
		"google_secret_manager_secret_iam_member":                 googleSecretManagerSecretIam,
		"google_secret_manager_secret_iam_policy":                 googleSecretManagerSecretIam,
		"google_secret_manager_secret_version":                    googleSecretManagerSecretVersion,
		"google_service_account":                                  googleServiceAccount,
		"google_service_account_iam_binding":                      googleServiceAccountIamBinding,
		"google_service_account_iam_member":                       googleServiceAccountIamMember,
		"google_service_account_iam_policy":                       googleServiceAccountIamPolicy,
		"google_service_account_key":                              googleServiceAccountKey,
		"google_service_networking_connection":                    googleServiceNetworkingConnection,
		"google_sourcerepo_repository":                            googleSourcerepoRepository,
		"google_sql_database":                                     googleSQLDatabase,
		"google_sql_database_instance":                            googleSQLDatabaseInstance,
		"google_sql_user":                                         googleSQLUser,
		"google_storage_bucket":                                   googleStorageBucket,
		"google_storage_bucket_acl":                               googleStorageBucketACL,
		"google_storage_bucket_iam_binding":                       googleStorageBucketIamBinding,
		"google_storage_bucket_object":                            googleStorageBucketObject,
		"google_storage_bucket_access_control":                    googleStorageBucketAccessControl,
		"google_storage_bucket_iam_member":                        googleStorageBucketIamMember,
		"google_storage_bucket_iam_policy":                        googleStorageBucketIamPolicy,
		"google_storage_default_object_access_control":            googleStorageDefaultObjectAccessControl,
		"google_storage_default_object_acl":                       googleStorageDefaultObjectACL,
		"google_storage_hmac_key":                                 googleStorageHmacKey,
		"google_storage_insights_report_config":                   googleStorageInsightsReportConfig,
		"google_storage_object_access_control":                    googleStorageObjectAccessControl,
		"google_cloudbuild_trigger":                               googleCloudbuildTrigger,
		"google_service_directory_endpoint":                       googleServiceDirectoryEndpoint,
		"google_service_directory_namespace":                      googleServiceDirectoryNamespace,
		"google_service_directory_namespace_iam_binding":          googleServiceDirectoryNamespaceIamBinding,
		"google_service_directory_namespace_iam_member":           googleServiceDirectoryNamespaceIamMember,
		"google_service_directory_namespace_iam_policy":           googleServiceDirectoryNamespaceIamPolicy,
		"google_service_directory_service":                        googleServiceDirectoryService,
		"google_service_directory_service_iam_binding":            googleServiceDirectoryServiceIamBinding,
		"google_service_directory_service_iam_member":             googleServiceDirectoryServiceIamMember,
		"google_service_directory_service_iam_policy":             googleServiceDirectoryServiceIamPolicy,
		"google_access_context_manager_access_level":              googleAccessContextManagerAccessLevel,
		"google_access_context_manager_access_levels":             googleAccessContextManagerAccessLevels,
		"google_access_context_manager_access_policy":             googleAccessContextManagerAccessPolicy,
		"google_access_context_manager_access_policy_iam_binding": googleAccessContextManagerAccessPolicyIam,
		"google_access_context_manager_access_policy_iam_member":  googleAccessContextManagerAccessPolicyIam,
		"google_access_context_manager_access_policy_iam_policy":  googleAccessContextManagerAccessPolicyIam,
		"google_access_context_manager_authorized_orgs_desc":      googleAccessContextManagerAuthorizedOrgsDesc,
		"google_access_context_manager_gcp_user_access_binding":   googleAccessContextManagerGcpUserAccessBinding,
		"google_access_context_manager_service_perimeter":         googleAccessContextManagerServicePerimeter,
		"google_access_context_manager_service_perimeters":        googleAccessContextManagerServicePerimeters,
		"google_alloydb_backup":                                   googleAlloydbBackup,
		"google_alloydb_cluster":                                  googleAlloydbCluster,
		"google_alloydb_instance":                                 googleAlloydbInstance,
		"google_alloydb_user":                                     googleAlloydbUser,
		"google_firebase_android_app":                             googleFirebaseAndroidApp,
		"google_firebase_apple_app":                               googleFirebaseAppleApp,
		"google_firebase_database_instance":                       googleFirebaseDatabaseInstance,
		"google_firebase_hosting_channel":                         googleFirebaseHostingSite,
		"google_firebase_hosting_custom_domain":                   googleFirebaseHostingSite,
		"google_firebase_hosting_release":                         googleFirebaseHostingSite,
		"google_firebase_hosting_site":                            googleFirebaseHostingSite,
		"google_firebase_hosting_version":                         googleFirebaseHostingSite,
		"google_firebase_project":                                 googleFirebaseProject,
		"google_firebase_storage_bucket":                          googleFirebaseStorageBucket,
		"google_firebase_web_app":                                 googleFirebaseWebApp,
		"google_firebaserules_release":                            googleFirebaserulesRelease,
		"google_firebaserules_ruleset":                            googleFirebaserulesRuleset,
		"google_bigtable_app_profile":                             googleBigtableAppProfile,
		"google_bigtable_gc_policy":                               googleBigtableGcPolicy,
		"google_api_gateway_api":                                  googleApiGatewayApi,
		"google_api_gateway_api_config":                           googleApiGatewayApiConfig,
		"google_api_gateway_api_config_iam_binding":               googleApiGatewayApiConfigIam,
		"google_api_gateway_api_config_iam_member":                googleApiGatewayApiConfigIam,
		"google_api_gateway_api_config_iam_policy":                googleApiGatewayApiConfigIam,
		"google_api_gateway_api_iam_binding":                      googleApiGatewayApiIam,
		"google_api_gateway_api_iam_member":                       googleApiGatewayApiIam,
		"google_api_gateway_api_iam_policy":                       googleApiGatewayApiIam,
		"google_api_gateway_gateway":                              googleApiGatewayGateway,
		"google_api_gateway_gateway_iam_binding":                  googleApiGatewayGatewayIam,
		"google_api_gateway_gateway_iam_member":                   googleApiGatewayGatewayIam,
		"google_api_gateway_gateway_iam_policy":                   googleApiGatewayGatewayIam,
	}

	return TFLookup[result]
}
