package handler

import (
	"context"
	task "task/proto"
)

type Task struct{}

func (t *Task) RecordTask(context context.Context, req *task.RecordRequst, rsp *task.RecordResponse) error {
	return nil
}
