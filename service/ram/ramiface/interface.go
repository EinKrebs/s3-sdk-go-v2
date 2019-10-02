// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ramiface provides an interface to enable mocking the AWS Resource Access Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ramiface

import (
	"github.com/aws/aws-sdk-go-v2/service/ram"
)

// ClientAPI provides an interface to enable mocking the
// ram.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // RAM.
//    func myFunc(svc ramiface.ClientAPI) bool {
//        // Make svc.AcceptResourceShareInvitation request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := ram.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        ramiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptResourceShareInvitation(input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AcceptResourceShareInvitationRequest(*ram.AcceptResourceShareInvitationInput) ram.AcceptResourceShareInvitationRequest

	AssociateResourceShareRequest(*ram.AssociateResourceShareInput) ram.AssociateResourceShareRequest

	CreateResourceShareRequest(*ram.CreateResourceShareInput) ram.CreateResourceShareRequest

	DeleteResourceShareRequest(*ram.DeleteResourceShareInput) ram.DeleteResourceShareRequest

	DisassociateResourceShareRequest(*ram.DisassociateResourceShareInput) ram.DisassociateResourceShareRequest

	EnableSharingWithAwsOrganizationRequest(*ram.EnableSharingWithAwsOrganizationInput) ram.EnableSharingWithAwsOrganizationRequest

	GetResourcePoliciesRequest(*ram.GetResourcePoliciesInput) ram.GetResourcePoliciesRequest

	GetResourceShareAssociationsRequest(*ram.GetResourceShareAssociationsInput) ram.GetResourceShareAssociationsRequest

	GetResourceShareInvitationsRequest(*ram.GetResourceShareInvitationsInput) ram.GetResourceShareInvitationsRequest

	GetResourceSharesRequest(*ram.GetResourceSharesInput) ram.GetResourceSharesRequest

	ListPendingInvitationResourcesRequest(*ram.ListPendingInvitationResourcesInput) ram.ListPendingInvitationResourcesRequest

	ListPrincipalsRequest(*ram.ListPrincipalsInput) ram.ListPrincipalsRequest

	ListResourcesRequest(*ram.ListResourcesInput) ram.ListResourcesRequest

	RejectResourceShareInvitationRequest(*ram.RejectResourceShareInvitationInput) ram.RejectResourceShareInvitationRequest

	TagResourceRequest(*ram.TagResourceInput) ram.TagResourceRequest

	UntagResourceRequest(*ram.UntagResourceInput) ram.UntagResourceRequest

	UpdateResourceShareRequest(*ram.UpdateResourceShareInput) ram.UpdateResourceShareRequest
}

var _ ClientAPI = (*ram.Client)(nil)
