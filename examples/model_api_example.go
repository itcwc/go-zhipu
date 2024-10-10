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

	postResponse, err := zhipu.BeCommonModel(mssage, token, t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(postResponse)
}
