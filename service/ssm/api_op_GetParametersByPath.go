// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve information about one or more parameters in a specific hierarchy.
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults. If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken. You can specify the NextToken in a subsequent call to get the
// next set of results.
func (c *Client) GetParametersByPath(ctx context.Context, params *GetParametersByPathInput, optFns ...func(*Options)) (*GetParametersByPathOutput, error) {
	if params == nil {
		params = &GetParametersByPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParametersByPath", params, optFns, addOperationGetParametersByPathMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParametersByPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersByPathInput struct {

	// The hierarchy for the parameter. Hierarchies start with a forward slash (/) and
	// end with the parameter name. A parameter name hierarchy can have a maximum of 15
	// levels. Here is an example of a hierarchy:
	// /Finance/Prod/IAD/WinServ2016/license33
	//
	// This member is required.
	Path *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// Filters to limit the request results. For GetParametersByPath, the following
	// filter Key names are supported: Type, KeyId, Label, and DataType. The following
	// Key values are not supported for GetParametersByPath: tag, Name, Path, and Tier.
	ParameterFilters []types.ParameterStringFilter

	// Retrieve all parameters within a hierarchy. If a user has access to a path, then
	// the user can access all levels of that path. For example, if a user has
	// permission to access path /a, then the user can also access /a/b. Even if a user
	// has explicitly been denied access in IAM for parameter /a/b, they can still call
	// the GetParametersByPath API action recursively for /a and view /a/b.
	Recursive bool

	// Retrieve all parameters in a hierarchy with their value decrypted.
	WithDecryption bool
}

type GetParametersByPathOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of parameters found in the specified hierarchy.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetParametersByPathMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetParametersByPath{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParametersByPath{}, middleware.After)
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
	if err = addOpGetParametersByPathValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersByPath(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetParametersByPath(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetParametersByPath",
	}
}
