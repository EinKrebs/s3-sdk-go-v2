// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the recovery points for a specified gateway. This operation is only
// supported in the cached volume gateway type. Each cache volume has one recovery
// point. A volume recovery point is a point in time at which all data of the
// volume is consistent and from which you can create a snapshot or clone a new
// cached volume from a source volume. To create a snapshot from a volume recovery
// point use the CreateSnapshotFromVolumeRecoveryPoint operation.
func (c *Client) ListVolumeRecoveryPoints(ctx context.Context, params *ListVolumeRecoveryPointsInput, optFns ...func(*Options)) (*ListVolumeRecoveryPointsOutput, error) {
	if params == nil {
		params = &ListVolumeRecoveryPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVolumeRecoveryPoints", params, optFns, addOperationListVolumeRecoveryPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVolumeRecoveryPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVolumeRecoveryPointsInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

type ListVolumeRecoveryPointsOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// An array of VolumeRecoveryPointInfo objects.
	VolumeRecoveryPointInfos []types.VolumeRecoveryPointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVolumeRecoveryPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumeRecoveryPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumeRecoveryPoints{}, middleware.After)
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
	if err = addOpListVolumeRecoveryPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumeRecoveryPoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVolumeRecoveryPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumeRecoveryPoints",
	}
}
