package main

import (
        "context"

        "github.com/corezoid/gitcall-go-runner/gitcall"
)

func usercode(ctx context.Context, data map[string]interface{}) error {

        data["foo"] = 1234

        return nil
}

func main() {
        gitcall.Handle(usercode)
}
