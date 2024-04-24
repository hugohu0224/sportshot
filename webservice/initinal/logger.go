package initinal

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
	}
	zap.ReplaceGlobals(logger)
}
