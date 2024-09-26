package taskDistributor

import (
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type TaskDistributor struct {
	Conn             *amqp.Connection
	Done             chan bool
	mutex            sync.Mutex
	wg               sync.WaitGroup
	ReciverQueueName []string
}

func NewTaskDistributor() (*TaskDistributor, error) {
	var taskDistributor TaskDistributor

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {

		return nil, err
	}

	taskDistributor.Conn = conn
	taskDistributor.Done = make(chan bool)

	return &taskDistributor, nil
}

func (taskDistributor *TaskDistributor) Start() {
	defer taskDistributor.Conn.Close()
	<-taskDistributor.Done
}

func (taskDistributor *TaskDistributor) SendQ(text string, queueName string) error {
	taskDistributor.mutex.Lock()
	defer taskDistributor.mutex.Unlock()
	time.Sleep(2 * time.Second)
	ch, err := taskDistributor.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(text),
		})

	return err
}
