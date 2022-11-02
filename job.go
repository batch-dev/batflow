package batflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type Job struct {
	Name  string
	Steps []Step
}

func RunJob(ctx workflow.Context, job Job) error {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 60 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	so := &workflow.SessionOptions{
		CreationTimeout:  time.Hour,
		ExecutionTimeout: time.Hour,
	}
	ctx, err := workflow.CreateSession(ctx, so)
	if err != nil {
		return err
	}
	defer workflow.CompleteSession(ctx)

	for _, step := range job.Steps {
		if err := workflow.ExecuteActivity(ctx, RunStep, step).Get(ctx, nil); err != nil {
			logger.Error("run step", "error", err, "name", step.Name)
			return err
		}
	}

	return nil
}
