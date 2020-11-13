// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the maintenance window targets or tasks that an
// instance is associated with.
func (c *Client) DescribeMaintenanceWindowsForTarget(ctx context.Context, params *DescribeMaintenanceWindowsForTargetInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowsForTargetOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowsForTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowsForTarget", params, optFns, addOperationDescribeMaintenanceWindowsForTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowsForTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowsForTargetInput struct {

	// The type of resource you want to retrieve information about. For example,
	// "INSTANCE".
	//
	// This member is required.
	ResourceType types.MaintenanceWindowResourceType

	// The instance ID or key/value pair to retrieve information about.
	//
	// This member is required.
	Targets []types.Target

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeMaintenanceWindowsForTargetOutput struct {

	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// Information about the maintenance window targets and tasks an instance is
	// associated with.
	WindowIdentities []types.MaintenanceWindowIdentityForTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMaintenanceWindowsForTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowsForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowsForTarget{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowsForTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowsForTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowsForTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowsForTarget",
	}
}
