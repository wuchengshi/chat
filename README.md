# 安装方法
本系统升级到golang1.12,请开启如下支持
```
#开启go mod支持
export GO111MODULE=on
#使用代理
export GOPROXY=https://goproxy.io

```

## 声明
> 课程参考 https://coding.imooc.com/class/339.html，本程序借鉴课程中思想，实现PC端的即时通讯  
> 
> 有问题可联系本人，微信 wucs_dd

## 1.项目配置，非常重要

### 1.1 数据库配置
修改config/config.ini 中数据库配置文件，数据库支持 MySQL 及 SQL Server，改为你自己的数据库以及密码
```ini
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName = chat
```

### 1.2 端口配置
```ini
HttpPort = :8080
```

### 1.3 建表语句
位于 model/chat.sql

+ `admin_user`：用户表，系统自动加载表 `admin_user` 里的用户，作为默认可添加好友
+ `community`：群信息表
+ `contact`：用户关系表
+ `record`：聊天记录表

### 1.4 页面入口地址
```
http://127.0.0.1:8080/chat.html
```

## 2.依赖包安装

使用go mod 自动处理安装包

# 演示截图

1. 添加好友
![添加好友](https://files.mdnice.com/user/16240/8042ea29-7cb0-41a7-86c3-6fe43b09ecc2.png)

2. 单聊
![单聊](https://files.mdnice.com/user/16240/511127e1-47e0-4b84-924c-7e39bfcde8b7.png)

3. 群聊
![群聊](https://files.mdnice.com/user/16240/7885af41-d133-43dc-b012-f91042d67315.png)

