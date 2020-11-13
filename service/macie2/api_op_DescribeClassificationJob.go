// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the status and settings for a classification job.
func (c *Client) DescribeClassificationJob(ctx context.Context, params *DescribeClassificationJobInput, optFns ...func(*Options)) (*DescribeClassificationJobOutput, error) {
	if params == nil {
		params = &DescribeClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClassificationJob", params, optFns, addOperationDescribeClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClassificationJobInput struct {

	// The unique identifier for the classification job.
	//
	// This member is required.
	JobId *string
}

type DescribeClassificationJobOutput struct {

	// The token that was provided to ensure the idempotency of the request to create
	// the job.
	ClientToken *string

	// The date and time, in UTC and extended ISO 8601 format, when the job was
	// created.
	CreatedAt *time.Time

	// The custom data identifiers that the job uses to analyze data.
	CustomDataIdentifierIds []string

	// The custom description of the job.
	Description *string

	// Specifies whether the job is configured to analyze all existing, eligible
	// objects immediately after it's created.
	InitialRun bool

	// The Amazon Resource Name (ARN) of the job.
	JobArn *string

	// The unique identifier for the job.
	JobId *string

	// The current status of the job. Possible values are:
	//
	// * CANCELLED - You cancelled
	// the job, or you paused the job and didn't resume it within 30 days of pausing
	// it.
	//
	// * COMPLETE - For a one-time job, Amazon Macie finished processing all the
	// data specified for the job. This value doesn't apply to recurring jobs.
	//
	// * IDLE
	// - For a recurring job, the previous scheduled run is complete and the next
	// scheduled run is pending. This value doesn't apply to one-time jobs.
	//
	// * PAUSED -
	// Amazon Macie started running the job but completion of the job would exceed one
	// or more quotas for your account.
	//
	// * RUNNING - For a one-time job, the job is in
	// progress. For a recurring job, a scheduled run is in progress.
	//
	// * USER_PAUSED -
	// You paused the job. If you don't resume the job within 30 days of pausing it,
	// the job will expire and be cancelled.
	JobStatus types.JobStatus

	// The schedule for running the job. Possible values are:
	//
	// * ONE_TIME - The job
	// runs only once.
	//
	// * SCHEDULED - The job runs on a daily, weekly, or monthly
	// basis. The scheduleFrequency property indicates the recurrence pattern for the
	// job.
	JobType types.JobType

	// The date and time, in UTC and extended ISO 8601 format, when the job last ran.
	LastRunTime *time.Time

	// The custom name of the job.
	Name *string

	// The S3 buckets that the job is configured to analyze, and the scope of that
	// analysis.
	S3JobDefinition *types.S3JobDefinition

	// The sampling depth, as a percentage, that determines the percentage of eligible
	// objects that the job analyzes.
	SamplingPercentage int32

	// The recurrence pattern for running the job. If the job is configured to run only
	// once, this value is null.
	ScheduleFrequency *types.JobScheduleFrequency

	// The number of times that the job has run and processing statistics for the job's
	// current run.
	Statistics *types.Statistics

	// A map of key-value pairs that specifies which tags (keys and values) are
	// associated with the classification job.
	Tags map[string]string

	// If the current status of the job is USER_PAUSED, specifies when the job was
	// paused and when the job will expire and be cancelled if it isn't resumed. This
	// value is present only if the value for jobStatus is USER_PAUSED.
	UserPausedDetails *types.UserPausedDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeClassificationJob{}, middleware.After)
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
	if err = addOpDescribeClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClassificationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DescribeClassificationJob",
	}
}
