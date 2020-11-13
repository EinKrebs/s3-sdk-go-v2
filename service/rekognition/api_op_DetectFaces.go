// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detects faces within an image that is provided as input. DetectFaces detects the
// 100 largest faces in the image. For each face detected, the operation returns
// face details. These details include a bounding box of the face, a confidence
// value (that the bounding box contains a face), and a fixed set of attributes
// such as facial landmarks (for example, coordinates of eye and mouth), presence
// of beard, sunglasses, and so on. The face-detection algorithm is most effective
// on frontal faces. For non-frontal or obscured faces, the algorithm might not
// detect the faces or might detect faces with lower confidence. You pass the input
// image either as base64-encoded image bytes or as a reference to an image in an
// Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition operations,
// passing image bytes is not supported. The image must be either a PNG or JPEG
// formatted file. This is a stateless API operation. That is, the operation does
// not persist any data. This operation requires permissions to perform the
// rekognition:DetectFaces action.
func (c *Client) DetectFaces(ctx context.Context, params *DetectFacesInput, optFns ...func(*Options)) (*DetectFacesOutput, error) {
	if params == nil {
		params = &DetectFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectFaces", params, optFns, addOperationDetectFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectFacesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// An array of facial attributes you want to be returned. This can be the default
	// list of attributes or all attributes. If you don't specify a value for
	// Attributes or if you specify ["DEFAULT"], the API returns the following subset
	// of facial attributes: BoundingBox, Confidence, Pose, Quality, and Landmarks. If
	// you provide ["ALL"], all facial attributes are returned, but the operation takes
	// longer to complete. If you provide both, ["ALL", "DEFAULT"], the service uses a
	// logical AND operator to determine which attributes to return (in this case, all
	// attributes).
	Attributes []types.Attribute
}

type DetectFacesOutput struct {

	// Details of each face found in the image.
	FaceDetails []types.FaceDetail

	// The value of OrientationCorrection is always null. If the input image is in
	// .jpeg format, it might contain exchangeable image file format (Exif) metadata
	// that includes the image's orientation. Amazon Rekognition uses this orientation
	// information to perform image correction. The bounding box coordinates are
	// translated to represent object locations after the orientation information in
	// the Exif metadata is used to correct the image orientation. Images in .png
	// format don't contain Exif metadata. Amazon Rekognition doesn’t perform image
	// correction for images in .png format and .jpeg images without orientation
	// information in the image Exif metadata. The bounding box coordinates aren't
	// translated and represent the object locations before the image is rotated.
	OrientationCorrection types.OrientationCorrection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectFaces{}, middleware.After)
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
	if err = addOpDetectFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectFaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectFaces",
	}
}
