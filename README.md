
## 简单使用方法
```golang
    package main
    
    import "github.com/xmkuban/logger"
    
    func main() {
    
    	logger.Error("test")
    }
```

如果需要更换日志实现可以调用  logger.SetLogger()

目前提供beego，cleanlog，seelog的实现,切换已实现的如下
```golang
logger.InitBeegoLogByConsole(5)
```