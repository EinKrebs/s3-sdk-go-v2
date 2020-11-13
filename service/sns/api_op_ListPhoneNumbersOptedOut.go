// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of phone numbers that are opted out, meaning you cannot send SMS
// messages to them. The results for ListPhoneNumbersOptedOut are paginated, and
// each page returns up to 100 phone numbers. If additional phone numbers are
// available after the first page of results, then a NextToken string will be
// returned. To receive the next page, you call ListPhoneNumbersOptedOut again
// using the NextToken string received from the previous call. When there are no
// more records to return, NextToken will be null.
func (c *Client) ListPhoneNumbersOptedOut(ctx context.Context, params *ListPhoneNumbersOptedOutInput, optFns ...func(*Options)) (*ListPhoneNumbersOptedOutOutput, error) {
	if params == nil {
		params = &ListPhoneNumbersOptedOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumbersOptedOut", params, optFns, addOperationListPhoneNumbersOptedOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumbersOptedOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutInput struct {

	// A NextToken string is used when you call the ListPhoneNumbersOptedOut action to
	// retrieve additional records that are available after the first page of results.
	NextToken *string
}

// The response from the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutOutput struct {

	// A NextToken string is returned when you call the ListPhoneNumbersOptedOut action
	// if additional records are available after the first page of results.
	NextToken *string

	// A list of phone numbers that are opted out of receiving SMS messages. The list
	// is paginated, and each page can contain up to 100 phone numbers.
	PhoneNumbers []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPhoneNumbersOptedOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPhoneNumbersOptedOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPhoneNumbersOptedOut{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbersOptedOut(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPhoneNumbersOptedOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListPhoneNumbersOptedOut",
	}
}
