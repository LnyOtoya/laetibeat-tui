package api

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// 初始化随机数生成器
func init() {
	rand.Seed(time.Now().UnixNano())
}

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

// AuthInfo 认证信息
type AuthInfo struct {
	Username string
	Password string
	Token    string
	Salt     string
}

// NewAuthInfo 创建新的认证信息
func NewAuthInfo(username, password string) *AuthInfo {
	salt := generateSalt(8)
	token := calculateAuthToken(password, salt)
	
	return &AuthInfo{
		Username: username,
		Password: password,
		Token:    token,
		Salt:     salt,
	}
}

// Refresh 刷新认证信息（生成新的盐值和令牌）
func (a *AuthInfo) Refresh() {
	a.Salt = generateSalt(8)
	a.Token = calculateAuthToken(a.Password, a.Salt)
}
