package ping

import (
	"context"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
)

func Pong(ctx context.Context, req *teaching_evaluate.PingRequest) (resp *teaching_evaluate.PingResponse, err error) {

	return &teaching_evaluate.PingResponse{
		Message:  "pong",
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
