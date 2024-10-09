package R

import (
	"github.com/gin-gonic/gin"
	"gvadmin_v3/core/global/E"
	"net/http"
)

// 通用响应信息
type CommonResp struct {
	Code int         `json:"code"`           //响应编码: 200 成功 500 错误 403 无操作权限 401 鉴权失败  -1  失败
	Msg  string      `json:"msg"`            //消息内容
	Data interface{} `json:"data,omitempty"` //数据内容
}

// 通用列表查询结果表单
type searchListResult struct {
	Total int64 `json:"total"`
	Rows  any   `json:"rows"`
}

type ApiResp struct {
	c *gin.Context
	r *CommonResp
}

// SuccessResp 返回一个成功的消息体
func SuccessResp(c *gin.Context) *ApiResp {
	msg := CommonResp{
		Code: E.SUCCESS,
		Msg:  "操作成功",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// ErrorResp 返回一个错误的消息体
func ErrorResp(c *gin.Context) *ApiResp {
	msg := CommonResp{
		Code: E.ERROR,
		Msg:  "操作失败",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// ForbiddenResp 返回一个拒绝访问的消息体
func ForbiddenResp(c *gin.Context) *ApiResp {
	msg := CommonResp{
		Code: E.FORBIDDEN,
		Msg:  "无操作权限",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// UnauthorizedResp JWT认证失败
func UnauthorizedResp(c *gin.Context) *ApiResp {
	msg := CommonResp{
		Code: E.UNAUTHORIZED,
		Msg:  "鉴权失败",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// SetMsg 设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.r.Msg = msg
	return resp
}

// SetCode 设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.r.Code = code
	return resp
}

// SetData 设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.r.Data = data
	return resp
}

// SetPageData 设置消息体的数据
func (resp *ApiResp) SetPageData(total int64, rows interface{}) *ApiResp {
	resp.r.Data = searchListResult{
		Total: total,
		Rows:  rows,
	}
	return resp
}

// WriteJsonExit 输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.c.Set("result", resp.r)
	resp.c.JSON(http.StatusOK, resp.r)
	resp.c.Abort()
}

// WriteErrJsonExit 输出json到客户端
func (resp *ApiResp) WriteErrJsonExit(errCode int) {
	resp.c.Set("result", resp.r)
	resp.c.JSON(errCode, resp.r)
	resp.c.Abort()
}
