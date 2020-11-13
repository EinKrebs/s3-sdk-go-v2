// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an existing cache subnet group.
func (c *Client) ModifyCacheSubnetGroup(ctx context.Context, params *ModifyCacheSubnetGroupInput, optFns ...func(*Options)) (*ModifyCacheSubnetGroupOutput, error) {
	if params == nil {
		params = &ModifyCacheSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCacheSubnetGroup", params, optFns, addOperationModifyCacheSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCacheSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ModifyCacheSubnetGroup operation.
type ModifyCacheSubnetGroupInput struct {

	// The name for the cache subnet group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	// Example: mysubnetgroup
	//
	// This member is required.
	CacheSubnetGroupName *string

	// A description of the cache subnet group.
	CacheSubnetGroupDescription *string

	// The EC2 subnet IDs for the cache subnet group.
	SubnetIds []string
}

type ModifyCacheSubnetGroupOutput struct {

	// Represents the output of one of the following operations:
	//
	// *
	// CreateCacheSubnetGroup
	//
	// * ModifyCacheSubnetGroup
	CacheSubnetGroup *types.CacheSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyCacheSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCacheSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCacheSubnetGroup{}, middleware.After)
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
	if err = addOpModifyCacheSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCacheSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCacheSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyCacheSubnetGroup",
	}
}
