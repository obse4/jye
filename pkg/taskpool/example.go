package taskpool

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 并发计算0+1+2+...+1000
// 演示怎么向协程池中添加带参数的函数任务
func ExampleNewPool() {
	start := time.Now()
	pool, _ := NewPool(func(option *Option) {
		// 限制最大并发数
		option.MaxWorkerNum = 1
	})

	var sum int32

	n := int(1e5)

	for i := 0; i < n; i++ {
		pool.Go(func(param ...interface{}) {
			ii := param[0].(int)
			atomic.AddInt32(&sum, int32(ii))
		}, i)
	}

	fmt.Println(sum)
	fmt.Println(time.Since(start).Nanoseconds())
}
