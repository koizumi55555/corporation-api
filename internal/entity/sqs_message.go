package entity

type SQSMessage struct {
	RequestID string
	Code      *string
	Message   *string
	SendTime  string
}
