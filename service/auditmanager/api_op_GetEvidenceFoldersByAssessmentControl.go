// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of evidence folders associated with a specified control of an
// assessment in AWS Audit Manager.
func (c *Client) GetEvidenceFoldersByAssessmentControl(ctx context.Context, params *GetEvidenceFoldersByAssessmentControlInput, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentControlOutput, error) {
	if params == nil {
		params = &GetEvidenceFoldersByAssessmentControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvidenceFoldersByAssessmentControl", params, optFns, addOperationGetEvidenceFoldersByAssessmentControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvidenceFoldersByAssessmentControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvidenceFoldersByAssessmentControlInput struct {

	// The identifier for the specified assessment.
	//
	// This member is required.
	AssessmentId *string

	// The identifier for the specified control.
	//
	// This member is required.
	ControlId *string

	// The identifier for the specified control set.
	//
	// This member is required.
	ControlSetId *string

	// Represents the maximum number of results per page, or per API request call.
	MaxResults *int32

	// The pagination token used to fetch the next set of results.
	NextToken *string
}

type GetEvidenceFoldersByAssessmentControlOutput struct {

	// The list of evidence folders returned by the
	// GetEvidenceFoldersByAssessmentControl API.
	EvidenceFolders []types.AssessmentEvidenceFolder

	// The pagination token used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEvidenceFoldersByAssessmentControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEvidenceFoldersByAssessmentControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEvidenceFoldersByAssessmentControl{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetEvidenceFoldersByAssessmentControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvidenceFoldersByAssessmentControl(options.Region), middleware.Before); err != nil {
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

// GetEvidenceFoldersByAssessmentControlAPIClient is a client that implements the
// GetEvidenceFoldersByAssessmentControl operation.
type GetEvidenceFoldersByAssessmentControlAPIClient interface {
	GetEvidenceFoldersByAssessmentControl(context.Context, *GetEvidenceFoldersByAssessmentControlInput, ...func(*Options)) (*GetEvidenceFoldersByAssessmentControlOutput, error)
}

var _ GetEvidenceFoldersByAssessmentControlAPIClient = (*Client)(nil)

// GetEvidenceFoldersByAssessmentControlPaginatorOptions is the paginator options
// for GetEvidenceFoldersByAssessmentControl
type GetEvidenceFoldersByAssessmentControlPaginatorOptions struct {
	// Represents the maximum number of results per page, or per API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEvidenceFoldersByAssessmentControlPaginator is a paginator for
// GetEvidenceFoldersByAssessmentControl
type GetEvidenceFoldersByAssessmentControlPaginator struct {
	options   GetEvidenceFoldersByAssessmentControlPaginatorOptions
	client    GetEvidenceFoldersByAssessmentControlAPIClient
	params    *GetEvidenceFoldersByAssessmentControlInput
	nextToken *string
	firstPage bool
}

// NewGetEvidenceFoldersByAssessmentControlPaginator returns a new
// GetEvidenceFoldersByAssessmentControlPaginator
func NewGetEvidenceFoldersByAssessmentControlPaginator(client GetEvidenceFoldersByAssessmentControlAPIClient, params *GetEvidenceFoldersByAssessmentControlInput, optFns ...func(*GetEvidenceFoldersByAssessmentControlPaginatorOptions)) *GetEvidenceFoldersByAssessmentControlPaginator {
	if params == nil {
		params = &GetEvidenceFoldersByAssessmentControlInput{}
	}

	options := GetEvidenceFoldersByAssessmentControlPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEvidenceFoldersByAssessmentControlPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEvidenceFoldersByAssessmentControlPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetEvidenceFoldersByAssessmentControl page.
func (p *GetEvidenceFoldersByAssessmentControlPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentControlOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetEvidenceFoldersByAssessmentControl(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetEvidenceFoldersByAssessmentControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "GetEvidenceFoldersByAssessmentControl",
	}
}
