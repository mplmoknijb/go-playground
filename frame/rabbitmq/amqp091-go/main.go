package main

import (
	"context"
	"fmt"
	"time"
)

const (
	KEY = "trace_id"
)

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, "111111")
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Golang梦工厂")
}

func main() {
	ProcessEnter(NewContextWithTraceID())
}
