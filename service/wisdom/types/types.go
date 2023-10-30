// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Configuration information for Amazon AppIntegrations to automatically ingest
// content.
type AppIntegrationsConfiguration struct {

	// The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use
	// for ingesting content.
	//   - For Salesforce (https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm)
	//   , your AppIntegrations DataIntegration must have an ObjectConfiguration if
	//   objectFields is not provided, including at least Id , ArticleNumber ,
	//   VersionNumber , Title , PublishStatus , and IsDeleted as source fields.
	//   - For ServiceNow (https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api)
	//   , your AppIntegrations DataIntegration must have an ObjectConfiguration if
	//   objectFields is not provided, including at least number , short_description ,
	//   sys_mod_count , workflow_state , and active as source fields.
	//   - For Zendesk (https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/)
	//   , your AppIntegrations DataIntegration must have an ObjectConfiguration if
	//   objectFields is not provided, including at least id , title , updated_at , and
	//   draft as source fields.
	//   - For SharePoint (https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index)
	//   , your AppIntegrations DataIntegration must have a FileConfiguration, including
	//   only file extensions that are among docx , pdf , html , htm , and txt .
	//   - For Amazon S3 (https://aws.amazon.com/s3/) , the ObjectConfiguration and
	//   FileConfiguration of your AppIntegrations DataIntegration must be null. The
	//   SourceURI of your DataIntegration must use the following format:
	//   s3://your_s3_bucket_name . The bucket policy of the corresponding S3 bucket
	//   must allow the Amazon Web Services principal app-integrations.amazonaws.com to
	//   perform s3:ListBucket , s3:GetObject , and s3:GetBucketLocation against the
	//   bucket.
	//
	// This member is required.
	AppIntegrationArn *string

	// The fields from the source that are made available to your agents in Wisdom.
	// Optional if ObjectConfiguration is included in the provided DataIntegration.
	//   - For Salesforce (https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm)
	//   , you must include at least Id , ArticleNumber , VersionNumber , Title ,
	//   PublishStatus , and IsDeleted .
	//   - For ServiceNow (https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api)
	//   , you must include at least number , short_description , sys_mod_count ,
	//   workflow_state , and active .
	//   - For Zendesk (https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/)
	//   , you must include at least id , title , updated_at , and draft .
	// Make sure to include additional fields. These fields are indexed and used to
	// source recommendations.
	ObjectFields []string

	noSmithyDocumentSerde
}

