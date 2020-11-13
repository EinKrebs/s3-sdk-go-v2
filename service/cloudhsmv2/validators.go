// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCopyBackupToRegion struct {
}

func (*validateOpCopyBackupToRegion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCopyBackupToRegion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CopyBackupToRegionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCopyBackupToRegionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateHsm struct {
}

func (*validateOpCreateHsm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateHsm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateHsmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateHsmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBackup struct {
}

func (*validateOpDeleteBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBackupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCluster struct {
}

func (*validateOpDeleteCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteHsm struct {
}

func (*validateOpDeleteHsm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteHsm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteHsmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteHsmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInitializeCluster struct {
}

func (*validateOpInitializeCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInitializeCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InitializeClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInitializeClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTags struct {
}

func (*validateOpListTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRestoreBackup struct {
}

func (*validateOpRestoreBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreBackupInput(input); err != nil {
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

func addOpCopyBackupToRegionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCopyBackupToRegion{}, middleware.After)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateHsmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateHsm{}, middleware.After)
}

func addOpDeleteBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBackup{}, middleware.After)
}

func addOpDeleteClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCluster{}, middleware.After)
}

func addOpDeleteHsmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteHsm{}, middleware.After)
}

func addOpInitializeClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInitializeCluster{}, middleware.After)
}

func addOpListTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTags{}, middleware.After)
}

func addOpRestoreBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreBackup{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
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

func validateOpCopyBackupToRegionInput(v *CopyBackupToRegionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CopyBackupToRegionInput"}
	if v.BackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupId"))
	}
	if v.DestinationRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationRegion"))
	}
	if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if v.HsmType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HsmType"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateHsmInput(v *CreateHsmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateHsmInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if v.AvailabilityZone == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZone"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBackupInput(v *DeleteBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBackupInput"}
	if v.BackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterInput(v *DeleteClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteHsmInput(v *DeleteHsmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteHsmInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInitializeClusterInput(v *InitializeClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InitializeClusterInput"}
	if v.ClusterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterId"))
	}
	if v.TrustAnchor == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrustAnchor"))
	}
	if v.SignedCert == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SignedCert"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsInput(v *ListTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsInput"}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreBackupInput(v *RestoreBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreBackupInput"}
	if v.BackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupId"))
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
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if v.TagList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagList"))
	} else if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
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
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if v.TagKeyList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeyList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
