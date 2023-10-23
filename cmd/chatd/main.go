package main

import (
	"bufio"
	"fmt"
	"github.com/q00/redpanda-chat/internal/redpanda"
	"github.com/q00/redpanda-chat/tools"
	"os"
	"strings"
)

func main() {
	topic := "chat-room"
	brokers := []string{"localhost:19092"}
	admin := tools.NewAdmin(brokers)
	defer admin.Close()
	if !admin.TopicExists(topic) {
		admin.CreateTopic(topic)
	}
	username := ""
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	producer := redpanda.NewProducer(brokers, topic)
	defer producer.Close()
	consumer := redpanda.NewConsumer(brokers, topic)
	defer consumer.Close()
	go consumer.PrintMessages()
	fmt.Println("Connected. Press Ctrl+C to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		producer.SendMessage(username, message)
	}
}
