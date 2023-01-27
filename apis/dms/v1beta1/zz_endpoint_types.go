/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ElasticsearchSettingsObservation struct {
}

type ElasticsearchSettingsParameters struct {

	// Endpoint for the OpenSearch cluster.
	// +kubebuilder:validation:Required
	EndpointURI *string `json:"endpointUri" tf:"endpoint_uri,omitempty"`

	// Maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster. Default is 300.
	// +kubebuilder:validation:Optional
	ErrorRetryDuration *float64 `json:"errorRetryDuration,omitempty" tf:"error_retry_duration,omitempty"`

	// Maximum percentage of records that can fail to be written before a full load operation stops. Default is 10.
	// +kubebuilder:validation:Optional
	FullLoadErrorPercentage *float64 `json:"fullLoadErrorPercentage,omitempty" tf:"full_load_error_percentage,omitempty"`

	// ARN of the IAM Role with permissions to write to the OpenSearch cluster.
	// +kubebuilder:validation:Required
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn" tf:"service_access_role_arn,omitempty"`
}

type EndpointObservation struct {

	// ARN for the endpoint.
	EndpointArn *string `json:"endpointArn,omitempty" tf:"endpoint_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EndpointParameters struct {

	// ARN for the certificate.
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Name of the endpoint database.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Configuration block for OpenSearch settings. See below.
	// +kubebuilder:validation:Optional
	ElasticsearchSettings []ElasticsearchSettingsParameters `json:"elasticsearchSettings,omitempty" tf:"elasticsearch_settings,omitempty"`

	// Type of endpoint. Valid values are source, target.
	// +kubebuilder:validation:Required
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// Type of engine for the endpoint. Valid values are aurora, aurora-postgresql, azuredb, db2, docdb, dynamodb, elasticsearch, kafka, kinesis, mariadb, mongodb, mysql, opensearch, oracle, postgres, redshift, s3, sqlserver, sybase. Please note that some of engine names are available only for target endpoint type (e.g. redshift).
	// +kubebuilder:validation:Required
	EngineName *string `json:"engineName" tf:"engine_name,omitempty"`

	// Additional attributes associated with the connection. For available attributes see Using Extra Connection Attributes with AWS Database Migration Service.
	// +kubebuilder:validation:Optional
	ExtraConnectionAttributes *string `json:"extraConnectionAttributes,omitempty" tf:"extra_connection_attributes,omitempty"`

	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for kms_key_arn, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Configuration block for Kafka settings. See below.
	// +kubebuilder:validation:Optional
	KafkaSettings []KafkaSettingsParameters `json:"kafkaSettings,omitempty" tf:"kafka_settings,omitempty"`

	// Configuration block for Kinesis settings. See below.
	// +kubebuilder:validation:Optional
	KinesisSettings []KinesisSettingsParameters `json:"kinesisSettings,omitempty" tf:"kinesis_settings,omitempty"`

	// Configuration block for MongoDB settings. See below.
	// +kubebuilder:validation:Optional
	MongodbSettings []MongodbSettingsParameters `json:"mongodbSettings,omitempty" tf:"mongodb_settings,omitempty"`

	// Password to be used to login to the endpoint database.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port used by the endpoint database.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	RedisSettings []RedisSettingsParameters `json:"redisSettings,omitempty" tf:"redis_settings,omitempty"`

	// Configuration block for Redshift settings. See below.
	// +kubebuilder:validation:Optional
	RedshiftSettings []RedshiftSettingsParameters `json:"redshiftSettings,omitempty" tf:"redshift_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for S3 settings. See below.
	// +kubebuilder:validation:Optional
	S3Settings []S3SettingsParameters `json:"s3Settings,omitempty" tf:"s3_settings,omitempty"`

	// SSL mode to use for the connection. Valid values are none, require, verify-ca, verify-full
	// +kubebuilder:validation:Optional
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`

	// ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in SecretsManagerSecret.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SecretsManagerAccessRoleArn *string `json:"secretsManagerAccessRoleArn,omitempty" tf:"secrets_manager_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate secretsManagerAccessRoleArn.
	// +kubebuilder:validation:Optional
	SecretsManagerAccessRoleArnRef *v1.Reference `json:"secretsManagerAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate secretsManagerAccessRoleArn.
	// +kubebuilder:validation:Optional
	SecretsManagerAccessRoleArnSelector *v1.Selector `json:"secretsManagerAccessRoleArnSelector,omitempty" tf:"-"`

	// Full ARN, partial ARN, or friendly name of the SecretsManagerSecret that contains the endpoint connection details. Supported only for engine_name as aurora, aurora-postgresql, mariadb, mongodb, mysql, oracle, postgres, redshift or sqlserver.
	// +kubebuilder:validation:Optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// Host name of the server.
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// ARN used by the service access IAM role for dynamodb endpoints.
	// +kubebuilder:validation:Optional
	ServiceAccessRole *string `json:"serviceAccessRole,omitempty" tf:"service_access_role,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// User name to be used to login to the endpoint database.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type KafkaSettingsObservation struct {
}

