package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	api "operationProject/pkg/api/github.com/example/path/gen"
	"os"
)

func main() {
	var input string

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed connection: %w", err))
	}
	for {
		fmt.Print("Введите числовой пример: ")
		fmt.Fscan(os.Stdin, &input)
		c := api.NewOperationClient(conn)
		res, err := c.Operation(context.Background(), &api.OpRequest{Input: input})
		if err != nil {
			log.Println(fmt.Errorf("failed response: %w", err))
		}
		log.Println("Result: ", res.GetResult())
	}
}
