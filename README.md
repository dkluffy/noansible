# noansible

一个模仿ansible的 golang 练习项目

## 基本了简易的ansible功能

- ssh
- 批量异步多host
- 任务嵌套(INCLUDE)
- 异步任务
- yaml格式配置

## 实现内置功能模块

- `plugin`:调用模块用以下格式
  ```yaml
  - name: aa
    plugin:
      mod: modname
      args: arg1,arg2,arg3... #统一逗号分隔,空格会被自动删除
  ```

- `file`:可以传文件
  * args: @/adir/a,/dst/s 等同 cp -rf /adir/a /dst/s
  * args: /adir/afile,/dst/dst.file 等同scp单个文件
  * args: ../fieldir/* ,/somepath/dst 等同 把fieldir下的所有文件copy到dst目录下
- `synctime`: 同步本地时间到远端
- `shell`: 发送shell命令到远程执行
- `include`: 包含子任务列表文件


## 使用说明

- 运行

```shell
#./noansible -h
Noansible @version= 1.0
Usage :
  -bs int
        SCP buffer size (default 1024)
  -i string
        Inventory file dir (default "inventory.yml")
  -log string
        Log file dir (default "output.log")
  -p string
        Inventory file dir (default "main.yml")

```

- include的文件要在同目录下，include的文件中除了`tasks`其他字段被忽略



## 编译

```shell
cd main
go build main.go
```

## TODO

- json支持
- telnet支持