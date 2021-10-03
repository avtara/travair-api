package users

import (
	"encoding/json"
	"fmt"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/google/uuid"
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

func (rq *repoQueue) EmailUsers(userID uuid.UUID, name, email, payloadType string) error {
	data := FromDomainUsers(userID , name, email, payloadType)
	dataJSON, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	_, err = rq.publishQueue.QueueDeclare(
		"travair:email",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(err)
	err = rq.publishQueue.Publish(
		"",
		"travair:email",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        dataJSON,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
