# noansible

一个模仿ansible的 golang 练习项目

## 简易的ansible功能

- ssh
- 批量异步多host
- 任务嵌套(INCLUDE)
- 异步任务
- yaml格式配置
- 变量模板`{{ avars }}` 渲染为 `vars`中对应的变量


## 实现内置功能模块
- `shell`: 发送shell命令到远程执行;`支持变量模板`
- `include`: 包含子任务列表文件
- `vars`：用于变量模板 `支持变量模板`
- `username`
- `hosts`：指定要执行的inventory的`groupname`

- `plugin`:调用模块用以下格式
  ```yaml
  - name: aa
    plugin:
      mod: modname
      args: arg1,arg2,arg3... #统一逗号分隔,空格会被自动删除，支持变量模板
  ```

- `plugin`/`file`:可以传文件
  * args: @/adir/a,/dst/s 等同 cp -rf /adir/a /dst/s
  * args: /adir/afile,/dst/dst.file 等同scp单个文件
  * args: ../fieldir/* ,/somepath/dst 等同 把fieldir下的所有文件copy到dst目录下
  
- `plugin`/`synctime`: 同步本地时间到远端
- `plugin`/`mod`:指定功能模块
- `plugin`/`args`:功能模块的参数，统一逗号分隔，`支持变量模板`
- `plugin`/`async`:异步执行任务

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
- 只有vars, shell 和  args支持 `支持变量模板`

```yaml
#inventory 格式
<groupname>:
  - <IPV4/IPV6 ADDR>,<PASSWORD>,[PORT:default=22]
  - <IPV4/IPV6 ADDR>,<PASSWORD>,[PORT:default=22]
```


## 编译

```shell
cd main
go build main.go
```

## TODO

- json支持
- telnet支持
- 从远端抓取文件
- 记录异步任务的执行结果（当前默认，异步任务成功）