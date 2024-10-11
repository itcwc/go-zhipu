# 重要！官方社区SDK已发布

* [官方社区Go SDK安装地址：https://github.com/yankeguo/zhipu](https://github.com/yankeguo/zhipu)
* [其他SDK：https://maas.aminer.cn/dev/api/libraries](https://maas.aminer.cn/dev/api/libraries)

本人始终推荐您使用官方SDK！！！

# go-zhipu

* 新版内容已更新完毕，参数详情请看[官方文档](https://maas.aminer.cn/dev/api)，欢迎提交bug~
* 首次写第三方扩展，有代码问题还请见谅。

***

## 前言

* go-智谱清言（go-zhipu）是基于 go 语言开发的智谱清言 API 接口包，主要用于智谱清言模型的调用，支持通用模型、图像大模型、超拟人大模型、向量模型、Batch API、模型微调、知识管理等功能。
* 详情见官方文档：[智谱清言官方文档](https://maas.aminer.cn/dev/api)
  
***

## 支持模型

* 通用模型
  * GLM-4
    * sse调用
    * 异步调用
    * 任务结果查询
  * GLM-4V
    * sse调用
  * GLM-3-Turbo
    * sse调用
    * 异步调用
    * 任务结果查询
* 图像大模型
* 超拟人大模型
  * CharacterGLM
    * sse调用
    * 异步调用
    * 任务结果查询
  * Emohaa
    * sse调用
    * 异步调用
    * 任务结果查询
* 向量模型
  * 创建向量模型同步请求

* 模型微调
  * 创建微调任务
  * 查询微调任务事件
  * 查询微调任务
  * 查询个人微调任务
  * 删除微调任务
  * 取消微调任务
  * 删除微调模型
* 搜索工具
  * Web-Search-Pro
* Batch API
  * 创建 Batch
  * 取消 Batch
  * 列出 Batch
  * 下载 Batch 结果
* 知识管理
  * 知识库管理
    * 创建知识库
    * 编辑知识库
    * 检索知识库列表
    * 删除知识库
    * 知识库使用量查询
  * 文件管理
    * 文件上传
    * 编辑知识库文件
    * 查询文件列表
    * 删除知识库文件
    * 查询知识库文件详情

***

## 安装使用

* 安装

```shell
  go get -u github.com/itcwc/go-zhipu
```

* 使用

```shell
  import "github.com/itcwc/go-zhipu"
```

***

## 示例

```go
package examples

import (
  "fmt"
  "time"

  zhipu "github.com/itcwc/go-zhipu/model_api"
  "github.com/itcwc/go-zhipu/utils"
)

func Example() {

  apiKey := "your api key"

  // token 填写你自己得或使用扩展中的方法生成
  expireAtTime := int64(1719803252) // token 过期时间
  token, _ := utils.GenerateToken(apiKey, expireAtTime)

  // token缓存处理等 。。。

  mssage := zhipu.PostParams{
    Model: "glm-3-turbo",
    Messages: []zhipu.Message{
      {
        Role:    "user",    // 消息的角色信息 详见文档
        Content: "content", // 消息内容
      },
    },
  }

  var t time.Duration = 60 // 请求等待时间 可不填 默认60秒

  postResponse, err := zhipu.BeCommonModel(expireAtTime, mssage, apiKey, token, t)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(postResponse)
}

```

***

## 其他模型调用

* 通用模型

    ```go
    // sse调用
    go-zhipu.model_api.BeCommonModel(postParams PostParams, token string, time ...time.Duration)
    // 异步调用
    go-zhipu.model_api.ModelAsynchronousCall(postParams PostParams, token string, time ...time.Duration)
    // 任务结果查询
    go-zhipu.model_api.ModelTaskResultQuery(id int, token string, time ...time.Duration)
    ```

* 图像大模型

  ```go
  go-zhipu.model_api.ImageLargeModel(prompt string, model string, userId string, token string, time ...time.Duration)
  ```

* 超拟人大模型

  ```go
  // 同步调用
  go-zhipu.model_api.SuperhumanoidModel(postParams PostSuperhumanoidParams, token string, time ...time.Duration)
  // 异步调用
  go-zhipu.model_api.SHMAsyncCall(postParams PostSuperhumanoidParams, token string, time ...time.Duration)
  ```

* 向量模型

  ```go
  go-zhipu.model_api.VectorModel(input string, model string, token string, time ...time.Duration)
  ```

* Batch API

  ```go
  // 创建 Batch
  go-zhipu.model_api.BatchAPICreate(postParams PostBatchParams, token string, time ...time.Duration)
  // 检索 Batch
  go-zhipu.model_api.BatchSearch(batchId int, token string, time ...time.Duration)
  // 取消 Batch
  go-zhipu.model_api.BatchCancel(batchId int, token string, time ...time.Duration)
  // 列出 Batch
  go-zhipu.model_api.BatchList(after string, limit int, token string, time ...time.Duration)
  // 下载 Batch 结果
  go-zhipu.model_api.BatchDownload(fileId int, token string, time ...time.Duration)
  ```

* 模型微调

  ```go
  // 创建微调任务
  go-zhipu.model_api.CreateModelFineTuning(trainingFile string, model string, token string, time ...time.Duration)
  // 查询微调任务事件
  go-zhipu.model_api.QueryModelFineTuningEvent(jobId int, after string, limit int, token string, time ...time.Duration)
  // 查询微调任务
  go-zhipu.model_api.QueryModelFineTuning(jobId int, after string, limit int, token string, time ...time.Duration)
  // 查询个人微调任务
  go-zhipu.model_api.QueryPersonalModelFineTuning(after string, limit int, token string, time ...time.Duration)
  // 删除微调任务
  go-zhipu.model_api.DeleteModelFineTuning(jobId int, token string, time ...time.Duration)
  // 取消微调任务
  go-zhipu.model_api.CancelModelFineTuning(jobId int, token string, time ...time.Duration)
  // 删除微调模型
  go-zhipu.model_api.DeleteModelFineTuningModel(fineTunedModel string, token string, time ...time.Duration)
  ```

* 搜索工具

```go
// Web-Search-Pro
go-zhipu.model_api.SearchTool(postParams PostSearchParams, token string, time ...time.Duration)
```

* 知识管理
  * 知识库管理

    ```go
    // 创建知识库
    go-zhipu.model_api.Knowledge(postParams PostKnowledgeParams, model string, token string, time ...time.Duration)
    // 编辑知识库
    go-zhipu.model_api.EditKnowledge(postParams PostKnowledgeItemParams, token string, time ...time.Duration)
    // 检索知识库列表
    go-zhipu.model_api.QueryKnowledgeList(page int, size int, token string, time ...time.Duration)
    // 删除知识库
    go-zhipu.model_api.DeleteKnowledge(knowledgeId string, token string, time ...time.Duration)
    // 知识库使用量查询
    go-zhipu.model_api.KnowledgeUsage(token string, time ...time.Duration)
    ```

  * 文件管理

    ```go
    // 文件管理
    go-zhipu.model_api.FileManagement(purpose string, model string, file *FileHeader, token string, time ...time.Duration)
    // 编辑知识库文件
    go-zhipu.model_api.EditKnowledgeFile(postParams KnowledgeFileParams, token string, time ...time.Duration)
    // 查询文件列表
    go-zhipu.model_api.QueryFileList(postParams QueryFileListParams, token string, time ...time.Duration)
    // 删除知识库文件
    go-zhipu.model_api.DeleteKnowledgeFile(id string, token string, time ...time.Duration)
    // 查询知识库文件详情
    go-zhipu.model_api.QueryKnowledgeFileDetail(id string, token string, time ...time.Duration)
    ```
    
## 软件贡献者

<a href="https://github.com/itcwc/go-zhipu/graphs/contributors">
    <img src="https://contrib.rocks/image?repo=itcwc/go-zhipu" />
</a>
  
由 [contrib.rocks](https://contrib.rocks) 自动生成。

## 联系我

* 邮箱：<it_cwc@qq.com>
