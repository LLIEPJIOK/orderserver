package service

import (
	"time"
)

func Retry(operation func() error, maxRetries int, baseDelay time.Duration) error {
	waitDelay := baseDelay

	var err error

	for range maxRetries {
		err = operation()
		if err != nil {
			return nil
		}

		time.Sleep(waitDelay)
		waitDelay *= 2
	}

	return err
}

func Timeout(operation func() error, timeout time.Duration) error {
	errCh := make(chan error)

	go func() {
		errCh <- operation()
	}()

	select {
	case err := <-errCh:
		return err

	case <-time.After(timeout):
		return NewErrTimeout()
	}
}

type DeadLetterQueue struct {
	messages []string
}

func NewDeadLetterQueue() *DeadLetterQueue {
	return &DeadLetterQueue{
		messages: make([]string, 0),
	}
}

func (dlq *DeadLetterQueue) AddMessage(msg string) {
	dlq.messages = append(dlq.messages, msg)
}

func (dlq *DeadLetterQueue) GetMessages(msg string) []string {
	msgCopy := make([]string, len(dlq.messages))
	copy(msgCopy, dlq.messages)

	return msgCopy
}

func ProcessWithDLQ(messages []string, operation func(msg string) error, dlq *DeadLetterQueue) {
	for _, msg := range messages {
		if err := operation(msg); err != nil {
			dlq.AddMessage(msg)
		}
	}
}
