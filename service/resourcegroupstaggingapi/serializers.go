// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpDescribeReportCreation struct {
}

func (*awsAwsjson11_serializeOpDescribeReportCreation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeReportCreation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeReportCreationInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.DescribeReportCreation")

	_ = input

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetComplianceSummary struct {
}

func (*awsAwsjson11_serializeOpGetComplianceSummary) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetComplianceSummary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetComplianceSummaryInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.GetComplianceSummary")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetComplianceSummaryInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetResources struct {
}

func (*awsAwsjson11_serializeOpGetResources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetResourcesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.GetResources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetTagKeys struct {
}

func (*awsAwsjson11_serializeOpGetTagKeys) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetTagKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTagKeysInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.GetTagKeys")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetTagKeysInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetTagValues struct {
}

func (*awsAwsjson11_serializeOpGetTagValues) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetTagValues) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTagValuesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.GetTagValues")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetTagValuesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpStartReportCreation struct {
}

func (*awsAwsjson11_serializeOpStartReportCreation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpStartReportCreation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartReportCreationInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.StartReportCreation")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentStartReportCreationInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpTagResources struct {
}

func (*awsAwsjson11_serializeOpTagResources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpTagResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourcesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.TagResources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentTagResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUntagResources struct {
}

func (*awsAwsjson11_serializeOpUntagResources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUntagResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourcesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ResourceGroupsTaggingAPI_20170126.UntagResources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUntagResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentGroupBy(v []types.GroupByAttribute, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsAwsjson11_serializeDocumentRegionFilterList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentResourceARNList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentResourceTypeFilterList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagFilter(v *types.TagFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Values != nil {
		ok := object.Key("Values")
		if err := awsAwsjson11_serializeDocumentTagValueList(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentTagFilterList(v []types.TagFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTagFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagKeyFilterList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagKeyListForUntag(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagValueList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTargetIdFilterList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeReportCreationInput(v *DescribeReportCreationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson11_serializeOpDocumentGetComplianceSummaryInput(v *GetComplianceSummaryInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GroupBy != nil {
		ok := object.Key("GroupBy")
		if err := awsAwsjson11_serializeDocumentGroupBy(v.GroupBy, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.PaginationToken != nil {
		ok := object.Key("PaginationToken")
		ok.String(*v.PaginationToken)
	}

	if v.RegionFilters != nil {
		ok := object.Key("RegionFilters")
		if err := awsAwsjson11_serializeDocumentRegionFilterList(v.RegionFilters, ok); err != nil {
			return err
		}
	}

	if v.ResourceTypeFilters != nil {
		ok := object.Key("ResourceTypeFilters")
		if err := awsAwsjson11_serializeDocumentResourceTypeFilterList(v.ResourceTypeFilters, ok); err != nil {
			return err
		}
	}

	if v.TagKeyFilters != nil {
		ok := object.Key("TagKeyFilters")
		if err := awsAwsjson11_serializeDocumentTagKeyFilterList(v.TagKeyFilters, ok); err != nil {
			return err
		}
	}

	if v.TargetIdFilters != nil {
		ok := object.Key("TargetIdFilters")
		if err := awsAwsjson11_serializeDocumentTargetIdFilterList(v.TargetIdFilters, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetResourcesInput(v *GetResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExcludeCompliantResources != nil {
		ok := object.Key("ExcludeCompliantResources")
		ok.Boolean(*v.ExcludeCompliantResources)
	}

	if v.IncludeComplianceDetails != nil {
		ok := object.Key("IncludeComplianceDetails")
		ok.Boolean(*v.IncludeComplianceDetails)
	}

	if v.PaginationToken != nil {
		ok := object.Key("PaginationToken")
		ok.String(*v.PaginationToken)
	}

	if v.ResourcesPerPage != nil {
		ok := object.Key("ResourcesPerPage")
		ok.Integer(*v.ResourcesPerPage)
	}

	if v.ResourceTypeFilters != nil {
		ok := object.Key("ResourceTypeFilters")
		if err := awsAwsjson11_serializeDocumentResourceTypeFilterList(v.ResourceTypeFilters, ok); err != nil {
			return err
		}
	}

	if v.TagFilters != nil {
		ok := object.Key("TagFilters")
		if err := awsAwsjson11_serializeDocumentTagFilterList(v.TagFilters, ok); err != nil {
			return err
		}
	}

	if v.TagsPerPage != nil {
		ok := object.Key("TagsPerPage")
		ok.Integer(*v.TagsPerPage)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetTagKeysInput(v *GetTagKeysInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.PaginationToken != nil {
		ok := object.Key("PaginationToken")
		ok.String(*v.PaginationToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetTagValuesInput(v *GetTagValuesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.PaginationToken != nil {
		ok := object.Key("PaginationToken")
		ok.String(*v.PaginationToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentStartReportCreationInput(v *StartReportCreationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3Bucket != nil {
		ok := object.Key("S3Bucket")
		ok.String(*v.S3Bucket)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentTagResourcesInput(v *TagResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARNList != nil {
		ok := object.Key("ResourceARNList")
		if err := awsAwsjson11_serializeDocumentResourceARNList(v.ResourceARNList, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson11_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUntagResourcesInput(v *UntagResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARNList != nil {
		ok := object.Key("ResourceARNList")
		if err := awsAwsjson11_serializeDocumentResourceARNList(v.ResourceARNList, ok); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		ok := object.Key("TagKeys")
		if err := awsAwsjson11_serializeDocumentTagKeyListForUntag(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}
