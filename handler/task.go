package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	task "task/proto"
)

type Task struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Task) Call(ctx context.Context, req *task.Request, rsp *task.Response) error {
	log.Info("Received Task.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Task) Stream(ctx context.Context, req *task.StreamingRequest, stream task.Task_StreamStream) error {
	log.Infof("Received Task.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&task.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Task) PingPong(ctx context.Context, stream task.Task_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&task.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
