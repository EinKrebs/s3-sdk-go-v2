// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Create an ApiKey resource. AWS CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-api-key.html)
func (c *Client) CreateApiKey(ctx context.Context, params *CreateApiKeyInput, optFns ...func(*Options)) (*CreateApiKeyOutput, error) {
	if params == nil {
		params = &CreateApiKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiKey", params, optFns, addOperationCreateApiKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create an ApiKey resource.
type CreateApiKeyInput struct {

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string

	// The description of the ApiKey.
	Description *string

	// Specifies whether the ApiKey can be used by callers.
	Enabled bool

	// Specifies whether (true) or not (false) the key identifier is distinct from the
	// created API key value. This parameter is deprecated and should not be used.
	GenerateDistinctId bool

	// DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.
	StageKeys []types.StageKey

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]string

	// Specifies a value of the API key.
	Value *string
}

// A resource that can be distributed to callers for executing Method resources
// that require an API key. API keys can be mapped to any Stage on any RestApi,
// which indicates that the callers with the API key can make requests to that
// stage. Use API Keys
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type CreateApiKeyOutput struct {

	// The timestamp when the API Key was created.
	CreatedDate *time.Time

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string

	// The description of the API Key.
	Description *string

	// Specifies whether the API Key can be used by callers.
	Enabled bool

	// The identifier of the API Key.
	Id *string

	// The timestamp when the API Key was last updated.
	LastUpdatedDate *time.Time

	// The name of the API Key.
	Name *string

	// A list of Stage resources that are associated with the ApiKey resource.
	StageKeys []string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// The value of the API Key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateApiKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateApiKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateApiKey",
	}
}
