package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/greet/greetpb"
	"log"
	"time"
)

func main() {
	// create a insecure connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connected with server: %s", err)
	}
	c := greetpb.NewGreetServiceClient(conn)
	defer conn.Close()

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dlo",
			LastName:  "Bagari",
		},
	}
	// set deadline
	// Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := c.Greet(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("deadline occur")
			}
		} else {
			log.Fatal(err)
		}
		return
	}
	fmt.Println(resp.Result)
}
