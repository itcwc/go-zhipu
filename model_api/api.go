package model_api

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/itcwc/go-zhipu/utils"
)

var v4url string = "https://open.bigmodel.cn/api/paas/v4/"

var v3url string = "https://open.bigmodel.cn/api/paas/v3/"

type PostParams struct {
	Model       string       `json:"model"`
	Messages    []Message    `json:"messages"`
	RequestId   *string      `json:"request_id,omitempty"`
	DoSample    *bool        `json:"do_sample,omitempty"`
	Stream      *bool        `json:"stream,omitempty"`
	Temperature *float64     `json:"temperature,omitempty"`
	TopP        *float64     `json:"top_p,omitempty"`
	Maxtokens   *int         `json:"max_tokens,omitempty"`
	Stop        *interface{} `json:"stop,omitempty"`
	Tools       *Tool        `json:"tools,omitempty"`
	ToolChoice  *string      `json:"tool_choice,omitempty"`
	UserId      *string      `json:"user_id,omitempty"`
}
type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  *ToolCalls `json:"tool_calls,omitempty"`
	ToolCallId *string    `json:"tool_call_id,omitempty"`
}

type Tool struct {
	Type      string       `json:"type"`
	Function  ToolFunction `json:"function,omitempty"`
	Retrieval Retrieval    `json:"retrieval,omitempty"`
	WebSearch WebSearch    `json:"web_search,omitempty"`
}

type ToolFunction struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Parameters  Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Location string `json:"location"`
	Unit     string `json:"unit"`
}
type Retrieval struct {
	KnowledgeId    string `json:"knowledge_id"`
	PromptTemplate string `json:"prompt_template"`
}
type WebSearch struct {
	Enable       string `json:"enable"`
	SearchQuery  string `json:"search_query"`
	SearchResult bool   `json:"search_result"`
}

type ToolCalls []struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

// type ToolCallFunction struct {
// 	Name      string `json:"name"`
// 	Arguments string `json:"arguments"`
// }

// 通用模型
func BeCommonModel(expireAtTime int64, postParams PostParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "chat/completions"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostImageParams struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	UserId string `json:"user_id"`
}

// 图像大模型
func ImageLargeModel(expireAtTime int64, prompt string, apiKey string, model string, userId string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "images/generations"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postParams := PostImageParams{
		Model:  model,
		Prompt: prompt,
		UserId: userId,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostSuperhumanoidParams struct {
	Model       string      `json:"model"`
	Messages    []Message   `json:"messages"`
	Meta        []Meta      `json:"meta"`
	RequestId   string      `json:"request_id"`
	DoSample    bool        `json:"do_sample"`
	Stream      bool        `json:"stream"`
	Temperature float64     `json:"temperature"`
	TopP        float64     `json:"top_p"`
	Maxtokens   int         `json:"max_tokens"`
	Stop        interface{} `json:"stop"`
	UserId      string      `json:"user_id"`
}
type Meta struct {
	UserInfo string `json:"user_info"`
	BotInfo  string `json:"bot_info"`
	BotName  string `json:"bot_name"`
	UserName string `json:"user_name"`
}

// 超拟人大模型
func SuperhumanoidModel(expireAtTime int64, postParams PostSuperhumanoidParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v3url + "model-api/charglm-3/sse-invoke"
	timeout := 60 * time.Second

	// 示例 POST 请求
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

type PostBatchParams struct {
	InputFileId      string      `json:"input_file_id"`
	Endpoint         string      `json:"endpoint"`
	CompletionWindow string      `json:"completion_window"`
	Metadata         interface{} `json:"metadata"`
}

// Batch API
func BatchAPI(expireAtTime int64, postParams PostBatchParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "batches"
	timeout := 60 * time.Second

	// 示例 POST 请求

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFineTuningParams struct {
	Model           string         `json:"model"`
	TrainingFile    string         `json:"training_file"`
	ValidationFile  string         `json:"validation_file"`
	Hyperparameters Hyperparameter `json:"hyperparameters"`
	Suffix          string         `json:"suffix"`
	RequestId       string         `json:"request_id"`
}

type Hyperparameter struct {
	LearningRateMultiplier string `json:"learning_rate_multiplier"`
	BatchSize              int    `json:"batch_size"`
	NEpochs                int    `json:"n_epochs"`
}

// 模型微调
func ModelFineTuning(expireAtTime int64, postParams PostFineTuningParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "fine_tuning/jobs"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostKnowledgeParams struct {
	EmbeddingId int    `json:"embedding_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 知识库管理
func Knowledge(expireAtTime int64, postParams PostKnowledgeParams, apiKey string, model string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "knowledge"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFileParams struct {
	File    *multipart.FileHeader `json:"file"`
	Purpose string                `json:"purpose"`
}

// 文件管理
func FileManagement(expireAtTime int64, postParams PostFileParams, apiKey string, model string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	// 示例用法
	apiURL := v4url + "files"
	timeout := 60 * time.Second

	// 示例 POST 请求
	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}
