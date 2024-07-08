package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larksheets "github.com/larksuite/oapi-sdk-go/v3/service/sheets/v3"
)

// 定义一个结构体来映射JSON数据
type Response struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

type ApiResponse struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}

type Data struct {
	Revision         int           `json:"revision"`
	SpreadsheetToken string        `json:"spreadsheetToken"`
	TotalCells       int           `json:"totalCells"`
	ValueRanges      []ValueRanges `json:"valueRanges"`
}

type ValueRanges struct {
	MajorDimension string          `json:"majorDimension"`
	Range          string          `json:"range"`
	Revision       int             `json:"revision"`
	Values         [][]interface{} `json:"values"`
}

type Item struct {
	Link          string `json:"link,omitempty"`
	MentionNotify bool   `json:"mentionNotify,omitempty"`
	MentionType   int    `json:"mentionType,omitempty"`
	Text          string `json:"text,omitempty"`
	Token         string `json:"token,omitempty"`
	Type          string `json:"type,omitempty"`
	Date          string `json:"date,omitempty"`
	Mark          string `json:"mark,omitempty"`
}

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
// 复制该 Demo 后, 需要将 "YOUR_APP_ID", "YOUR_APP_SECRET" 替换为自己应用的 APP_ID, APP_SECRET.
// 以下示例代码是根据 API 调试台参数自动生成，如果存在代码问题，请在 API 调试台填上相关必要参数后再使用
func main() {
	token := getNewToken()
	readData(token)
	// getToken(client)
	// createExcel(client, "test999")

}
func getNewToken() string {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

	// 请求体数据
	data := map[string]string{
		"app_id":     "cli_a603ea18833ad00e",
		"app_secret": "oTJQknTuwUUywRvmtDyv6edKD3aCnxPi",
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return ""
	}

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// 设置Content-Type请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建一个HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	// 打印响应体
	fmt.Println(string(body))
	var respToken Response
	err = json.Unmarshal([]byte(string(body)), &respToken)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	// 打印格式化后的结果
	// fmt.Printf("Code: %d\nExpire: %d\nMessage: %s\nTenant Access Token: %s\n", respToken.Code, respToken.Expire, respToken.Msg, respToken.TenantAccessToken)
	// 打印响应体
	return respToken.TenantAccessToken

}
func readData(token string) {
	url := "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/N5Wts8V9Wh3gXJtyxPvcDbMZnJc/values_batch_get?ranges=5cc663!A2:D&valueRenderOption=FormattedValue&dateTimeRenderOption=FormattedString"

	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置Authorization请求头
	req.Header.Set("Authorization", "Bearer "+token)

	// 创建一个HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// fmt.Println(string(body))
	var respExcelData ApiResponse
	err = json.Unmarshal([]byte(string(body)), &respExcelData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	var Message []Item
	for _, value := range respExcelData.Data.ValueRanges[0].Values {
		jsonData, err := json.Marshal(value[1])

		if err != nil {
			log.Fatalf("Error marshalling data: %v", err)
		}

		// 解析 JSON 数据到结构体
		var result interface{}
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 根据结果类型进行处理
		switch v := result.(type) {
		case string:
			for mesindex, _ := range Message {
				Message[mesindex].Mark = v
			}
			fmt.Println("处理字符串:", v)
		case []interface{}:
			var parsedData []Item
			err = json.Unmarshal(jsonData, &parsedData)
			if err != nil {
				log.Fatalf("Error unmarshalling data: %v", err)
			}
			parsedData[0].Date = value[3].(string)
			// 使用 parsedData
			// fmt.Printf("Parsed Data: %+v,%s,%s,%s\n", parsedData[0], parsedData[0].Link, parsedData[0].Text, parsedData[0].Token)
			Message = append(Message, parsedData[0])
			fmt.Println("处理切片:", v)
		default:
			fmt.Println("其他类型:", v)
		}

		// var parsedData []Item

		// err = json.Unmarshal(jsonData, &parsedData)
		// if err != nil {
		// 	log.Fatalf("Error unmarshalling data: %v", err)
		// }
		// parsedData[0].Date = value[3].(string)
		// // 使用 parsedData
		// // fmt.Printf("Parsed Data: %+v,%s,%s,%s\n", parsedData[0], parsedData[0].Link, parsedData[0].Text, parsedData[0].Token)
		// Message = append(Message, parsedData[0])
		// fmt.Println("Range: ", value[1])
	}
	for _, mes := range Message {
		fmt.Println(mes.Link, mes)
	}
	// fmt.Printf("Parsed JSON: %+v\n", respExcelData)

}

func createExcel(client *lark.Client, filename string) {
	req := larksheets.NewCreateSpreadsheetReqBuilder().
		Spreadsheet(larksheets.NewSpreadsheetBuilder().
			Title(filename).
			FolderToken(`X17CfQ9TDlzyyudn2QtccXX1ngb`).
			Build()).
		Build()

	resp, err := client.Sheets.Spreadsheet.Create(context.Background(), req, larkcore.WithUserAccessToken("u-eJ65h4XU58bbOBlvpoSw1C0l24T1417Hg0w0hgww2dEN"))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	// fmt.Println(larkcore.Prettify(resp))
}
