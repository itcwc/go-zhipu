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

type ToolCalls []struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type Tool struct {
	Type     string `json:"type"`
	Function struct {
		Name        string    `json:"name"`
		Description string    `json:"arguments"`
		Parameters  Parameter `json:"parameters,omitempty"`
	} `json:"function"`
	Retrieval Retrieval `json:"retrieval,omitempty"`
	WebSearch WebSearch `json:"web_search,omitempty"`
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
	Enable       *string `json:"enable,omitempty"`
	SearchQuery  *string `json:"search_query,omitempty"`
	SearchResult *bool   `json:"search_result,omitempty"`
}

// BeCommonModel 通用模型 sse调用
func BeCommonModel(expireAtTime int64, postParams PostParams, apiKey string, timeout ...time.Duration) (map[string]interface{}, error) {
	var t time.Duration
	if len(timeout) > 0 {
		t = timeout[0]
	} else {
		t = 60 * time.Second
	}
	token, _ := utils.GenerateToken(apiKey, expireAtTime)
	apiURL := v4url + "chat/completions"
	postResponse, err := utils.Post(apiURL, token, postParams, t)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 异步调用
func ModelAsynchronousCall(expireAtTime int64, postParams PostParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "async/chat/completions"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 任务结果查询
func ModelTaskResultQuery(expireAtTime int64, id int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "async-result/" + strconv.Itoa(id)
	timeout := 60 * time.Second

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
func ImageLargeModel(expireAtTime int64, prompt string, apiKey string, model string, userId string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "images/generations"
	timeout := 60 * time.Second

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
func SuperhumanoidModel(expireAtTime int64, postParams PostSuperhumanoidParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "chat/completions"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 异步调用
func SHMAsyncCall(expireAtTime int64, postParams PostSuperhumanoidParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "async/chat/completions"
	timeout := 60 * time.Second

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

	apiURL := v4url + "embeddings"
	timeout := 60 * time.Second

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
func BatchAPICreate(expireAtTime int64, postParams PostBatchParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "batches"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 检索 Batch GET
func BatchSearch(expireAtTime int64, batchId int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "batches/" + strconv.Itoa(batchId)
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 取消 Batch POST
func BatchCancel(expireAtTime int64, batchId int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "batches/" + strconv.Itoa(batchId) + "/cancel"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, nil, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 列出 Batch GET
func BatchList(expireAtTime int64, after string, limit int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "batches?after=" + after + "&limit=" + strconv.Itoa(limit)
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 下载 Batch 结果 GET
func BatchDownload(expireAtTime int64, fileId int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "files/" + strconv.Itoa(fileId) + "/content"
	timeout := 60 * time.Second

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
func CreateModelFineTuning(expireAtTime int64, postParams PostFineTuningParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询微调任务事件 GET
func QueryModelFineTuningEvent(expireAtTime int64, jobId int, after string, limit int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId) + "/events?limit=" + strconv.Itoa(limit) + "&after=" + after
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询微调任务 GET
func QueryModelFineTuning(expireAtTime int64, jobId int, after string, limit int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId) + "/events?limit=" + strconv.Itoa(limit) + "&after=" + after
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询个人微调任务 GET
func QueryPersonalModelFineTuning(expireAtTime int64, after string, limit int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs/?limit=" + strconv.Itoa(limit) + "&after=" + after
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除微调任务 DELETE
func DeleteModelFineTuning(expireAtTime int64, jobId int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs/" + strconv.Itoa(jobId)
	timeout := 60 * time.Second

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 取消微调任务 POST
func CancelModelFineTuning(expireAtTime int64, jobId int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/jobs" + strconv.Itoa(jobId) + "/cancel"
	timeout := 60 * time.Second

	postResponse, err := utils.Post(apiURL, token, nil, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除微调模型 DELETE
func DeleteModelFineTuningModel(expireAtTime int64, fineTunedModel string, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "fine_tuning/fine_tuned_models/" + fineTunedModel
	timeout := 60 * time.Second

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
func Knowledge(expireAtTime int64, postParams PostKnowledgeParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "knowledge"
	timeout := 60 * time.Second

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
func EditKnowledge(expireAtTime int64, postParams PostKnowledgeItemParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "knowledge/" + postParams.KnowledgeId
	timeout := 60 * time.Second

	postResponse, err := utils.Put(apiURL, token, postParams, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 检索知识库列表 GET
func QueryKnowledgeList(expireAtTime int64, page int, size int, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "knowledge?page=" + strconv.Itoa(page) + "&size=" + strconv.Itoa(size)
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除知识库 DELETE
func DeleteKnowledge(expireAtTime int64, knowledgeId string, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "knowledge/" + knowledgeId
	timeout := 60 * time.Second

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 知识库使用量查询 GET
func KnowledgeUsage(expireAtTime int64, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "knowledge/capacity"
	timeout := 60 * time.Second

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
func FileManagement(expireAtTime int64, postParams PostFileParams, apiKey string) (map[string]interface{}, error) {

	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "files"
	timeout := 60 * time.Second

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
func EditKnowledgeFile(expireAtTime int64, postParams KnowledgeFileParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "document/" + strconv.Itoa(postParams.KnowledgeType)
	timeout := 60 * time.Second

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
func QueryFileList(expireAtTime int64, postParams QueryFileListParams, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

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
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 删除知识库文件 DELETE
func DeleteKnowledgeFile(expireAtTime int64, id string, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "document/" + id
	timeout := 60 * time.Second

	postResponse, err := utils.Delete(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}

// 查询知识库文件详情 GET
func QueryKnowledgeFileDetail(expireAtTime int64, id string, apiKey string) (map[string]interface{}, error) {
	token, _ := utils.GenerateToken(apiKey, expireAtTime)

	apiURL := v4url + "document/" + id
	timeout := 60 * time.Second

	postResponse, err := utils.Get(apiURL, token, timeout)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	return postResponse, nil
}
