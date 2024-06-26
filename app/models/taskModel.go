package models

type Task struct {
	TaskId    string `json:"userId"`
	CreatedBY string `json:"createdBy"`
	TaskName  string `json:"taskName"`
	Username  string `json:"username"`
}
