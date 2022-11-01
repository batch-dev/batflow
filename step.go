package batflow

import (
	"context"
	"os"
	"os/exec"

	"go.temporal.io/sdk/activity"
)

type Step struct {
	Name string
	Run  string
}

func RunStep(ctx context.Context, step Step) error {
	logger := activity.GetLogger(ctx)

	cmd := exec.CommandContext(ctx, "sh", "-c", step.Run)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Info("run step", "error", err, "name", step.Name)
		return err
	}

	return nil
}
