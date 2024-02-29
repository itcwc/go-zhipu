# go-zhipu
* 当前未完成，待续...

***
## 前言
  * go智谱清言api
  * 详情见官方文档：https://maas.aminer.cn/dev/api
  
***
## 支持模型
  * 通用模型
    * GLM-4
    * GLM-4V
    * GLM-3-Turbo
  * 图像大模型
  * 超拟人大模型
  * 向量模型
  * 模型微调
  * 文件管理
    * 文件上传
    * 查询文件列表

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
  import "github.com/itcwc/go-zhipu/model_api"

  ...

	expireAtTime := int64(1640995200) // token 过期时间
	mssage := []model_api.Messages{
		{
			Role:    "user",    // 消息的角色信息 详见文档
			Content: "content", // 消息内容
		},
	}
	apiKey := "your api key"
	model := "glm-3-turbo"

	postResponse, err := model_api.BeCommonModel(expireAtTime, mssage, apiKey, model)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(postResponse)
```

***

## 其他
  * 通用模型
    ```go
      go-zhipu.model_api.BeCommonModel(expireAtTime int64, mssage []Messages, apiKey string, model string)
    ```
  * 图像大模型
    ```go
      go-zhipu.model_api.ImageLargeModel(expireAtTime int64, prompt string, apiKey string, model string)
    ```
  * 超拟人大模型
    ```go
      go-zhipu.model_api.SuperhumanoidModel(expireAtTime int64, meta []Meta, prompt []Prompt, apiKey string)
    ```
  * 向量模型
    ```go
      go-zhipu.model_api.VectorModel(expireAtTime int64, input string, apiKey string, model string)
    ```
  * 模型微调
    ```go
      go-zhipu.model_api.ModelFineTuning(expireAtTime int64, trainingFile string, apiKey string, model string)
    ```
  * 文件管理
    ```go
      go-zhipu.model_api.FileManagement(expireAtTime int64, purpose string, apiKey string, model string, file *FileHeader)
    ```