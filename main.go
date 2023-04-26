package main

import (
	"net/http"
	"wecomGPT/config"
	"wecomGPT/handler"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const ADDR = "0.0.0.0:5000"

func main() {
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Info("程序启动")
	err := config.LoadConfig()
	if err != nil {
		log.Warn("没有找到配置文件，尝试读取环境变量")
	}
	log.Infof("加载配置文件: %#v", config.GetConfig())

	engine := gin.Default()
	err = registryHandler(engine)
	if err != nil {
		log.Errorf("registryHandler err: %v", err)
		return
	}

	if err := engine.Run(); err != nil {
		log.Errorf("engine.Run err: %v", err)
		return
	}
}

func registryHandler(serverMux *gin.Engine) error {
	serverMux.GET("/api/v1/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})
	serverMux.POST("/api/v1/chat", handler.QuestionHandler)
	return nil
}