// Information about the assistant association.
type AssistantAssociationData struct {

	// The Amazon Resource Name (ARN) of the Wisdom assistant.
	//
	// This member is required.
	AssistantArn *string

	// The Amazon Resource Name (ARN) of the assistant association.
	//
	// This member is required.
	AssistantAssociationArn *string

	// The identifier of the assistant association.
	//
	// This member is required.
	AssistantAssociationId *string

	// The identifier of the Wisdom assistant.
	//
	// This member is required.
	AssistantId *string

	// A union type that currently has a single argument, the knowledge base ID.
	//
	// This member is required.
	AssociationData AssistantAssociationOutputData

	// The type of association.
	//
	// This member is required.
	AssociationType AssociationType

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The data that is input into Wisdom as a result of the assistant association.
//
// The following types satisfy this interface:
//
//	AssistantAssociationInputDataMemberKnowledgeBaseId
type AssistantAssociationInputData interface {
	isAssistantAssociationInputData()
}

// The identifier of the knowledge base.
type AssistantAssociationInputDataMemberKnowledgeBaseId struct {
	Value string

	noSmithyDocumentSerde
}

func (*AssistantAssociationInputDataMemberKnowledgeBaseId) isAssistantAssociationInputData() {}

// The data that is output as a result of the assistant association.
//
// The following types satisfy this interface:
//
//	AssistantAssociationOutputDataMemberKnowledgeBaseAssociation
type AssistantAssociationOutputData interface {
	isAssistantAssociationOutputData()
}

// The knowledge base where output data is sent.
type AssistantAssociationOutputDataMemberKnowledgeBaseAssociation struct {
	Value KnowledgeBaseAssociationData

	noSmithyDocumentSerde
}

func (*AssistantAssociationOutputDataMemberKnowledgeBaseAssociation) isAssistantAssociationOutputData() {
}

// Summary information about the assistant association.
type AssistantAssociationSummary struct {

	// The Amazon Resource Name (ARN) of the Wisdom assistant.
	//
	// This member is required.
	AssistantArn *string

	// The Amazon Resource Name (ARN) of the assistant association.
	//
	// This member is required.
	AssistantAssociationArn *string

	// The identifier of the assistant association.
	//
	// This member is required.
	AssistantAssociationId *string

	// The identifier of the Wisdom assistant.
	//
	// This member is required.
	AssistantId *string

	// The association data.
	//
	// This member is required.
	AssociationData AssistantAssociationOutputData

	// The type of association.
	//
	// This member is required.
	AssociationType AssociationType

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The assistant data.
type AssistantData struct {

	// The Amazon Resource Name (ARN) of the Wisdom assistant.
	//
	// This member is required.
	AssistantArn *string

	// The identifier of the Wisdom assistant.
	//
	// This member is required.
	AssistantId *string

	// The name.
	//
	// This member is required.
	Name *string

	// The status of the assistant.
	//
	// This member is required.
	Status AssistantStatus

	// The type of assistant.
	//
	// This member is required.
	Type AssistantType

	// The description.
	Description *string

	// The configuration information for the Wisdom assistant integration.
	IntegrationConfiguration *AssistantIntegrationConfiguration

	// The configuration information for the customer managed key used for encryption.
	// This KMS key must have a policy that allows kms:CreateGrant and kms:DescribeKey
	// permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom
	// with chat, the key policy must also allow kms:Decrypt , kms:GenerateDataKey* ,
	// and kms:DescribeKey permissions to the connect.amazonaws.com service principal.
	// For more information about setting up a customer managed key for Wisdom, see
	// Enable Amazon Connect Wisdom for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html)
	// .
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The configuration information for the Wisdom assistant integration.
type AssistantIntegrationConfiguration struct {

	// The Amazon Resource Name (ARN) of the integrated Amazon SNS topic used for
	// streaming chat messages.
	TopicIntegrationArn *string

	noSmithyDocumentSerde
}

// Summary information about the assistant.
type AssistantSummary struct {

	// The Amazon Resource Name (ARN) of the Wisdom assistant.
	//
	// This member is required.
	AssistantArn *string

	// The identifier of the Wisdom assistant.
	//
	// This member is required.
	AssistantId *string

	// The name of the assistant.
	//
	// This member is required.
	Name *string

	// The status of the assistant.
	//
	// This member is required.
	Status AssistantStatus

	// The type of the assistant.
	//
	// This member is required.
	Type AssistantType

	// The description of the assistant.
	Description *string

	// The configuration information for the Wisdom assistant integration.
	IntegrationConfiguration *AssistantIntegrationConfiguration

	// The configuration information for the customer managed key used for encryption.
	// This KMS key must have a policy that allows kms:CreateGrant and kms:DescribeKey
	// permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom
	// with chat, the key policy must also allow kms:Decrypt , kms:GenerateDataKey* ,
	// and kms:DescribeKey permissions to the connect.amazonaws.com service principal.
	// For more information about setting up a customer managed key for Wisdom, see
	// Enable Amazon Connect Wisdom for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html)
	// .
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Information about the content.
type ContentData struct {

	// The Amazon Resource Name (ARN) of the content.
	//
	// This member is required.
	ContentArn *string

	// The identifier of the content.
	//
	// This member is required.
	ContentId *string

	// The media type of the content.
	//
	// This member is required.
	ContentType *string

	// The Amazon Resource Name (ARN) of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseId *string

	// A key/value map to store attributes without affecting tagging or
	// recommendations. For example, when synchronizing data between an external system
	// and Wisdom, you can store an external version identifier as metadata to utilize
	// for determining drift.
	//
	// This member is required.
	Metadata map[string]string

	// The name of the content.
	//
	// This member is required.
	Name *string

	// The identifier of the content revision.
	//
	// This member is required.
	RevisionId *string

	// The status of the content.
	//
	// This member is required.
	Status ContentStatus

	// The title of the content.
	//
	// This member is required.
	Title *string

	// The URL of the content.
	//
	// This member is required.
	Url *string

	// The expiration time of the URL as an epoch timestamp.
	//
	// This member is required.
	UrlExpiry *time.Time

	// The URI of the content.
	LinkOutUri *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Reference information about the content.
type ContentReference struct {

	// The Amazon Resource Name (ARN) of the content.
	ContentArn *string

	// The identifier of the content.
	ContentId *string

	// The Amazon Resource Name (ARN) of the knowledge base.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

// Summary information about the content.
type ContentSummary struct {

	// The Amazon Resource Name (ARN) of the content.
	//
	// This member is required.
	ContentArn *string

	// The identifier of the content.
	//
	// This member is required.
	ContentId *string

	// The media type of the content.
	//
	// This member is required.
	ContentType *string

	// The Amazon Resource Name (ARN) of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseId *string

	// A key/value map to store attributes without affecting tagging or
	// recommendations. For example, when synchronizing data between an external system
	// and Wisdom, you can store an external version identifier as metadata to utilize
	// for determining drift.
	//
	// This member is required.
	Metadata map[string]string

	// The name of the content.
	//
	// This member is required.
	Name *string

	// The identifier of the revision of the content.
	//
	// This member is required.
	RevisionId *string

	// The status of the content.
	//
	// This member is required.
	Status ContentStatus

	// The title of the content.
	//
	// This member is required.
	Title *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The document.
type Document struct {

	// A reference to the content resource.
	//
	// This member is required.
	ContentReference *ContentReference

	// The excerpt from the document.
	Excerpt *DocumentText

	// The title of the document.
	Title *DocumentText

	noSmithyDocumentSerde
}

// The text of the document.
type DocumentText struct {

	// Highlights in the document text.
	Highlights []Highlight

	// Text in the document.
	Text *string

	noSmithyDocumentSerde
}

// A search filter.
type Filter struct {

	// The field on which to filter.
	//
	// This member is required.
	Field FilterField

	// The operator to use for comparing the field’s value with the provided value.
	//
	// This member is required.
	Operator FilterOperator

	// The desired field value on which to filter.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Offset specification to describe highlighting of document excerpts for
// rendering search results and recommendations.
type Highlight struct {

	// The offset for the start of the highlight.
	BeginOffsetInclusive int32

	// The offset for the end of the highlight.
	EndOffsetExclusive int32

	noSmithyDocumentSerde
}

// Association information about the knowledge base.
type KnowledgeBaseAssociationData struct {

	// The Amazon Resource Name (ARN) of the knowledge base.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

// Information about the knowledge base.
type KnowledgeBaseData struct {

	// The Amazon Resource Name (ARN) of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The type of knowledge base.
	//
	// This member is required.
	KnowledgeBaseType KnowledgeBaseType

	// The name of the knowledge base.
	//
	// This member is required.
	Name *string

	// The status of the knowledge base.
	//
	// This member is required.
	Status KnowledgeBaseStatus

	// The description.
	Description *string

	// An epoch timestamp indicating the most recent content modification inside the
	// knowledge base. If no content exists in a knowledge base, this value is unset.
	LastContentModificationTime *time.Time

	// Information about how to render the content.
	RenderingConfiguration *RenderingConfiguration

	// The configuration information for the customer managed key used for encryption.
	// This KMS key must have a policy that allows kms:CreateGrant and kms:DescribeKey
	// permissions to the IAM identity using the key to invoke Wisdom. For more
	// information about setting up a customer managed key for Wisdom, see Enable
	// Amazon Connect Wisdom for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html)
	// .
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// Source configuration information about the knowledge base.
	SourceConfiguration SourceConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about the knowledge base.
type KnowledgeBaseSummary struct {

	// The Amazon Resource Name (ARN) of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseArn *string

	// The identifier of the knowledge base.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The type of knowledge base.
	//
	// This member is required.
	KnowledgeBaseType KnowledgeBaseType

	// The name of the knowledge base.
	//
	// This member is required.
	Name *string

	// The status of the knowledge base summary.
	//
	// This member is required.
	Status KnowledgeBaseStatus

	// The description of the knowledge base.
	Description *string

	// Information about how to render the content.
	RenderingConfiguration *RenderingConfiguration

	// The configuration information for the customer managed key used for encryption.
	// This KMS key must have a policy that allows kms:CreateGrant and kms:DescribeKey
	// permissions to the IAM identity using the key to invoke Wisdom. For more
	// information about setting up a customer managed key for Wisdom, see Enable
	// Amazon Connect Wisdom for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html)
	// .
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration

	// Configuration information about the external data source.
	SourceConfiguration SourceConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An error occurred when creating a recommendation.
type NotifyRecommendationsReceivedError struct {

	// A recommendation is causing an error.
	Message *string

	// The identifier of the recommendation that is in error.
	RecommendationId *string

	noSmithyDocumentSerde
}

// Data associated with the QUERY RecommendationTriggerType.
type QueryRecommendationTriggerData struct {

	// The text associated with the recommendation trigger.
	Text *string

	noSmithyDocumentSerde
}

// Information about the recommendation.
type RecommendationData struct {

	// The recommended document.
	//
	// This member is required.
	Document *Document

	// The identifier of the recommendation.
	//
	// This member is required.
	RecommendationId *string

	// The relevance level of the recommendation.
	RelevanceLevel RelevanceLevel

	// The relevance score of the recommendation.
	RelevanceScore float64

	// The type of recommendation.
	Type RecommendationType

	noSmithyDocumentSerde
}

// A recommendation trigger provides context on the event that produced the
// referenced recommendations. Recommendations are only referenced in
// recommendationIds by a single RecommendationTrigger.
type RecommendationTrigger struct {

	// A union type containing information related to the trigger.
	//
	// This member is required.
	Data RecommendationTriggerData

	// The identifier of the recommendation trigger.
	//
	// This member is required.
	Id *string

	// The identifiers of the recommendations.
	//
	// This member is required.
	RecommendationIds []string

	// The source of the recommendation trigger.
	//   - ISSUE_DETECTION: The corresponding recommendations were triggered by a
	//   Contact Lens issue.
	//   - RULE_EVALUATION: The corresponding recommendations were triggered by a
	//   Contact Lens rule.
	//
	// This member is required.
	Source RecommendationSourceType

	// The type of recommendation trigger.
	//
	// This member is required.
	Type RecommendationTriggerType

	noSmithyDocumentSerde
}

// A union type containing information related to the trigger.
//
// The following types satisfy this interface:
//
//	RecommendationTriggerDataMemberQuery
type RecommendationTriggerData interface {
	isRecommendationTriggerData()
}

// Data associated with the QUERY RecommendationTriggerType.
type RecommendationTriggerDataMemberQuery struct {
	Value QueryRecommendationTriggerData

	noSmithyDocumentSerde
}

func (*RecommendationTriggerDataMemberQuery) isRecommendationTriggerData() {}

// Information about how to render the content.
type RenderingConfiguration struct {

	// A URI template containing exactly one variable in ${variableName} format. This
	// can only be set for EXTERNAL knowledge bases. For Salesforce, ServiceNow, and
	// Zendesk, the variable must be one of the following:
	//   - Salesforce: Id , ArticleNumber , VersionNumber , Title , PublishStatus , or
	//   IsDeleted
	//   - ServiceNow: number , short_description , sys_mod_count , workflow_state , or
	//   active
	//   - Zendesk: id , title , updated_at , or draft
	// The variable is replaced with the actual value for a piece of content when
	// calling GetContent (https://docs.aws.amazon.com/wisdom/latest/APIReference/API_GetContent.html)
	// .
	TemplateUri *string

	noSmithyDocumentSerde
}

// Information about the result.
type ResultData struct {

	// The document.
	//
	// This member is required.
	Document *Document

	// The identifier of the result data.
	//
	// This member is required.
	ResultId *string

	// The relevance score of the results.
	RelevanceScore float64

	noSmithyDocumentSerde
}

// The search expression.
type SearchExpression struct {

	// The search expression filters.
	//
	// This member is required.
	Filters []Filter

	noSmithyDocumentSerde
}

// The configuration information for the customer managed key used for encryption.
type ServerSideEncryptionConfiguration struct {

	// The customer managed key used for encryption. For more information about
	// setting up a customer managed key for Wisdom, see Enable Amazon Connect Wisdom
	// for your instance (https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html)
	// . For information about valid ID values, see Key identifiers (KeyId) (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id)
	// .
	KmsKeyId *string

	noSmithyDocumentSerde
}

// Information about the session.
type SessionData struct {

	// The name of the session.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the session.
	//
	// This member is required.
	SessionArn *string

	// The identifier of the session.
	//
	// This member is required.
	SessionId *string

	// The description of the session.
	Description *string

	// The configuration information for the session integration.
	IntegrationConfiguration *SessionIntegrationConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The configuration information for the session integration.
type SessionIntegrationConfiguration struct {

	// The Amazon Resource Name (ARN) of the integrated Amazon SNS topic used for
	// streaming chat messages.
	TopicIntegrationArn *string

	noSmithyDocumentSerde
}

// Summary information about the session.
type SessionSummary struct {

	// The Amazon Resource Name (ARN) of the Wisdom assistant.
	//
	// This member is required.
	AssistantArn *string

	// The identifier of the Wisdom assistant.
	//
	// This member is required.
	AssistantId *string

	// The Amazon Resource Name (ARN) of the session.
	//
	// This member is required.
	SessionArn *string

	// The identifier of the session.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

// Configuration information about the external data source.
//
// The following types satisfy this interface:
//
//	SourceConfigurationMemberAppIntegrations
type SourceConfiguration interface {
	isSourceConfiguration()
}

// Configuration information for Amazon AppIntegrations to automatically ingest
// content.
type SourceConfigurationMemberAppIntegrations struct {
	Value AppIntegrationsConfiguration

	noSmithyDocumentSerde
}

func (*SourceConfigurationMemberAppIntegrations) isSourceConfiguration() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAssistantAssociationInputData()  {}
func (*UnknownUnionMember) isAssistantAssociationOutputData() {}
func (*UnknownUnionMember) isRecommendationTriggerData()      {}
func (*UnknownUnionMember) isSourceConfiguration()            {}
