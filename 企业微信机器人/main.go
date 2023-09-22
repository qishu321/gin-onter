package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	webhookURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=XXXXXXXXXXXXXXXXXXXXXX"
	imageDir   = "static/"
)

func main() {
	for {
		sendRandomImageMessage()
		time.Sleep(1 * time.Minute)
	}
}

func sendRandomImageMessage() {
	// 1. 读取所有图片文件
	imageFiles, err := ioutil.ReadDir(imageDir)
	if err != nil {
		fmt.Println("Error reading image directory:", err)
		return
	}

	// 2. 随机选择一张图片
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(imageFiles))
	selectedImageFile := imageFiles[randomIndex]

	// 3. 打开选定的图片文件
	imagePath := imageDir + selectedImageFile.Name()
	imageFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer imageFile.Close()

	// 4. 将图片内容读取为base64编码
	imageBytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return
	}
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	// 5. 计算图片内容的MD5值
	md5Sum := md5.Sum(imageBytes)
	md5Value := fmt.Sprintf("%x", md5Sum)

	// 6. 构建消息体
	requestBody := fmt.Sprintf(`{
        "msgtype": "image",
        "image": {
            "base64": "%s",
            "md5": "%s"
        }
    }`, base64Image, md5Value)

	// 7. 发送HTTP POST请求到企业微信机器人的Webhook地址
	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(requestBody))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// 8. 处理响应
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Random image message sent successfully! Image: %s\n", selectedImageFile.Name())
	} else {
		fmt.Printf("Failed to send random image message. Status code: %d\n", resp.StatusCode)
	}
}
