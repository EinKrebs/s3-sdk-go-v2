// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the batch inference jobs that have been performed off of a
// solution version.
func (c *Client) ListBatchInferenceJobs(ctx context.Context, params *ListBatchInferenceJobsInput, optFns ...func(*Options)) (*ListBatchInferenceJobsOutput, error) {
	if params == nil {
		params = &ListBatchInferenceJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchInferenceJobs", params, optFns, addOperationListBatchInferenceJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchInferenceJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchInferenceJobsInput struct {

	// The maximum number of batch inference job results to return in each page. The
	// default value is 100.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution version from which the batch
	// inference jobs were created.
	SolutionVersionArn *string
}

type ListBatchInferenceJobsOutput struct {

	// A list containing information on each job that is returned.
	BatchInferenceJobs []types.BatchInferenceJobSummary

	// The token to use to retreive the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBatchInferenceJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBatchInferenceJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBatchInferenceJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchInferenceJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListBatchInferenceJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListBatchInferenceJobs",
	}
}
