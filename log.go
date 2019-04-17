package hippo

import (
//    "flag"
    "time"

    "github.com/golang/glog"
    "github.com/szuecs/gin-glog"
    "github.com/gin-gonic/gin"
)


type LogT struct {
	Info func(args ...interface{})
	Warn func(args ...interface{})
	Error  func(args ...interface{})
}

var Log = LogT{
	Info: glog.Info,
	Warn: glog.Warning,
	Error: glog.Error,
}

func InitLoggging(router *gin.Engine, config Configuration) {
//	IsDevMode
//	flag.Parse()
	router.Use(ginglog.Logger(3 * time.Second))
}
