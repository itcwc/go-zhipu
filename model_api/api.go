package model_api

import (
	"fmt"
	"net/textproto"
	"time"

	"github.com/itcwc/go-zhipu/utils"
)

var v4url string = "https://open.bigmodel.cn/api/paas/v4/"

var v3url string = "https://open.bigmodel.cn/api/paas/v3/"

type PostParams struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}
type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 通用模型
func BeCommonModel(expireAtTime int64, mssage []Messages, apiKey string, model string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "chat/completions"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostParams{
		Model:    model,
		Messages: mssage,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostImageParams struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// 图像大模型
func ImageLargeModel(expireAtTime int64, prompt string, apiKey string, model string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "images/generations"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostImageParams{
		Model:  model,
		Prompt: prompt,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostSuperhumanoidParams struct {
	Prompt []Prompt `json:"prompt"`
	Meta   []Meta   `json:"meta"`
}
type Prompt struct {
	Role    string `json:"prompt"`
	Content string `json:"content"`
}
type Meta struct {
	UserInfo string `json:"user_info"`
	BotInfo  string `json:"bot_info"`
	BotName  string `json:"bot_name"`
	UserName string `json:"user_name"`
}

// 超拟人大模型
func SuperhumanoidModel(expireAtTime int64, meta []Meta, prompt []Prompt, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v3url + "model-api/charglm-3/sse-invoke"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostSuperhumanoidParams{
		Prompt: prompt,
		Meta:   meta,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostVectorParams struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// 向量模型
func VectorModel(expireAtTime int64, input string, apiKey string, model string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "mbeddings"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostVectorParams{
		Input: input,
		Model: model,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFineTuningParams struct {
	Model        string `json:"model"`
	TrainingFile string `json:"training_file"`
}

// 模型微调
func ModelFineTuning(expireAtTime int64, trainingFile string, apiKey string, model string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "fine_tuning/jobs"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostFineTuningParams{
		Model:        model,
		TrainingFile: trainingFile,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFileParams struct {
	File    *FileHeader `json:"file"`
	Purpose string      `json:"purpose"`
}

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content   []byte
	tmpfile   string
	tmpoff    int64
	tmpshared bool
}

// 文件管理
func FileManagement(expireAtTime int64, purpose string, apiKey string, model string, file *FileHeader) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "files"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostFileParams{
		File:    file,
		Purpose: purpose,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}
