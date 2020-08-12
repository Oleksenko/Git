package main

import (
	"context"
	"os"

	"io/ioutil"

	"fmt"

	"github.com/corezoid/gitcall-go-runner/gitcall"
)

func usercode(_ context.Context, data map[string]interface{}) error {
	n, ok := data["name"]
	if !ok {
		return fmt.Errorf("add unique name to the data")
	}

	f := fmt.Sprintf("trigger-%s.txt", n)
	if fileExists(f) {
		return nil
	}

	if err := ioutil.WriteFile(f, []byte("1"), 0644); err != nil {
		return err
	}

	os.Exit(1)

	return nil
}

func main() {
	gitcall.Handle(usercode)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
