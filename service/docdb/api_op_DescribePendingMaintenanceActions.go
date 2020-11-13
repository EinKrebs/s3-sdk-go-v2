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

// Returns a list of resources (for example, instances) that have at least one
// pending maintenance action.
func (c *Client) DescribePendingMaintenanceActions(ctx context.Context, params *DescribePendingMaintenanceActionsInput, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) {
	if params == nil {
		params = &DescribePendingMaintenanceActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePendingMaintenanceActions", params, optFns, addOperationDescribePendingMaintenanceActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePendingMaintenanceActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribePendingMaintenanceActions.
type DescribePendingMaintenanceActionsInput struct {

	// A filter that specifies one or more resources to return pending maintenance
	// actions for. Supported filters:
	//
	// * db-cluster-id - Accepts cluster identifiers
	// and cluster Amazon Resource Names (ARNs). The results list includes only pending
	// maintenance actions for the clusters identified by these ARNs.
	//
	// * db-instance-id
	// - Accepts instance identifiers and instance ARNs. The results list includes only
	// pending maintenance actions for the DB instances identified by these ARNs.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The ARN of a resource to return pending maintenance actions for.
	ResourceIdentifier *string
}

// Represents the output of DescribePendingMaintenanceActions.
type DescribePendingMaintenanceActionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maintenance actions to be applied.
	PendingMaintenanceActions []types.ResourcePendingMaintenanceActions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePendingMaintenanceActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribePendingMaintenanceActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribePendingMaintenanceActions{}, middleware.After)
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
	if err = addOpDescribePendingMaintenanceActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribePendingMaintenanceActions",
	}
}
