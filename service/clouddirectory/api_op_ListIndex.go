// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists objects attached to the specified index.
func (c *Client) ListIndex(ctx context.Context, params *ListIndexInput, optFns ...func(*Options)) (*ListIndexOutput, error) {
	if params == nil {
		params = &ListIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIndex", params, optFns, addOperationListIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndexInput struct {

	// The ARN of the directory that the index exists in.
	//
	// This member is required.
	DirectoryArn *string

	// The reference to the index to list.
	//
	// This member is required.
	IndexReference *types.ObjectReference

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of objects in a single page to retrieve from the index during
	// a request. For more information, see Amazon Cloud Directory Limits
	// (http://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html).
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// Specifies the ranges of indexed values that you want to query.
	RangesOnIndexedValues []types.ObjectAttributeRange
}

type ListIndexOutput struct {

	// The objects and indexed values attached to the index.
	IndexAttachments []types.IndexAttachment

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndex{}, middleware.After)
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
	if err = addOpListIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListIndex",
	}
}
