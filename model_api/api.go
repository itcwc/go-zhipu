package model_api

import (
	"fmt"
	"mime/multipart"
	"net/url"
	"strconv"
	"time"

	"github.com/itcwc/go-zhipu/utils"
)

var v4url string = "https://open.bigmodel.cn/api/paas/v4/"

// var v3url string = "https://open.bigmodel.cn/api/paas/v3/"

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
	Tools       []Tool       `json:"tools,omitempty"`
	ToolChoice  *string      `json:"tool_choice,omitempty"`
	UserId      *string      `json:"user_id,omitempty"`
}
type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  *ToolCalls `json:"tool_calls,omitempty"`
	ToolCallId *string    `json:"tool_call_id,omitempty"`
}

type ToolCalls []struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type Tool struct {
	Type      string        `json:"type"`
	Function  *ToolFunction `json:"function,omitempty"`
	Retrieval *Retrieval    `json:"retrieval,omitempty"`
	WebSearch *WebSearch    `json:"web_search,omitempty"`
}

type ToolFunction struct {
	Name        string    `json:"name"`
	Description string    `json:"arguments"`
	Parameters  Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

type Property struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

type Retrieval struct {
	KnowledgeId    string  `json:"knowledge_id"`
	PromptTemplate *string `json:"prompt_template,omitempty"`
}
type WebSearch struct {
	Enable       *bool   `json:"enable,omitempty"`
	SearchQuery  *string `json:"search_query,omitempty"`
	SearchResult *bool   `json:"search_result,omitempty"`
}

// BeCommonModel 通用模型 sse调用
func BeCommonModel(postParams PostParams, token string, time ...time.Duration) (map[string]interface{}, error) {

	timeout := utils.GetTimeout(time...)

	apiURL := v4url + "chat/completions"
	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 异步调用
func ModelAsynchronousCall(postParams PostParams, token string, time ...time.Duration) (map[string]interface{}, error) {

	timeout := utils.GetTimeout(time...)

	apiURL := v4url + "async/chat/completions"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 任务结果查询
func ModelTaskResultQuery(id int, token string, time ...time.Duration) (map[string]interface{}, error) {
	timeout := utils.GetTimeout(time...)

	apiURL := v4url + "async-result/" + strconv.Itoa(id)
	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostImageParams struct {
	Model  string  `json:"model"`
	Prompt string  `json:"prompt"`
	UserId *string `json:"user_id,omitempty"`
}

// 图像大模型
func ImageLargeModel(prompt string, model string, userId string, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "images/generations"

	postParams := PostImageParams{
		Model:  model,
		Prompt: prompt,
		UserId: &userId,
	}

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostSuperhumanoidParams struct {
	Model       string       `json:"model"`
	Messages    []Message    `json:"messages"`
	Meta        []Meta       `json:"meta"`
	RequestId   *string      `json:"request_id,omitempty"`
	DoSample    *bool        `json:"do_sample,omitempty"`
	Stream      *bool        `json:"stream,omitempty"`
	Temperature *float64     `json:"temperature,omitempty"`
	TopP        *float64     `json:"top_p,omitempty"`
	Maxtokens   *int         `json:"max_tokens,omitempty"`
	Stop        *interface{} `json:"stop,omitempty"`
	UserId      *string      `json:"user_id,omitempty"`
}
type Meta struct {
	UserInfo string `json:"user_info"`
	BotInfo  string `json:"bot_info"`
	BotName  string `json:"bot_name"`
	UserName string `json:"user_name"`
}

// 超拟人大模型 同步调用
func SuperhumanoidModel(postParams PostSuperhumanoidParams, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "chat/completions"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 异步调用
func SHMAsyncCall(postParams PostSuperhumanoidParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "async/chat/completions"

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
func VectorModel(input string, model string, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "embeddings"

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

// Batch API 创建 Batch
func BatchAPICreate(postParams PostBatchParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "batches"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 检索 Batch GET
func BatchSearch(batchId int, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "batches/" + strconv.Itoa(batchId)

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 取消 Batch POST
func BatchCancel(batchId int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "batches/" + strconv.Itoa(batchId) + "/cancel"

	postResponse, err := utils.Post(apiURL, token, nil, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 列出 Batch GET
func BatchList(after string, limit int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "batches?after=" + after + "&limit=" + strconv.Itoa(limit)

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 下载 Batch 结果 GET
func BatchDownload(fileId int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "files/" + strconv.Itoa(fileId) + "/content"

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFineTuningParams struct {
	Model           string          `json:"model"`
	TrainingFile    string          `json:"training_file"`
	ValidationFile  *string         `json:"validation_file,omitempty"`
	Hyperparameters *Hyperparameter `json:"hyperparameters,omitempty"`
	Suffix          *string         `json:"suffix,omitempty"`
	RequestId       *string         `json:"request_id,omitempty"`
}

type Hyperparameter struct {
	LearningRateMultiplier string `json:"learning_rate_multiplier"`
	BatchSize              string `json:"batch_size"`
	NEpochs                string `json:"n_epochs"`
}

// 模型微调 创建微调任务
func CreateModelFineTuning(postParams PostFineTuningParams, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostSearchParams struct {
	ToolType  string    `json:"tool_type"`
	Messages  []Message `json:"messages"`
	RequestId *string   `json:"request_id"`
	Stream    *bool     `json:"stream"`
}

// 搜索工具
func SearchTool(postParams PostSearchParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "tools"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询微调任务事件 GET
func QueryModelFineTuningEvent(jobId int, after string, limit int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId) + "/events?limit=" + strconv.Itoa(limit) + "&after=" + after

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询微调任务 GET
func QueryModelFineTuning(jobId int, after string, limit int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId) + "/events?limit=" + strconv.Itoa(limit) + "&after=" + after

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询个人微调任务 GET
func QueryPersonalModelFineTuning(after string, limit int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs/?limit=" + strconv.Itoa(limit) + "&after=" + after

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除微调任务 DELETE
func DeleteModelFineTuning(jobId int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId)

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 取消微调任务 POST
func CancelModelFineTuning(jobId int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/jobs" + strconv.Itoa(jobId) + "/cancel"

	postResponse, err := utils.Post(apiURL, token, nil, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除微调模型 DELETE
func DeleteModelFineTuningModel(fineTunedModel string, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "fine_tuning/fine_tuned_models/" + fineTunedModel

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostKnowledgeParams struct {
	EmbeddingId int     `json:"embedding_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// 知识库管理 创建知识库 POST
func Knowledge(postParams PostKnowledgeParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "knowledge"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostKnowledgeItemParams struct {
	KnowledgeId string  `json:"knowledge_id"`
	EmbeddingId string  `json:"embedding_id"`
	Name        *string `json:"name,omitempty"`
	Content     *string `json:"content,omitempty"`
}

// 编辑知识库 PUT
func EditKnowledge(postParams PostKnowledgeItemParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "knowledge/" + postParams.KnowledgeId

	postResponse, err := utils.Put(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 检索知识库列表 GET
func QueryKnowledgeList(page int, size int, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "knowledge?page=" + strconv.Itoa(page) + "&size=" + strconv.Itoa(size)

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除知识库 DELETE
func DeleteKnowledge(knowledgeId string, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "knowledge/" + knowledgeId

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 知识库使用量查询 GET
func KnowledgeUsage(apiKey string, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "knowledge/capacity"

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type PostFileParams struct {
	File    *multipart.FileHeader `json:"file,omitempty"`
	Purpose string                `json:"purpose"`
}

// 文件管理 文件上传
func FileManagement(postParams PostFileParams, token string, time ...time.Duration) (map[string]interface{}, error) {

	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "files"

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type KnowledgeFileParams struct {
	DocumentId      string  `json:"document_id"`
	KnowledgeType   int     `json:"knowledge_type"`
	CustomSeparator *string `json:"custom_separator,omitempty"`
	SentenceSize    *int    `json:"sentence_size,omitempty"`
}

// 编辑知识库文件 PUT
func EditKnowledgeFile(postParams KnowledgeFileParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "document/" + strconv.Itoa(postParams.KnowledgeType)

	postResponse, err := utils.Put(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

type QueryFileListParams struct {
	Purpose     string  `json:"purpose"`
	KnowledgeId *string `json:"knowledge_id,omitempty"`
	Page        *int    `json:"page,omitempty"`
	Limit       *int    `json:"limit,omitempty"`
	After       *string `json:"after,omitempty"`
	Order       *string `json:"order,omitempty"`
}

// 查询文件列表 GET
func QueryFileList(postParams QueryFileListParams, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	// 将postParams转换为url.Values以便附加到URL
	params := url.Values{}
	if postParams.Purpose != "" {
		params.Set("purpose", postParams.Purpose)
	}
	if postParams.KnowledgeId != nil {
		params.Set("knowledge_id", *postParams.KnowledgeId)
	}
	if postParams.Page != nil {
		params.Set("page", fmt.Sprint(*postParams.Page))
	}
	if postParams.Limit != nil {
		params.Set("limit", fmt.Sprint(*postParams.Limit))
	}
	if postParams.After != nil {
		params.Set("after", *postParams.After)
	}
	if postParams.Order != nil {
		params.Set("order", *postParams.Order)
	}

	// 将查询参数附加到URL
	apiURL := fmt.Sprintf("%s/files?%s", v4url, params.Encode())

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除知识库文件 DELETE
func DeleteKnowledgeFile(id string, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "document/" + id

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询知识库文件详情 GET
func QueryKnowledgeFileDetail(id string, token string, time ...time.Duration) (map[string]interface{}, error) {
	// token, _ := utils.GenerateToken(apiKey, expireAtTime)
	timeout := utils.GetTimeout(time...)
	// timeout := 60 * time.Second

	apiURL := v4url + "document/" + id

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}
