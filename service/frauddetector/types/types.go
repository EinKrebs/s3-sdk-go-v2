// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Provides the error of the batch create variable API.
type BatchCreateVariableError struct {

	// The error code.
	Code int32

	// The error message.
	Message *string

	// The name.
	Name *string
}

// Provides the error of the batch get variable API.
type BatchGetVariableError struct {

	// The error code.
	Code int32

	// The error message.
	Message *string

	// The error name.
	Name *string
}

// The model training validation messages.
type DataValidationMetrics struct {

	// The field-specific model training validation messages.
	FieldLevelMessages []FieldValidationMessage

	// The file-specific model training validation messages.
	FileLevelMessages []FileValidationMessage
}

// The detector.
type Detector struct {

	// The detector ARN.
	Arn *string

	// Timestamp of when the detector was created.
	CreatedTime *string

	// The detector description.
	Description *string

	// The detector ID.
	DetectorId *string

	// The name of the event type.
	EventTypeName *string

	// Timestamp of when the detector was last updated.
	LastUpdatedTime *string
}

// The summary of the detector version.
type DetectorVersionSummary struct {

	// The detector version description.
	Description *string

	// The detector version ID.
	DetectorVersionId *string

	// Timestamp of when the detector version was last updated.
	LastUpdatedTime *string

	// The detector version status.
	Status DetectorVersionStatus
}

// The entity details.
type Entity struct {

	// The entity ID. If you do not know the entityId, you can pass unknown, which is
	// areserved string literal.
	//
	// This member is required.
	EntityId *string

	// The entity type.
	//
	// This member is required.
	EntityType *string
}

// The entity type details.
type EntityType struct {

	// The entity type ARN.
	Arn *string

	// Timestamp of when the entity type was created.
	CreatedTime *string

	// The entity type description.
	Description *string

	// Timestamp of when the entity type was last updated.
	LastUpdatedTime *string

	// The entity type name.
	Name *string
}

// The event type details.
type EventType struct {

	// The entity type ARN.
	Arn *string

	// Timestamp of when the event type was created.
	CreatedTime *string

	// The event type description.
	Description *string

	// The event type entity types.
	EntityTypes []string

	// The event type event variables.
	EventVariables []string

	// The event type labels.
	Labels []string

	// Timestamp of when the event type was last updated.
	LastUpdatedTime *string

	// The event type name.
	Name *string
}

// Details for the external events data used for model version training.
type ExternalEventsDetail struct {

	// The ARN of the role that provides Amazon Fraud Detector access to the data
	// location.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The Amazon S3 bucket location for the data.
	//
	// This member is required.
	DataLocation *string
}

// The Amazon SageMaker model.
type ExternalModel struct {

	// The model ARN.
	Arn *string

	// Timestamp of when the model was last created.
	CreatedTime *string

	// The input configuration.
	InputConfiguration *ModelInputConfiguration

	// The role used to invoke the model.
	InvokeModelEndpointRoleArn *string

	// Timestamp of when the model was last updated.
	LastUpdatedTime *string

	// The Amazon SageMaker model endpoints.
	ModelEndpoint *string

	// The Amazon Fraud Detector status for the external model endpoint
	ModelEndpointStatus ModelEndpointStatus

	// The source of the model.
	ModelSource ModelSource

	// The output configuration.
	OutputConfiguration *ModelOutputConfiguration
}

// The message details.
type FieldValidationMessage struct {

	// The message content.
	Content *string

	// The field name.
	FieldName *string

	// The message ID.
	Identifier *string

	// The message title.
	Title *string

	// The message type.
	Type *string
}

// The message details.
type FileValidationMessage struct {

	// The message content.
	Content *string

	// The message title.
	Title *string

	// The message type.
	Type *string
}

// The KMS key details.
type KMSKey struct {

	// The encryption key ARN.
	KmsEncryptionKeyArn *string
}

// The label details.
type Label struct {

	// The label ARN.
	Arn *string

	// Timestamp of when the event type was created.
	CreatedTime *string

	// The label description.
	Description *string

	// Timestamp of when the label was last updated.
	LastUpdatedTime *string

	// The label name.
	Name *string
}

