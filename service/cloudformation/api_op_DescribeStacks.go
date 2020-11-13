// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description for the specified stack; if no stack name was specified,
// then it returns the description for all the stacks created. If the stack does
// not exist, an AmazonCloudFormationException is returned.
func (c *Client) DescribeStacks(ctx context.Context, params *DescribeStacksInput, optFns ...func(*Options)) (*DescribeStacksOutput, error) {
	if params == nil {
		params = &DescribeStacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStacks", params, optFns, addOperationDescribeStacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStacks action.
type DescribeStacksInput struct {

	// A string that identifies the next page of stacks that you want to retrieve.
	NextToken *string

	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	// * Running stacks: You can specify either the stack's
	// name or its unique stack ID.
	//
	// * Deleted stacks: You must specify the unique
	// stack ID.
	//
	// Default: There is no default value.
	StackName *string
}

// The output for a DescribeStacks action.
type DescribeStacksOutput struct {

	// If the output exceeds 1 MB in size, a string that identifies the next page of
	// stacks. If no additional page exists, this value is null.
	NextToken *string

	// A list of stack structures.
	Stacks []types.Stack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStacks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStacks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStacks",
	}
}
