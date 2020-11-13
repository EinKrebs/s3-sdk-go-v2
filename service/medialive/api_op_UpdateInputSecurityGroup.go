// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update an Input Security Group's Whilelists.
func (c *Client) UpdateInputSecurityGroup(ctx context.Context, params *UpdateInputSecurityGroupInput, optFns ...func(*Options)) (*UpdateInputSecurityGroupOutput, error) {
	if params == nil {
		params = &UpdateInputSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInputSecurityGroup", params, optFns, addOperationUpdateInputSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInputSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to update some combination of the Input Security Group name and the
// IPv4 CIDRs the Input Security Group should allow.
type UpdateInputSecurityGroupInput struct {

	// The id of the Input Security Group to update.
	//
	// This member is required.
	InputSecurityGroupId *string

	// A collection of key-value pairs.
	Tags map[string]string

	// List of IPv4 CIDR addresses to whitelist
	WhitelistRules []types.InputWhitelistRuleCidr
}

// Placeholder documentation for UpdateInputSecurityGroupResponse
type UpdateInputSecurityGroupOutput struct {

	// An Input Security Group
	SecurityGroup *types.InputSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInputSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInputSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInputSecurityGroup{}, middleware.After)
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
	if err = addOpUpdateInputSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInputSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInputSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateInputSecurityGroup",
	}
}
