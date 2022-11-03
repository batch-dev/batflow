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

	childCtx, cancel := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)
	completedJobs := make(map[string]Job)
	jobsTotal := len(wf.Jobs)
	var childErr error

	for {
		if len(completedJobs) == jobsTotal {
			logger.Info("completed workflow", "name", wf.Name)
			break
		}
		if childCtx.Err() != nil {
			logger.Info("cancelled workflow", "error", childCtx.Err(), "name", wf.Name)
			break
		}
		if len(wf.Jobs) == 0 {
			// Wait a job to complete.
			selector.Select(ctx)
		}

	NextNewJob:
		for id, job := range wf.Jobs {
			// Check whether need jobs have been completed.
			if len(job.Needs) != 0 {
				for _, nid := range job.Needs {
					if _, ok := completedJobs[nid]; !ok {
						continue NextNewJob
					}
				}
			}

			future := workflow.ExecuteChildWorkflow(ctx, RunJob, job)
			selector.AddFuture(future, func(id string, job Job) func(workflow.Future) {
				return func(f workflow.Future) {
					if err := f.Get(childCtx, nil); err != nil {
						cancel()
						childErr = err
						return
					}
					completedJobs[id] = job
				}
			}(id, job))

			// Remove started job.
			delete(wf.Jobs, id)
		}

		// Wait a job to complete.
		selector.Select(ctx)
	}

	return childErr
}
