package main

import (
	"Week02/internal/handler"
	"fmt"
	"log"
)

func main() {
	h, err := handler.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := h.GetNodes()
	if err != nil {
		fmt.Printf("%+v\r\n", err)
	}
	fmt.Println(nodes)
}
