package model

import "github.com/gin-gonic/gin"

type Response struct {
	Code    string
	Message string
	Data    interface{}
}

func (m *Response) ToH() map[string]interface{} {
	return gin.H{
		"code":    m.Code,
		"message": m.Message,
		"data":    m.Data,
	}
}
