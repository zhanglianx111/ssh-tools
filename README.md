# ssh-tools
**ssh-tools**将主机信息存储在`sqlite3`数据库中，`login`通过存储在数据库中到主机信息登陆主机。

## 使用方法
```shell
ssh登陆主机工具

Usage:
  use: [command]

Available Commands:
  add         将主机信息添加到数据库中
  completion  Generate the autocompletion script for the specified shell
  delete      删除数据库中的主机信息
  help        Help about any command
  init        初始化数据库
  login       登陆主机
  update      更新数据库中到主机信息
  version     打印版本信息

Flags:
  -h, --help   help for use:

Use "use: [command] --help" for more information about a command.

```



## 编译
```shell 
$ cd ssh-tools
$ go build -ldflags "-w -s -X 'github.com/zhanglianx111/ssh-tools/cmd.Author=username' 
    -X 'github.com/zhanglianx111/ssh-tools/cmd.GoVersion=`go version`' 
    -X 'github.com/zhanglianx111/ssh-tools/cmd.CommitID=`git log --pretty=format:%h -1`'"
```
