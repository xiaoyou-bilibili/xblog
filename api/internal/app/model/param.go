// Package model @Description 请求参数
// @Author 小游
// @Date 2021/01/19
package model

import "xBlog/internal/db"

/*
validate 常用参数
validate:"required" 表示这个值必填
validate:"omitempty" 如果数据为空就忽略
validate:"len=0" 指定长度
validate:"eq=0" 必须为某个值
validate:"gt=0" 大于
validate:"gte=0" 大于等于
validate:"lt=0" 小于
validate:"lte=0" 小于等于
validate:"min=1" 最小值
validate:"max=2" 最大值
validate:"oneof=5 7 9" 必须为其中的一个
validate:"alphanum" 字符串值是只包含 ASCII 字母数字字符
validate:"alphaunicode" 字符串值是否只包含 unicode 字符
validate:"alphanumunicode" 字符串值是否只包含 unicode 字母数字字符
validate:"lowercase" 是否只包含小写字母
validate:"uppercase" 只包含大写字母
validate:"email" 为电子邮件
validate:"url" 是否为合法的url
validate:"contains=@" 字符串包含@
validate:"ip" 是否为有效的ip地址
validate:"datetime" 是否为时间
*/

// Comment 评论信息参数
type Comment struct {
	Email   string `json:"email" validate:"email"`      // 邮件地址
	Content string `json:"content" validate:"required"` // 评论内容
	Name    string `json:"name" validate:"required"`    // 评论者昵称
	UserID  int    `json:"user_id"`                     // 用户id
	Parent  int    `json:"parent"`                      // 评论父节点
	Site    string `json:"site"`                        // 评论者网站
	Avatar  string `json:"avatar"`                      // 评论者头像
	Hang    string `json:"hang"`                        // 头像挂件
	Level   string `json:"level"`                       // 用户等级
	Uid     string `json:"uid"`                         // 评论者uid
}

// PostStatus 文章状态参数
type PostStatus struct {
	Good       bool `json:"good"`       // 文章是否点赞
	Collection bool `json:"collection"` // 文章是否收藏
}

// WechatComment 微信小程序评论
type WechatComment struct {
	OpenID  string `json:"openid" validate:"required"`
	Content string `json:"content" validate:"required"` // 评论内容
	Name    string `json:"name" validate:"required"`    // 评论者昵称
	Parent  int    `json:"parent"`                      // 评论父节点
	Avatar  string `json:"avatar" validate:"required"`  // 评论者头像
}

// Advice 意见反馈
type Advice struct {
	Concat  string `json:"concat"`                      // 联系方式
	Content string `json:"content" validate:"required"` // 反馈内容
}

// AddUserParam 用户注册
type AddUserParam struct {
	Username string `json:"username" validate:"required"` // 用户名
	Nickname string `json:"nickname" validate:"required"` // 昵称
	Password string `json:"password" validate:"required"` // 密码
	Email    string `json:"email" validate:"required"`    // 邮箱
}

// UserPutUserParam 修改用户个人信息
type UserPutUserParam struct {
	Avatar       string `json:"avatar"`       // 头像
	Hang         string `json:"hang"`         // 头像挂件
	Nickname     string `json:"nickname"`     // 昵称
	Email        string `json:"email"`        // 邮箱
	Sign         string `json:"sign"`         // 签名
	Subscription string `json:"subscription"` // 是否订阅
	OldPassword  string `json:"old_password"` // 旧密码
	NewPassword  string `json:"new_password"` // 新密码
}

// UserLoginParam 用户登录
type UserLoginParam struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}

// UserResetPasswordEmail 用户发送重置密码的邮件
type UserResetPasswordEmail struct {
	Email  string `json:"email"`  // 邮箱
	Option string `json:"option"` // 操作
}

// UserResetPasswordToken 用户通过token来重置密码
type UserResetPasswordToken struct {
	UserID   string `json:"user_id" form:"id" validate:"required"`        // 用户id
	Token    string `json:"token" form:"token" validate:"required"`       // 用户token
	Password string `json:"password" form:"password" validate:"required"` // 用户密码
}

// UserGetUserParam 判断用户名是否存在
type UserGetUserParam struct {
	User  string `json:"user"`  // 用户名
	Email string `json:"email"` // 邮箱
}

// UserPostCodeParam 手机端获取验证码参数
type UserPostCodeParam struct {
	Email  string `json:"email" validate:"required"`  // 邮箱地址
	Option string `json:"option" validate:"required"` // 执行的操作
}

