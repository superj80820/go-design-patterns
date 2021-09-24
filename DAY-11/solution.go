package main

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type RequestID struct{}

func RequestHandler(ctx context.Context) {
	fmt.Printf("request ID is %s\n", ctx.Value(RequestID{}))

	// do something
}

func main() {
	ctx := context.Background()
	go RequestHandler(context.WithValue(ctx, RequestID{}, uuid.NewV4().String()))
	go RequestHandler(context.WithValue(ctx, RequestID{}, uuid.NewV4().String()))
	go RequestHandler(context.WithValue(ctx, RequestID{}, uuid.NewV4().String()))

	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