// The label schema.
type LabelSchema struct {

	// The label mapper maps the Amazon Fraud Detector supported model classification
	// labels (FRAUD, LEGIT) to the appropriate event type labels. For example, if
	// "FRAUD" and "LEGIT" are Amazon Fraud Detector supported labels, this mapper
	// could be: {"FRAUD" => ["0"], "LEGIT" => ["1"]} or {"FRAUD" => ["false"], "LEGIT"
	// => ["true"]} or {"FRAUD" => ["fraud", "abuse"], "LEGIT" => ["legit", "safe"]}.
	// The value part of the mapper is a list, because you may have multiple label
	// variants from your event type for a single Amazon Fraud Detector label.
	//
	// This member is required.
	LabelMapper map[string][]string
}

// Model performance metrics data points.
type MetricDataPoint struct {

	// The false positive rate. This is the percentage of total legitimate events that
	// are incorrectly predicted as fraud.
	Fpr *float32

	// The percentage of fraud events correctly predicted as fraudulent as compared to
	// all events predicted as fraudulent.
	Precision *float32

	// The model threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled as
	// fraud.
	Threshold *float32

	// The true positive rate. This is the percentage of total fraud the model detects.
	// Also known as capture rate.
	Tpr *float32
}

// The model.
type Model struct {

	// The ARN of the model.
	Arn *string

	// Timestamp of when the model was created.
	CreatedTime *string

	// The model description.
	Description *string

	// The name of the event type.
	EventTypeName *string

	// Timestamp of last time the model was updated.
	LastUpdatedTime *string

	// The model ID.
	ModelId *string

	// The model type.
	ModelType ModelTypeEnum
}

// A pre-formed Amazon SageMaker model input you can include if your detector
// version includes an imported Amazon SageMaker model endpoint with pass-through
// input configuration.
type ModelEndpointDataBlob struct {

	// The byte buffer of the Amazon SageMaker model endpoint input data blob.
	ByteBuffer []byte

	// The content type of the Amazon SageMaker model endpoint input data blob.
	ContentType *string
}

// The Amazon SageMaker model input configuration.
type ModelInputConfiguration struct {

	// The event variables.
	//
	// This member is required.
	UseEventVariables *bool

	// Template for constructing the CSV input-data sent to SageMaker. At
	// event-evaluation, the placeholders for variable-names in the template will be
	// replaced with the variable values before being sent to SageMaker.
	CsvInputTemplate *string

	// The event type name.
	EventTypeName *string

	// The format of the model input configuration. The format differs depending on if
	// it is passed through to SageMaker or constructed by Amazon Fraud Detector.
	Format ModelInputDataFormat

	// Template for constructing the JSON input-data sent to SageMaker. At
	// event-evaluation, the placeholders for variable names in the template will be
	// replaced with the variable values before being sent to SageMaker.
	JsonInputTemplate *string
}

// Provides the Amazon Sagemaker model output configuration.
type ModelOutputConfiguration struct {

	// The format of the model output configuration.
	//
	// This member is required.
	Format ModelOutputDataFormat

	// A map of CSV index values in the SageMaker response to the Amazon Fraud Detector
	// variables.
	CsvIndexToVariableMap map[string]string

	// A map of JSON keys in response from SageMaker to the Amazon Fraud Detector
	// variables.
	JsonKeyToVariableMap map[string]string
}

// The fraud prediction scores.
type ModelScores struct {

	// The model version.
	ModelVersion *ModelVersion

	// The model's fraud prediction scores.
	Scores map[string]float32
}

// The model version.
type ModelVersion struct {

	// The model ID.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType ModelTypeEnum

	// The model version number.
	//
	// This member is required.
	ModelVersionNumber *string

	// The model version ARN.
	Arn *string
}

