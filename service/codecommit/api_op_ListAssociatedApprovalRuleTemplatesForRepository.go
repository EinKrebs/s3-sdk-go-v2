// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all approval rule templates that are associated with a specified
// repository.
func (c *Client) ListAssociatedApprovalRuleTemplatesForRepository(ctx context.Context, params *ListAssociatedApprovalRuleTemplatesForRepositoryInput, optFns ...func(*Options)) (*ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	if params == nil {
		params = &ListAssociatedApprovalRuleTemplatesForRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociatedApprovalRuleTemplatesForRepository", params, optFns, addOperationListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociatedApprovalRuleTemplatesForRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedApprovalRuleTemplatesForRepositoryInput struct {

	// The name of the repository for which you want to list all associated approval
	// rule templates.
	//
	// This member is required.
	RepositoryName *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type ListAssociatedApprovalRuleTemplatesForRepositoryOutput struct {

	// The names of all approval rule templates associated with the repository.
	ApprovalRuleTemplateNames []string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
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
	if err = addOpListAssociatedApprovalRuleTemplatesForRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListAssociatedApprovalRuleTemplatesForRepository",
	}
}
