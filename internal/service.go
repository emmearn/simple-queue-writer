package internal

import (
	"errors"
	"simple-queue-writer/internal/config"
	"simple-queue-writer/internal/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	ErrBadParam = errors.New("bad_parameter")
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

	isValid := util.IsEmailValid(email)
	if !isValid {
		s.l.ErrorLogger.Printf("Invalid eMail: %s", email)
		return ErrBadParam
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(s.cfg.AWS.Region),
	}))
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
