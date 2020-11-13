// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables you to provide feedback to Amazon Kendra to improve the performance of
// the service.
func (c *Client) SubmitFeedback(ctx context.Context, params *SubmitFeedbackInput, optFns ...func(*Options)) (*SubmitFeedbackOutput, error) {
	if params == nil {
		params = &SubmitFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitFeedback", params, optFns, addOperationSubmitFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitFeedbackInput struct {

	// The identifier of the index that was queried.
	//
	// This member is required.
	IndexId *string

	// The identifier of the specific query for which you are submitting feedback. The
	// query ID is returned in the response to the operation.
	//
	// This member is required.
	QueryId *string

	// Tells Amazon Kendra that a particular search result link was chosen by the user.
	ClickFeedbackItems []types.ClickFeedback

	// Provides Amazon Kendra with relevant or not relevant feedback for whether a
	// particular item was relevant to the search.
	RelevanceFeedbackItems []types.RelevanceFeedback
}

type SubmitFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSubmitFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitFeedback{}, middleware.After)
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
	if err = addOpSubmitFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitFeedback(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "SubmitFeedback",
	}
}
