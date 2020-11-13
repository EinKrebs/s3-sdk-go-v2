// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified tags from the specified resources. When you specify a tag
// key, the action removes both that key and its associated value. The operation
// succeeds even if you attempt to remove tags from a resource that were already
// removed. Note the following:
//
// * To remove tags from a resource, you need the
// necessary permissions for the service that the resource belongs to as well as
// permissions for removing tags. For more information, see this list
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
// *
// You can only tag resources that are located in the specified Region for the AWS
// account.
func (c *Client) UntagResources(ctx context.Context, params *UntagResourcesInput, optFns ...func(*Options)) (*UntagResourcesOutput, error) {
	if params == nil {
		params = &UntagResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagResources", params, optFns, addOperationUntagResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagResourcesInput struct {

	// A list of ARNs. An ARN (Amazon Resource Name) uniquely identifies a resource.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	ResourceARNList []string

	// A list of the tag keys that you want to remove from the specified resources.
	//
	// This member is required.
	TagKeys []string
}

type UntagResourcesOutput struct {

	// Details of resources that could not be untagged. An error code, status code, and
	// error message are returned for each failed item.
	FailedResourcesMap map[string]types.FailureInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUntagResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUntagResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUntagResources{}, middleware.After)
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
	if err = addOpUntagResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "UntagResources",
	}
}
