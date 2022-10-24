package batflow

const BatflowTaskQueue = "BATFLOW_TASK_QUEUE"

type ExecInput struct {
	Name string
	Args []string
}

type ExecOutput struct {
	Stdout   string
	Stderr   string
	ExitCode int
}
