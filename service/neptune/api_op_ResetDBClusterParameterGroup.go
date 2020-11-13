// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a DB cluster parameter group to the default value. To
// reset specific parameters submit a list of the following: ParameterName and
// ApplyMethod. To reset the entire DB cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters. When resetting
// the entire group, dynamic parameters are updated immediately and static
// parameters are set to pending-reboot to take effect on the next DB instance
// restart or RebootDBInstance request. You must call RebootDBInstance for every DB
// instance in your DB cluster that you want the updated static parameter to apply
// to.
func (c *Client) ResetDBClusterParameterGroup(ctx context.Context, params *ResetDBClusterParameterGroupInput, optFns ...func(*Options)) (*ResetDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ResetDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetDBClusterParameterGroup", params, optFns, addOperationResetDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetDBClusterParameterGroupInput struct {

	// The name of the DB cluster parameter group to reset.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameter names in the DB cluster parameter group to reset to the
	// default values. You can't use this parameter if the ResetAllParameters parameter
	// is set to true.
	Parameters []types.Parameter

	// A value that is set to true to reset all parameters in the DB cluster parameter
	// group to their default values, and false otherwise. You can't use this parameter
	// if there is a list of parameter names specified for the Parameters parameter.
	ResetAllParameters bool
}

type ResetDBClusterParameterGroupOutput struct {

	// The name of the DB cluster parameter group. Constraints:
	//
	// * Must be 1 to 255
	// letters or numbers.
	//
	// * First character must be a letter
	//
	// * Cannot end with a
	// hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase
	// string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResetDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResetDBClusterParameterGroup{}, middleware.After)
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
	if err = addOpResetDBClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetDBClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetDBClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ResetDBClusterParameterGroup",
	}
}
