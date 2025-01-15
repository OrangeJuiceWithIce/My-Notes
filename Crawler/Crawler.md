# Crawler

---

## 分类
##### 1.通用网络爬虫
会爬取互联网上的所有内容
##### 2.聚焦网络爬虫
会对内容进行筛选
##### 3.增量式网络爬虫
只抓取新增或变化的网页部分

---

## 网络请求
### Request
#### 1.请求行
Method:请求方式  
URL
#### 2.请求头
Host:请求来源
Content-Type
#### 3.请求体
请求参数
### Response
#### 1.状态行
状态码：表示本次请求是否成功(404)
#### 2.响应头
Server:处理请求的服务器软件信息
Content-Type
#### 3.响应体

---

## url的编码解码
1.parse.quote()：编码  
2.parse.unquote()：解码
```python
from urllib import parse

q_wd='XXX'
q_ASCII=parse.quote(q_wd)
```
3.网页的编码方式：response.encoding属性和response.apparent_encoding属性  
**diff:**
前者是http响应头声明的编码方式，后者是网页内容解析出的编码方式

---

## Requests库
请求方法  
.request():构造请求对象  
.get():获取网页  
.head()：获取网页头信息  
.post():提交POST请求  
.put()：提交PUT请求  
.patch()：提交局部修改请求  
.delete():提交删除请求  

---

## 反爬
### User-Agent 用户代理
1.请求头中包含User-Agent  
2.User-Agent包含了请求方的操作系统、CPU类型、浏览器版本等信息(适应)  
3.在使用request发送的请求头中，User-Agent标明为python-requests  
**自定义请求头**
往请求方法中传参
```python
import requests

url="XXX"
header={}
response=requests.get(url,headers=header)
```
4.所有的网站的UA是一样的  
5.可以通过构建代理池(list)，随机作为header,降低相同UA的访问频率，降低网站封杀的可能  
6.fake_useragent库可以随机生成UA  
```python
import fake_useragent

ua=fake_useragent.UserAgent()  //实例化对象

print(ua.chrome)
print(ua.ie)
print(ua.firefox)
print(ua.safari)
```
### referer
引荐网页的url.若不符合预期，会被拦截，建议在header中加入referer

---

## 文本数据分析
### 1.re正则表达式(python内置模块)
##### 普通字符
包括字母、数字、下划线
##### 元字符
样例string
```
tfirst,文明六偶尔好玩
f4
48
```
|元字符|English|含义|搜索例子|匹配结果|
|---|---|---|---|---|
|.||除换行符外的字符
|\w|word|普通字符|
|\s|space|空白字符|
|\d|digit|数字|
|\b|word boundary|单词边界|t\b|tfirs**t**|
|\n|newline|换行符|
|^||行首|^f.|f4|
|$||行尾|.4$|f4|
|\W||非普通字符|
|\S||非空白字符|
|\D||非数字字符|
|a\|b||a或b|t(first\|second)|tfirst|
|[abc]||a、b、c中的任意一个字符|t[fgh]irst|tfirst|
|[^abc]||a、b、c以外的任意一个字符|t[^g]irst|tfrist|
##### 量词
|量词|含义|
|---|---|
|{n}|重复n次|
|{n,}|>=n次|
|{n,m}|[n,m]次|
|+|>=1次|
|*|>=0次|
|?|0或1次|
##### 贪婪匹配和惰性匹配
在量词后面加?为惰性匹配：匹配尽可能短的字符串，数量较多
反之为贪婪匹配
#### re模块
**re.findall(正则表达式,待匹配字符串)**：返回一个list，但列表不适合存放大量数据  
**re.finditer(正则表达式，待匹配字符串)**:返回迭代器，每个匹配项都是一个Match对象，包括span和内容，可以用group()方法获取内容
**re.search(正则表达式,待匹配字符串)**:返回第一个匹配项，如果没有匹配项，返回None  
**re.match(正则表达式,待匹配字符串)**:有^开头限制，如果匹配失败，返回None  
**re.compile(正则表达式,flags值)**:将正则表达式编译成Pattern对象，可以多次调用
|flags值|含义|
|---|---|
|0|无扩展|
|re.A|元字符只匹配ASCII码|
|re.I|忽略大小写|
|re.S|元字符.可以匹配换行符|
|re.M|元字符^和$可以匹配每行的开头和结尾|
```python
import re
re_obj=re.compile("[a-z]+")

talk_list_1=re_obj.findall(待匹配字符串)
```
#### re爬取
用括号圈出自己想要提取的部分  
PS:使用finditer()时，若想提取多个部分，需要在括号内开头用?P<组名>定义组名，然后在后续的操作中使用group('组名')提取
```python
//list版本
import re
import requests
import fake_useragent

url="https://movie.douban.com/top250"
ua=fake_useragent.UserAgent()
header={"User-Agent":ua.random}

response=requests.get(url,headers=header)
print(response.text)
re_obj=re.compile('<span class="title">(.*?)</span>.*?<span class="title">&nbsp;/&nbsp;.*?</span>',flags=re.S)
html=response.text
re_list=re_obj.findall(html)
print(re_list)
response.close()
```
```python
//迭代器版本
re_obj=re.compile('<span class="title">(?P<组名1>.*?)</span>.*?<span class="title">&nbsp;/&nbsp;.*?</span>',flags=re.S)
...
film_list=re_obj.finditer(html)
...
print(film_list.group('组名1'))
```

