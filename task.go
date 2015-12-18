package main

import (
	"io"
	"os/exec"
)

type TaskConfig struct {
	command  string
	argArray []string
}

type TaskResult struct {
	state  int
	stdout io.Reader
	stderr io.Reader
}

func TaskRun(taskConfig *TaskConfig) (*TaskResult, error) {
	cmd := exec.Command(taskConfig.command, taskConfig.argArray...)
	taskResult := new(TaskResult)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	taskResult.stdout = stdout
	taskResult.stderr = stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return taskResult, nil
}
