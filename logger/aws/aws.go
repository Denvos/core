package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type AWS struct {
	client   *cloudwatchlogs.CloudWatchLogs
	group    string
	stream   string
	sequence *string
}

func New(region, group, stream string) (*AWS, error) {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	client := cloudwatchlogs.New(sess)
	return &AWS{
		client: client,
		group:  group,
		stream: stream,
	}, nil
}

func (a *AWS) Write(p []byte) (int, error) {
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(a.group),
		LogStreamName: aws.String(a.stream),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(string(p)),
				Timestamp: aws.Int64(aws.TimeUnixMilliNow()),
			},
		},
		SequenceToken: a.sequence,
	}
	resp, err := a.client.PutLogEvents(input)
	if err != nil {
		return 0, err
	}
	if resp.NextSequenceToken != nil {
		a.sequence = resp.NextSequenceToken
	}
	return len(p), nil
}
