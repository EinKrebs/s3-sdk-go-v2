// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines a job to ingest data to IoT SiteWise from Amazon S3. For more
// information, see Create a bulk import job (CLI) (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/CreateBulkImportJob.html)
// in the Amazon Simple Storage Service User Guide. Before you create a bulk import
// job, you must enable IoT SiteWise warm tier or IoT SiteWise cold tier. For more
// information about how to configure storage settings, see PutStorageConfiguration (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_PutStorageConfiguration.html)
// . Bulk import is designed to store historical data to IoT SiteWise. It does not
// trigger computations or notifications on IoT SiteWise warm or cold tier storage.
func (c *Client) CreateBulkImportJob(ctx context.Context, params *CreateBulkImportJobInput, optFns ...func(*Options)) (*CreateBulkImportJobOutput, error) {
	if params == nil {
		params = &CreateBulkImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBulkImportJob", params, optFns, c.addOperationCreateBulkImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBulkImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBulkImportJobInput struct {

	// The Amazon S3 destination where errors associated with the job creation request
	// are saved.
	//
	// This member is required.
	ErrorReportLocation *types.ErrorReportLocation

	// The files in the specified Amazon S3 bucket that contain your data.
	//
	// This member is required.
	Files []types.File

	// Contains the configuration information of a job, such as the file format used
	// to save data in Amazon S3.
	//
	// This member is required.
	JobConfiguration *types.JobConfiguration

	// The unique name that helps identify the job request.
	//
	// This member is required.
	JobName *string

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the IAM role that allows IoT SiteWise to read Amazon S3 data.
	//
	// This member is required.
	JobRoleArn *string

	// If set to true, ingest new data into IoT SiteWise storage. Measurements with
	// notifications, metrics and transforms are computed. If set to false, historical
	// data is ingested into IoT SiteWise as is.
	AdaptiveIngestion *bool

	// If set to true, your data files is deleted from S3, after ingestion into IoT
	// SiteWise storage.
	DeleteFilesAfterImport *bool

	noSmithyDocumentSerde
}

type CreateBulkImportJobOutput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The unique name that helps identify the job request.
	//
	// This member is required.
	JobName *string

	// The status of the bulk import job can be one of following values:
	//   - PENDING – IoT SiteWise is waiting for the current bulk import job to finish.
	//   - CANCELLED – The bulk import job has been canceled.
	//   - RUNNING – IoT SiteWise is processing your request to import your data from
	//   Amazon S3.
	//   - COMPLETED – IoT SiteWise successfully completed your request to import data
	//   from Amazon S3.
	//   - FAILED – IoT SiteWise couldn't process your request to import data from
	//   Amazon S3. You can use logs saved in the specified error report location in
	//   Amazon S3 to troubleshoot issues.
	//   - COMPLETED_WITH_FAILURES – IoT SiteWise completed your request to import data
	//   from Amazon S3 with errors. You can use logs saved in the specified error report
	//   location in Amazon S3 to troubleshoot issues.
	//
	// This member is required.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBulkImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBulkImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBulkImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBulkImportJob"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opCreateBulkImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateBulkImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBulkImportJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opCreateBulkImportJobMiddleware struct {
}

func (*endpointPrefix_opCreateBulkImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateBulkImportJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateBulkImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateBulkImportJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateBulkImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBulkImportJob",
	}
}