---

### 2.bs4(python第三方库)
用于从HTML或XML文档中提取指定数据
|解析器|使用方法|
|---|---|
|Python标准库|BeautifulSoup(html,"html.parser")|
|lxml HTML|BeautifulSoup(html,"lxml")|
|lxml XML|BeautifulSoup(html,"xml")|
|html5lib|BeautifulSoup(html,"html5lib")|

BeautifulSoup对象.findall(标签名，筛选条件)  
筛选结果.get("属性名")，返回属性值
```python
from bs4 import BeautifulSoup(html,"lxml")
...
re_obj=BeautifulSoup(html,"html.parser")
attr={"class":"xxx"}
results=re_obj.findall('a',attrs=attr)

for i in results:
    print(i.get("title"))
    print(i.get("href"))
```

--- 

### 3.Xpath(lxml库对其提供了良好的支持)
将标签看作节点树
|表达式|作用|
|---|---|
|nodename|选取所有名为nodename的节点|
|/|从根节点中选取|
|currentnode//或//|从当前节点的所有后代中选取|
|.|选取当前节点|
|..|选取当前节点的父节点|
|@|选取属性|
|*|通配符|
|/text()|获取节点的文本|
```python
from lxml import etree
...
page=etree.HTML(html)  //返回Element对象
results=page.xpath('//*[@id="content"]/p[3]')  //返回匹配结果
for i in results:
    i.strip().replace('\n','')  //去除文本两端的空白，以及换行符
    ...
```
> XPath表达式可以在开发者模式中，打开检查元素功能，点击目标元素，右键复制Xpath  

>路径中的tbody标签要删去

---