type KafkaSettingsParameters struct {

	// Kafka broker location. Specify in the form broker-hostname-or-ip:port.
	// +kubebuilder:validation:Required
	Broker *string `json:"broker" tf:"broker,omitempty"`

	// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output. Default is false.
	// +kubebuilder:validation:Optional
	IncludeControlDetails *bool `json:"includeControlDetails,omitempty" tf:"include_control_details,omitempty"`

	// Include NULL and empty columns for records migrated to the endpoint. Default is false.
	// +kubebuilder:validation:Optional
	IncludeNullAndEmpty *bool `json:"includeNullAndEmpty,omitempty" tf:"include_null_and_empty,omitempty"`

	// Shows the partition value within the Kafka message output unless the partition type is schema-table-type. Default is false.
	// +kubebuilder:validation:Optional
	IncludePartitionValue *bool `json:"includePartitionValue,omitempty" tf:"include_partition_value,omitempty"`

	// Includes any data definition language (DDL) operations that change the table in the control data, such as rename-table, drop-table, add-column, drop-column, and rename-column. Default is false.
	// +kubebuilder:validation:Optional
	IncludeTableAlterOperations *bool `json:"includeTableAlterOperations,omitempty" tf:"include_table_alter_operations,omitempty"`

	// Provides detailed transaction information from the source database. This information includes a commit timestamp, a log position, and values for transaction_id, previous transaction_id, and transaction_record_id (the record offset within a transaction). Default is false.
	// +kubebuilder:validation:Optional
	IncludeTransactionDetails *bool `json:"includeTransactionDetails,omitempty" tf:"include_transaction_details,omitempty"`

	// Output format for the records created on the endpoint. Message format is JSON (default) or JSON_UNFORMATTED (a single line with no tab).
	// +kubebuilder:validation:Optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// Maximum size in bytes for records created on the endpoint Default is 1,000,000.
	// +kubebuilder:validation:Optional
	MessageMaxBytes *float64 `json:"messageMaxBytes,omitempty" tf:"message_max_bytes,omitempty"`

	// Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the no_hex_prefix endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.
	// +kubebuilder:validation:Optional
	NoHexPrefix *bool `json:"noHexPrefix,omitempty" tf:"no_hex_prefix,omitempty"`

	// Prefixes schema and table names to partition values, when the partition type is primary-key-type. Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. Default is false.
	// +kubebuilder:validation:Optional
	PartitionIncludeSchemaTable *bool `json:"partitionIncludeSchemaTable,omitempty" tf:"partition_include_schema_table,omitempty"`

	// ARN for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
	// +kubebuilder:validation:Optional
	SSLCACertificateArn *string `json:"sslCaCertificateArn,omitempty" tf:"ssl_ca_certificate_arn,omitempty"`

	// ARN of the client certificate used to securely connect to a Kafka target endpoint.
	// +kubebuilder:validation:Optional
	SSLClientCertificateArn *string `json:"sslClientCertificateArn,omitempty" tf:"ssl_client_certificate_arn,omitempty"`

	// ARN for the client private key used to securely connect to a Kafka target endpoint.
	// +kubebuilder:validation:Optional
	SSLClientKeyArn *string `json:"sslClientKeyArn,omitempty" tf:"ssl_client_key_arn,omitempty"`

	// Password for the client private key used to securely connect to a Kafka target endpoint.
	// +kubebuilder:validation:Optional
	SSLClientKeyPasswordSecretRef *v1.SecretKeySelector `json:"sslClientKeyPasswordSecretRef,omitempty" tf:"-"`

	// Secure password you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	// +kubebuilder:validation:Optional
	SaslPasswordSecretRef *v1.SecretKeySelector `json:"saslPasswordSecretRef,omitempty" tf:"-"`

	// Secure user name you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	// +kubebuilder:validation:Optional
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS). Options include ssl-encryption, ssl-authentication, and sasl-ssl. sasl-ssl requires sasl_username and sasl_password.
	// +kubebuilder:validation:Optional
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`

	// Kafka topic for migration. Default is kafka-default-topic.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type KinesisSettingsObservation struct {
}

type KinesisSettingsParameters struct {

	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output. Default is false.
	// +kubebuilder:validation:Optional
	IncludeControlDetails *bool `json:"includeControlDetails,omitempty" tf:"include_control_details,omitempty"`

	// Include NULL and empty columns in the target. Default is false.
	// +kubebuilder:validation:Optional
	IncludeNullAndEmpty *bool `json:"includeNullAndEmpty,omitempty" tf:"include_null_and_empty,omitempty"`

	// Shows the partition value within the Kinesis message output, unless the partition type is schema-table-type. Default is false.
	// +kubebuilder:validation:Optional
	IncludePartitionValue *bool `json:"includePartitionValue,omitempty" tf:"include_partition_value,omitempty"`

	// Includes any data definition language (DDL) operations that change the table in the control data. Default is false.
	// +kubebuilder:validation:Optional
	IncludeTableAlterOperations *bool `json:"includeTableAlterOperations,omitempty" tf:"include_table_alter_operations,omitempty"`

	// Provides detailed transaction information from the source database. Default is false.
	// +kubebuilder:validation:Optional
	IncludeTransactionDetails *bool `json:"includeTransactionDetails,omitempty" tf:"include_transaction_details,omitempty"`

	// Output format for the records created. Default is json. Valid values are json and json-unformatted (a single line with no tab).
	// +kubebuilder:validation:Optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// Prefixes schema and table names to partition values, when the partition type is primary-key-type. Default is false.
	// +kubebuilder:validation:Optional
	PartitionIncludeSchemaTable *bool `json:"partitionIncludeSchemaTable,omitempty" tf:"partition_include_schema_table,omitempty"`

	// ARN of the IAM Role with permissions to write to the Kinesis data stream.
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`

	// ARN of the Kinesis data stream.
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type MongodbSettingsObservation struct {
}

