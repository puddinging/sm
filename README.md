# sm
## 一个用户管理用户服务器连接的小工具
## 配置存储路径
`~/.sm/config.json`
例如
```
{
    "serverList": [
        {
            "alias": "local",
            "ip": "192.168.0.1",
            "username": "jiefeng",
            "password": "welcome1"
        }
    ]
}
```
## 如何使用
1. `list` 命令，列出来所有的服务器连接信息
```
$ sm list              
list filePath /Users/jiefeng/.sm/config.json 
Alias: local      IP: 0.0.0.0  
--------------
Alias: local1     IP: 192.168.0.1    
--------------
```
2. `open` 命令，使用-a指定 list 命令列出的 `Alias` 即可。可在当前命令行打开 `ssh` 链接
```
sm open -a local
```
3. `add` 命令，在配置文件里边添加一个新的服务器连接信息
```
sm add -a local -i 192.168.0.1 -u jiefeng -p welcome1
```
## 注意
1. `Linux/macos` 需安装 `sshpass` 工具
2. `Windows` 需安装 `PuTTY/KiTTY` 工具
