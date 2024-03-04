package examples

import (
	"fmt"

	"github.com/itcwc/go-zhipu/model_api"
)

func ModelApiExample() {

	expireAtTime := int64(1640995200) // token 过期时间
	apiKey := "your api key"

	postParams := model_api.PostParams{
		Model: "glm-3-turbo",
		Messages: []model_api.Messages{
			{
				Role:    "user",   // 消息的角色信息 详见文档
				Content: "提问消息内容", // 消息内容
			},
		},
	}

	postResponse, err := model_api.BeCommonModel(expireAtTime, postParams, apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(postResponse)
}
