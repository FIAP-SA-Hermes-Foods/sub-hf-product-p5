package broker

import (
	"encoding/json"
	l "sub-hf-product-p5/external/logger"
	ps "sub-hf-product-p5/external/strings"
	sqsBroker "sub-hf-product-p5/internal/core/broker"
	pBroker "sub-hf-product-p5/internal/core/domain/broker"
	"sub-hf-product-p5/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ pBroker.ProductBroker = (*productBroker)(nil)

type productBroker struct {
	queueURL string
	broker   sqsBroker.SQSBroker
}

func NewProductBroker(broker sqsBroker.SQSBroker, queueURL string) *productBroker {
	return &productBroker{broker: broker, queueURL: queueURL}
}

func (p *productBroker) GetProductByID(input dto.ProductBroker) (*dto.ProductBroker, error) {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	nMaxMsg := 5
	waitTime := 30

	inSub := &sqs.ReceiveMessageInput{
		QueueUrl:            &p.queueURL,
		MaxNumberOfMessages: int32(nMaxMsg),
		WaitTimeSeconds:     int32(waitTime),
	}

	result, err := p.broker.Sub(inSub)
	if err != nil {
		return nil, err
	}

	msgs := make([]dto.ProductBroker, 0)

	if result != nil {
		for _, msg := range result.Messages {
			if msg.Body == nil && len(*msg.Body) == 0 {
				continue
			}

			msgBody := string(*msg.Body)
			l.Infof(input.MessageID, "Message received: ", " | ", msgBody)

			msgObj := dto.ProductBroker{}

			if err := json.Unmarshal([]byte(msgBody), msgBody); err != nil {
				return nil, err
			}

			if msgObj.MessageID == input.MessageID {
				msgs = append(msgs, msgObj)

				inDeleteMsg := sqs.DeleteMessageInput{
					QueueUrl:      &p.queueURL,
					ReceiptHandle: msg.ReceiptHandle,
				}

				if _, err := p.broker.Delete(&inDeleteMsg); err != nil {
					return nil, err
				}
			}
		}
	}

	var out = new(dto.ProductBroker)
	if len(msgs) > 0 {
		out = &msgs[0]
	}

	l.Infof(input.MessageID, "Message received with success: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (p *productBroker) SaveProduct(input dto.ProductBroker) (*dto.ProductBroker, error) {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	nMaxMsg := 5
	waitTime := 30

	inSub := &sqs.ReceiveMessageInput{
		QueueUrl:            &p.queueURL,
		MaxNumberOfMessages: int32(nMaxMsg),
		WaitTimeSeconds:     int32(waitTime),
	}

	result, err := p.broker.Sub(inSub)
	if err != nil {
		return nil, err
	}

	msgs := make([]dto.ProductBroker, 0)

	if result != nil {
		for _, msg := range result.Messages {
			if msg.Body == nil && len(*msg.Body) == 0 {
				continue
			}

			msgBody := string(*msg.Body)
			l.Infof(input.MessageID, "Message received: ", " | ", msgBody)

			msgObj := dto.ProductBroker{}

			if err := json.Unmarshal([]byte(msgBody), msgBody); err != nil {
				return nil, err
			}

			if msgObj.MessageID == input.MessageID {
				msgs = append(msgs, msgObj)

				inDeleteMsg := sqs.DeleteMessageInput{
					QueueUrl:      &p.queueURL,
					ReceiptHandle: msg.ReceiptHandle,
				}

				if _, err := p.broker.Delete(&inDeleteMsg); err != nil {
					return nil, err
				}
			}
		}
	}

	var out = new(dto.ProductBroker)
	if len(msgs) > 0 {
		out = &msgs[0]
	}

	l.Infof(input.MessageID, "Message received with success: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (p *productBroker) UpdateProductByID(input dto.ProductBroker) (*dto.ProductBroker, error) {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	nMaxMsg := 5
	waitTime := 30

	inSub := &sqs.ReceiveMessageInput{
		QueueUrl:            &p.queueURL,
		MaxNumberOfMessages: int32(nMaxMsg),
		WaitTimeSeconds:     int32(waitTime),
	}

	result, err := p.broker.Sub(inSub)
	if err != nil {
		return nil, err
	}

	msgs := make([]dto.ProductBroker, 0)

	if result != nil {
		for _, msg := range result.Messages {
			if msg.Body == nil && len(*msg.Body) == 0 {
				continue
			}

			msgBody := string(*msg.Body)
			l.Infof(input.MessageID, "Message received: ", " | ", msgBody)

			msgObj := dto.ProductBroker{}

			if err := json.Unmarshal([]byte(msgBody), msgBody); err != nil {
				return nil, err
			}

			if msgObj.MessageID == input.MessageID {
				msgs = append(msgs, msgObj)

				inDeleteMsg := sqs.DeleteMessageInput{
					QueueUrl:      &p.queueURL,
					ReceiptHandle: msg.ReceiptHandle,
				}

				if _, err := p.broker.Delete(&inDeleteMsg); err != nil {
					return nil, err
				}
			}
		}
	}

	var out = new(dto.ProductBroker)
	if len(msgs) > 0 {
		out = &msgs[0]
	}

	l.Infof(input.MessageID, "Message received with success: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (p *productBroker) GetProductByCategory(input dto.ProductBroker) (*dto.ProductBroker, error) {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	nMaxMsg := 5
	waitTime := 30

	inSub := &sqs.ReceiveMessageInput{
		QueueUrl:            &p.queueURL,
		MaxNumberOfMessages: int32(nMaxMsg),
		WaitTimeSeconds:     int32(waitTime),
	}

	result, err := p.broker.Sub(inSub)
	if err != nil {
		return nil, err
	}

	msgs := make([]dto.ProductBroker, 0)

	if result != nil {
		for _, msg := range result.Messages {
			if msg.Body == nil && len(*msg.Body) == 0 {
				continue
			}

			msgBody := string(*msg.Body)
			l.Infof(input.MessageID, "Message received: ", " | ", msgBody)

			msgObj := dto.ProductBroker{}

			if err := json.Unmarshal([]byte(msgBody), msgBody); err != nil {
				return nil, err
			}

			if msgObj.MessageID == input.MessageID {
				msgs = append(msgs, msgObj)

				inDeleteMsg := sqs.DeleteMessageInput{
					QueueUrl:      &p.queueURL,
					ReceiptHandle: msg.ReceiptHandle,
				}

				if _, err := p.broker.Delete(&inDeleteMsg); err != nil {
					return nil, err
				}
			}
		}
	}

	var out = new(dto.ProductBroker)
	if len(msgs) > 0 {
		out = &msgs[0]
	}

	l.Infof(input.MessageID, "Message received with success: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (p *productBroker) DeleteProductByID(input dto.ProductBroker) (*dto.ProductBroker, error) {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	nMaxMsg := 5
	waitTime := 30

	inSub := &sqs.ReceiveMessageInput{
		QueueUrl:            &p.queueURL,
		MaxNumberOfMessages: int32(nMaxMsg),
		WaitTimeSeconds:     int32(waitTime),
	}

	result, err := p.broker.Sub(inSub)
	if err != nil {
		return nil, err
	}

	msgs := make([]dto.ProductBroker, 0)

	if result != nil {
		for _, msg := range result.Messages {
			if msg.Body == nil && len(*msg.Body) == 0 {
				continue
			}

			msgBody := string(*msg.Body)
			l.Infof(input.MessageID, "Message received: ", " | ", msgBody)

			msgObj := dto.ProductBroker{}

			if err := json.Unmarshal([]byte(msgBody), msgBody); err != nil {
				return nil, err
			}

			if msgObj.MessageID == input.MessageID {
				msgs = append(msgs, msgObj)

				inDeleteMsg := sqs.DeleteMessageInput{
					QueueUrl:      &p.queueURL,
					ReceiptHandle: msg.ReceiptHandle,
				}

				if _, err := p.broker.Delete(&inDeleteMsg); err != nil {
					return nil, err
				}
			}
		}
	}

	var out = new(dto.ProductBroker)
	if len(msgs) > 0 {
		out = &msgs[0]
	}

	l.Infof(input.MessageID, "Message received with success: ", " | ", ps.MarshalString(out))
	return out, nil
}
