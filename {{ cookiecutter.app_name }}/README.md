# 嘉为蓝鲸监控 exporter 开发指南

1. 环境准备
    - 拉取标准化 exporter，修改配置
    - 修改 common/define.go 中 NameSpace 和 Version
2. 开始开发
    - 自定义 main 函数中的命令行参数
    - 自定义 config/config.go 中的配置结构体
3. 日志使用
    - 日志实现轮转，24 小时轮转或大小达到 100M 轮转，日志保留 10 天
    - 通过 logger.GetStdLogger()使用
    - 日志分为 error，info，debug 三个级别

# 业务开发

业务逻辑主要在于 collector 目录下

- exporter.go 提供 exporter 初始化，prometheus 的接口等。
- collectXxx 为具体的抓取数据函数
- scrapefn.go 中为抓取数据函数的集合。

主要逻辑为： 当 prometheus 访问 exporter 接口时，触发 collector 中 scrape 函数，开始采集数据并将数据发送到通道中。

## 采集逻辑

采集逻辑视对象而定

## 使用说明

### 插件功能

采集器去访问系统的对外api，查询元数据信息。

### 版本支持：

linux实测：centos7

windows实测：Windows Server 2012R2

**组件支持版本：**

理论上支持：组件XXX 21, 19, 18, 12, or 11.2

实测支持：组件XXX 11g

**是否支持远程采集:**

是

### 使用指引

登陆到系统上面进行相关操作

**注：xxx指标指标需要配置后才可以采集到；**

**注：xxx相关指标需要配置后才可以采集到；**




### 参数说明


| 参数名                | 含义                | 是否必填 | 使用举例                          |
|--------------------|-------------------|------|-------------------------------|
| web.listen-address | exporter监听id及端口地址 | 是    | http://127.0.0.1/9601/metrics |
| user               | 用户                | 是    | user                          |
| password           | 密码                | 是    | 123456                        |


*参数为空时，应当填上""(英文双引号)*

*用户名或密码中含有特殊字符时，应当加上""(英文双引号)*


### 指标列表

| 指标名称                         | 指标描述                                  | 单位   | 维度  |
|------------------------------|---------------------------------------|------|-----|
| std_exporter_up              | 插件运行状态     （1 == active ， 0 == error） | none |     |
| std_exporter_status          | 组件状态     （1==OPEN，else ==0）           | none |     |
| std_exporter_run_health_time | 组件运行时长                                | s    |     |

### 版本日志
#### std_exporter 1.0.0
