package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Post 发送HTTP POST请求并获取响应结果
func Post(apiURL, token string, params interface{}, timeout time.Duration) (map[string]interface{}, error) {
	// 将请求参数转换为 JSON 格式
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("编码请求参数失败: %v", err)
	}
	fmt.Println(string(jsonParams))

	// 创建请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头的 Authorization 字段
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", token)

	// 创建一个具有指定超时时间的 HTTP 客户端
	client := http.Client{Timeout: timeout}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// 读取响应体内容
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	// 解析响应结果
	if err != nil {
		return nil, fmt.Errorf("解码响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		// 返回响应结果
		return response, nil
	}
	// 处理非 200 状态码的情况
	return nil, fmt.Errorf("意外响应状态码: %v", response)
}

func Stream(apiURL, token string, params interface{}, timeout time.Duration) (*http.Response, error) {
	// 将请求参数转换为 JSON 格式
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	// 创建 POST 请求对象
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, err
	}

	// 设置请求头部信息
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// 创建具有指定超时时间的 HTTP 客户端
	client := http.Client{Timeout: timeout}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}

	// 检查响应状态码是否为 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("意外响应状态码: %s", resp.Status)
	}

	return resp, nil
}

func Get(apiURL, token string, timeout time.Duration) (map[string]interface{}, error) {
	// 创建一个具有指定超时时间的 HTTP 客户端
	client := http.Client{Timeout: timeout}

	// 创建 GET 请求对象
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头部信息
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码是否为 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("意外响应状态码: %s", resp.Status)
	}

	// 解码响应结果
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("解码响应失败: %v", err)
	}

	return response, nil
}
