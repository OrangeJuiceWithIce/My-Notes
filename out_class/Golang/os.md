# os
```go
envs := os.Environ()   //返回[]string
path := os.Getenv("PATH")   //返回指定环境变量的值
dir=os.Getwd()   //返回当前工作目录

os.Chdir("/tmp")    //改变当前工作目录
os.Mkdir("./test", 0755)   //创建目录,第二个参数是权限
os,MkdirAll("./test/test2", 0755)   //创建多级目录
os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)   //打开文件,第一个参数是文件名,第二个参数是打开模式,分别是不存在就创建,读写模式,追加模式,第三个参数是权限
```