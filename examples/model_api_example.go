package examples

import (
	"fmt"

	"github.com/itcwc/go-zhipu/model_api"
)

func ModelApiExample() {

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
}
