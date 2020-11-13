// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides the list of entities of a given type.
func (c *Client) ListEntities(ctx context.Context, params *ListEntitiesInput, optFns ...func(*Options)) (*ListEntitiesOutput, error) {
	if params == nil {
		params = &ListEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntities", params, optFns, addOperationListEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesInput struct {

	// The catalog related to the request. Fixed value: AWSMarketplace
	//
	// This member is required.
	Catalog *string

	// The type of entities to retrieve.
	//
	// This member is required.
	EntityType *string

	// An array of filter objects. Each filter object contains two attributes,
	// filterName and filterValues.
	FilterList []types.Filter

	// Specifies the upper limit of the elements on a single page. If a value isn't
	// provided, the default value is 20.
	MaxResults *int32

	// The value of the next token, if it exists. Null if there are no more results.
	NextToken *string

	// An object that contains two attributes, SortBy and SortOrder.
	Sort *types.Sort
}

type ListEntitiesOutput struct {

	// Array of EntitySummary object.
	EntitySummaryList []types.EntitySummary

	// The value of the next token if it exists. Null if there is no more result.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEntities{}, middleware.After)
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
	if err = addOpListEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntities(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "ListEntities",
	}
}
