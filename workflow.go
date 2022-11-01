package batflow

import (
	"go.temporal.io/sdk/workflow"
)

type Workflow struct {
	Name string
	Jobs map[string]Job
}

func RunWorkflow(ctx workflow.Context, wf Workflow) error {
	logger := workflow.GetLogger(ctx)

	for _, job := range wf.Jobs {
		if err := workflow.ExecuteChildWorkflow(ctx, RunJob, job).Get(ctx, nil); err != nil {
			logger.Info("run job", "error", err, "name", job.Name)
		}
	}

	return nil
}