type MongodbSettingsParameters struct {

	// Authentication mechanism to access the MongoDB source endpoint. Default is default.
	// +kubebuilder:validation:Optional
	AuthMechanism *string `json:"authMechanism,omitempty" tf:"auth_mechanism,omitempty"`

	// Authentication database name. Not used when auth_type is no. Default is admin.
	// +kubebuilder:validation:Optional
	AuthSource *string `json:"authSource,omitempty" tf:"auth_source,omitempty"`

	// Authentication type to access the MongoDB source endpoint. Default is password.
	// +kubebuilder:validation:Optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// Number of documents to preview to determine the document organization. Use this setting when nesting_level is set to one. Default is 1000.
	// +kubebuilder:validation:Optional
	DocsToInvestigate *string `json:"docsToInvestigate,omitempty" tf:"docs_to_investigate,omitempty"`

	// Document ID. Use this setting when nesting_level is set to none. Default is false.
	// +kubebuilder:validation:Optional
	ExtractDocID *string `json:"extractDocId,omitempty" tf:"extract_doc_id,omitempty"`

	// Specifies either document or table mode. Default is none. Valid values are one (table mode) and none (document mode).
	// +kubebuilder:validation:Optional
	NestingLevel *string `json:"nestingLevel,omitempty" tf:"nesting_level,omitempty"`
}

