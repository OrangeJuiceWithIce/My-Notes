# Shell

## 语法

### 以#!/bin/bash开头
#!声明脚本由什么shell解释，否则使用默认shell

### chmod +x xxx.sh
给予脚本执行权限

### 变量
|behavior|example|
|---|---|
|定义|很寻常num=10|
|引用|$变量名|
|清除变量|unset 变量名|
|从键盘读取值到变量|read -p "我是提示词" 变量名1 变量名2|
|声明只读变量|readonly num=10|

### 条件判断
```
if [ command ];then
   符合该条件执行的语句
elif [ command ];then
   符合该条件执行的语句
else
   符合该条件执行的语句
fi
```

```
[ num1 -eq num2 ]      num1 和 num2 两数相等为真 , =
[ num1 -ne num2 ]      num1 和 num2 两数不等为真 ,!=
[ num1 -gt num2 ]      num1 大于 num1 为真 , >
[ num1 -ge num2 ]      num1 大于等于num2 为真, >=
[ num1 -lt num2 ]      num1 小于n um2 为真 , <
[ num1 -le num2 ]      num1 小于等于 num2 为真, <=
```