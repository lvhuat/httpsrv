package httpsrv

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvhuat/code"
)

// overloadChecker 抑制过载
type overloadChecker struct {
	curTime  int64
	mutex    sync.Mutex
	curCount int32
	LimitCnt int32
	Service  string
}

// Check 检查是否过载
func (overloadChecker *overloadChecker) Check(ctx *gin.Context) code.Error {
	if overloadChecker.LimitCnt <= 0 {
		return nil
	}

	timeNow := time.Now().Unix()
	overloadChecker.mutex.Lock()
	defer overloadChecker.mutex.Unlock()

	if timeNow > overloadChecker.curTime {
		overloadChecker.curTime = timeNow
		overloadChecker.curCount = 1
		return nil
	}

	if overloadChecker.curCount >= overloadChecker.LimitCnt {
		return code.NewMcodef("OVERLOAD_DENIED", "Check Snow Protect failed,service = %v", overloadChecker.Service)
	}

	overloadChecker.curCount++
	return nil
}

// OverloadChecker 过载检测
type OverloadChecker interface {
	Check(*gin.Context) code.Error
}

// NewDefaultOverloadChecker 创建一个内置默认
func NewDefaultOverloadChecker(cnt int32, service string) OverloadChecker {
	return &overloadChecker{
		LimitCnt: cnt,
		Service:  service,
	}
}
