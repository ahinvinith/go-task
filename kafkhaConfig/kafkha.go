package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type KafkaConfig struct {
	Address    string `envconfig:"KAFKA_ADDRESS"required:"true"`
	Port       int    `envconfig:"KAFKA_PORT"default:"9092"`
	WriteTopic string `envconfig:"KAFKA_WRITETOPIC"default:"topic1"`
	ReadTopic  string `split_words:"true"`
}

func main() {
	r, err := kafkaconfig()
	if err != nil {
		log.Fatal("fail")
	}
	log.Fatal(r)
}
func kafkaconfig() (*KafkaConfig, error) {
	kafkaconfig := &KafkaConfig{}
	err := envconfig.Process("KAFKA", kafkaconfig)
	if err != nil {
		fmt.Printf("Kafka config parameter error : %v", err)
		return nil, err
	}
	fmt.Printf("kafka config: %+v \n", kafkaconfig)
	return kafkaconfig, nil
}
