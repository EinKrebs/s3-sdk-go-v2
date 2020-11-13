// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of findings that match the specified criteria.
func (c *Client) GetFindings(ctx context.Context, params *GetFindingsInput, optFns ...func(*Options)) (*GetFindingsOutput, error) {
	if params == nil {
		params = &GetFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindings", params, optFns, addOperationGetFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsInput struct {

	// The finding attributes used to define a condition to filter the returned
	// findings. You can filter by up to 10 finding attributes. For each attribute, you
	// can provide up to 20 filter values. Note that in the available filter fields,
	// WorkflowState is deprecated. To search for a finding based on its workflow
	// status, use WorkflowStatus.
	Filters *types.AwsSecurityFindingFilters

	// The maximum number of findings to return.
	MaxResults int32

	// The token that is required for pagination. On your first call to the GetFindings
	// operation, set the value of this parameter to NULL. For subsequent calls to the
	// operation, to continue listing data, set the value of this parameter to the
	// value returned from the previous response.
	NextToken *string

	// The finding attributes used to sort the list of returned findings.
	SortCriteria []types.SortCriterion
}

type GetFindingsOutput struct {

	// The findings that matched the filters specified in the request.
	//
	// This member is required.
	Findings []types.AwsSecurityFinding

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetFindings",
	}
}
