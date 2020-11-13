// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the cors configuration information set for the bucket. To use this
// operation, you must have permission to perform the s3:GetBucketCORS action. By
// default, the bucket owner has this permission and can grant it to others. For
// more information about cors, see  Enabling Cross-Origin Resource Sharing
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html). The following
// operations are related to GetBucketCors:
//
// * PutBucketCors
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html)
//
// *
// DeleteBucketCors
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html)
func (c *Client) GetBucketCors(ctx context.Context, params *GetBucketCorsInput, optFns ...func(*Options)) (*GetBucketCorsOutput, error) {
	if params == nil {
		params = &GetBucketCorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketCors", params, optFns, addOperationGetBucketCorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketCorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketCorsInput struct {

	// The bucket name for which to get the cors configuration.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketCorsOutput struct {

	// A set of origins and methods (cross-origin access that you want to allow). You
	// can add up to 100 rules to the configuration.
	CORSRules []types.CORSRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketCorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketCors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketCors{}, middleware.After)
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
	if err = addOpGetBucketCorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketCors(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketCors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketCors",
	}
}
