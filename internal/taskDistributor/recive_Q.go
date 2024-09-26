package taskDistributor

import (
	"fmt"
	"time"
)

func (taskDistributor *TaskDistributor) Consumer(queueName string) error {
	taskDistributor.mutex.Lock()
	for _, v := range taskDistributor.ReciverQueueName {
		if v == queueName {
			return nil
		}
	}
	taskDistributor.ReciverQueueName = append(taskDistributor.ReciverQueueName, queueName)
	taskDistributor.mutex.Unlock()

	ch, err := taskDistributor.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil)

	go func() {
		for d := range msgs {
			taskDistributor.mutex.Lock()
			time.Sleep(3 * time.Second)
			fmt.Printf("Received a message: %s\n", d.Body)
			fmt.Print("\n______________\n")
			taskDistributor.mutex.Unlock()
		}
	}()

	<-taskDistributor.Done
	return nil
}
