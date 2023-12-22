package queue

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/koizumi55555/corporation-api/config"
	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"

	"github.com/koizumi55555/corporation-api/internal/entity"
	"github.com/koizumi55555/corporation-api/internal/usecase"
	"github.com/koizumi55555/corporation-api/pkg/logger"

	"github.com/AlekSi/pointer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type queue struct {
	l    *logger.Logger
	conf *config.QueueConfig
}

func NewQueueUsecase(l *logger.Logger, conf *config.QueueConfig) usecase.Queue {
	return &queue{l: l, conf: conf}
}

func (q *queue) SendMessage(
	ctx context.Context, appErr apierr.ApiErrF,
) apierr.ApiErrF {
	// クライアント生成
	awsConfig := &aws.Config{Region: aws.String(q.conf.Region)}
	session := session.Must(session.NewSession(awsConfig))
	if q.conf.Endpoint != "" {
		session.Config.Endpoint = aws.String(q.conf.Endpoint)
	}
	sqsSvc := sqs.New(session)

	// メッセージをJSON文字列に変換
	messageBody := entity.SQSMessage{
		RequestID: uuid.NewString(),
		Code:      pointer.ToString(strconv.Itoa(int(appErr.StatusCode()))),
		Message:   pointer.ToString(appErr.Error().ErrorMessage),
		SendTime:  time.Now().String(),
	}

	msgJson, err := json.Marshal(messageBody)
	if err != nil {
		q.l.Error("failed to marshal input. %s", err.Error())
		return apierr.ErrorCodeInternalServerError{}
	}

	// 送信するメッセージを作成
	sendMessage := sqs.SendMessageInput{
		MessageBody: aws.String(string(msgJson)),
		QueueUrl:    aws.String(q.conf.Url),
	}

	// メッセージを送信
	_, sendError := sqsSvc.SendMessage(&sendMessage)
	if sendError != nil {
		q.l.Error("failed to enqueue message to sqs: ", sendError.Error())
		q.l.Errorf("fail to send message detail: %#v", sendMessage)

		return apierr.ErrorCodeInternalServerError{}
	}

	return nil
}
