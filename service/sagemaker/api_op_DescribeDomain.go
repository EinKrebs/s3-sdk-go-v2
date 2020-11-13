// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The description of the domain.
func (c *Client) DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) {
	if params == nil {
		params = &DescribeDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomain", params, optFns, addOperationDescribeDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string
}

type DescribeDomainOutput struct {

	// Specifies the VPC used for non-EFS traffic. The default value is
	// PublicInternetOnly.
	//
	// * PublicInternetOnly - Non-EFS traffic is through a VPC
	// managed by Amazon SageMaker, which allows direct internet access
	//
	// * VpcOnly -
	// All Studio traffic is through the specified VPC and subnets
	AppNetworkAccessType types.AppNetworkAccessType

	// The domain's authentication mode.
	AuthMode types.AuthMode

	// The creation time.
	CreationTime *time.Time

	// Settings which are applied to all UserProfile in this domain, if settings are
	// not explicitly specified in a given UserProfile.
	DefaultUserSettings *types.UserSettings

	// The domain's Amazon Resource Name (ARN).
	DomainArn *string

	// The domain ID.
	DomainId *string

	// The domain name.
	DomainName *string

	// The failure reason.
	FailureReason *string

	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string

	// The AWS Key Management Service encryption key ID.
	HomeEfsFileSystemKmsKeyId *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string

	// The status.
	Status types.DomainStatus

	// The VPC subnets that Studio uses for communication.
	SubnetIds []string

	// The domain's URL.
	Url *string

	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for
	// communication.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDomain{}, middleware.After)
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
	if err = addOpDescribeDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeDomain",
	}
}
