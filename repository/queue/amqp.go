package queue

import (
	"encoding/json"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
)

type repoQueue struct {
	publishQueue *amqp.Channel
}

func NewRepoAMPQ(pubAmpq *amqp.Channel) queue.Repository {
	return &repoQueue{
		publishQueue: pubAmpq,
	}
}

func (rq *repoQueue) EmailUsers(userID uuid.UUID, name, email, payloadType string) {
	data := FromDomainUsers(userID , name, email, payloadType)
	dataJSON, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = rq.publishQueue.QueueDeclare(
		"travair:email",
		false,
		false,
		false,
		false,
		nil,
	)
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
		log.Fatal(err)
	}
}
