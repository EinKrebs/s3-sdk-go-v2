// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The newer BatchGetDeploymentTargets should be used instead because it works with
// all compute types. ListDeploymentInstances throws an exception if it is used
// with a compute platform other than EC2/On-premises or AWS Lambda. Lists the
// instance for a deployment associated with the IAM user or AWS account.
func (c *Client) ListDeploymentInstances(ctx context.Context, params *ListDeploymentInstancesInput, optFns ...func(*Options)) (*ListDeploymentInstancesOutput, error) {
	if params == nil {
		params = &ListDeploymentInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentInstances", params, optFns, addOperationListDeploymentInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeploymentInstances operation.
type ListDeploymentInstancesInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// A subset of instances to list by status:
	//
	// * Pending: Include those instances
	// with pending deployments.
	//
	// * InProgress: Include those instances where
	// deployments are still in progress.
	//
	// * Succeeded: Include those instances with
	// successful deployments.
	//
	// * Failed: Include those instances with failed
	// deployments.
	//
	// * Skipped: Include those instances with skipped deployments.
	//
	// *
	// Unknown: Include those instances with deployments in an unknown state.
	InstanceStatusFilter []types.InstanceStatus

	// The set of instances in a blue/green deployment, either those in the original
	// environment ("BLUE") or those in the replacement environment ("GREEN"), for
	// which you want to view instance information.
	InstanceTypeFilter []types.InstanceType

	// An identifier returned from the previous list deployment instances call. It can
	// be used to return the next set of deployment instances in the list.
	NextToken *string
}

// Represents the output of a ListDeploymentInstances operation.
type ListDeploymentInstancesOutput struct {

	// A list of instance IDs.
	InstancesList []string

	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list deployment instances call to return the next
	// set of deployment instances in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeploymentInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentInstances{}, middleware.After)
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
	if err = addOpListDeploymentInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeploymentInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeploymentInstances",
	}
}
