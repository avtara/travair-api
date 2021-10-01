package queue

import (
	"encoding/json"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/streadway/amqp"
)

type repoQueue struct {
	publishQueue *amqp.Channel
}

func NewRepoAMPQ(pubAmpq *amqp.Channel) queue.Repository {
	return &repoQueue{
		publishQueue: pubAmpq,
	}
}

func (rq *repoQueue) Publish(name string,raw interface{}) error {
	data, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	_, err = rq.publishQueue.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	err = rq.publishQueue.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}

	defer rq.publishQueue.Close()
	return nil
}
