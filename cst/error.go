package cst

const (
	UNKNOW_ERROR = "unknow_error" //未知错误

	JSON_ERROR = "json_error" //json错误

	VERIFY_ERROR = "verify_error" //参数校验错误

	DAO_ERROR = "dao_error" //数据库操作错误

	NUMBER_EXIST = "number_exist" //序列号已存在

	STRAIN_EXIST = "strain_exist" //品系名已存在

	NUMBER_NULL = "number_null" //空序列号

	USER_EXIST = "user_exist" //用户已存在

	PW_ERROR = "pw_error" //用户名或者密码错误

	TOKEN_EXPIRE = "token_expire" //token过期

	NEED_LOGIN = "needLogin" //需要登陆

	NO_TOKEN = "no_token" //没TOKEN

	NO_LOGIN = "no_login" //未登陆

	USER_NOT_EXIST = "user_not_exist" //用户已存在

)