## 存储
### .txt
```python
with open("文件路径",'工作模式',encoding='编码方式') as f:
    f.writelines('xxx')
```
### .csv
```python
import pandas

 columnName=['书名','作者']
 data=[]
 for book,author in zip(book_list,authors):  //zip()可以将两个列表合并为一个元组列表，做到同时遍历
    data.append([book,author])
 df=pandas.DataFrame(data,columns=columnName)

 df.to_csv("books.csv",index=True,encoding="utf-8")
```
### Pymysql
""":三引号sql语句可以换行而不是\n，且可以包含引号""
```python
import pymysql

class CreateSQL():
    def __init__(self,host,port,user,pw,sql_name):
        self.db=pymysql.connnect(
            host=host,
            port=port,
            user=user,
            password=pw,
            charset='utf8mb4'
        )
        self.cursor=self.db.cursor()
        self.sql_name=sql_name

        sql="""create database if not exists %s 
        default charset utf8 
        default collate utf8_general_ci;"""%sql_name  //创建新数据库的sql语句
        self.cursor.execute(sql)

    def save_data(self,data,table):
        self.cursor.execute('use %s;'%self.sql_name)  //切换到新数据库
        sql_table='''create table if not exists %s(
        name varchar(50),
        author varchar(50)
        );'''%table  //创建新表的sql语句
        self.cursor.execute(sql_table)

        sql='insert into '+table+'(name,author) values(%s %s);'
        self.cursor.executemany(sql,data)  //插入数据
        self.db.commit()  //提交事务

    def sel_data(self,table)
        sql_select="select * from %s;"%table
        self.cursor.execute(sql_select)
        a=self.cursor.fetchall()
        for i in a:
            print(i)

    def close_(self):
        self.cursor.close()
        slef.db.close()
```

---

## 模拟登录
### 1.Cookie插入Header
### 2.Cookie传入requests.get()
---

## Some Practice
### 1.爬取完整页面
```python
from urllib import response
import urllib.request

url="..."
response=urllib.request.urlopen(url)
result=response.read().decode("gbk")  //html的head的charset属性规定了编码方式
```
### 2.爬取图片
```python
import requests
import fake_useragent
import os  //处理文件的模块

url="..."
ua=fake_useragent.UserAgent()
header={"User-Agent":ua.random}

response=requests.get(url,headers=header)

os.makedirs("images",exist_ok=True)  //创建存储图片文件夹
imgs_data=response.json()['data']

for i,data in enumerate(img_data):
    if 'hoverURL' in data:
        img_url=data['hoverURL']
        img_data=requests.get(img_url,headers=header).content
        img_path="images/"+str(i)+'.jpg'
        with open(img_path,'wb') as fp:
            fp.write(img_data)
            print("one image saved")
```
### 3.自动翻页
循环修改url中的**pn参数**即可  
url中的键值对可以以字典形式存储，通过get()方法的params参数传入  
可以用ctrl+R(pycharm环境中)打开正则替换，使params符合字典格式key带''并,分割  
response.close()关闭连接  
```python
import requests
import os
import fake_useragent
import json

url="https://image.baidu.com/search/acjson"

ua=fake_useragent.UserAgent()
header={"User-Agent":ua.random}

os.makedirs("images",exist_ok=True)
for j in range(10):
    param={
'tn': 'resultjson_com',
'logid': '9564341554286606630',
'ipn': 'rj',
'ct': '201326592',
'is': '',
'fp': 'result',
'fr': '',
'word': 'AI美女',
'queryWord': 'AI美女',
'cl': '2',
'lm': '-1',
'ie': 'utf-8',
'oe': 'utf-8',
'adpicid': '',
'st': '-1',
'z': '',
'ic': '',
'hd': '',
'latest': '',
'copyright': '',
's': '',
'se': '',
'tab': '',
'width': '',
'height': '',
'face': '0',
'istype': '2',
'qc': '',
'nc': '1',
'expermode': '',
'nojc': '',
'isAsync': '',
'pn':str(30*j),
'rn': '30',
'gsm': '3c',
'1736824796163': '',
    }
    response = requests.get(url, headers=header,params=param)
    imgs_data = json.loads(response.text, strict=False)['data']
    print(imgs_data)

    for i,data in enumerate(imgs_data):
        if 'thumbURL' in data:
            img_url=data['thumbURL']
            img_data=requests.get(img_url,headers=header).content
            img_path="images/"+str(30*j+i)+'.jpg'
            with open(img_path,'wb') as fp:
                fp.write(img_data)
                print("one image saved")
```