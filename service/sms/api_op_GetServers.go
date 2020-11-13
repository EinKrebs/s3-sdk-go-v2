// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the servers in your server catalog. Before you can describe your
// servers, you must import them using ImportServerCatalog.
func (c *Client) GetServers(ctx context.Context, params *GetServersInput, optFns ...func(*Options)) (*GetServersOutput, error) {
	if params == nil {
		params = &GetServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServers", params, optFns, addOperationGetServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServersInput struct {

	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The server addresses.
	VmServerAddressList []types.VmServerAddress
}

type GetServersOutput struct {

	// The time when the server was last modified.
	LastModifiedOn *time.Time

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The status of the server catalog.
	ServerCatalogStatus types.ServerCatalogStatus

	// Information about the servers.
	ServerList []types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetServers",
	}
}
