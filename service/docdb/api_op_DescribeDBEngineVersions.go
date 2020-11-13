// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the available engines.
func (c *Client) DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBEngineVersions", params, optFns, addOperationDescribeDBEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBEngineVersions.
type DescribeDBEngineVersionsInput struct {

	// The name of a specific parameter group family to return details for.
	// Constraints:
	//
	// * If provided, must match an existing DBParameterGroupFamily.
	DBParameterGroupFamily *string

	// Indicates that only the default version of the specified engine or engine and
	// major version combination is returned.
	DefaultOnly bool

	// The database engine to return.
	Engine *string

	// The database engine version to return. Example: 3.6.0
	EngineVersion *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// If this parameter is specified and the requested engine supports the
	// CharacterSetName parameter for CreateDBInstance, the response includes a list of
	// supported character sets for each engine version.
	ListSupportedCharacterSets *bool

	// If this parameter is specified and the requested engine supports the TimeZone
	// parameter for CreateDBInstance, the response includes a list of supported time
	// zones for each engine version.
	ListSupportedTimezones *bool

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Represents the output of DescribeDBEngineVersions.
type DescribeDBEngineVersionsOutput struct {

	// Detailed information about one or more engine versions.
	DBEngineVersions []types.DBEngineVersion

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBEngineVersions{}, middleware.After)
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
	if err = addOpDescribeDBEngineVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBEngineVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBEngineVersions",
	}
}
