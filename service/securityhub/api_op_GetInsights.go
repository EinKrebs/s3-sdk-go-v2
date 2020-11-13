// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists and describes insights for the specified insight ARNs.
func (c *Client) GetInsights(ctx context.Context, params *GetInsightsInput, optFns ...func(*Options)) (*GetInsightsOutput, error) {
	if params == nil {
		params = &GetInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsights", params, optFns, addOperationGetInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightsInput struct {

	// The ARNs of the insights to describe. If you do not provide any insight ARNs,
	// then GetInsights returns all of your custom insights. It does not return any
	// managed insights.
	InsightArns []string

	// The maximum number of items to return in the response.
	MaxResults int32

	// The token that is required for pagination. On your first call to the GetInsights
	// operation, set the value of this parameter to NULL. For subsequent calls to the
	// operation, to continue listing data, set the value of this parameter to the
	// value returned from the previous response.
	NextToken *string
}

type GetInsightsOutput struct {

	// The insights returned by the operation.
	//
	// This member is required.
	Insights []types.Insight

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsights{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsights(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetInsights",
	}
}
