package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"wecomGPT/openai"
	"wecomGPT/wecom"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// REQUEST_PARSING_FAILED
var ErrRequestParseFailed = errors.New("参数解析失败")

func Handle(msg string) *string {
	requestText := strings.TrimSpace(msg)
	reply, err := openai.Completions(requestText)
	if err != nil {
		log.Error(err)
	}
	return reply
}

func QuestionHandler(ctx *gin.Context) {
	var request QuestionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithError(400, ErrRequestParseFailed)
		return
	}
	fmt.Printf("%#v\n", request)
	// var reply string = "Hello"
	reply := Handle(request.Content)
	// 发送企微消息
	err = wecom.SendMsg(ctx, request.Content, *reply)
	if err != nil {
		log.Errorf("QuestionHandler.wecom.SendMsg err: %v", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"reply": reply,
	})
}
