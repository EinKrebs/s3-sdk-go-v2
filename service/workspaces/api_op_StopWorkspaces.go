// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops the specified WorkSpaces. You cannot stop a WorkSpace unless it has a
// running mode of AutoStop and a state of AVAILABLE, IMPAIRED, UNHEALTHY, or
// ERROR.
func (c *Client) StopWorkspaces(ctx context.Context, params *StopWorkspacesInput, optFns ...func(*Options)) (*StopWorkspacesOutput, error) {
	if params == nil {
		params = &StopWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopWorkspaces", params, optFns, addOperationStopWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopWorkspacesInput struct {

	// The WorkSpaces to stop. You can specify up to 25 WorkSpaces.
	//
	// This member is required.
	StopWorkspaceRequests []types.StopRequest
}

type StopWorkspacesOutput struct {

	// Information about the WorkSpaces that could not be stopped.
	FailedRequests []types.FailedWorkspaceChangeRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopWorkspaces{}, middleware.After)
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
	if err = addOpStopWorkspacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopWorkspaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "StopWorkspaces",
	}
}
