// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAdvertiseByoipCidr struct {
}

func (*validateOpAdvertiseByoipCidr) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAdvertiseByoipCidr) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AdvertiseByoipCidrInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAdvertiseByoipCidrInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAccelerator struct {
}

func (*validateOpCreateAccelerator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAcceleratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEndpointGroup struct {
}

func (*validateOpCreateEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateListener struct {
}

func (*validateOpCreateListener) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateListenerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAccelerator struct {
}

func (*validateOpDeleteAccelerator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAcceleratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEndpointGroup struct {
}

func (*validateOpDeleteEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteListener struct {
}

func (*validateOpDeleteListener) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteListenerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeprovisionByoipCidr struct {
}

func (*validateOpDeprovisionByoipCidr) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeprovisionByoipCidr) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeprovisionByoipCidrInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeprovisionByoipCidrInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAcceleratorAttributes struct {
}

func (*validateOpDescribeAcceleratorAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAcceleratorAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAcceleratorAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAcceleratorAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAccelerator struct {
}

func (*validateOpDescribeAccelerator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAcceleratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeEndpointGroup struct {
}

func (*validateOpDescribeEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeListener struct {
}

func (*validateOpDescribeListener) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeListenerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEndpointGroups struct {
}

func (*validateOpListEndpointGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEndpointGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEndpointGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEndpointGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListListeners struct {
}

func (*validateOpListListeners) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListListeners) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListListenersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListListenersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpProvisionByoipCidr struct {
}

func (*validateOpProvisionByoipCidr) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpProvisionByoipCidr) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ProvisionByoipCidrInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpProvisionByoipCidrInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAcceleratorAttributes struct {
}

func (*validateOpUpdateAcceleratorAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAcceleratorAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAcceleratorAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAcceleratorAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAccelerator struct {
}

func (*validateOpUpdateAccelerator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAcceleratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEndpointGroup struct {
}

func (*validateOpUpdateEndpointGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEndpointGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateListener struct {
}

func (*validateOpUpdateListener) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateListenerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpWithdrawByoipCidr struct {
}

func (*validateOpWithdrawByoipCidr) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpWithdrawByoipCidr) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*WithdrawByoipCidrInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpWithdrawByoipCidrInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAdvertiseByoipCidrValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAdvertiseByoipCidr{}, middleware.After)
}

func addOpCreateAcceleratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAccelerator{}, middleware.After)
}

func addOpCreateEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEndpointGroup{}, middleware.After)
}

func addOpCreateListenerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateListener{}, middleware.After)
}

func addOpDeleteAcceleratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAccelerator{}, middleware.After)
}

func addOpDeleteEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEndpointGroup{}, middleware.After)
}

func addOpDeleteListenerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteListener{}, middleware.After)
}

func addOpDeprovisionByoipCidrValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeprovisionByoipCidr{}, middleware.After)
}

func addOpDescribeAcceleratorAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAcceleratorAttributes{}, middleware.After)
}

func addOpDescribeAcceleratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAccelerator{}, middleware.After)
}

func addOpDescribeEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeEndpointGroup{}, middleware.After)
}

func addOpDescribeListenerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeListener{}, middleware.After)
}

func addOpListEndpointGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEndpointGroups{}, middleware.After)
}

func addOpListListenersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListListeners{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpProvisionByoipCidrValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpProvisionByoipCidr{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateAcceleratorAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAcceleratorAttributes{}, middleware.After)
}

func addOpUpdateAcceleratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAccelerator{}, middleware.After)
}

func addOpUpdateEndpointGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEndpointGroup{}, middleware.After)
}

func addOpUpdateListenerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateListener{}, middleware.After)
}

func addOpWithdrawByoipCidrValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpWithdrawByoipCidr{}, middleware.After)
}

func validateCidrAuthorizationContext(v *types.CidrAuthorizationContext) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CidrAuthorizationContext"}
	if v.Message == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Message"))
	}
	if v.Signature == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Signature"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTags(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tags"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAdvertiseByoipCidrInput(v *AdvertiseByoipCidrInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AdvertiseByoipCidrInput"}
	if v.Cidr == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cidr"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAcceleratorInput(v *CreateAcceleratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAcceleratorInput"}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.IdempotencyToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdempotencyToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEndpointGroupInput(v *CreateEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEndpointGroupInput"}
	if v.EndpointGroupRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointGroupRegion"))
	}
	if v.IdempotencyToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdempotencyToken"))
	}
	if v.ListenerArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListenerArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateListenerInput(v *CreateListenerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateListenerInput"}
	if v.IdempotencyToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdempotencyToken"))
	}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if v.PortRanges == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PortRanges"))
	}
	if len(v.Protocol) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Protocol"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAcceleratorInput(v *DeleteAcceleratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAcceleratorInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEndpointGroupInput(v *DeleteEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEndpointGroupInput"}
	if v.EndpointGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteListenerInput(v *DeleteListenerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteListenerInput"}
	if v.ListenerArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListenerArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeprovisionByoipCidrInput(v *DeprovisionByoipCidrInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeprovisionByoipCidrInput"}
	if v.Cidr == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cidr"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAcceleratorAttributesInput(v *DescribeAcceleratorAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAcceleratorAttributesInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAcceleratorInput(v *DescribeAcceleratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAcceleratorInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeEndpointGroupInput(v *DescribeEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeEndpointGroupInput"}
	if v.EndpointGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeListenerInput(v *DescribeListenerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeListenerInput"}
	if v.ListenerArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListenerArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEndpointGroupsInput(v *ListEndpointGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEndpointGroupsInput"}
	if v.ListenerArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListenerArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListListenersInput(v *ListListenersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListListenersInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpProvisionByoipCidrInput(v *ProvisionByoipCidrInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProvisionByoipCidrInput"}
	if v.Cidr == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cidr"))
	}
	if v.CidrAuthorizationContext == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CidrAuthorizationContext"))
	} else if v.CidrAuthorizationContext != nil {
		if err := validateCidrAuthorizationContext(v.CidrAuthorizationContext); err != nil {
			invalidParams.AddNested("CidrAuthorizationContext", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAcceleratorAttributesInput(v *UpdateAcceleratorAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAcceleratorAttributesInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAcceleratorInput(v *UpdateAcceleratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAcceleratorInput"}
	if v.AcceleratorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceleratorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEndpointGroupInput(v *UpdateEndpointGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEndpointGroupInput"}
	if v.EndpointGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateListenerInput(v *UpdateListenerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateListenerInput"}
	if v.ListenerArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListenerArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpWithdrawByoipCidrInput(v *WithdrawByoipCidrInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "WithdrawByoipCidrInput"}
	if v.Cidr == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cidr"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
