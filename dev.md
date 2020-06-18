# noansible

## 功能

- 读取YAML，inventory也用YAML
- [ x ] SSH to exec cmd
- scp 上传文件 https://www.jianshu.com/p/f9d6dfefb63d
- ~~类似ansible，可以添加模块，模块可以独立编译(plugin模块)~~ windows还不支持
- 只能通过exec调用或者rpc来实现了（暂时不做）
- 只能在args字段里做文章
- 异步任务
- 并发多台机器（不一定要）