package main

import "github.com/rodrigoaasm/truck-monitoring/file-processor/cmd/app"

func main() {
	kafkaConsumer := app.CreateApp()
	if err := kafkaConsumer.Run(); err != nil {
		return
	}
}
