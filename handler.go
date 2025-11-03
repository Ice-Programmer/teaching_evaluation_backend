package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"teaching_evaluate_backend/handler/ping"
	teaching_evaluate "teaching_evaluate_backend/kitex_gen/teaching_evaluate"
)

// TeachingEvaluateServiceImpl implements the last service interface defined in the IDL.
type TeachingEvaluateServiceImpl struct{}

// Ping implements the TeachingEvaluateServiceImpl interface.
func (s *TeachingEvaluateServiceImpl) Ping(ctx context.Context, req *teaching_evaluate.PingRequest) (resp *teaching_evaluate.PingResponse, err error) {
	resp, err = ping.Pong(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "ping-pong failed: %v", err)
		return nil, err
	}
	return resp, nil
}
