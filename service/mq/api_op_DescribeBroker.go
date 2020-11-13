// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the specified broker.
func (c *Client) DescribeBroker(ctx context.Context, params *DescribeBrokerInput, optFns ...func(*Options)) (*DescribeBrokerOutput, error) {
	if params == nil {
		params = &DescribeBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBroker", params, optFns, addOperationDescribeBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBrokerInput struct {

	// The name of the broker. This value must be unique in your AWS account, 1-50
	// characters long, must contain only letters, numbers, dashes, and underscores,
	// and must not contain whitespaces, brackets, wildcard characters, or special
	// characters.
	//
	// This member is required.
	BrokerId *string
}

type DescribeBrokerOutput struct {

	// The authentication strategy used to secure the broker.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. Enables automatic upgrades to new minor versions for brokers, as
	// Apache releases the versions. The automatic upgrades occur during the
	// maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade bool

	// The Amazon Resource Name (ARN) of the broker.
	BrokerArn *string

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// A list of information about allocated brokers.
	BrokerInstances []types.BrokerInstance

	// The name of the broker. This value must be unique in your AWS account, 1-50
	// characters long, must contain only letters, numbers, dashes, and underscores,
	// and must not contain whitespaces, brackets, wildcard characters, or special
	// characters.
	BrokerName *string

	// The status of the broker.
	BrokerState types.BrokerState

	// The list of all revisions for the specified configuration.
	Configurations *types.Configurations

	// The time when the broker was created.
	Created *time.Time

	// Required. The deployment mode of the broker.
	DeploymentMode types.DeploymentMode

	// Encryption options for the broker.
	EncryptionOptions *types.EncryptionOptions

	// Required. The type of broker engine. Note: Currently, Amazon MQ supports only
	// ACTIVEMQ.
	EngineType types.EngineType

	// The version of the broker engine. For a list of supported engine versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// The broker's instance type.
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataOutput

	// The list of information about logs currently enabled and pending to be deployed
	// for the specified broker.
	Logs *types.LogsSummary

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The authentication strategy that will be applied when the broker is rebooted.
	PendingAuthenticationStrategy types.AuthenticationStrategy

	// The version of the broker engine to upgrade to. For a list of supported engine
	// versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	PendingEngineVersion *string

	// The host instance type of the broker to upgrade to. For a list of supported
	// instance types, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide//broker.html#broker-instance-types
	PendingHostInstanceType *string

	// The metadata of the LDAP server that will be used to authenticate and authorize
	// connections to the broker once it is rebooted.
	PendingLdapServerMetadata *types.LdapServerMetadataOutput

	// The list of pending security groups to authorize connections to brokers.
	PendingSecurityGroups []string

	// Required. Enables connections from applications outside of the VPC that hosts
	// the broker's subnets.
	PubliclyAccessible bool

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	// The broker's storage type.
	StorageType types.BrokerStorageType

	// The list of groups (2 maximum) that define which subnets and IP ranges the
	// broker can use from different Availability Zones. A SINGLE_INSTANCE deployment
	// requires one subnet (for example, the default subnet). An
	// ACTIVE_STANDBY_MULTI_AZ deployment requires two subnets.
	SubnetIds []string

	// The list of all tags associated with this broker.
	Tags map[string]string

	// The list of all ActiveMQ usernames for the specified broker.
	Users []types.UserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBroker{}, middleware.After)
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
	if err = addOpDescribeBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBroker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeBroker",
	}
}
