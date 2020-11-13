// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates an AWS Elastic Beanstalk configuration template, associated with a
// specific Elastic Beanstalk application. You define application configuration
// settings in a configuration template. You can then use the configuration
// template to deploy different versions of the application with the same
// configuration settings. Templates aren't associated with any environment. The
// EnvironmentName response element is always null. Related Topics
//
// *
// DescribeConfigurationOptions
//
// * DescribeConfigurationSettings
//
// *
// ListAvailableSolutionStacks
func (c *Client) CreateConfigurationTemplate(ctx context.Context, params *CreateConfigurationTemplateInput, optFns ...func(*Options)) (*CreateConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationTemplate", params, optFns, addOperationCreateConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create a configuration template.
type CreateConfigurationTemplateInput struct {

	// The name of the Elastic Beanstalk application to associate with this
	// configuration template.
	//
	// This member is required.
	ApplicationName *string

	// The name of the configuration template. Constraint: This name must be unique per
	// application.
	//
	// This member is required.
	TemplateName *string

	// An optional description for this configuration.
	Description *string

	// The ID of an environment whose settings you want to use to create the
	// configuration template. You must specify EnvironmentId if you don't specify
	// PlatformArn, SolutionStackName, or SourceConfiguration.
	EnvironmentId *string

	// Option values for the Elastic Beanstalk configuration, such as the instance
	// type. If specified, these values override the values obtained from the solution
	// stack or the source configuration template. For a complete list of Elastic
	// Beanstalk configuration options, see Option Values
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in
	// the AWS Elastic Beanstalk Developer Guide.
	OptionSettings []types.ConfigurationOptionSetting

	// The Amazon Resource Name (ARN) of the custom platform. For more information, see
	// Custom Platforms
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html)
	// in the AWS Elastic Beanstalk Developer Guide. If you specify PlatformArn, then
	// don't specify SolutionStackName.
	PlatformArn *string

	// The name of an Elastic Beanstalk solution stack (platform version) that this
	// configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7
	// Java 7. A solution stack specifies the operating system, runtime, and
	// application server for a configuration template. It also determines the set of
	// configuration options as well as the possible and default values. For more
	// information, see Supported Platforms
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	// in the AWS Elastic Beanstalk Developer Guide. You must specify SolutionStackName
	// if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration. Use the
	// ListAvailableSolutionStacks
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ListAvailableSolutionStacks.html)
	// API to obtain a list of available solution stacks.
	SolutionStackName *string

	// An Elastic Beanstalk configuration template to base this one on. If specified,
	// Elastic Beanstalk uses the configuration values from the specified configuration
	// template to create a new configuration. Values specified in OptionSettings
	// override any values obtained from the SourceConfiguration. You must specify
	// SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or
	// SolutionStackName. Constraint: If both solution stack name and source
	// configuration are specified, the solution stack of the source configuration
	// template must match the specified solution stack name.
	SourceConfiguration *types.SourceConfiguration

	// Specifies the tags applied to the configuration template.
	Tags []types.Tag
}

// Describes the settings for a configuration set.
type CreateConfigurationTemplateOutput struct {

	// The name of the application associated with this configuration set.
	ApplicationName *string

	// The date (in UTC time) when this configuration set was created.
	DateCreated *time.Time

	// The date (in UTC time) when this configuration set was last modified.
	DateUpdated *time.Time

	// If this configuration set is associated with an environment, the
	// DeploymentStatus parameter indicates the deployment status of this configuration
	// set:
	//
	// * null: This configuration is not associated with a running
	// environment.
	//
	// * pending: This is a draft configuration that is not deployed to
	// the associated environment but is in the process of deploying.
	//
	// * deployed: This
	// is the configuration that is currently deployed to the associated running
	// environment.
	//
	// * failed: This is a draft configuration that failed to
	// successfully deploy.
	DeploymentStatus types.ConfigurationDeploymentStatus

	// Describes this configuration set.
	Description *string

	// If not null, the name of the environment for this configuration set.
	EnvironmentName *string

	// A list of the configuration options and their values in this configuration set.
	OptionSettings []types.ConfigurationOptionSetting

	// The ARN of the platform version.
	PlatformArn *string

	// The name of the solution stack this configuration set uses.
	SolutionStackName *string

	// If not null, the name of the configuration template for this configuration set.
	TemplateName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateConfigurationTemplate{}, middleware.After)
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
	if err = addOpCreateConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreateConfigurationTemplate",
	}
}
