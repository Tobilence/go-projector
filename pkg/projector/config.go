package projector

import (
	"fmt"
	"os"
	"path/filepath"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation
	Pwd       string
	Config    string
}

func getPwd(opts Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getOperation(opts Opts) Operation {
	if len(opts.Args) == 0 {
		return Print
	}

	if opts.Args[0] == "add" {
		return Add
	}

	if opts.Args[0] == "rem" {
		return Remove
	}
	return Print
}

func getArgs(opts Opts) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}
	operation := getOperation(opts)
	if operation == Add {
		if len(opts.Args) != 3 {
			return nil, fmt.Errorf("add requires 2 arguments but received %v", len(opts.Args)-1)
		}
		return opts.Args[1:], nil
	}

	if operation == Remove {
		if len(opts.Args) != 2 {
			return nil, fmt.Errorf("remove requires 1 argument but received %v", len(opts.Args)-1)
		}
		return opts.Args[1:], nil
	}

	if len(opts.Args) > 1 {
		return nil, fmt.Errorf("print requires 0 or 1 arguments but received %v", len(opts.Args)-1)
	}
	return opts.Args, nil
}

func getConfig(opts Opts) (string, error) {
	BASE_CONFIG_PATH := "/Users/tobias/Documents/Projects/courses/ts-go-rust/go"
	if opts.Config != "" {
		return opts.Pwd, nil
	}
	path := filepath.Join(BASE_CONFIG_PATH, "conf.json")
	return path, nil
}

func NewConfig(opts Opts) (*Config, error) {
	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}
	config, err := getConfig(opts)
	if err != nil {
		return nil, err
	}
	operation := getOperation(opts)

	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Pwd:       pwd,
		Config:    config,
		Args:      args,
		Operation: operation,
	}, nil
}
