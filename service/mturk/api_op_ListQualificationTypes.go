// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListQualificationTypes operation returns a list of Qualification types,
// filtered by an optional search term.
func (c *Client) ListQualificationTypes(ctx context.Context, params *ListQualificationTypesInput, optFns ...func(*Options)) (*ListQualificationTypesOutput, error) {
	if params == nil {
		params = &ListQualificationTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQualificationTypes", params, optFns, addOperationListQualificationTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQualificationTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQualificationTypesInput struct {

	// Specifies that only Qualification types that a user can request through the
	// Amazon Mechanical Turk web site, such as by taking a Qualification test, are
	// returned as results of the search. Some Qualification types, such as those
	// assigned automatically by the system, cannot be requested directly by users. If
	// false, all Qualification types, including those managed by the system, are
	// considered. Valid values are True | False.
	//
	// This member is required.
	MustBeRequestable *bool

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// Specifies that only Qualification types that the Requester created are returned.
	// If false, the operation returns all Qualification types.
	MustBeOwnedByCaller *bool

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A text query against all of the searchable attributes of Qualification types.
	Query *string
}

type ListQualificationTypesOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of Qualification types on this page in the filtered results list,
	// equivalent to the number of types this operation returns.
	NumResults *int32

	// The list of QualificationType elements returned by the query.
	QualificationTypes []types.QualificationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQualificationTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQualificationTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQualificationTypes{}, middleware.After)
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
	if err = addOpListQualificationTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQualificationTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListQualificationTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListQualificationTypes",
	}
}
