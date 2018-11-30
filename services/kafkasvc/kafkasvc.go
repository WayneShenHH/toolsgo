package kafkasvc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg        sync.WaitGroup
	kafkaHost = "localhost:9092"
)

func consume() {
	consumer, err := sarama.NewConsumer([]string{kafkaHost}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("testGo")

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("testGo", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}

func produce() {
	config := getConfig()
	producer, err := sarama.NewSyncProducer([]string{kafkaHost}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic:     "testGo",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		// 生产消息
		inputReader := bufio.NewReader(os.Stdin)
		value, err = inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		value = strings.Replace(value, "\n", "", -1)
		msg.Value = sarama.ByteEncoder(value)
		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}
}
func getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	return config
}
func createTopic(topic string) error {
	config := getConfig()
	admin, err := sarama.NewClusterAdmin([]string{kafkaHost}, config)
	err = admin.CreateTopic(topic, &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false)
	return err
}
