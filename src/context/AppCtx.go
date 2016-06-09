package context

import (
	"src/context/AppCtx"
	"src/context/tool"
)

/*
*上下文 包括 redis的初始化
*数据库的链接
*是否是调试模式
*日志
 */
type AppContext struct {
	Redis     *redis.Pool
	debugMode bool
	logger    tool.ILogger
}