type RedisSettingsObservation struct {
}

type RedisSettingsParameters struct {

	// The password provided with the auth-role and auth-token options of the AuthType setting for a Redis target endpoint.
	// +kubebuilder:validation:Optional
	AuthPasswordSecretRef *v1.SecretKeySelector `json:"authPasswordSecretRef,omitempty" tf:"-"`

	// Authentication type to access the MongoDB source endpoint. Default is password.
	// +kubebuilder:validation:Required
	AuthType *string `json:"authType" tf:"auth_type,omitempty"`

	// The username provided with the auth-role option of the AuthType setting for a Redis target endpoint.
	// +kubebuilder:validation:Optional
	AuthUserName *string `json:"authUserName,omitempty" tf:"auth_user_name,omitempty"`

	// Port used by the endpoint database.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	// +kubebuilder:validation:Optional
	SSLCACertificateArn *string `json:"sslCaCertificateArn,omitempty" tf:"ssl_ca_certificate_arn,omitempty"`

	// The plaintext option doesn't provide Transport Layer Security (TLS) encryption for traffic between endpoint and database. Options include plaintext, ssl-encryption. The default is ssl-encryption.
	// +kubebuilder:validation:Optional
	SSLSecurityProtocol *string `json:"sslSecurityProtocol,omitempty" tf:"ssl_security_protocol,omitempty"`

	// Host name of the server.
	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`
}

type RedshiftSettingsObservation struct {
}

type RedshiftSettingsParameters struct {

	// Custom S3 Bucket Object prefix for intermediate storage.
	// +kubebuilder:validation:Optional
	BucketFolder *string `json:"bucketFolder,omitempty" tf:"bucket_folder,omitempty"`

	// Custom S3 Bucket name for intermediate storage.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The server-side encryption mode that you want to encrypt your intermediate .csv object files copied to S3. Defaults to SSE_S3. Valid values are SSE_S3 and SSE_KMS.
	// +kubebuilder:validation:Optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// If you set encryptionMode to SSE_KMS, set this parameter to the Amazon Resource Name (ARN) for the AWS KMS key.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionKMSKeyID *string `json:"serverSideEncryptionKmsKeyId,omitempty" tf:"server_side_encryption_kms_key_id,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role with permissions to read from or write to the S3 Bucket for intermediate storage.
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`
}

type S3SettingsObservation struct {
}

type S3SettingsParameters struct {

	// Whether to add column name information to the .csv output file. Default is false.
	// +kubebuilder:validation:Optional
	AddColumnName *bool `json:"addColumnName,omitempty" tf:"add_column_name,omitempty"`

	// S3 object prefix.
	// +kubebuilder:validation:Optional
	BucketFolder *string `json:"bucketFolder,omitempty" tf:"bucket_folder,omitempty"`

	// S3 bucket name.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Predefined (canned) access control list for objects created in an S3 bucket. Valid values include NONE, PRIVATE, PUBLIC_READ, PUBLIC_READ_WRITE, AUTHENTICATED_READ, AWS_EXEC_READ, BUCKET_OWNER_READ, and BUCKET_OWNER_FULL_CONTROL. Default is NONE.
	// +kubebuilder:validation:Optional
	CannedACLForObjects *string `json:"cannedAclForObjects,omitempty" tf:"canned_acl_for_objects,omitempty"`

	// Whether to write insert and update operations to .csv or .parquet output files. Default is false.
	// +kubebuilder:validation:Optional
	CdcInsertsAndUpdates *bool `json:"cdcInsertsAndUpdates,omitempty" tf:"cdc_inserts_and_updates,omitempty"`

	// Whether to write insert operations to .csv or .parquet output files. Default is false.
	// +kubebuilder:validation:Optional
	CdcInsertsOnly *bool `json:"cdcInsertsOnly,omitempty" tf:"cdc_inserts_only,omitempty"`

	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3. Default is 60.
	// +kubebuilder:validation:Optional
	CdcMaxBatchInterval *float64 `json:"cdcMaxBatchInterval,omitempty" tf:"cdc_max_batch_interval,omitempty"`

	// Minimum file size condition as defined in kilobytes to output a file to Amazon S3. Default is 32000. NOTE: Previously, this setting was measured in megabytes but now represents kilobytes. Update configurations accordingly.
	// +kubebuilder:validation:Optional
	CdcMinFileSize *float64 `json:"cdcMinFileSize,omitempty" tf:"cdc_min_file_size,omitempty"`

	// Folder path of CDC files. For an S3 source, this setting is required if a task captures change data; otherwise, it's optional. If cdc_path is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. Supported in AWS DMS versions 3.4.2 and later.
	// +kubebuilder:validation:Optional
	CdcPath *string `json:"cdcPath,omitempty" tf:"cdc_path,omitempty"`

	// Set to compress target files. Default is NONE. Valid values are GZIP and NONE.
	// +kubebuilder:validation:Optional
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	// Delimiter used to separate columns in the source files. Default is ,.
	// +kubebuilder:validation:Optional
	CsvDelimiter *string `json:"csvDelimiter,omitempty" tf:"csv_delimiter,omitempty"`

	// String to use for all columns not included in the supplemental log.
	// +kubebuilder:validation:Optional
	CsvNoSupValue *string `json:"csvNoSupValue,omitempty" tf:"csv_no_sup_value,omitempty"`

	// String to as null when writing to the target.
	// +kubebuilder:validation:Optional
	CsvNullValue *string `json:"csvNullValue,omitempty" tf:"csv_null_value,omitempty"`

	// Delimiter used to separate rows in the source files. Default is \n.
	// +kubebuilder:validation:Optional
	CsvRowDelimiter *string `json:"csvRowDelimiter,omitempty" tf:"csv_row_delimiter,omitempty"`

	// Output format for the files that AWS DMS uses to create S3 objects. Valid values are csv and parquet. Default is csv.
	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Size of one data page in bytes. Default is 1048576 (1 MiB).
	// +kubebuilder:validation:Optional
	DataPageSize *float64 `json:"dataPageSize,omitempty" tf:"data_page_size,omitempty"`

	// Date separating delimiter to use during folder partitioning. Valid values are SLASH, UNDERSCORE, DASH, and NONE. Default is SLASH.
	// +kubebuilder:validation:Optional
	DatePartitionDelimiter *string `json:"datePartitionDelimiter,omitempty" tf:"date_partition_delimiter,omitempty"`

	// Partition S3 bucket folders based on transaction commit dates. Default is false.
	// +kubebuilder:validation:Optional
	DatePartitionEnabled *bool `json:"datePartitionEnabled,omitempty" tf:"date_partition_enabled,omitempty"`

	// Date format to use during folder partitioning. Use this parameter when date_partition_enabled is set to true. Valid values are YYYYMMDD, YYYYMMDDHH, YYYYMM, MMYYYYDD, and DDMMYYYY. Default is YYYYMMDD.
	// +kubebuilder:validation:Optional
	DatePartitionSequence *string `json:"datePartitionSequence,omitempty" tf:"date_partition_sequence,omitempty"`

	// Maximum size in bytes of an encoded dictionary page of a column. Default is 1048576 (1 MiB).
	// +kubebuilder:validation:Optional
	DictPageSizeLimit *float64 `json:"dictPageSizeLimit,omitempty" tf:"dict_page_size_limit,omitempty"`

	// Whether to enable statistics for Parquet pages and row groups. Default is true.
	// +kubebuilder:validation:Optional
	EnableStatistics *bool `json:"enableStatistics,omitempty" tf:"enable_statistics,omitempty"`

	// Type of encoding to use. Value values are rle_dictionary, plain, and plain_dictionary. Default is rle_dictionary.
	// +kubebuilder:validation:Optional
	EncodingType *string `json:"encodingType,omitempty" tf:"encoding_type,omitempty"`

	// Server-side encryption mode that you want to encrypt your .csv or .parquet object files copied to S3. Valid values are SSE_S3 and SSE_KMS. Default is SSE_S3.
	// +kubebuilder:validation:Optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// JSON document that describes how AWS DMS should interpret the data.
	// +kubebuilder:validation:Optional
	ExternalTableDefinition *string `json:"externalTableDefinition,omitempty" tf:"external_table_definition,omitempty"`

	// When this value is set to 1, DMS ignores the first row header in a .csv file. Default is 0.
	// +kubebuilder:validation:Optional
	IgnoreHeaderRows *float64 `json:"ignoreHeaderRows,omitempty" tf:"ignore_header_rows,omitempty"`

	// Deprecated. This setting has no effect. Will be removed in a future version.
	// This setting has no effect, is deprecated, and will be removed in a future version
	// +kubebuilder:validation:Optional
	IgnoreHeadersRow *float64 `json:"ignoreHeadersRow,omitempty" tf:"ignore_headers_row,omitempty"`

	// Whether to enable a full load to write INSERT operations to the .csv output files only to indicate how the rows were added to the source database. Default is false.
	// +kubebuilder:validation:Optional
	IncludeOpForFullLoad *bool `json:"includeOpForFullLoad,omitempty" tf:"include_op_for_full_load,omitempty"`

	// Maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load. Valid values are from 1 to 1048576. Default is 1048576 (1 GB).
	// +kubebuilder:validation:Optional
	MaxFileSize *float64 `json:"maxFileSize,omitempty" tf:"max_file_size,omitempty"`

	// - Specifies the precision of any TIMESTAMP column values written to an S3 object file in .parquet format. Default is false.
	// +kubebuilder:validation:Optional
	ParquetTimestampInMillisecond *bool `json:"parquetTimestampInMillisecond,omitempty" tf:"parquet_timestamp_in_millisecond,omitempty"`

	// Version of the .parquet file format. Default is parquet-1-0. Valid values are parquet-1-0 and parquet-2-0.
	// +kubebuilder:validation:Optional
	ParquetVersion *string `json:"parquetVersion,omitempty" tf:"parquet_version,omitempty"`

	// Whether DMS saves the transaction order for a CDC load on the S3 target specified by cdc_path. Default is false.
	// +kubebuilder:validation:Optional
	PreserveTransactions *bool `json:"preserveTransactions,omitempty" tf:"preserve_transactions,omitempty"`

	// For an S3 source, whether each leading double quotation mark has to be followed by an ending double quotation mark. Default is true.
	// +kubebuilder:validation:Optional
	Rfc4180 *bool `json:"rfc4180,omitempty" tf:"rfc_4180,omitempty"`

	// Number of rows in a row group. Default is 10000.
	// +kubebuilder:validation:Optional
	RowGroupLength *float64 `json:"rowGroupLength,omitempty" tf:"row_group_length,omitempty"`

	// If you set encryptionMode to SSE_KMS, set this parameter to the ARN for the AWS KMS key.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionKMSKeyID *string `json:"serverSideEncryptionKmsKeyId,omitempty" tf:"server_side_encryption_kms_key_id,omitempty"`

	// ARN of the IAM Role with permissions to read from or write to the S3 Bucket.
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`

	// Column to add with timestamp information to the endpoint data for an Amazon S3 target.
	// +kubebuilder:validation:Optional
	TimestampColumnName *string `json:"timestampColumnName,omitempty" tf:"timestamp_column_name,omitempty"`

	// Whether to use csv_no_sup_value for columns not included in the supplemental log.
	// +kubebuilder:validation:Optional
	UseCsvNoSupValue *bool `json:"useCsvNoSupValue,omitempty" tf:"use_csv_no_sup_value,omitempty"`

	// When set to true, uses the task start time as the timestamp column value instead of the time data is written to target.
	// For full load, when set to true, each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.
	// When set to false, the full load timestamp in the timestamp column increments with the time data arrives at the target. Default is false.
	// +kubebuilder:validation:Optional
	UseTaskStartTimeForFullLoadTimestamp *bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" tf:"use_task_start_time_for_full_load_timestamp,omitempty"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API. Provides a DMS (Data Migration Service) endpoint resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
