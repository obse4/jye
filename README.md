# jye

**自用工具包**
> golang小白，计划摘录和自封装业务中常用的工具。不建议直接使用。

## pkg/taskpool
- 非阻塞协程池
- 原地址[github.com/q191201771/naza](https://github.com/q191201771/naza)

## util/strval
- 任意格式转string类型
- 会有额外的性能开销，不建议频繁使用

## util/jyetimer
- string格式日期转unix时间戳

## util/jyehttp
- http请求常用封装