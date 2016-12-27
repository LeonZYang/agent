falcon-agent
===

This is a windows monitor agent.

## 


## Installation

It is a golang classic project

```
# set $GOPATH and $GOROOT
git clone https://github.com/LeonZYang/agent.git
cd agent
go build -o falcon-agent.exe
C:\workspace\src\github.com\LeonZYang\agent\falcon-agent.exe -c C:\workspace\src\github.com\LeonZYang\agent\cfg.json > C:\workspace\src\github.com\LeonZYang\agent\var\var.log 2>&1
```

## Configuration

- heartbeat: heartbeat server rpc address
- transfer: transfer rpc address
- ignore: the metrics should ignore


## 说明
- 数据采集用的是gopsutil（修改了部分代码），具体请查看看https://github.com/shirou/gopsutil
- agent不提供插件、http和缓存功能
- 目前监控项只有alive,cpu,disk,mem,net,load
- 编译环境需要go-1.5
  go 1.5有一个bug，修改src/net/interface_windows.go,大概在第50行，修改完重新编译go即可
  修改 return iia[:iilen-1], nil 为 return iia[:iilen], nil
- 网卡名称不支持中文(上报Transfer后数据无法显示，具体原因没细查，好像可以修改graph表字符集为utf-8)

