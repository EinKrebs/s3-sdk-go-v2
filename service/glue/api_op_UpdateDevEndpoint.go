// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDevEndpointInput struct {
	_ struct{} `type:"structure"`

	// The map of arguments to add the map of arguments used to configure the DevEndpoint.
	//
	// Valid arguments are:
	//
	//    * "--enable-glue-datacatalog": ""
	//
	//    * "GLUE_PYTHON_VERSION": "3"
	//
	//    * "GLUE_PYTHON_VERSION": "2"
	//
	// You can specify a version of Python support for development endpoints by
	// using the Arguments parameter in the CreateDevEndpoint or UpdateDevEndpoint
	// APIs. If no arguments are provided, the version defaults to Python 2.
	AddArguments map[string]string `type:"map"`

	// The list of public keys for the DevEndpoint to use.
	AddPublicKeys []string `type:"list"`

	// Custom Python or Java libraries to be loaded in the DevEndpoint.
	CustomLibraries *DevEndpointCustomLibraries `type:"structure"`

	// The list of argument keys to be deleted from the map of arguments used to
	// configure the DevEndpoint.
	DeleteArguments []string `type:"list"`

	// The list of public keys to be deleted from the DevEndpoint.
	DeletePublicKeys []string `type:"list"`

	// The name of the DevEndpoint to be updated.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`

	// The public key for the DevEndpoint to use.
	PublicKey *string `type:"string"`

	// True if the list of custom libraries to be loaded in the development endpoint
	// needs to be updated, or False if otherwise.
	UpdateEtlLibraries *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateDevEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDevEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDevEndpointInput"}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDevEndpointOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDevEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDevEndpoint = "UpdateDevEndpoint"

// UpdateDevEndpointRequest returns a request value for making API operation for
// AWS Glue.
//
// Updates a specified development endpoint.
//
//    // Example sending a request using UpdateDevEndpointRequest.
//    req := client.UpdateDevEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateDevEndpoint
func (c *Client) UpdateDevEndpointRequest(input *UpdateDevEndpointInput) UpdateDevEndpointRequest {
	op := &aws.Operation{
		Name:       opUpdateDevEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDevEndpointInput{}
	}

	req := c.newRequest(op, input, &UpdateDevEndpointOutput{})
	return UpdateDevEndpointRequest{Request: req, Input: input, Copy: c.UpdateDevEndpointRequest}
}

// UpdateDevEndpointRequest is the request type for the
// UpdateDevEndpoint API operation.
type UpdateDevEndpointRequest struct {
	*aws.Request
	Input *UpdateDevEndpointInput
	Copy  func(*UpdateDevEndpointInput) UpdateDevEndpointRequest
}

// Send marshals and sends the UpdateDevEndpoint API request.
func (r UpdateDevEndpointRequest) Send(ctx context.Context) (*UpdateDevEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDevEndpointResponse{
		UpdateDevEndpointOutput: r.Request.Data.(*UpdateDevEndpointOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDevEndpointResponse is the response type for the
// UpdateDevEndpoint API operation.
type UpdateDevEndpointResponse struct {
	*UpdateDevEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDevEndpoint request.
func (r *UpdateDevEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
