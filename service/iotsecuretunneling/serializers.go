// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsecuretunneling

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpCloseTunnel struct {
}

func (*awsAwsjson11_serializeOpCloseTunnel) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCloseTunnel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CloseTunnelInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.CloseTunnel")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCloseTunnelInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpDescribeTunnel struct {
}

func (*awsAwsjson11_serializeOpDescribeTunnel) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTunnel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTunnelInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.DescribeTunnel")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeTunnelInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpListTagsForResource struct {
}

func (*awsAwsjson11_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpListTunnels struct {
}

func (*awsAwsjson11_serializeOpListTunnels) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTunnels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTunnelsInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.ListTunnels")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTunnelsInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpOpenTunnel struct {
}

func (*awsAwsjson11_serializeOpOpenTunnel) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpOpenTunnel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*OpenTunnelInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.OpenTunnel")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentOpenTunnelInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpTagResource struct {
}

func (*awsAwsjson11_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpUntagResource struct {
}

func (*awsAwsjson11_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("IoTSecuredTunneling.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsAwsjson11_serializeDocumentDestinationConfig(v *types.DestinationConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Services != nil {
		ok := object.Key("services")
		if err := awsAwsjson11_serializeDocumentServiceList(v.Services, ok); err != nil {
			return err
		}
	}

	if v.ThingName != nil {
		ok := object.Key("thingName")
		ok.String(*v.ThingName)
	}

	return nil
}

func awsAwsjson11_serializeDocumentServiceList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentTimeoutConfig(v *types.TimeoutConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxLifetimeTimeoutMinutes != 0 {
		ok := object.Key("maxLifetimeTimeoutMinutes")
		ok.Integer(v.MaxLifetimeTimeoutMinutes)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentCloseTunnelInput(v *CloseTunnelInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Delete {
		ok := object.Key("delete")
		ok.Boolean(v.Delete)
	}

	if v.TunnelId != nil {
		ok := object.Key("tunnelId")
		ok.String(*v.TunnelId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeTunnelInput(v *DescribeTunnelInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TunnelId != nil {
		ok := object.Key("tunnelId")
		ok.String(*v.TunnelId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTunnelsInput(v *ListTunnelsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != 0 {
		ok := object.Key("maxResults")
		ok.Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	if v.ThingName != nil {
		ok := object.Key("thingName")
		ok.String(*v.ThingName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentOpenTunnelInput(v *OpenTunnelInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if v.DestinationConfig != nil {
		ok := object.Key("destinationConfig")
		if err := awsAwsjson11_serializeDocumentDestinationConfig(v.DestinationConfig, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsAwsjson11_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	if v.TimeoutConfig != nil {
		ok := object.Key("timeoutConfig")
		if err := awsAwsjson11_serializeDocumentTimeoutConfig(v.TimeoutConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsAwsjson11_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.TagKeys != nil {
		ok := object.Key("tagKeys")
		if err := awsAwsjson11_serializeDocumentTagKeyList(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}
