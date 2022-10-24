package batflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ExecWorkflow(ctx workflow.Context, input ExecInput) (ExecOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var output ExecOutput
	err := workflow.ExecuteActivity(ctx, Exec, input).Get(ctx, &output)
	return output, err
}
