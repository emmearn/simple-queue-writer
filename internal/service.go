package internal

import (
	"simple-queue-writer/internal/config"
	"simple-queue-writer/internal/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Servicer interface {
	SendToQueue(desc string) error
}

type Service struct {
	l   *util.Logger
	cfg *config.Config
}

func NewService(l *util.Logger, cfg *config.Config) (*Service, error) {
	return &Service{
		l:   l,
		cfg: cfg,
	}, nil
}

func (s *Service) SendToQueue(email string) error {

	sess := session.Must(session.NewSession())
	sqsClient := sqs.New(sess)

	result, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &s.cfg.AWS.SQS.QueueURL,
		MessageBody: aws.String(email),
	})
	if err != nil {
		return err
	}

	s.l.InfoLogger.Printf("Successfully sent! %s", *result.MessageId)

	return nil
}
