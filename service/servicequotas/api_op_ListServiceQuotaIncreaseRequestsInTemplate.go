// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the quota increase requests in the template.
func (c *Client) ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, params *ListServiceQuotaIncreaseRequestsInTemplateInput, optFns ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	if params == nil {
		params = &ListServiceQuotaIncreaseRequestsInTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceQuotaIncreaseRequestsInTemplate", params, optFns, addOperationListServiceQuotaIncreaseRequestsInTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceQuotaIncreaseRequestsInTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceQuotaIncreaseRequestsInTemplateInput struct {

	// Specifies the AWS Region for the quota that you want to use.
	AwsRegion *string

	// (Optional) Limits the number of results that you want to include in the
	// response. If you don't include this parameter, the response defaults to a value
	// that's specific to the operation. If additional items exist beyond the specified
	// maximum, the NextToken element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the call to the operation to
	// get the next part of the results. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int32

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available. In a
	// subsequent call, set it to the value of the previous call's NextToken response
	// to indicate where the output should continue from.
	NextToken *string

	// The identifier for a service. When performing an operation, use the ServiceCode
	// to specify a particular service.
	ServiceCode *string
}

type ListServiceQuotaIncreaseRequestsInTemplateOutput struct {

	// If present in the response, this value indicates there's more output available
	// that what's included in the current response. This can occur even when the
	// response includes no values at all, such as when you ask for a filtered view of
	// a very long list. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to continue processing and get the next part of
	// the output. You should repeat this until the NextToken response element comes
	// back empty (as null).
	NextToken *string

	// Returns the list of values of the quota increase request in the template.
	ServiceQuotaIncreaseRequestInTemplateList []types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListServiceQuotaIncreaseRequestsInTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListServiceQuotaIncreaseRequestsInTemplate",
	}
}
