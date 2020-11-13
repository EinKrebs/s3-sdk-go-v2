// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a DataSource that includes metadata and data file information, as well
// as the current status of the DataSource. GetDataSource provides results in
// normal or verbose format. The verbose format adds the schema description and the
// list of files pointed to by the DataSource to the normal format.
func (c *Client) GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) {
	if params == nil {
		params = &GetDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSource", params, optFns, addOperationGetDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSourceInput struct {

	// The ID assigned to the DataSource at creation.
	//
	// This member is required.
	DataSourceId *string

	// Specifies whether the GetDataSource operation should return DataSourceSchema. If
	// true, DataSourceSchema is returned. If false, DataSourceSchema is not returned.
	Verbose bool
}

// Represents the output of a GetDataSource operation and describes a DataSource.
type GetDataSourceOutput struct {

	// The parameter is true if statistics need to be generated from the observation
	// data.
	ComputeStatistics bool

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the DataSource, normalized and scaled on computation resources.
	// ComputeTime is only available if the DataSource is in the COMPLETED state and
	// the ComputeStatistics is set to true.
	ComputeTime *int64

	// The time that the DataSource was created. The time is expressed in epoch time.
	CreatedAt *time.Time

	// The AWS user account from which the DataSource was created. The account type can
	// be either an AWS root account or an AWS Identity and Access Management (IAM)
	// user account.
	CreatedByIamUser *string

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	DataLocationS3 *string

	// A JSON string that represents the splitting and rearrangement requirement used
	// when this DataSource was created.
	DataRearrangement *string

	// The total size of observations in the data files.
	DataSizeInBytes *int64

	// The ID assigned to the DataSource at creation. This value should be identical to
	// the value of the DataSourceId in the request.
	DataSourceId *string

	// The schema used by all of the data files of this DataSource. Note: This
	// parameter is provided as part of the verbose format.
	DataSourceSchema *string

	// The epoch time when Amazon Machine Learning marked the DataSource as COMPLETED
	// or FAILED. FinishedAt is only available when the DataSource is in the COMPLETED
	// or FAILED state.
	FinishedAt *time.Time

	// The time of the most recent edit to the DataSource. The time is expressed in
	// epoch time.
	LastUpdatedAt *time.Time

	// A link to the file containing logs of CreateDataSourceFrom* operations.
	LogUri *string

	// The user-supplied description of the most recent details about creating the
	// DataSource.
	Message *string

	// A user-supplied name or description of the DataSource.
	Name *string

	// The number of data files referenced by the DataSource.
	NumberOfFiles *int64

	// The datasource details that are specific to Amazon RDS.
	RDSMetadata *types.RDSMetadata

	// Describes the DataSource details specific to Amazon Redshift.
	RedshiftMetadata *types.RedshiftMetadata

	// The Amazon Resource Name (ARN) of an AWS IAM Role
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html#roles-about-termsandconcepts),
	// such as the following: arn:aws:iam::account:role/rolename.
	RoleARN *string

	// The epoch time when Amazon Machine Learning marked the DataSource as INPROGRESS.
	// StartedAt isn't available if the DataSource is in the PENDING state.
	StartedAt *time.Time

	// The current status of the DataSource. This element can have one of the following
	// values:
	//
	// * PENDING - Amazon ML submitted a request to create a DataSource.
	//
	// *
	// INPROGRESS - The creation process is underway.
	//
	// * FAILED - The request to create
	// a DataSource did not run to completion. It is not usable.
	//
	// * COMPLETED - The
	// creation process completed successfully.
	//
	// * DELETED - The DataSource is marked
	// as deleted. It is not usable.
	Status types.EntityStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataSource{}, middleware.After)
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
	if err = addOpGetDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "GetDataSource",
	}
}
