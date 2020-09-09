# 短网址服务 GoDwz

GoDwz 使用 go语言编写，router部分使用gin框架，并使用lru算法缓存热点数据，具有极致性能

## Contents

## 安装

1. 第一步 下载代码

2. 拷贝配置文件并修改

```cp app.ini.bak  app.ini
   根据实际情况编辑配置文件
```

3. 启动程序,注意端口是不是已经被占用
nohup ./dwz &
## 快速开始

## API 例子
以下用127.0.0.1代替域名或者ip,默认端口是80,n是new的简写,url=后面原始长网址

* 举例 你可以访问 http://127.0.0.1/n?url=http://www.baidu.com
获得返回值 类似 {"dwz":"http://127.0.0.1/2RBIc3"} dwz代表返回值 dwz是短网址的拼音缩写
* 访问http://127.0.0.1/2RBIc3 就会301 跳转到 http://www.baidu.com

```
感谢 gin框架,httpRouter

