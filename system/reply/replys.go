package reply

const (
	Success = iota
	Error
	DataNoExist
	ParamError
	LoginDisable
	LoginError
	AccessNoPrivilege
	ParamIncorrect
	InvalidToken
	RegisterAdmin
	AccountExist
	AccountNotExist
	PasswordError
	SyncPostsError
	UpdatePasswordError
	FileTypeError
	ImportFileError
	DatabaseSqlParseError
)

var Errors = map[int]string{
	Success:               "操作成功",
	Error:                 "操作失败",
	DataNoExist:           "该数据不存在",
	ParamError:            "参数错误",
	LoginDisable:          "账户已被禁用",
	LoginError:            "登录失败，用户名或密码错误",
	AccessNoPrivilege:     "不具备访问权限",
	ParamIncorrect:        "传入参数有误",
	InvalidToken:          "token解析失败",
	RegisterAdmin:         "注册失败",
	AccountExist:          "账号已存在",
	AccountNotExist:       "用户不存在",
	PasswordError:         "密码错误",
	SyncPostsError:        "同步文章失败",
	UpdatePasswordError:   "密码修改失败",
	FileTypeError:         "文件类型错误",
	ImportFileError:       "文件导入失败",
	DatabaseSqlParseError: "数据库解析异常",
}