// UserPostAddParam 手机端用户注册
type UserPostAddParam struct {
	UserName string `json:"username" validate:"required"` // 用户名
	Nickname string `json:"nickname" validate:"required"` // 昵称
	Password string `json:"password" validate:"required"` // 密码
	Email    string `json:"email" validate:"required"`    // 邮箱
	Code     string `json:"code" validate:"required"`     // 验证码
}

// UserPutAppPasswordParam 手机端用户重置密码
type UserPutAppPasswordParam struct {
	Email    string `json:"email" validate:"required"`    // 邮箱地址
	Password string `json:"password" validate:"required"` // 新密码
	Code     string `json:"code" validate:"required"`     // 验证码
}

// AdminUpdatePost 管理员更新文章
type AdminUpdatePost struct {
	Title    string   `json:"title"`    // 文章标题
	Html     string   `json:"html"`     // HTML 内容
	Status   string   `json:"status"`   // 文章状态
	Password string   `json:"password"` // 文章密码
	Md       string   `json:"md"`       // markdown数据
	IsTop    string   `json:"is_top"`   // 文章是否置顶(因为要判断参数是否存在，所以需要使用string类型)
	IsDraft  string   `json:"is_draft"` // 是否为草稿
	Delete   string   `json:"delete"`   // 是否删除这个文章
	Category []int    `json:"category"` // 文章分类
	Tags     []string `json:"tags"`     // 文章标签
	Parent   string   `json:"parent"`   // 文档父节点
}

// AdminAddPost 管理员添加文章
type AdminAddPost struct {
	Title    string   `json:"title" validate:"required"`  // 文章标题
	Html     string   `json:"html" validate:"required"`   // 文章内容
	Status   string   `json:"status" validate:"required"` // 文章状态
	Password string   `json:"password"`                   // 文章密码
	Md       string   `json:"md"`                         // markdown数据
	IsTop    string   `json:"is_top"`                     // 文章是否置顶(因为要判断参数是否存在，所以需要使用string类型)
	IsDraft  string   `json:"is_draft"`                   // 是否为草稿
	Delete   string   `json:"delete"`                     // 是否删除这个文章
	Category []int    `json:"category"`                   // 文章分类
	Tags     []string `json:"tags"`                       // 文章标签
}

// AdminAddDiary 管理员添加日记
type AdminAddDiary struct {
	Html     string `json:"html" validate:"required"`   // 文章内容
	Status   string `json:"status" validate:"required"` // 文章状态
	Password string `json:"password"`                   // 文章密码
	Md       string `json:"md"`                         // markdown数据
	IsDraft  bool   `json:"is_draft"`                   // 是否为草稿
}

// AdminAddDoc 管理员添加文档
type AdminAddDoc struct {
	Title  string `json:"title" validate:"required"` // 文档标题
	Parent int    `json:"parent"`                    // 父节点
}

// AdminAddCategoryParm 管理员添加分类
type AdminAddCategoryParm struct {
	Name   string `json:"name" validate:"required"` // 标签名
	Parent int    `json:"parent"`                   // 父节点
	Type   string `json:"type" validate:"required"` // 类型 category 分类 tag 标签
}

// AdminUpdateCategoryParm 管理员更新分类
type AdminUpdateCategoryParm struct {
	Name   string `json:"name"`   // 标签名
	Parent string `json:"parent"` //父节点,为string是为了避免为0导致误判
	Type   string `json:"type"`   // 类型 category 分类 tag 标签
}

// AdminUpdateComments 管理员更新评论
type AdminUpdateComments struct {
	Agree string `json:"agree"` // 是否同意 0 拒绝 1 同意 (使用string是为了避免误判)
}

// AdminUpdateUser 管理员更新用户数据
type AdminUpdateUser struct {
	Username     string `json:"username"`     // 用户名
	Nickname     string `json:"nickname"`     // 昵称
	Email        string `json:"email"`        // 用户邮箱
	Status       string `json:"status"`       // 用户状态
	Subscription string `json:"subscription"` // 用户是否订阅
	Identity     string `json:"identity"`     // 用户权限
}

// AdminUpdateOption 管理员更新站点设置
type AdminUpdateOption struct {
	Key   string `json:"key"`   // 设置关键词
	Value string `json:"value"` // 设置的值（这个为字符串，布尔类型的也是字符串true或者false）
	Type  string `json:"type"`  // 设置的类型（string int bool) 三种类型
}

// AdminPutSideParam 管理员侧边栏设置
type AdminPutSideParam struct {
	Left  []db.SettingAdminSideDetail `json:"left"`  // 左边插件信息
	Right []db.SettingAdminSideDetail `json:"right"` // 右边插件信息
}
