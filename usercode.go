package main

import (
	"context"
        "os"

	"github.com/corezoid/gitcall-go-runner/gitcall"
)

func usercode(_ context.Context, data map[string]interface{}) error {
       os.Exit(1)

	return nil
}

func main() {
	gitcall.Handle(usercode)
}
