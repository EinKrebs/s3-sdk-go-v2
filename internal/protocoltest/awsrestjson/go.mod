module github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.13
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.20
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.14
	github.com/aws/smithy-go v1.13.2
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/
