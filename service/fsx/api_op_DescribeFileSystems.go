// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx file systems, if a FileSystemIds
// value is provided for that file system. Otherwise, it returns descriptions of
// all file systems owned by your AWS account in the AWS Region of the endpoint
// that you're calling. When retrieving all file system descriptions, you can
// optionally specify the MaxResults parameter to limit the number of descriptions
// in a response. If more file system descriptions remain, Amazon FSx returns a
// NextToken value in the response. In this case, send a later request with the
// NextToken request parameter set to the value of NextToken from the last
// response. This action is used in an iterative process to retrieve a list of your
// file system descriptions. DescribeFileSystems is called first without a
// NextTokenvalue. Then the action continues to be called with the NextToken
// parameter set to the value of the last NextToken value until a response has no
// NextToken. When using this action, keep the following in mind:
//
// * The
// implementation might return fewer than MaxResults file system descriptions while
// still including a NextToken value.
//
// * The order of file systems returned in the
// response of one DescribeFileSystems call and the order of file systems returned
// across the responses of a multicall iteration is unspecified.
func (c *Client) DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
	if params == nil {
		params = &DescribeFileSystemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFileSystems", params, optFns, addOperationDescribeFileSystemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFileSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DescribeFileSystems operation.
type DescribeFileSystemsInput struct {

	// IDs of the file systems whose descriptions you want to retrieve (String).
	FileSystemIds []string

	// Maximum number of file systems to return in the response (integer). This
	// parameter value must be greater than 0. The number of items that Amazon FSx
	// returns is the minimum of the MaxResults parameter specified in the request and
	// the service's internal maximum number of items per page.
	MaxResults *int32

	// Opaque pagination token returned from a previous DescribeFileSystems operation
	// (String). If a token present, the action continues the list from where the
	// returning call left off.
	NextToken *string
}

// The response object for DescribeFileSystems operation.
type DescribeFileSystemsOutput struct {

	// An array of file system descriptions.
	FileSystems []types.FileSystem

	// Present if there are more file systems than returned in the response (String).
	// You can use the NextToken value in the later request to fetch the descriptions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFileSystemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFileSystems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFileSystems{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFileSystems(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFileSystems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeFileSystems",
	}
}
