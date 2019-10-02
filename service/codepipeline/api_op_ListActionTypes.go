// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListActionTypes action.
type ListActionTypesInput struct {
	_ struct{} `type:"structure"`

	// Filters the list of action types to those created by a specified entity.
	ActionOwnerFilter ActionOwner `locationName:"actionOwnerFilter" type:"string" enum:"true"`

	// An identifier that was returned from the previous list action types call,
	// which can be used to return the next set of action types in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActionTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActionTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActionTypesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListActionTypes action.
type ListActionTypesOutput struct {
	_ struct{} `type:"structure"`

	// Provides details of the action types.
	//
	// ActionTypes is a required field
	ActionTypes []ActionType `locationName:"actionTypes" type:"list" required:"true"`

	// If the amount of returned information is significantly large, an identifier
	// is also returned. It can be used in a subsequent list action types call to
	// return the next set of action types in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActionTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListActionTypes = "ListActionTypes"

// ListActionTypesRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Gets a summary of all AWS CodePipeline action types associated with your
// account.
//
//    // Example sending a request using ListActionTypesRequest.
//    req := client.ListActionTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/ListActionTypes
func (c *Client) ListActionTypesRequest(input *ListActionTypesInput) ListActionTypesRequest {
	op := &aws.Operation{
		Name:       opListActionTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListActionTypesInput{}
	}

	req := c.newRequest(op, input, &ListActionTypesOutput{})
	return ListActionTypesRequest{Request: req, Input: input, Copy: c.ListActionTypesRequest}
}

// ListActionTypesRequest is the request type for the
// ListActionTypes API operation.
type ListActionTypesRequest struct {
	*aws.Request
	Input *ListActionTypesInput
	Copy  func(*ListActionTypesInput) ListActionTypesRequest
}

// Send marshals and sends the ListActionTypes API request.
func (r ListActionTypesRequest) Send(ctx context.Context) (*ListActionTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActionTypesResponse{
		ListActionTypesOutput: r.Request.Data.(*ListActionTypesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListActionTypesRequestPaginator returns a paginator for ListActionTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListActionTypesRequest(input)
//   p := codepipeline.NewListActionTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListActionTypesPaginator(req ListActionTypesRequest) ListActionTypesPaginator {
	return ListActionTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListActionTypesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListActionTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListActionTypesPaginator struct {
	aws.Pager
}

func (p *ListActionTypesPaginator) CurrentPage() *ListActionTypesOutput {
	return p.Pager.CurrentPage().(*ListActionTypesOutput)
}

// ListActionTypesResponse is the response type for the
// ListActionTypes API operation.
type ListActionTypesResponse struct {
	*ListActionTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActionTypes request.
func (r *ListActionTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
