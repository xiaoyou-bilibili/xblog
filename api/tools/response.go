// Package tools @Description 返回结果相关
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

/**
http 状态码说明
200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
202 Accepted - [*]：表示一个请求已经进入后台排队（异步任务）
204 NO CONTENT - [DELETE]：用户删除数据成功。
400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。
500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。
*/

// Errors api错误处理(多错误)
type Errors struct {
	Messages []string `json:"messages"` // 错误内容
}

// Error api错误处理(单错误)
type Error struct {
	Message string `json:"message"` // 错误内容
}

type Response struct {
}

// GlobalResponse 全局的response对象
var GlobalResponse = Response{}

// 错误信息处理(返回单个或多个错误)
func errorProcess(data string, msg ...string) interface{} {
	if len(msg) == 1 {
		return Error{Message: msg[0]}
	} else if len(msg) > 1 {
		return Errors{Messages: msg}
	} else {
		return Error{Message: data}
	}
}

// ResponseOk 200返回数据成功
func (response *Response) ResponseOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// ResponseCreated 201创建数据成功
func (response *Response) ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// ResponseNoContent 204删除数据成功(删除数据一般不需要返回数据)
func (response *Response) ResponseNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, "")
}

// ResponseBadRequest 400错误，一般是用户输入的参数有问题
func (response *Response) ResponseBadRequest(c *gin.Context, msg ...string) {
	c.JSON(http.StatusBadRequest, errorProcess("参数错误", msg...))
}

// ResponseUnauthorized 401错误，用户没有权限
func (response *Response) ResponseUnauthorized(c *gin.Context, msg ...string) {
	c.JSON(http.StatusUnauthorized, errorProcess("权限不足,请登录", msg...))
}

// ResponseForbidden 403错误，用户禁止访问
func (response *Response) ResponseForbidden(c *gin.Context, msg ...string) {
	c.JSON(http.StatusForbidden, errorProcess("禁止访问", msg...))
}

// ResponseNotFound 404错误，没有这个资源
func (response *Response) ResponseNotFound(c *gin.Context, msg ...string) {
	c.JSON(http.StatusNotFound, errorProcess("内容不存在", msg...))
}

// ResponseUnProcessEntity 422错误，用户在创建对象的时候发生错误
func (response *Response) ResponseUnProcessEntity(c *gin.Context, msg ...string) {
	c.JSON(http.StatusNotFound, errorProcess("创建对象失败", msg...))
}

// ResponseServerError 500错误，服务器错误
func (response *Response) ResponseServerError(c *gin.Context, msg ...string) {
	c.JSON(http.StatusInternalServerError, errorProcess("服务器错误", msg...))
}

// ResponseByte 返回原始的字节数据
func (response *Response) ResponseByte(c *gin.Context, contentType string, data []byte) {
	c.Data(http.StatusOK, contentType, data)
}

// ResponseHtml 返回html对象
func (response *Response) ResponseHtml(c *gin.Context, data string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, data)
}

// 实例化验证对象
var validate = validator.New()

// ValidatorParam 校验参数是否正确(true表示参数非法)
func ValidatorParam(c *gin.Context, param interface{}) bool {
	// 先验证基本的数据接口是否匹配
	if c.Bind(param) != nil && c.BindQuery(param) != nil {
		GlobalResponse.ResponseBadRequest(c)
		return false
	} else if err := validate.Struct(param); err != nil { // 再验证格式是否正确
		GlobalResponse.ResponseBadRequest(c, "参数不符合格式")
		return true
	} else { // 验证通过后返回空
		return false
	}
}

// JudgeParams 判断用户的参数是否都输入了
func JudgeParams(params ...string) bool {
	//遍历参数
	for _, v := range params {
		if v == "" {
			return true
		}
	}
	return false
}
