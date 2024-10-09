package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gvadmin_v3/app/system/model"
	"gvadmin_v3/core/config"
	"gvadmin_v3/core/db"
	"gvadmin_v3/core/global/E"
	"gvadmin_v3/core/global/R"
	"gvadmin_v3/core/log"
	"gvadmin_v3/core/queue"
	"gvadmin_v3/core/util"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func LogTo() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 获取请求参数
		var param string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				log.Instance().Error("copy request body Failed." + err.Error())
				err = nil
			}
			rb, _ := ioutil.ReadAll(bf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			param = string(rb)
		}

		// 继续执行
		c.Next()

		// 结束时间
		endTime := time.Now()

		// OPTION不记录
		if c.Request.Method == http.MethodOptions {
			return
		}

		// 构造日志
		var tmp model.SysOperLog
		userId, _ := c.Get("userId")
		uid := util.AnyToInt(userId)
		userName, _ := c.Get("userName")
		uname, _ := util.AnyToStr(userName)
		deptId, _ := c.Get("deptId")
		did := util.AnyToInt(deptId)
		//user := service.FindUserById(uid)
		//dept := service.GetDeptInfoById(strconv.Itoa(user.DeptId))

		tmp.UserId = uid
		//tmp.Title = title
		//tmp.BusinessType = strconv.Itoa(businessType)
		tmp.BusinessName = "" //暂不写
		tmp.RequestMethod = c.Request.Method
		tmp.OperatorType = "0" //访问渠道 0-默认 1-PC 2-APP
		tmp.OperName = uname
		tmp.DeptId = did
		//tmp.DeptName = dept.DeptName

		tmp.OperUrl = c.Request.URL.Path
		tmp.OperIp = util.GetClientIp(c.Request)
		tmp.OperLocation = "" //暂不写
		tmp.OperParam = param
		tmp.OperTime = time.Now()
		tmp.LatencyTime = strconv.FormatInt(endTime.Sub(startTime).Milliseconds(), 10) + "ms"
		tmp.UserAgent = c.Request.Header.Get("User-Agent")

		// 处理返回结果
		outBody, _ := c.Get("result")
		respBody := outBody.(*R.CommonResp)
		respData, err := util.AnyToStr(respBody.Data)
		if err != nil {
			log.Instance().Error("Parse Data Failed..." + err.Error())
		}
		if len(respData) > 128 {
			tmp.JsonResult = respData[0:127] //暂不写完,仅写255
		} else {
			tmp.JsonResult = respData
		}

		if respBody.Code == 500 {
			tmp.Status = "1"
		} else {
			tmp.Status = "0"
		}

		tmp.Msg = respBody.Msg

		//写日志
		//switch config.Instance().ZapLog.SaveMode {
		//case "file":
		//	logToFile(tmp)
		//case "db":
		//	logToDB(tmp)
		//case "both":
		//	logToFile(tmp)
		//	logToDB(tmp)
		//default:
		//	logToFile(tmp)
		//}

		//写入队列
		msg, err := json.Marshal(tmp)
		if err != nil {
			log.Instance().Error("Marshal OperLog Failed..." + err.Error())
		}
		err = queue.Instance().Publish(E.TopicOperLog, string(msg))
		if err != nil {
			log.Instance().Error("Push Msg Failed..." + err.Error())
		}
	}
}

// 操作日志入库
func logToDB(operLog model.SysOperLog) {
	if err := db.Instance().Model(model.SysOperLog{}).Create(&operLog).Error; err != nil {
		log.Instance().Error("Insert OperLog Failed..." + err.Error())
	}
}

// 操作日志入文件
func logToFile(operLog model.SysOperLog) {
	log.Instance().Info("OperLog Success...", zap.Any("operLog", operLog))
}

func WriteTo(tmp string) {
	var sysOperLog model.SysOperLog
	err := json.Unmarshal([]byte(tmp), &sysOperLog)
	if err != nil {
		log.Instance().Error("UnMarshal OperLog Failed..." + err.Error())
	}
	switch config.Instance().ZapLog.SaveMode {
	case "file":
		logToFile(sysOperLog)
	case "db":
		logToDB(sysOperLog)
	case "both":
		logToFile(sysOperLog)
		logToDB(sysOperLog)
	default:
		logToFile(sysOperLog)
	}
}