// The details of the model version.
type ModelVersionDetail struct {

	// The model version ARN.
	Arn *string

	// The timestamp when the model was created.
	CreatedTime *string

	// The event details.
	ExternalEventsDetail *ExternalEventsDetail

	// The timestamp when the model was last updated.
	LastUpdatedTime *string

	// The model ID.
	ModelId *string

	// The model type.
	ModelType ModelTypeEnum

	// The model version number.
	ModelVersionNumber *string

	// The status of the model version.
	Status *string

	// The training data schema.
	TrainingDataSchema *TrainingDataSchema

	// The model version training data source.
	TrainingDataSource TrainingDataSourceEnum

	// The training results.
	TrainingResult *TrainingResult
}

// The outcome.
type Outcome struct {

	// The outcome ARN.
	Arn *string

	// The timestamp when the outcome was created.
	CreatedTime *string

	// The outcome description.
	Description *string

	// The timestamp when the outcome was last updated.
	LastUpdatedTime *string

	// The outcome name.
	Name *string
}

// A rule.
type Rule struct {

	// The detector for which the rule is associated.
	//
	// This member is required.
	DetectorId *string

	// The rule ID.
	//
	// This member is required.
	RuleId *string

	// The rule version.
	//
	// This member is required.
	RuleVersion *string
}

// The details of the rule.
type RuleDetail struct {

	// The rule ARN.
	Arn *string

	// The timestamp of when the rule was created.
	CreatedTime *string

	// The rule description.
	Description *string

	// The detector for which the rule is associated.
	DetectorId *string

	// The rule expression.
	Expression *string

	// The rule language.
	Language Language

	// Timestamp of the last time the rule was updated.
	LastUpdatedTime *string

	// The rule outcomes.
	Outcomes []string

	// The rule ID.
	RuleId *string

	// The rule version.
	RuleVersion *string
}

// The rule results.
type RuleResult struct {

	// The outcomes of the matched rule, based on the rule execution mode.
	Outcomes []string

	// The rule ID that was matched, based on the rule execution mode.
	RuleId *string
}

// A key and value pair.
type Tag struct {

	// A tag key.
	//
	// This member is required.
	Key *string

	// A value assigned to a tag key.
	//
	// This member is required.
	Value *string
}

// The training data schema.
type TrainingDataSchema struct {

	// The label schema.
	//
	// This member is required.
	LabelSchema *LabelSchema

	// The training data schema variables.
	//
	// This member is required.
	ModelVariables []string
}

// The training metric details.
type TrainingMetrics struct {

	// The area under the curve. This summarizes true positive rate (TPR) and false
	// positive rate (FPR) across all possible model score thresholds. A model with no
	// predictive power has an AUC of 0.5, whereas a perfect model has a score of 1.0.
	Auc *float32

	// The data points details.
	MetricDataPoints []MetricDataPoint
}

// The training result details.
type TrainingResult struct {

	// The validation metrics.
	DataValidationMetrics *DataValidationMetrics

	// The training metric details.
	TrainingMetrics *TrainingMetrics
}

// The variable.
type Variable struct {

	// The ARN of the variable.
	Arn *string

	// The time when the variable was created.
	CreatedTime *string

	// The data source of the variable.
	DataSource DataSource

	// The data type of the variable. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	DataType DataType

	// The default value of the variable.
	DefaultValue *string

	// The description of the variable.
	Description *string

	// The time when variable was last updated.
	LastUpdatedTime *string

	// The name of the variable.
	Name *string

	// The variable type of the variable. Valid Values: AUTH_CODE | AVS |
	// BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY |
	// BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN |
	// CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL |
	// FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER |
	// PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 |
	// SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE |
	// SHIPPING_STATE | SHIPPING_ZIP | USERAGENT
	VariableType *string
}

// A variable in the list of variables for the batch create variable request.
type VariableEntry struct {

	// The data source of the variable.
	DataSource *string

	// The data type of the variable.
	DataType *string

	// The default value of the variable.
	DefaultValue *string

	// The description of the variable.
	Description *string

	// The name of the variable.
	Name *string

	// The type of the variable. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	// Valid Values: AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 |
	// BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE |
	// BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS |
	// FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID |
	// PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 |
	// SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME |
	// SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT
	VariableType *string
}
