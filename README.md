# go-zhipu

* 新版本内容未完成，谨慎使用~

***

## 前言

* go智谱清言api
* 详情见官方文档：<https://maas.aminer.cn/dev/api>
  
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
* Batch API
  * 创建 Batch
  * 取消 Batch
  * 列出 Batch
  * 下载 Batch 结果
* 模型微调
  * 创建微调任务
  * 查询微调任务事件
  * 查询微调任务
  * 查询个人微调任务
  * 删除微调任务
  * 取消微调任务
  * 删除微调模型
* 知识管理
  * 知识库管理
    * 创建知识库
    * 编辑知识库
    * 检索知识库列表
    * 删除知识库
    * 知识库使用量查询
  * 文件管理
    * 文件上传
    * 知识库响应示例
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

zhipu "github.com/itcwc/go-zhipu/model_api"
)

func Example() {
  expireAtTime := int64(1719803252) // token 过期时间
  mssage := zhipu.PostParams{
    Model: "glm-3-turbo",
    Messages: []zhipu.Message{
      {
        Role:    "user",    // 消息的角色信息 详见文档
        Content: "content", // 消息内容
      },
    },
  }
  
  apiKey := "your api key"

  postResponse, err := zhipu.BeCommonModel(expireAtTime, mssage, apiKey)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(postResponse)
}
```

***

## 其他

* 通用模型

    ```go
      go-zhipu.model_api.BeCommonModel(expireAtTime int64, postParams PostParams, apiKey string)
    ```

* 图像大模型

  ```go
    go-zhipu.model_api.ImageLargeModel(expireAtTime int64, prompt string, apiKey string, model string)
  ```

* 超拟人大模型

  ```go
    go-zhipu.model_api.SuperhumanoidModel(expireAtTime int64, postParams PostSuperhumanoidParams, apiKey string)
  ```

* 向量模型

  ```go
    go-zhipu.model_api.VectorModel(expireAtTime int64, input string, apiKey string, model string)
  ```

* Batch API

  ```go
    go-zhipu.model_api.BatchAPI(expireAtTime int64, postParams PostBatchParams, apiKey string)
  ```

* 模型微调

  ```go
    go-zhipu.model_api.ModelFineTuning(expireAtTime int64, trainingFile string, apiKey string, model string)
  ```

* 知识管理
  * 知识库管理

    ```go
      go-zhipu.model_api.Knowledge(expireAtTime int64, postParams PostKnowledgeParams, apiKey string, model string)
    ```

  * 文件管理

    ```go
      go-zhipu.model_api.FileManagement(expireAtTime int64, purpose string, apiKey string, model string, file *FileHeader)
    ```
