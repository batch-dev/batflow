package cmd

import (
	"github.com/batch-dev/batflow"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func getRunnerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner",
		Short: "Serve Batflow Runner",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runner()
		},
	}
	return cmd
}

func runner() error {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()

	w := worker.New(c, batflow.BatflowTaskQueue, worker.Options{
		EnableSessionWorker: true,
	})
	w.RegisterWorkflow(batflow.RunWorkflow)
	w.RegisterWorkflow(batflow.RunJob)
	w.RegisterActivity(batflow.RunStep)
	err = w.Run(worker.InterruptCh())
	return err
}
