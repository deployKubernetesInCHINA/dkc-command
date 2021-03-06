#### 说明

dkc项目旨帮助kubernetes小白在centos 7.9.2009上快速安装kubernetes v1.20.7

安装剧本使用kuberspray项目: version 2.16.0

所需要的k8s安装文件都通过download命令下载

#### 使用场景

1. 服务器系统为centos 7.9.2009 [下载地址](http://isoredirect.centos.org/centos/7/isos/x86_64/)

1. 服务器中任意节点作为dkc运行节点 
![](./images/dkc（master1）.png)

#### 功能设计

![](./images/dkc.png)

#### 使用方式


1. 在release页面下载dkc.tar.gz, 解压, 
2. 下载k8s相关的离线文件,
```shell script
./dkc download cache
```
3. 访问l[公网web地址](http://dkc.k8sfans.com:5555)编辑拓扑文件,下载hosts.yaml,并放在inventory目录

![](./images/page.png)

```shell script
#检查环境信息
./dkc prepare show
```

4. 安装kubernetes
```shell script
./dkc install k8s
```
5. 增加节点
```shell script
./dkc node add --hostname <hostname> --ip <ip>
```
6. 删除节点
```shell script
./dkc node del --hostname <hostname> 
```
7. 检查kubernetes
```shell script
./dkc status k8s
```

8. 删除kubernetes
```shell script
./dkc uninstall k8s
```