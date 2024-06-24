# go-zhipu

* 当前已完成，参数详情请看文档，欢迎提交bug...

***

## 前言

* go智谱清言api
* 详情见官方文档：<https://maas.aminer.cn/dev/api>
  
***

## 支持模型

* 通用模型
  * GLM-4
  * GLM-4V
  * GLM-3-Turbo
* 图像大模型
* 超拟人大模型
* 向量模型
* Batch API
* 模型微调
* 知识管理
  * 知识库管理
  * 文件管理

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
