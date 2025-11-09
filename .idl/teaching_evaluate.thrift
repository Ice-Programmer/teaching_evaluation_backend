namespace go teaching_evaluate
include 'base.thrift'

struct PingRequest {
	255: required base.Base Base
}

struct PingResponse {
	1:   required string        message
	255: required base.BaseResp BaseResp
}

// ======================================= Student Class ======================================= //
struct CreateClassRequest {
	1:   required string    classNumber
	255: required base.Base Base
}

struct CreateClassResponse {
	1:            i64           id
	255: required base.BaseResp BaseResp
}


// ======================================= User ======================================= //
enum UserRole {
	Student = 1
	Admin   = 2
}

struct UserInfo {
	1:  i64      id
	2:  string   name
	3:  UserRole role
	4:  i64      createAt
}

struct UserLoginRequest {
	1:            string    userAccount
	2:            string    userPassword
	255: required base.Base Base
}

struct UserLoginResponse {
	1:            UserInfo      userInfo
	2:            string        token
	3:            i64           expireAt
	255: required base.BaseResp BaseResp
}

service TeachingEvaluateService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping", api.serializer="json")

    // ======================================= Student Class ======================================= //
    CreateClassResponse CreateClass(1: CreateClassRequest req) (api.post="/api/v1/itmo/teaching/class/create", api.serializer="json")

    // ======================================= User  ======================================= //
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/api/v1/itmo/teaching/user/login", api.serializer="json")

} (agw.preserve_base="true", agw.js_conv="str")