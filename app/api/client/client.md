### 1. "用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `-`
- Response: `UserInfoRes`

2. request definition



3. response definition



```golang
type UserInfoRes struct {
	UserId int64 `json:"userId"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}
```

### 2. "店铺信息"

1. route definition

- Url: /v1/client/store/info
- Method: POST
- Request: `StoreInfoReq`
- Response: `StoreInfoRes`

2. request definition



```golang
type StoreInfoReq struct {
	StoreId int64 `json:"storeId,omitempty"`
}
```


3. response definition



```golang
type StoreInfoRes struct {
	StoreId int64 `json:"storeId"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Contacts int64 `json:"contacts"`
}
```

### 3. "店铺列表"

1. route definition

- Url: /v1/client/store/list
- Method: POST
- Request: `-`
- Response: `StoreListRes`

2. request definition



3. response definition



```golang
type StoreListRes struct {
	Total int64 `json:"total"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Current int64 `json:"current"`
	Rows interface{} `json:"rows"`
}
```

### 4. "店铺会员列表"

1. route definition

- Url: /v1/client/store/user/list
- Method: POST
- Request: `StoreUsersReq`
- Response: `StoreUsersRes`

2. request definition



```golang
type StoreUsersReq struct {
	StoreId int64 `json:"storeId,omitempty"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
```


3. response definition



```golang
type StoreUsersRes struct {
	Total int64 `json:"total"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Current int64 `json:"current"`
	Rows interface{} `json:"rows"`
}
```

### 5. "登录"

1. route definition

- Url: /v1/client/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginRes`

2. request definition



```golang
type LoginReq struct {
	Mobile string `json:"mobile,omitempty"`
	Password string `json:"password,omitempty"`
}
```


3. response definition



```golang
type LoginRes struct {
	UserId int64 `json:"userId"`
	Token string `json:"token"`
}
```

### 6. "注册"

1. route definition

- Url: /v1/client/user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterRes`

2. request definition



```golang
type RegisterReq struct {
	Mobile string `json:"mobile,omitempty"`
	Name string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
```


3. response definition



```golang
type RegisterRes struct {
	UserId int64 `json:"userId"`
	Token string `json:"token"`
}
```

