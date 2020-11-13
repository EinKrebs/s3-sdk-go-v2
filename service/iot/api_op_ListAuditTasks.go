// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the Device Defender audits that have been performed during a given time
// period.
func (c *Client) ListAuditTasks(ctx context.Context, params *ListAuditTasksInput, optFns ...func(*Options)) (*ListAuditTasksOutput, error) {
	if params == nil {
		params = &ListAuditTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditTasks", params, optFns, addOperationListAuditTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditTasksInput struct {

	// The end of the time period.
	//
	// This member is required.
	EndTime *time.Time

	// The beginning of the time period. Audit information is retained for a limited
	// time (90 days). Requesting a start time prior to what is retained results in an
	// "InvalidRequestException".
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// A filter to limit the output to audits with the specified completion status: can
	// be one of "IN_PROGRESS", "COMPLETED", "FAILED", or "CANCELED".
	TaskStatus types.AuditTaskStatus

	// A filter to limit the output to the specified type of audit: can be one of
	// "ON_DEMAND_AUDIT_TASK" or "SCHEDULED__AUDIT_TASK".
	TaskType types.AuditTaskType
}

type ListAuditTasksOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// The audits that were performed during the specified time period.
	Tasks []types.AuditTaskMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAuditTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditTasks{}, middleware.After)
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
	if err = addOpListAuditTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditTasks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAuditTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuditTasks",
	}
}
