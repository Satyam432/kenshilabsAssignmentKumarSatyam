package models

//Here the model is defined for task

type Task struct {
	TaskId     string `json:"taskId"`
	CreatedBY  string `json:"createdBy"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
}

type CreateTask struct {
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
}
