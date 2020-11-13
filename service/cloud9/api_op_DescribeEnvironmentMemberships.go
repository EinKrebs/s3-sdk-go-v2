// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about environment members for an AWS Cloud9 development
// environment.
func (c *Client) DescribeEnvironmentMemberships(ctx context.Context, params *DescribeEnvironmentMembershipsInput, optFns ...func(*Options)) (*DescribeEnvironmentMembershipsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentMemberships", params, optFns, addOperationDescribeEnvironmentMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEnvironmentMembershipsInput struct {

	// The ID of the environment to get environment member information about.
	EnvironmentId *string

	// The maximum number of environment members to get information about.
	MaxResults *int32

	// During a previous call, if there are more than 25 items in the list, only the
	// first 25 items are returned, along with a unique string called a next token. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The type of environment member permissions to get information about. Available
	// values include:
	//
	// * owner: Owns the environment.
	//
	// * read-only: Has read-only
	// access to the environment.
	//
	// * read-write: Has read-write access to the
	// environment.
	//
	// If no value is specified, information about all environment
	// members are returned.
	Permissions []types.Permissions

	// The Amazon Resource Name (ARN) of an individual environment member to get
	// information about. If no value is specified, information about all environment
	// members are returned.
	UserArn *string
}

type DescribeEnvironmentMembershipsOutput struct {

	// Information about the environment members for the environment.
	Memberships []types.EnvironmentMember

	// If there are more than 25 items in the list, only the first 25 items are
	// returned, along with a unique string called a next token. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEnvironmentMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEnvironmentMemberships{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentMemberships(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "DescribeEnvironmentMemberships",
	}
}
