// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the proxy sessions for the specified Amazon Chime Voice Connector.
func (c *Client) ListProxySessions(ctx context.Context, params *ListProxySessionsInput, optFns ...func(*Options)) (*ListProxySessionsOutput, error) {
	if params == nil {
		params = &ListProxySessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProxySessions", params, optFns, addOperationListProxySessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProxySessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProxySessionsInput struct {

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The proxy session status.
	Status types.ProxySessionStatus
}

type ListProxySessionsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The proxy session details.
	ProxySessions []types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProxySessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProxySessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProxySessions{}, middleware.After)
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
	if err = addOpListProxySessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProxySessions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListProxySessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListProxySessions",
	}
}
