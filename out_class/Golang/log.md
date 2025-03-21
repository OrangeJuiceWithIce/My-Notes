# 日志
---
## log包
```go
log.Print("Hello")  //普通打印
log.Printf("Hello,%s","alice")  //格式化打印
log.Println("Hello")    //打印换行

log.Fatalln("Hello")    //打印并退出程序

log.Panicln("Hello")    //打印日志及其详细信息并退出
```
### 设置
```go
log.SetPrefix("Test:")  //设置前缀
log.SetFLags(log.Ldate|log.Ltime|log.Lmicroseconds)  //设置日志格式,默认是Ldate|Ltime

f,err := os.OpenFile("PATH",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)  //打开日志文件
if err != nil{
    log.Fatalln(err)
}
log.SetOutput(f)  //设置输出位置
```