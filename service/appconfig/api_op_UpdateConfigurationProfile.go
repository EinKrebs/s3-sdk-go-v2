// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a configuration profile.
func (c *Client) UpdateConfigurationProfile(ctx context.Context, params *UpdateConfigurationProfileInput, optFns ...func(*Options)) (*UpdateConfigurationProfileOutput, error) {
	if params == nil {
		params = &UpdateConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationProfile", params, optFns, addOperationUpdateConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConfigurationProfileInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the configuration profile.
	//
	// This member is required.
	ConfigurationProfileId *string

	// A description of the configuration profile.
	Description *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri.
	RetrievalRoleArn *string

	// A list of methods for validating the configuration.
	Validators []types.Validator
}

type UpdateConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile description.
	Description *string

	// The configuration profile ID.
	Id *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri.
	RetrievalRoleArn *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfigurationProfile{}, middleware.After)
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
	if err = addOpUpdateConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "UpdateConfigurationProfile",
	}
}
