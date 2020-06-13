package core

import (
	"bytes"
)

type ExecResult struct {
	stdOut, stdErr bytes.Buffer
}
