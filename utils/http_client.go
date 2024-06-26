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

// Get 发送HTTP GET请求并获取响应结果
func Get(apiURL, token string, timeout time.Duration) (map[string]interface{}, error) {
	// 创建请求
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	// 创建一个具有指定超时时间的 HTTP 客户端
	client := http.Client{
		Timeout: timeout,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	var response map[string]interface{}
	// 解析响应结果
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("解码响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		// 返回响应结果
		return response, nil
	}
	// 处理非 200 状态码的情况
	return nil, fmt.Errorf("意外响应状态码: %d, 响应内容: %v", resp.StatusCode, response)
}

// Delete 发送HTTP DELETE请求并获取响应结果
func Delete(apiURL, token string, timeout time.Duration) (map[string]interface{}, error) {
	// 创建请求
	req, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建DELETE请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	// 创建一个具有指定超时时间的 HTTP 客户端
	client := http.Client{Timeout: timeout}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送DELETE请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取DELETE响应体失败: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("解码DELETE响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusOK {
		// 对于DELETE请求，204（No Content）是更常见的成功状态码
		return map[string]interface{}{"message": "Resource deleted successfully"}, nil
	}
	// 处理非预期状态码的情况
	return nil, fmt.Errorf("意外的DELETE响应状态码: %v, 响应内容: %v", resp.StatusCode, response)
}

// Put 发送HTTP PUT请求并获取响应结果
func Put(apiURL, token string, params interface{}, timeout time.Duration) (map[string]interface{}, error) {
	// 将请求参数转换为 JSON 格式
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("编码请求参数失败: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", token)

	// 创建一个具有指定超时时间的 HTTP 客户端
	client := &http.Client{
		Timeout: timeout,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("解码响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		return response, nil
	}
	// 处理非 200 状态码的情况
	return nil, fmt.Errorf("意外响应状态码: %d, 响应内容: %v", resp.StatusCode, response)
}
