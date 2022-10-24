package batflow

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
)

func Exec(ctx context.Context, input ExecInput) (ExecOutput, error) {
	cmd := exec.CommandContext(ctx, input.Name, input.Args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		var exiterr *exec.ExitError
		if errors.As(err, &exiterr) {
			return ExecOutput{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: exiterr.ExitCode()}, err
		}
		return ExecOutput{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: 255}, err
	}
	return ExecOutput{Stdout: stdout.String(), Stderr: stderr.String()}, nil
}
