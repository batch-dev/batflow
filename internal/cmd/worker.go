package cmd

import (
	"github.com/batch-dev/batflow"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func getWorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker",
		Short: "Serve Batflow Worker",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server()
		},
	}
	return cmd
}

func server() error {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()

	w := worker.New(c, batflow.BatflowTaskQueue, worker.Options{})
	w.RegisterWorkflow(batflow.ExecWorkflow)
	w.RegisterActivity(batflow.Exec)
	err = w.Run(worker.InterruptCh())
	return err
}
