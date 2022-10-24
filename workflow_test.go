package batflow

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	input := ExecInput{Name: "hostname"}
	ouput := ExecOutput{Stdout: "node01\n"}
	env.OnActivity(Exec, mock.Anything, input).Return(ouput, nil)
	env.ExecuteWorkflow(ExecWorkflow, input)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result ExecOutput
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, ouput, result)
}
