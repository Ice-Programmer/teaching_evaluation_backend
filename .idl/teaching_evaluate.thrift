namespace go teaching_evaluate
include 'base.thrift'

struct PingRequest {
	255: required base.Base Base
}

struct PingResponse {
	1:   required string        message
	255: required base.BaseResp BaseResp
}

service TeachingEvaluateService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping", api.serializer="json")
} (agw.preserve_base="true", agw.js_conv="str")