# install guide 安装指南





## 1. 安装目录与权限
一般来说, AAA 应用可以安装在任意目录下, 该目录下需要对运行 aaa 的 linux 系统用户有写入当前目录的权限

aaa 将创建  ./log/ 子目录并写入运行日志 
配置文件需要与可执行文件在同一目录, 权限设置为当前角色只读即可



## 2. 配置文件

配置文件名为 fast-config.json
文件格式为 json , 请务必保障 json 格式正确
```
{
  "Name": "Aaa Server",
  "Version": "0.1.2",
  "Debug": true,
  "UmsConfig": {
    "ActiveAuthUri": "http://50.7.101.250/rpc/active",
    "RegisterAuthUri": "http://50.7.101.250/rpc/auth",
    "PlayAuthUri": "http://127.0.0.1"
  },
  "AaaConfig": {
    "ServerPort": ":8095",
    "EpgGslb": [
      " http://50.7.101.250"
    ],
    "VodGslb": [
      "192.168.2.1"
    ],
    "LiveGslb": [
      "23.237.64.186:43000"
    ]
  }
}
```

其中 UmsConfig 配置是到 ums 用户管理子系统的访问地址
AaaConfig 配置的是 当前 AAA 服务器的运行参数与配置


## 3. 优化处理

taskset 指定 aaa 运行核心
