// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the policies that are directly attached to the specified target root,
// organizational unit (OU), or account. You must specify the policy type that you
// want included in the returned list. Always check the NextToken response
// parameter for a null value when calling a List* operation. These operations can
// occasionally return an empty set of results even when there are more results
// available. The NextToken response parameter value is null only when there are no
// more results to display. This operation can be called only from the
// organization's management account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) ListPoliciesForTarget(ctx context.Context, params *ListPoliciesForTargetInput, optFns ...func(*Options)) (*ListPoliciesForTargetOutput, error) {
	if params == nil {
		params = &ListPoliciesForTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPoliciesForTarget", params, optFns, addOperationListPoliciesForTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesForTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesForTargetInput struct {

	// The type of policy that you want to include in the returned list. You must
	// specify one of the following values:
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	// *
	// BACKUP_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	// *
	// SERVICE_CONTROL_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	// *
	// TAG_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	Filter types.PolicyType

	// The unique identifier (ID) of the root, organizational unit, or account whose
	// policies you want to list. The regex pattern (http://wikipedia.org/wiki/regex)
	// for a target ID string requires one of the following:
	//
	// * Root - A string that
	// begins with "r-" followed by from 4 to 32 lowercase letters or digits.
	//
	// *
	// Account - A string that consists of exactly 12 digits.
	//
	// * Organizational unit
	// (OU) - A string that begins with "ou-" followed by from 4 to 32 lowercase
	// letters or digits (the ID of the root that the OU is in). This string is
	// followed by a second "-" dash and from 8 to 32 additional lowercase letters or
	// digits.
	//
	// This member is required.
	TargetId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListPoliciesForTargetOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// The list of policies that match the criteria in the request.
	Policies []types.PolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPoliciesForTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPoliciesForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPoliciesForTarget{}, middleware.After)
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
	if err = addOpListPoliciesForTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPoliciesForTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPoliciesForTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListPoliciesForTarget",
	}
}
