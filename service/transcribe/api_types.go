// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes the input media file in a transcription request.
type Media struct {
	_ struct{} `type:"structure"`

	// The S3 location of the input media file. The URI must be in the same region
	// as the API endpoint that you are calling. The general form is:
	//
	// https://s3-<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3-us-east-1.amazonaws.com/examplebucket/example.mp4
	//
	// https://s3-us-east-1.amazonaws.com/examplebucket/mediadocs/example.mp4
	//
	// For more information about S3 object names, see Object Keys (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide.
	MediaFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Media) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Media) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Media"}
	if s.MediaFileUri != nil && len(*s.MediaFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MediaFileUri", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides optional settings for the StartTranscriptionJob operation.
type Settings struct {
	_ struct{} `type:"structure"`

	// Instructs Amazon Transcribe to process each audio channel separately and
	// then merge the transcription output of each channel into a single transcription.
	//
	// Amazon Transcribe also produces a transcription of each item detected on
	// an audio channel, including the start time and end time of the item and alternative
	// transcriptions of the item including the confidence that Amazon Transcribe
	// has in the transcription.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	ChannelIdentification *bool `type:"boolean"`

	// The maximum number of speakers to identify in the input audio. If there are
	// more speakers in the audio than this number, multiple speakers will be identified
	// as a single speaker. If you specify the MaxSpeakerLabels field, you must
	// set the ShowSpeakerLabels field to true.
	MaxSpeakerLabels *int64 `min:"2" type:"integer"`

	// Determines whether the transcription job uses speaker recognition to identify
	// different speakers in the input audio. Speaker recognition labels individual
	// speakers in the audio file. If you set the ShowSpeakerLabels field to true,
	// you must also set the maximum number of speaker labels MaxSpeakerLabels field.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	ShowSpeakerLabels *bool `type:"boolean"`

	// The name of a vocabulary to use when processing the transcription job.
	VocabularyName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Settings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Settings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Settings"}
	if s.MaxSpeakerLabels != nil && *s.MaxSpeakerLabels < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxSpeakerLabels", 2))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Identifies the location of a transcription.
type Transcript struct {
	_ struct{} `type:"structure"`

	// The location where the transcription is stored.
	//
	// Use this URI to access the transcription. If you specified an S3 bucket in
	// the OutputBucketName field when you created the job, this is the URI of that
	// bucket. If you chose to store the transcription in Amazon Transcribe, this
	// is a shareable URL that provides secure access to that location.
	TranscriptFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Transcript) String() string {
	return awsutil.Prettify(s)
}

// Describes an asynchronous transcription job that was created with the StartTranscriptionJob
// operation.
type TranscriptionJob struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// A timestamp that shows when the job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, this field contains information
	// about why the job failed.
	//
	// The FailureReason field can contain one of the following values:
	//
	//    * Unsupported media format - The media format specified in the MediaFormat
	//    field of the request isn't valid. See the description of the MediaFormat
	//    field for a list of valid values.
	//
	//    * The media format provided does not match the detected media format -
	//    The media format of the audio file doesn't match the format specified
	//    in the MediaFormat field in the request. Check the media format of your
	//    media file and make sure that the two values match.
	//
	//    * Invalid sample rate for audio file - The sample rate specified in the
	//    MediaSampleRateHertz of the request isn't valid. The sample rate must
	//    be between 8000 and 48000 Hertz.
	//
	//    * The sample rate provided does not match the detected sample rate - The
	//    sample rate in the audio file doesn't match the sample rate specified
	//    in the MediaSampleRateHertz field in the request. Check the sample rate
	//    of your media file and make sure that the two values match.
	//
	//    * Invalid file size: file size too large - The size of your audio file
	//    is larger than Amazon Transcribe can process. For more information, see
	//    Limits (https://docs.aws.amazon.com/transcribe/latest/dg/limits-guidelines.html#limits)
	//    in the Amazon Transcribe Developer Guide.
	//
	//    * Invalid number of channels: number of channels too large - Your audio
	//    contains more channels than Amazon Transcribe is configured to process.
	//    To request additional channels, see Amazon Transcribe Limits (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits-amazon-transcribe)
	//    in the Amazon Web Services General Reference.
	FailureReason *string `type:"string"`

	// The language code for the input speech.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// An object that describes the input media for the transcription job.
	Media *Media `type:"structure"`

	// The format of the input media file.
	MediaFormat MediaFormat `type:"string" enum:"true"`

	// The sample rate, in Hertz, of the audio track in the input media file.
	MediaSampleRateHertz *int64 `min:"8000" type:"integer"`

	// Optional settings for the transcription job. Use these settings to turn on
	// speaker recognition, to set the maximum number of speakers that should be
	// identified and to specify a custom vocabulary to use when processing the
	// transcription job.
	Settings *Settings `type:"structure"`

	// An object that describes the output of the transcription job.
	Transcript *Transcript `type:"structure"`

	// The name of the transcription job.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The status of the transcription job.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s TranscriptionJob) String() string {
	return awsutil.Prettify(s)
}

// Provides a summary of information about a transcription job.
type TranscriptionJobSummary struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// A timestamp that shows when the job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, a description of the error.
	FailureReason *string `type:"string"`

	// The language code for the input speech.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Indicates the location of the output of the transcription job.
	//
	// If the value is CUSTOMER_BUCKET then the location is the S3 bucket specified
	// in the outputBucketName field when the transcription job was started with
	// the StartTranscriptionJob operation.
	//
	// If the value is SERVICE_BUCKET then the output is stored by Amazon Transcribe
	// and can be retrieved using the URI in the GetTranscriptionJob response's
	// TranscriptFileUri field.
	OutputLocationType OutputLocationType `type:"string" enum:"true"`

	// The name of the transcription job.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The status of the transcription job. When the status is COMPLETED, use the
	// GetTranscriptionJob operation to get the results of the transcription.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s TranscriptionJobSummary) String() string {
	return awsutil.Prettify(s)
}

// Provides information about a custom vocabulary.
type VocabularyInfo struct {
	_ struct{} `type:"structure"`

	// The language code of the vocabulary entries.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary was last modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the vocabulary.
	VocabularyName *string `min:"1" type:"string"`

	// The processing state of the vocabulary. If the state is READY you can use
	// the vocabulary in a StartTranscriptionJob request.
	VocabularyState VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s VocabularyInfo) String() string {
	return awsutil.Prettify(s)
}
