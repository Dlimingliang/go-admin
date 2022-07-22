package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/services"
	"github.com/Dlimingliang/go-admin/utils"
)

var operationRecordService = services.ServiceGroupApp.OperationRecordService

func OperationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body []byte
		var userId int
		//获取请求参数
		if ctx.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				global.GaLog.Error("read body from request error:", zap.Error(err))
			} else {
				ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := ctx.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		//获取请求人
		claims, _ := utils.GetClaims(ctx)
		if claims.Id != 0 {
			userId = claims.Id
		}
		//构建记录
		record := model.OperationRecord{
			Ip:     ctx.ClientIP(),
			Method: ctx.Request.Method,
			Path:   ctx.Request.URL.Path,
			Agent:  ctx.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = writer
		now := time.Now()
		ctx.Next()
		latency := time.Since(now)
		record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = ctx.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()
		if _, err := operationRecordService.CreateOperationRecord(record); err != nil {
			global.GaLog.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
