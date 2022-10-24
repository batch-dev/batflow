package cmd

import (
	"context"
	"fmt"

	"github.com/batch-dev/batflow"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
)

func getExecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "Exec command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return exec(args)
		},
	}
	return cmd
}

func exec(args []string) error {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "exec",
		TaskQueue: batflow.BatflowTaskQueue,
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, batflow.ExecWorkflow, batflow.ExecInput{
		Name: args[0],
		Args: args[1:],
	})
	if err != nil {
		return err
	}
	var output batflow.ExecOutput
	err = we.Get(context.Background(), &output)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", output)
	return nil
}
