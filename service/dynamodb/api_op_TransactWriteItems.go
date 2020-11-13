// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// TransactWriteItems is a synchronous write operation that groups up to 25 action
// requests. These actions can target items in different tables, but not in
// different AWS accounts or Regions, and no two actions can target the same item.
// For example, you cannot both ConditionCheck and Update the same item. The
// aggregate size of the items in the transaction cannot exceed 4 MB. The actions
// are completed atomically so that either all of them succeed, or all of them
// fail. They are defined by the following objects:
//
// * Put  Initiates a PutItem
// operation to write a new item. This structure specifies the primary key of the
// item to be written, the name of the table to write it in, an optional condition
// expression that must be satisfied for the write to succeed, a list of the item's
// attributes, and a field indicating whether to retrieve the item's attributes if
// the condition is not met.
//
// * Update  Initiates an UpdateItem operation to
// update an existing item. This structure specifies the primary key of the item to
// be updated, the name of the table where it resides, an optional condition
// expression that must be satisfied for the update to succeed, an expression that
// defines one or more attributes to be updated, and a field indicating whether to
// retrieve the item's attributes if the condition is not met.
//
// * Delete 
// Initiates a DeleteItem operation to delete an existing item. This structure
// specifies the primary key of the item to be deleted, the name of the table where
// it resides, an optional condition expression that must be satisfied for the
// deletion to succeed, and a field indicating whether to retrieve the item's
// attributes if the condition is not met.
//
// * ConditionCheck  Applies a condition
// to an item that is not being modified by the transaction. This structure
// specifies the primary key of the item to be checked, the name of the table where
// it resides, a condition expression that must be satisfied for the transaction to
// succeed, and a field indicating whether to retrieve the item's attributes if the
// condition is not met.
//
// DynamoDB rejects the entire TransactWriteItems request if
// any of the following is true:
//
// * A condition in one of the condition expressions
// is not met.
//
// * An ongoing operation is in the process of updating the same
// item.
//
// * There is insufficient provisioned capacity for the transaction to be
// completed.
//
// * An item size becomes too large (bigger than 400 KB), a local
// secondary index (LSI) becomes too large, or a similar validation error occurs
// because of changes made by the transaction.
//
// * The aggregate size of the items
// in the transaction exceeds 4 MB.
//
// * There is a user error, such as an invalid
// data format.
func (c *Client) TransactWriteItems(ctx context.Context, params *TransactWriteItemsInput, optFns ...func(*Options)) (*TransactWriteItemsOutput, error) {
	if params == nil {
		params = &TransactWriteItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TransactWriteItems", params, optFns, addOperationTransactWriteItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TransactWriteItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TransactWriteItemsInput struct {

	// An ordered array of up to 25 TransactWriteItem objects, each of which contains a
	// ConditionCheck, Put, Update, or Delete object. These can operate on items in
	// different tables, but the tables must reside in the same AWS account and Region,
	// and no two of them can operate on the same item.
	//
	// This member is required.
	TransactItems []types.TransactWriteItem

	// Providing a ClientRequestToken makes the call to TransactWriteItems idempotent,
	// meaning that multiple identical calls have the same effect as one single call.
	// Although multiple identical calls using the same client request token produce
	// the same result on the server (no side effects), the responses to the calls
	// might not be the same. If the ReturnConsumedCapacity> parameter is set, then the
	// initial TransactWriteItems call returns the amount of write capacity units
	// consumed in making the changes. Subsequent TransactWriteItems calls with the
	// same client token return the number of read capacity units consumed in reading
	// the item. A client request token is valid for 10 minutes after the first request
	// that uses it is completed. After 10 minutes, any request with the same client
	// token is treated as a new request. Do not resubmit the same request with the
	// same client token for more than 10 minutes, or the result might not be
	// idempotent. If you submit a request with the same client token but a change in
	// other parameters within the 10-minute idempotency window, DynamoDB returns an
	// IdempotentParameterMismatch exception.
	ClientRequestToken *string

	// Determines the level of detail about provisioned throughput consumption that is
	// returned in the response:
	//
	// * INDEXES - The response includes the aggregate
	// ConsumedCapacity for the operation, together with ConsumedCapacity for each
	// table and secondary index that was accessed. Note that some operations, such as
	// GetItem and BatchGetItem, do not access any indexes at all. In these cases,
	// specifying INDEXES will only return ConsumedCapacity information for
	// table(s).
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity
	// for the operation.
	//
	// * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	// Determines whether item collection metrics are returned. If set to SIZE, the
	// response includes statistics about item collections (if any), that were modified
	// during the operation and are returned in the response. If set to NONE (the
	// default), no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics
}

type TransactWriteItemsOutput struct {

	// The capacity units consumed by the entire TransactWriteItems operation. The
	// values of the list are ordered according to the ordering of the TransactItems
	// request parameter.
	ConsumedCapacity []types.ConsumedCapacity

	// A list of tables that were processed by TransactWriteItems and, for each table,
	// information about any item collections that were affected by individual
	// UpdateItem, PutItem, or DeleteItem operations.
	ItemCollectionMetrics map[string][]types.ItemCollectionMetrics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTransactWriteItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpTransactWriteItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpTransactWriteItems{}, middleware.After)
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
	if err = addIdempotencyToken_opTransactWriteItemsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpTransactWriteItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTransactWriteItems(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpTransactWriteItems struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpTransactWriteItems) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpTransactWriteItems) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*TransactWriteItemsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *TransactWriteItemsInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opTransactWriteItemsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpTransactWriteItems{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opTransactWriteItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "TransactWriteItems",
	}
}
