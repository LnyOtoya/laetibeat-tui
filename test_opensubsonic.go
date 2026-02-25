package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 生成随机盐值
func generateSalt(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, length)
	for i := range salt {
		salt[i] = charset[rand.Intn(len(charset))]
	}
	return string(salt)
}

// 计算认证哈希
func calculateAuthToken(password, salt string) string {
	hash := md5.Sum([]byte(password + salt))
	return fmt.Sprintf("%x", hash)
}

// 初始化随机数生成器
func init() {
	rand.Seed(time.Now().UnixNano())
}

// OpenSubsonicClient 客户端结构体
type OpenSubsonicClient struct {
	BaseURL string
	Username string
	Password string
	ClientID string
	APIVersion string
}

// Response 基础响应结构
type Response struct {
	SubsonicResponse struct {
		Status string `json:"status"`
		Version string `json:"version"`
		Type string `json:"type,omitempty"`
		ServerVersion string `json:"serverVersion,omitempty"`
		OpenSubsonic bool `json:"openSubsonic,omitempty"`
	} `json:"subsonic-response"`
}

// MusicFolder 音乐文件夹
type MusicFolder struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// MusicFoldersResponse 音乐文件夹响应
type MusicFoldersResponse struct {
	SubsonicResponse struct {
		Status string `json:"status"`
		Version string `json:"version"`
		MusicFolders struct {
			MusicFolder []MusicFolder `json:"musicFolder"`
		} `json:"musicFolders"`
	} `json:"subsonic-response"`
}

// NewClient 创建新客户端
func NewClient(baseURL, username, password, clientID string) *OpenSubsonicClient {
	return &OpenSubsonicClient{
		BaseURL: baseURL,
		Username: username,
		Password: password,
		ClientID: clientID,
		APIVersion: "1.16.1",
	}
}

// buildURL 构建请求URL
func (c *OpenSubsonicClient) buildURL(endpoint string, params map[string]string) string {
	// 确保URL有协议前缀
	baseURL := c.BaseURL
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
	
	base, err := url.Parse(baseURL + "/rest/" + endpoint)
	if err != nil {
		panic(fmt.Sprintf("Invalid URL: %v", err))
	}
	q := base.Query()
	q.Set("u", c.Username)
	
	// 生成随机盐值并计算认证哈希
	salt := generateSalt(8)
	token := calculateAuthToken(c.Password, salt)
	
	q.Set("t", token)
	q.Set("s", salt)
	q.Set("v", c.APIVersion)
	q.Set("c", c.ClientID)
	q.Set("f", "json")

	for k, v := range params {
		q.Set(k, v)
	}

	base.RawQuery = q.Encode()
	return base.String()
}

// Ping 测试连接
func (c *OpenSubsonicClient) Ping() error {
	url := c.buildURL("ping.view", nil)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	if response.SubsonicResponse.Status != "ok" {
		return fmt.Errorf("ping failed: %s", response.SubsonicResponse.Status)
	}

	fmt.Println("Ping successful!")
	fmt.Printf("Server: %s\n", response.SubsonicResponse.Type)
	fmt.Printf("Server Version: %s\n", response.SubsonicResponse.ServerVersion)
	fmt.Printf("OpenSubsonic: %t\n", response.SubsonicResponse.OpenSubsonic)
	return nil
}

// GetMusicFolders 获取音乐文件夹
func (c *OpenSubsonicClient) GetMusicFolders() error {
	url := c.buildURL("getMusicFolders.view", nil)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response MusicFoldersResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	if response.SubsonicResponse.Status != "ok" {
		return fmt.Errorf("getMusicFolders failed: %s", response.SubsonicResponse.Status)
	}

	fmt.Println("Music Folders:")
	for _, folder := range response.SubsonicResponse.MusicFolders.MusicFolder {
		fmt.Printf("- ID: %s, Name: %s\n", folder.ID, folder.Name)
	}
	return nil
}

func main() {
	// 获取用户输入
	var baseURL, username, password string

	fmt.Println("请输入OpenSubsonic服务器信息：")
	fmt.Print("服务器URL (默认: http://localhost:4040): ")
	fmt.Scanln(&baseURL)
	if baseURL == "" {
		baseURL = "http://localhost:4040"
	}

	fmt.Print("用户名 (默认: admin): ")
	fmt.Scanln(&username)
	if username == "" {
		username = "admin"
	}

	fmt.Print("密码 (默认: admin): ")
	fmt.Scanln(&password)
	if password == "" {
		password = "admin"
	}

	clientID := "laetibeat-tui"

	client := NewClient(baseURL, username, password, clientID)

	fmt.Println("Testing OpenSubsonic API...")
	fmt.Println("==================================")

	// 测试ping
	if err := client.Ping(); err != nil {
		fmt.Printf("Ping error: %v\n", err)
	} else {
		fmt.Println("Ping test passed!")
	}

	fmt.Println("==================================")

	// 测试getMusicFolders
	if err := client.GetMusicFolders(); err != nil {
		fmt.Printf("GetMusicFolders error: %v\n", err)
	} else {
		fmt.Println("GetMusicFolders test passed!")
	}

	fmt.Println("==================================")
	fmt.Println("API test completed!")
}
