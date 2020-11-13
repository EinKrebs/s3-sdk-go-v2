// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the RestApis resources for your collection.
func (c *Client) GetRestApis(ctx context.Context, params *GetRestApisInput, optFns ...func(*Options)) (*GetRestApisOutput, error) {
	if params == nil {
		params = &GetRestApisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRestApis", params, optFns, addOperationGetRestApisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRestApisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to list existing RestApis defined for your collection.
type GetRestApisInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string
}

// Contains references to your APIs and links that guide you in how to interact
// with your collection. A collection offers a paginated view of your APIs. Create
// an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetRestApisOutput struct {

	// The current page of elements from this collection.
	Items []types.RestApi

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRestApisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRestApis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRestApis{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRestApis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRestApis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetRestApis",
	}
}
