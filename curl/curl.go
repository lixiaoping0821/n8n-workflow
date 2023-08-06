package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"n8n-workflow/config"
	"n8n-workflow/pkg"
	"net/http"
	"net/url"
	"os/exec"
)

func HttpRequest(c *gin.Context) {
	// 获取请求的URL
	reqURL := c.Request.URL
	// 解析URL
	parsedURL, err := url.ParseRequestURI(reqURL.String())
	fmt.Println("请求的url1111:", parsedURL.String())
	var apiUrl string
	var body []byte
	if parsedURL.String() == "/api/v1/webhook/execute" {
		var webHookData map[string]interface{}
		if err := c.ShouldBindJSON(&webHookData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("请求的参数1:%s\n", webHookData)
		fmt.Printf("请求的webhookid :%s\n", webHookData["webhookid"])
		apiUrl = "http://" + config.Conf.N8n.Host + ":" + config.Conf.N8n.Port + "/" + config.Conf.N8n.WebHook + "/" + webHookData["webhookid"].(string)
		delete(webHookData, "webhookid")
		body, _ = json.Marshal(webHookData)
	} else {
		apiUrl = "http://" + config.Conf.N8n.Host + ":" + config.Conf.N8n.Port + parsedURL.String()
		fmt.Printf("请求的方法:%s\n", c.Request.Method)
		// 读取请求体
		body, _ = io.ReadAll(c.Request.Body)
	}
	fmt.Printf("请求的url:%s\n", apiUrl)
	fmt.Printf("请求的参数:%s\n", body)
	req, err := http.NewRequest(c.Request.Method, apiUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("8888", err)
		return
	}
	// 设置请求头，根据需要添加认证信息等
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-N8N-API-KEY", config.Conf.N8n.Apikey)
	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	status := 500
	if err != nil {
		fmt.Println("999", err)
		res := pkg.Response{
			Status: status,
			Msg:    pkg.GetMsg(status),
			Error:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println("Response Status:", resp.Status)
	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		res := pkg.Response{
			Status: status,
			Msg:    pkg.GetMsg(status),
			Error:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	fmt.Printf("Response Body:%#v\n", string(respBody))

	//打印原始json数据
	rawResponse := json.RawMessage(respBody)
	status = 200
	res := pkg.Response{
		Status: status,
		Msg:    pkg.GetMsg(status),
		Data:   rawResponse,
	}
	c.JSON(http.StatusOK, res)
}

func ExecRequest(c *gin.Context) {
	fmt.Printf("请求的id:%s\n", c.Param("id"))
	cmd := exec.Command("n8n", "execute", "--id="+c.Param("id"))

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}

	fmt.Println("命令输出:", string(output))
}
func TestJSON() {
	data := `{"name":"John Doe","age":30,"email":"johndoe@example.com","hobby":[{"test":"test","test1":"test1"}]}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("解析JSON数据出错:", err)
		return
	}

	name := result["name"].(string)
	age := result["age"].(float64)
	email := result["email"].(string)
	hobby := result["hobby"].([]interface{})

	fmt.Printf("json解析:%#v", result)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Email:", email)
	fmt.Println("Hobby:", hobby[0])
	hobbyMap := hobby[0].(map[string]interface{})
	test := hobbyMap["test"].(string)
	fmt.Println("test:", test)
}
