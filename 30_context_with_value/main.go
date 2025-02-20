package main

import "context"

type contextKey string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKey("token"), "password")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value(contextKey("token")).(string)
	println(token)
}
