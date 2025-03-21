# etcd
存储少量的、重要的信息
## 应用场景
### 服务注册与发现
1.server向注册中心注册服务  
2.client向注册中心订阅服务时，获取列表缓存本地  
3.server在下线服务前，通知注册中心  
4.注册中心通知client，更新本地缓存  
5.server完成剩余工作，正式关闭连接

## 架构
1.write ahead log(WAL)  
先写日志(顺序追加)，再执行写操作  
2.raft共识算法
—leader选举:节点有三种状态leader、condidate、follower。如果follower长时间没有收到leader的心跳，则转换为candidate，竞选leader
—日志复制:leader将日志新entry复制到其他节点，收到大多数节点的确认后才会提交entry,并通知其他节点提交entry
3.每个key-value都是B+树，存储历史版本
## 基本指令
```
.\etcdctl.exe put key value  //存储
.\etcdctl.exe get key        //获取
.\etcdctl.exe watch key      //监听
.\etcdctl.exe del key        //删除

.\etcdctl.exe get key -w json  //获取版本信息
.\etcdctl.exe get key --rev=xx  //获取指定历史版本值

.\etcdctl.exe txn -i  //事务操作
$compares:
条件(value("key") = "value"或mod("key") = "xx"或create("key") = "xx")
$success requests(get, put, delete):
多行操作
$failure requests(get, put, delete):
多行操作
```
|版本名称|作用域|含义|
|---|---|---|
|revision|集群|任意key被修改都会自增|
|create_revision|特定key|创建时的版本号|
|mode_revision|特定key|key被修改自增|
|version|特定key|key被修改次数|
## 租约lease
**ttl**(time to live):租约的有效期，租约到期后，etcd会自动删除所有绑定该租约的key-value
```
.\etcdctl.exe lease grant 10  //创建一个租约，有效期10秒
$返回lease ID
.\etcdctl.exe put key value --lease=leaseID  //绑定租约到key-value
```
## 搭建集群
使用goreman工具
1.创建一个参数文件夹Procfile.learner
```
etcd1: etcd --name infra1 --listen-client-urls http://127.0.0.1:12379 --advertise-client-urls http://127.0.0.1:12379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd2: etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd3: etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
```
|代码部分|含义|
|---|---|
|--name|节点名称|
|--listen-client-urls|监听客户端请求的URL,一般为2379|
|--advertise-client-urls|建议客户端请求的URL|
|--listen-peer-urls|监听集群内节点间通信的URL,一般为2380|
|--initial-cluster|接口地址列表 节点1=xxx,节点2=xxx,节点3=xxx...|
|--initial-cluster-state|new代表新建的,existing代表已存在的|
2.通过goreman -f ./Procfile.learner start启动集群
```
etcdctl --endpoints=某一个listen-client-url member list  //查看集群成员
goreman run stop etcd1  //关闭某个节点
goreman run restart etcd1  //重启某个节点

etcdctl put key value --endpoints=xxx  //从某一节点写入数据
etcdctl get key --endpoints=xxx  //从某一节点读取数据,所有节点都可以读到
```