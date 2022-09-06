package server

import "strconv"

type LANG string

const (
	ZH LANG = `zh-CN`
	EN LANG = `en-US`
)

// Code 状态码别名, 减少本地
//
//go:generate stringer -type=Code
type Code int32

// 状态码
const (
	SUCCESS                        Code = 0    //0: 成功
	ErrUserPassword                Code = 1109 //@field 1109: 账号或密码有误,请重试
	ErrSession                     Code = 8001 //8001: 无效会话或会话已过期
	ErrDefListSearchNotSupport     Code = 8101 //8101: 该资源不支持通用查找
	ErrDefListSearchAbnormalField  Code = 8102 //8102: 资源过滤字段异常
	ErrDefListSearchTempleNotFound Code = 8103 //8103: 资源列表查询模版未找到
	ErrDefFirstXNotSupport         Code = 8104 //8104: 该资源不支持通用查看
	ErrModNotFound                 Code = 8105 //8105: 模块不存在
	ErrPatchSettings               Code = 8106 //8106: 修改模块配置失败
	ErrDeleteInnerMod              Code = 8107 //8107: 只能删除外置模块
	ErrCleanRoutes                 Code = 8108 //8008: 清理API路由表异常
	ErrCleanMenus                  Code = 8109 //8009: 清理菜单表异常
	ErrDelModule                   Code = 8110 //8010: 删除模块信息错误
	ErrRouteNotFound               Code = 8111 //8111: 请求接口不存在
	ErrInnerRouteNotFound          Code = 8112 //8112: 请求接口不存在
	ErrArgsURLScheme               Code = 8113 //8113: 解析URL异常找不到协议
	ErrArgsRangeType               Code = 8114 //8114: 区间类型非数字
	ErrArgsRangeCMDTYPE            Code = 8115 //8115: 非法区间比较符
	ErrArgsRangeNum                Code = 8116 //8116: 区间参数缺失
	ErrArgsRange                   Code = 8117 //8117: 区间范围异常
	ErrItemXNotFound               Code = 8118 //8118: 资源未找到
	ErrInnerServer                 Code = 8119 //8119: 内部服务异常
	ErrAssetValueNoConfig          Code = 8120 //8120: 资产服务无配置
	ErrCodeNotFound                Code = 8121 //8121: 命令码不存在
	ErrLoadSettings                Code = 8122 //8122: 加载配置失败
	ErrNoSettingsFound             Code = 8123 //8123: 没有相应的配置
	ErrParseSettings               Code = 8124 //8124: 解析相应配置异常
	ErrOpenRpcClient               Code = 8125 //8125: 创建RPC客户端失败
	ErrNoEtcdConfig                Code = 8126 //8126: ETCD配置为空
	ErrDuplicateKey                Code = 8127 //8127: 记录值重复
	ErrSessionTimeout              Code = 8128 //8128: 会话过期



	CacheKeyUserID = `/user/users/user_id/`

	CacheKeyUserOauthList = `/user/oauth_users/user_id/`

	CacheKeyBaseUserID = `/user/base_users/user_id/`

	CommonUA = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36`
)

// Resp 标准化返回
type Resp struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Code) String() string {
	return strconv.Itoa(int(r))
}

// Locale 多语言化返回
func (r Code) Tr(lang LANG) *Resp {
	//TODO
	return &Resp{Code: r, Message: `format`}
}

// With 设置Resp data
func (r *Resp) With(data interface{}) *Resp {
	r.Data = data
	return r
}
