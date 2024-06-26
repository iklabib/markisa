package model

import (
	"bytes"
)

// run stage model
type RunResult struct {
	Message  string         `json:"message"`
	Builds   []CompileError `json:"builds"`
	Tests    []TestResult   `json:"tests"`
	ExitCode int            `json:"exit_code"`
}

type Submission struct {
	Type    string `json:"type"`
	Src     string `json:"src"`
	SrcTest string `json:"src_test"`
}

type SandboxExecResult struct {
	Error  error
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}

type CompileError struct {
	Filename string
	Message  string
	Line     int
	Column   int
}

type TestResult struct {
	Status string // PASS or FAILED
	Name   string
	Output string
	Order  int
}
