package main

import (
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"
	kafkaservices "www.github.com/TalesPalma/kafkaServices"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fasthttp.ListenAndServe(":8080", HandlerFastHttp)
	}()

	go func() {
		defer wg.Done()
		fasthttp.ListenAndServe(":8081", fastHttpHandler)
	}()

	wg.Wait()
}

func HandlerFastHttp(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, kafkaOrderGetMsg())
}

func fastHttpHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there sua uri foi %q", ctx.RequestURI())
}

func kafkaOrderGetMsg() string {
	order := kafkaservices.NewOrderProcessor("order-processor")
	return order.GetMessages()
}
