// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all of the email identities that are associated with your
// Amazon Pinpoint account. An identity can be either an email address or a domain.
// This operation returns identities that are verified as well as those that
// aren't.
func (c *Client) ListEmailIdentities(ctx context.Context, params *ListEmailIdentitiesInput, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) {
	if params == nil {
		params = &ListEmailIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEmailIdentities", params, optFns, addOperationListEmailIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEmailIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list all of the email identities associated with your Amazon
// Pinpoint account. This list includes identities that you've already verified,
// identities that are unverified, and identities that were verified in the past,
// but are no longer verified.
type ListEmailIdentitiesInput struct {

	// A token returned from a previous call to ListEmailIdentities to indicate the
	// position in the list of identities.
	NextToken *string

	// The number of results to show in a single call to ListEmailIdentities. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 0, and can be no
	// more than 1000.
	PageSize *int32
}

// A list of all of the identities that you've attempted to verify for use with
// Amazon Pinpoint, regardless of whether or not those identities were successfully
// verified.
type ListEmailIdentitiesOutput struct {

	// An array that includes all of the identities associated with your Amazon
	// Pinpoint account.
	EmailIdentities []types.IdentityInfo

	// A token that indicates that there are additional configuration sets to list. To
	// view additional configuration sets, issue another request to
	// ListEmailIdentities, and pass this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEmailIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEmailIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEmailIdentities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEmailIdentities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEmailIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListEmailIdentities",
	}
}
