// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified BGP peer on the specified virtual interface with the
// specified customer address and ASN. You cannot delete the last BGP peer from a
// virtual interface.
func (c *Client) DeleteBGPPeer(ctx context.Context, params *DeleteBGPPeerInput, optFns ...func(*Options)) (*DeleteBGPPeerOutput, error) {
	if params == nil {
		params = &DeleteBGPPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBGPPeer", params, optFns, addOperationDeleteBGPPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBGPPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBGPPeerInput struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	Asn int32

	// The ID of the BGP peer.
	BgpPeerId *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string
}

type DeleteBGPPeerOutput struct {

	// The virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBGPPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBGPPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBGPPeer{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBGPPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBGPPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteBGPPeer",
	}
}
