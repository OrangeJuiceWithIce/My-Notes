# Docker
## 概念
### 容器Container
1.容器有各自的操作系统  
2.通过Docker引擎共用主机的内核
### 镜像Image
1.提供容器运行时所需的程序、库、配置、依赖等文件
2.不包含动态数据  
> **Image-Container≈类-实例**  
### 定制镜像文件Dockerfile
1.制作镜像的"菜谱"(操作系统、应用版本等)  
2.可以在其它镜像基础上进行定制
### 工作流
Dockerfile(docker build)——>Image(docker run)——>Container
## Node.js的Dockerfile示例
1.Dockerfile:逐步执行，缓存机制
```
FROM node:18-alpine3.15  //指定基础镜像
WORKDIR /MyApp  //设置工作目录
COPY package.json .  //复制package.json到镜像当前目录中
Run npm install  //根据json文件安装依赖

COPY . .  //复制当前目录所有文件到镜像当前目录中

EXPOSE 3000  //声明要暴露的端口，没有实际作用
CMD ["node", "app.js"]  //启动命令
```
2.dockerignore:指明忽略的部分
```
node_modules  //安装依赖时会自动生成
Dockerfile
dockerignore
```
>docker build( -t my-node-app) .  //构建镜像,括号代表命名

>docker images  //查看本地镜像

>docker tag 镜像ID(可以只取前几位) 镜像名：版本号  //给镜像打标签

>docker rmi 镜像名字  //删除镜像

>docker run -d -v 主机目录:容器目录:ro --p -v 不同步的容器目录 主机端口:容器端口 --name 容器名 镜像名字
//-d代表后台运行，防止占用当前cmd窗口;
-v代表文件同步，一方变动另一方跟随,ro使主机目录只读;
-p代表端口映射,访问主机端口时会自动映射到容器端口;
--name指定名称，不加则随机分配

>docker stop 容器名  //停止容器

>docker rm -fv 容器名  //删除容器和volumn

>docker ps  //查看运行中的容器
>docker ps -a  //查看所有容器

3.docker-compose.yml:定义多个容器的配置
>docker -compose up -d --build  //后台执行本地docker-compose.yml文件,--build表示镜像若有改动则重新构建
>dokcer -compose down -v  //删除所有容器和volumn

```
version:'3.8'  //版本
services:
  node-app:
    build: .  //指定Dockerfile所在目录
    ports:
      - "3000:3000"  //端口映射
    volumes:
      - ./:/MyApp:ro  //同步目录，主机只读不修改
      - /app/node_modules  //不同步目录
```