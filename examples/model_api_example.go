package examples

import (
	"fmt"

	itcwc "github.com/itcwc/go-zhipu/model_api"
)

func example() {
	expireAtTime := int64(1719803252) // token 过期时间
	mssage := itcwc.PostParams{
		Model: "glm-3-turbo",
		Message: []itcwc.Messages{
			{
				Role:    "user",    // 消息的角色信息 详见文档
				Content: "content", // 消息内容
			},
		},
	}
	apiKey := "your api key"

	postResponse, err := itcwc.BeCommonModel(expireAtTime, mssage, apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(postResponse)
}
