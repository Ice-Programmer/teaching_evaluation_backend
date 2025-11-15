package utils

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Obj2JsonStr(ctx context.Context, obj interface{}) string {
	objJson, err := json.Marshal(obj)
	if err != nil {
		klog.CtxErrorf(ctx, "Obj2JsonStr err: %v", err)
	}
	return string(objJson)
}
