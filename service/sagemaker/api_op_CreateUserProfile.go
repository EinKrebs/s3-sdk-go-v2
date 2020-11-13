// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a user profile. A user profile represents a single user within a domain,
// and is the main way to reference a "person" for the purposes of sharing,
// reporting, and other user-oriented features. This entity is created when a user
// onboards to Amazon SageMaker Studio. If an administrator invites a person by
// email or imports them from SSO, a user profile is automatically created. A user
// profile is the primary holder of settings for an individual user and has a
// reference to the user's private Amazon Elastic File System (EFS) home directory.
func (c *Client) CreateUserProfile(ctx context.Context, params *CreateUserProfileInput, optFns ...func(*Options)) (*CreateUserProfileOutput, error) {
	if params == nil {
		params = &CreateUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserProfile", params, optFns, addOperationCreateUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserProfileInput struct {

	// The ID of the associated Domain.
	//
	// This member is required.
	DomainId *string

	// A name for the UserProfile.
	//
	// This member is required.
	UserProfileName *string

	// A specifier for the type of value specified in SingleSignOnUserValue. Currently,
	// the only supported value is "UserName". If the Domain's AuthMode is SSO, this
	// field is required. If the Domain's AuthMode is not SSO, this field cannot be
	// specified.
	SingleSignOnUserIdentifier *string

	// The username of the associated AWS Single Sign-On User for this UserProfile. If
	// the Domain's AuthMode is SSO, this field is required, and must match a valid
	// username of a user in your directory. If the Domain's AuthMode is not SSO, this
	// field cannot be specified.
	SingleSignOnUserValue *string

	// Each tag consists of a key and an optional value. Tag keys must be unique per
	// resource.
	Tags []types.Tag

	// A collection of settings.
	UserSettings *types.UserSettings
}

type CreateUserProfileOutput struct {

	// The user profile Amazon Resource Name (ARN).
	UserProfileArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserProfile{}, middleware.After)
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
	if err = addOpCreateUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateUserProfile",
	}
}
