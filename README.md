# module2

Go module test package

```shell
# go1.11beta2默认在gopath里按旧规则运行，为达到测试目的，请将代码放到gopath外

# 第一次运行时，使用了错误的版本
jd@DESKTOP-D5AMTAJ:module2$ go1.11beta2 run main.go
go: finding github.com/donnol/module/a latest
go: finding github.com/donnol/module v0.1.1
go: downloading github.com/donnol/module v0.1.1
go: github.com/donnol/module@v0.1.1: parsing go.mod: unexpected module path "jdlau.com/module"
go: error loading module requirements

# 第二次运行时，使用了正确的版本
jd@DESKTOP-D5AMTAJ:module2$ go1.11beta2 run main.go
go: finding github.com/donnol/module/a latest
go: finding github.com/donnol/module v1.0.0
go: downloading github.com/donnol/module v1.0.0
hello, 1

# 再次直接运行
jd@DESKTOP-D5AMTAJ:module2$ go1.11beta2 run main.go
hello, 1

# 目前包缓存放在$GOPATH/src/mod目录里，结构如下：
# |- cache -- 缓存
#       |- download -- 里面保存了源码压缩文件和文件hash值
#       |- vcs -- 版本控制工具有关的文件
# |- github.com -- 源码
#       |- donnol
#           |- module@v0.1.1 -- 版本v0.1.1的源码
#           |- module@v1.0.0 -- 版本v1.0.0的源码
```
