namespace go teaching_evaluate
include "base.thrift"

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

struct EditClassRequest {
	1:   required i64       id
	2:            string    classNumber
	255: required base.Base Base
}

struct EditClassResponse {
	255: required base.BaseResp BaseResp
}

struct QueryClassCondition {
	1: optional i64       id
	2: optional list<i64> ids
	3: optional string    searchText
}

struct QueryClassRequest {
	1:   optional QueryClassCondition condition
	2:            i32                 pageSize
	3:            i32                 pageNum
	255: required base.Base           Base
}

struct ClassInfo {
	1:  i64    id
	2:  string classNumber
	3:  i64    createTime
	4:  i64    createOpId
	5:  string createOpName
	6:  i64    updateTime
	7:  i64    updateOpId
	8:  string updateOpName
}

struct QueryClassResponse {
	1:            list<ClassInfo> classInfoList
	2:            i64             total
	255: required base.BaseResp   BaseResp
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
    CreateClassResponse CreateClass(1: CreateClassRequest req) (api.post="/api/v1/itmo/teaching/admin/class/create", api.serializer="json")
    EditClassResponse EditClass(1: EditClassRequest req) (api.post="/api/v1/itmo/teaching/admin/class/edit", api.serializer="json")
    QueryClassResponse QueryClass(1: QueryClassRequest req) (api.post="/api/v1/itmo/teaching/admin/class/query", api.serializer="json")

    // ======================================= User  ======================================= //
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/api/v1/itmo/teaching/evaluation/user/login", api.serializer="json")

} (agw.preserve_base="true", agw.js_conv="str")