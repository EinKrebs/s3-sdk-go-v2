// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request parameters represent the input of a SQL statement over an array
// of data.
type BatchExecuteStatementInput struct {
	_ struct{} `type:"structure"`

	// The name of the database.
	Database *string `locationName:"database" type:"string"`

	// The parameter set for the batch operation.
	ParameterSets [][]SqlParameter `locationName:"parameterSets" type:"list"`

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" min:"11" type:"string" required:"true"`

	// The name of the database schema.
	Schema *string `locationName:"schema" type:"string"`

	// The name or ARN of the secret that enables access to the DB cluster.
	//
	// SecretArn is a required field
	SecretArn *string `locationName:"secretArn" min:"11" type:"string" required:"true"`

	// The SQL statement to run.
	//
	// Sql is a required field
	Sql *string `locationName:"sql" type:"string" required:"true"`

	// The identifier of a transaction that was started by using the BeginTransaction
	// operation. Specify the transaction ID of the transaction that you want to
	// include the SQL statement in.
	//
	// If the SQL statement is not part of a transaction, don't set this parameter.
	TransactionId *string `locationName:"transactionId" type:"string"`
}

// String returns the string representation
func (s BatchExecuteStatementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchExecuteStatementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchExecuteStatementInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 11 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 11))
	}

	if s.SecretArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretArn"))
	}
	if s.SecretArn != nil && len(*s.SecretArn) < 11 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretArn", 11))
	}

	if s.Sql == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sql"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchExecuteStatementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Database != nil {
		v := *s.Database

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "database", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParameterSets != nil {
		v := s.ParameterSets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "parameterSets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls1 := ls0.List()
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddFields(v2)
			}
			ls1.End()
		}
		ls0.End()

	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecretArn != nil {
		v := *s.SecretArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "secretArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Sql != nil {
		v := *s.Sql

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sql", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TransactionId != nil {
		v := *s.TransactionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transactionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The response elements represent the output of a SQL statement over an array
// of data.
type BatchExecuteStatementOutput struct {
	_ struct{} `type:"structure"`

	// The execution results of each batch entry.
	UpdateResults []UpdateResult `locationName:"updateResults" type:"list"`
}

// String returns the string representation
func (s BatchExecuteStatementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchExecuteStatementOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UpdateResults != nil {
		v := s.UpdateResults

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "updateResults", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchExecuteStatement = "BatchExecuteStatement"

// BatchExecuteStatementRequest returns a request value for making API operation for
// AWS RDS DataService.
//
// Runs a batch SQL statement over an array of data.
//
// You can run bulk update and insert operations for multiple records using
// a DML statement with different parameter sets. Bulk operations can provide
// a significant performance improvement over individual insert and update operations.
//
// If a call isn't part of a transaction because it doesn't include the transactionID
// parameter, changes that result from the call are committed automatically.
//
//    // Example sending a request using BatchExecuteStatementRequest.
//    req := client.BatchExecuteStatementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/BatchExecuteStatement
func (c *Client) BatchExecuteStatementRequest(input *BatchExecuteStatementInput) BatchExecuteStatementRequest {
	op := &aws.Operation{
		Name:       opBatchExecuteStatement,
		HTTPMethod: "POST",
		HTTPPath:   "/BatchExecute",
	}

	if input == nil {
		input = &BatchExecuteStatementInput{}
	}

	req := c.newRequest(op, input, &BatchExecuteStatementOutput{})
	return BatchExecuteStatementRequest{Request: req, Input: input, Copy: c.BatchExecuteStatementRequest}
}

// BatchExecuteStatementRequest is the request type for the
// BatchExecuteStatement API operation.
type BatchExecuteStatementRequest struct {
	*aws.Request
	Input *BatchExecuteStatementInput
	Copy  func(*BatchExecuteStatementInput) BatchExecuteStatementRequest
}

// Send marshals and sends the BatchExecuteStatement API request.
func (r BatchExecuteStatementRequest) Send(ctx context.Context) (*BatchExecuteStatementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchExecuteStatementResponse{
		BatchExecuteStatementOutput: r.Request.Data.(*BatchExecuteStatementOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchExecuteStatementResponse is the response type for the
// BatchExecuteStatement API operation.
type BatchExecuteStatementResponse struct {
	*BatchExecuteStatementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchExecuteStatement request.
func (r *BatchExecuteStatementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
