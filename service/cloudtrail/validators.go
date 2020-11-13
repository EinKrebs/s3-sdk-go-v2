// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAddTags struct {
}

func (*validateOpAddTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTrail struct {
}

func (*validateOpCreateTrail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTrail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTrailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTrailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTrail struct {
}

func (*validateOpDeleteTrail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTrail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTrailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTrailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEventSelectors struct {
}

func (*validateOpGetEventSelectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEventSelectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEventSelectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEventSelectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInsightSelectors struct {
}

func (*validateOpGetInsightSelectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInsightSelectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInsightSelectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInsightSelectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTrail struct {
}

func (*validateOpGetTrail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTrail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTrailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTrailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTrailStatus struct {
}

func (*validateOpGetTrailStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTrailStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTrailStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTrailStatusInput(input); err != nil {
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

type validateOpLookupEvents struct {
}

func (*validateOpLookupEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpLookupEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*LookupEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpLookupEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutEventSelectors struct {
}

func (*validateOpPutEventSelectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutEventSelectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutEventSelectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutEventSelectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutInsightSelectors struct {
}

func (*validateOpPutInsightSelectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutInsightSelectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutInsightSelectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutInsightSelectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveTags struct {
}

func (*validateOpRemoveTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartLogging struct {
}

func (*validateOpStartLogging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartLogging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartLoggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartLoggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopLogging struct {
}

func (*validateOpStopLogging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopLogging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopLoggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopLoggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTrail struct {
}

func (*validateOpUpdateTrail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTrail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateTrailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateTrailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTags{}, middleware.After)
}

func addOpCreateTrailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTrail{}, middleware.After)
}

func addOpDeleteTrailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTrail{}, middleware.After)
}

func addOpGetEventSelectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEventSelectors{}, middleware.After)
}

func addOpGetInsightSelectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInsightSelectors{}, middleware.After)
}

func addOpGetTrailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTrail{}, middleware.After)
}

func addOpGetTrailStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTrailStatus{}, middleware.After)
}

func addOpListTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTags{}, middleware.After)
}

func addOpLookupEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpLookupEvents{}, middleware.After)
}

func addOpPutEventSelectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutEventSelectors{}, middleware.After)
}

func addOpPutInsightSelectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutInsightSelectors{}, middleware.After)
}

func addOpRemoveTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveTags{}, middleware.After)
}

func addOpStartLoggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartLogging{}, middleware.After)
}

func addOpStopLoggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopLogging{}, middleware.After)
}

func addOpUpdateTrailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTrail{}, middleware.After)
}

func validateLookupAttribute(v *types.LookupAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LookupAttribute"}
	if v.AttributeValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeValue"))
	}
	if len(v.AttributeKey) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLookupAttributesList(v []types.LookupAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LookupAttributesList"}
	for i := range v {
		if err := validateLookupAttribute(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
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
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagsList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagsList"}
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

func validateOpAddTagsInput(v *AddTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsInput"}
	if v.TagsList != nil {
		if err := validateTagsList(v.TagsList); err != nil {
			invalidParams.AddNested("TagsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTrailInput(v *CreateTrailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTrailInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.TagsList != nil {
		if err := validateTagsList(v.TagsList); err != nil {
			invalidParams.AddNested("TagsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTrailInput(v *DeleteTrailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTrailInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEventSelectorsInput(v *GetEventSelectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEventSelectorsInput"}
	if v.TrailName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrailName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInsightSelectorsInput(v *GetInsightSelectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInsightSelectorsInput"}
	if v.TrailName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrailName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTrailInput(v *GetTrailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTrailInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTrailStatusInput(v *GetTrailStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTrailStatusInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
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
	if v.ResourceIdList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpLookupEventsInput(v *LookupEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LookupEventsInput"}
	if v.LookupAttributes != nil {
		if err := validateLookupAttributesList(v.LookupAttributes); err != nil {
			invalidParams.AddNested("LookupAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutEventSelectorsInput(v *PutEventSelectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutEventSelectorsInput"}
	if v.TrailName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrailName"))
	}
	if v.EventSelectors == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventSelectors"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutInsightSelectorsInput(v *PutInsightSelectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutInsightSelectorsInput"}
	if v.InsightSelectors == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightSelectors"))
	}
	if v.TrailName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrailName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveTagsInput(v *RemoveTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveTagsInput"}
	if v.TagsList != nil {
		if err := validateTagsList(v.TagsList); err != nil {
			invalidParams.AddNested("TagsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartLoggingInput(v *StartLoggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartLoggingInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopLoggingInput(v *StopLoggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopLoggingInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateTrailInput(v *UpdateTrailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTrailInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
