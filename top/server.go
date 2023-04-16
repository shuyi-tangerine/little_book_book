package top

import "context"

type Server interface {
	// IsBlock 同步启动是否会阻塞
	IsBlock(ctx context.Context) (isBlock bool)
	// Start 开始提供服务
	Start(ctx context.Context) (err error)
	// AsyncStart 异步启动
	AsyncStart(ctx context.Context)
}
