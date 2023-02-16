package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"natsPublisher/generator"
	"natsPublisher/model"
)

func Publish(order *model.Order) {
	sc, err := stan.Connect("L0", "publisher")

	if err != nil {
		fmt.Printf("ERROR: publisher can't connect to nats: %v\n", err)
		return
	}

	defer sc.Close()

	data, _ := json.Marshal(order)
	sc.Publish("L0", data)
}

func main() {
	order := generator.GenerateOrder()
	Publish(order)
}
