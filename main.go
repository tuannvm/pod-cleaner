package main

import (
	"log"

	"github.com/honestbee/pod-cleaner/pkg/kube"
)

func main() {
	client := kube.New()
	err := client.DeleteEvictPods()
	if err != nil {
		log.Fatal(err)
	}
}
