// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Lambda functions, with the version-specific configuration of
// each. Lambda returns up to 50 functions per call. Set FunctionVersion to ALL to
// include all published versions of each function in addition to the unpublished
// version. To get more information about a function or version, use GetFunction.
func (c *Client) ListFunctions(ctx context.Context, params *ListFunctionsInput, optFns ...func(*Options)) (*ListFunctionsOutput, error) {
	if params == nil {
		params = &ListFunctionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctions", params, optFns, addOperationListFunctionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionsInput struct {

	// Set to ALL to include entries for all published versions of each function.
	FunctionVersion types.FunctionVersion

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// For Lambda@Edge functions, the AWS Region of the master function. For example,
	// us-east-1 filters the list of functions to only include Lambda@Edge functions
	// replicated from a master function in US East (N. Virginia). If specified, you
	// must set FunctionVersion to ALL.
	MasterRegion *string

	// The maximum number of functions to return.
	MaxItems *int32
}

// A list of Lambda functions.
type ListFunctionsOutput struct {

	// A list of Lambda functions.
	Functions []types.FunctionConfiguration

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFunctionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFunctions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctions",
	}
}
