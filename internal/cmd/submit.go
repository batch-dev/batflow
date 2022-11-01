package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/batch-dev/batflow"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v3"
)

func getSubmitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit workflow",
		RunE: func(cmd *cobra.Command, args []string) error {
			return submit(args)
		},
	}
	return cmd
}

func submit(args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	var wf batflow.Workflow
	if err := yaml.NewDecoder(f).Decode(&wf); err != nil {
		return err
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "exec",
		TaskQueue: batflow.BatflowTaskQueue,
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, batflow.RunWorkflow, wf)
	if err != nil {
		return err
	}
	fmt.Printf("start workflow %s: %s\n", we.GetID(), we.GetRunID())
	return nil
}
