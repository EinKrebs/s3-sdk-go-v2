// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables group metrics for the specified Auto Scaling group. For more
// information, see Monitoring Your Auto Scaling Groups and Instances
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) EnableMetricsCollection(ctx context.Context, params *EnableMetricsCollectionInput, optFns ...func(*Options)) (*EnableMetricsCollectionOutput, error) {
	if params == nil {
		params = &EnableMetricsCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableMetricsCollection", params, optFns, addOperationEnableMetricsCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The granularity to associate with the metrics to collect. The only valid value
	// is 1Minute.
	//
	// This member is required.
	Granularity *string

	// Specifies which group-level metrics to start collecting. You can specify one or
	// more of the following metrics:
	//
	// * GroupMinSize
	//
	// * GroupMaxSize
	//
	// *
	// GroupDesiredCapacity
	//
	// * GroupInServiceInstances
	//
	// * GroupPendingInstances
	//
	// *
	// GroupStandbyInstances
	//
	// * GroupTerminatingInstances
	//
	// * GroupTotalInstances
	//
	// The
	// instance weighting feature supports the following additional metrics:
	//
	// *
	// GroupInServiceCapacity
	//
	// * GroupPendingCapacity
	//
	// * GroupStandbyCapacity
	//
	// *
	// GroupTerminatingCapacity
	//
	// * GroupTotalCapacity
	//
	// If you omit this parameter, all
	// metrics are enabled.
	Metrics []string
}

type EnableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableMetricsCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableMetricsCollection{}, middleware.After)
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
	if err = addOpEnableMetricsCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableMetricsCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableMetricsCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "EnableMetricsCollection",
	}
}
