package aws_helper

import (
	"github.com/alphagov/dynolocker/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Returns an AWS session object for the given region, ensuring that the credentials are available
func CreateAwsSession(awsRegion string) (*session.Session, error) {
	session, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(awsRegion)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, errors.WithStackTraceAndPrefix(err, "Error intializing session")
	}

	_, err = session.Config.Credentials.Get()
	if err != nil {
		return nil, errors.WithStackTraceAndPrefix(err, "Error finding AWS credentials")
	}

	return session, nil
}
