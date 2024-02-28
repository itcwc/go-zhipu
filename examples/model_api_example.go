package examples

import "github.com/itcwc/go-zhipu/model_api"

func ModelApiExample() {
	apiKey := "your api key"
	model := "glm-3-turbo"
	mssage := []model_api.Messages{
		{
			Role:"user", // 消息的角色信息 详见文档
			Content: "content" // 消息内容
		}
	}

	postResponse, err := model_api.BeCommonModel(expireAtTime, mssage, apiKey, model)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	response.Data(c, postResponse)
}
