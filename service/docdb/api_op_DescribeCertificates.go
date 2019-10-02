// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The user-supplied certificate identifier. If this parameter is specified,
	// information for only the specified certificate is returned. If this parameter
	// is omitted, a list of up to MaxRecords certificates is returned. This parameter
	// is not case sensitive.
	//
	// Constraints
	//
	//    * Must match an existing CertificateIdentifier.
	CertificateIdentifier *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeCertificates
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints:
	//
	//    * Minimum: 20
	//
	//    * Maximum: 100
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificatesInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of certificates for this AWS account.
	Certificates []Certificate `locationNameList:"Certificate" type:"list"`

	// An optional pagination token provided if the number of records retrieved
	// is greater than MaxRecords. If this parameter is specified, the marker specifies
	// the next record in the list. Including the value of Marker in the next call
	// to DescribeCertificates results in the next page of certificates.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificates = "DescribeCertificates"

// DescribeCertificatesRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Returns a list of certificate authority (CA) certificates provided by Amazon
// RDS for this AWS account.
//
//    // Example sending a request using DescribeCertificatesRequest.
//    req := client.DescribeCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/DescribeCertificates
func (c *Client) DescribeCertificatesRequest(input *DescribeCertificatesInput) DescribeCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificatesInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificatesOutput{})
	return DescribeCertificatesRequest{Request: req, Input: input, Copy: c.DescribeCertificatesRequest}
}

// DescribeCertificatesRequest is the request type for the
// DescribeCertificates API operation.
type DescribeCertificatesRequest struct {
	*aws.Request
	Input *DescribeCertificatesInput
	Copy  func(*DescribeCertificatesInput) DescribeCertificatesRequest
}

// Send marshals and sends the DescribeCertificates API request.
func (r DescribeCertificatesRequest) Send(ctx context.Context) (*DescribeCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificatesResponse{
		DescribeCertificatesOutput: r.Request.Data.(*DescribeCertificatesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificatesResponse is the response type for the
// DescribeCertificates API operation.
type DescribeCertificatesResponse struct {
	*DescribeCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificates request.
func (r *DescribeCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
