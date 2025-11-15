package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/handler/class"
	"teaching_evaluate_backend/handler/ping"
	"teaching_evaluate_backend/handler/user"
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

// UserLogin implements the TeachingEvaluateServiceImpl interface.
func (s *TeachingEvaluateServiceImpl) UserLogin(ctx context.Context, req *teaching_evaluate.UserLoginRequest) (resp *teaching_evaluate.UserLoginResponse, err error) {
	resp, err = user.UserLogin(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "UserLogin err: %v", err)
		resp = &teaching_evaluate.UserLoginResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// CreateClass implements the TeachingEvaluateServiceImpl interface.
func (s *TeachingEvaluateServiceImpl) CreateClass(ctx context.Context, req *teaching_evaluate.CreateClassRequest) (resp *teaching_evaluate.CreateClassResponse, err error) {
	resp, err = class.CreateClass(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "create-class failed: %v", err)
		resp = &teaching_evaluate.CreateClassResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// EditClass implements the TeachingEvaluateServiceImpl interface.
func (s *TeachingEvaluateServiceImpl) EditClass(ctx context.Context, req *teaching_evaluate.EditClassRequest) (resp *teaching_evaluate.EditClassResponse, err error) {
	resp, err = class.EditClass(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "edit-class failed: %v", err)
		resp = &teaching_evaluate.EditClassResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// QueryClass implements the TeachingEvaluateServiceImpl interface.
func (s *TeachingEvaluateServiceImpl) QueryClass(ctx context.Context, req *teaching_evaluate.QueryClassRequest) (resp *teaching_evaluate.QueryClassResponse, err error) {
	resp, err = class.QueryClass(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "query-class failed: %v", err)
		resp = &teaching_evaluate.QueryClassResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}
