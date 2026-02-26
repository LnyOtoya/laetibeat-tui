package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/LnyOtoya/laetibeat-tui/internal/models"
)

// APIClient OpenSubsonic API客户端接口
type APIClient interface {
	Ping() error
	GetMusicFolders() ([]models.MusicFolder, error)
	GetArtists() ([]models.Artist, error)
	GetArtist(artistID string) (*models.Artist, error)
	GetAlbums(artistID string) ([]models.Album, error)
	GetAlbum(albumID string) (*models.Album, error)
	GetSongs(albumID string) ([]models.Song, error)
	GetSong(songID string) (*models.Song, error)
	GetSongStreamURL(songID string) (string, error)
	Search(query string, artistCount, albumCount, songCount int) (*models.SearchResult, error)
}

// Client OpenSubsonic API客户端实现
type Client struct {
	BaseURL    string
	AuthInfo   *AuthInfo
	ClientID   string
	APIVersion string
	HTTPClient *http.Client
}

// NewClient 创建新的API客户端
func NewClient(baseURL, username, password, clientID string) *Client {
	return &Client{
		BaseURL:    baseURL,
		AuthInfo:   NewAuthInfo(username, password),
		ClientID:   clientID,
		APIVersion: "1.16.1",
		HTTPClient: &http.Client{},
	}
}

// buildURL 构建请求URL
func (c *Client) buildURL(endpoint string, params map[string]string) string {
	// 确保BaseURL有正确的协议
	baseURL := c.BaseURL
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}

	// 构建完整的URL
	base, _ := url.Parse(baseURL + "/rest/" + endpoint)
	q := base.Query()

	// 添加认证参数
	q.Set("u", c.AuthInfo.Username)
	q.Set("t", c.AuthInfo.Token)
	q.Set("s", c.AuthInfo.Salt)
	q.Set("v", c.APIVersion)
	q.Set("c", c.ClientID)
	q.Set("f", "json")

	// 添加额外参数
	for k, v := range params {
		q.Set(k, v)
	}

	base.RawQuery = q.Encode()
	return base.String()
}

// sendRequest 发送HTTP请求
func (c *Client) sendRequest(endpoint string, params map[string]string, result interface{}) error {
	// 构建URL
	requestURL := c.buildURL(endpoint, params)

	// 发送GET请求
	resp, err := c.HTTPClient.Get(requestURL)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	// 解析响应
	var apiResponse struct {
		SubsonicResponse struct {
			Status  string `json:"status"`
			Version string `json:"version"`
			Error   *struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error,omitempty"`
			// 其他字段将被嵌入到result中
		} `json:"subsonic-response"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return errors.Wrap(err, "failed to unmarshal API response")
	}

	// 检查响应状态
	if apiResponse.SubsonicResponse.Status != "ok" {
		if apiResponse.SubsonicResponse.Error != nil {
			return errors.Errorf("API error: %s (code: %d)",
				apiResponse.SubsonicResponse.Error.Message,
				apiResponse.SubsonicResponse.Error.Code)
		}
		return errors.Errorf("API error: unknown error (status: %s)", apiResponse.SubsonicResponse.Status)
	}

	// 如果需要解析具体结果
	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return errors.Wrap(err, "failed to unmarshal result")
		}
	}

	return nil
}

// Ping 测试连接
func (c *Client) Ping() error {
	return c.sendRequest("ping.view", nil, nil)
}

// GetMusicFolders 获取音乐文件夹
func (c *Client) GetMusicFolders() ([]models.MusicFolder, error) {
	var response struct {
		SubsonicResponse struct {
			MusicFolders struct {
				MusicFolder []models.MusicFolder `json:"musicFolder"`
			} `json:"musicFolders"`
		} `json:"subsonic-response"`
	}

	if err := c.sendRequest("getMusicFolders.view", nil, &response); err != nil {
		return nil, err
	}

	return response.SubsonicResponse.MusicFolders.MusicFolder, nil
}

// GetArtists 获取艺术家列表
func (c *Client) GetArtists() ([]models.Artist, error) {
	var response struct {
		SubsonicResponse struct {
			Artists struct {
				Index []struct {
					Name   string          `json:"name"`
					Artist []models.Artist `json:"artist"`
				} `json:"index"`
			} `json:"artists"`
		} `json:"subsonic-response"`
	}

	if err := c.sendRequest("getArtists.view", nil, &response); err != nil {
		return nil, err
	}

	// 收集所有艺术家
	var artists []models.Artist
	for _, index := range response.SubsonicResponse.Artists.Index {
		artists = append(artists, index.Artist...)
	}

	return artists, nil
}

// GetArtist 获取艺术家详情
func (c *Client) GetArtist(artistID string) (*models.Artist, error) {
	var response struct {
		SubsonicResponse struct {
			Artist models.Artist `json:"artist"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"id": artistID,
	}

	if err := c.sendRequest("getArtist.view", params, &response); err != nil {
		return nil, err
	}

	return &response.SubsonicResponse.Artist, nil
}

// GetAlbums 获取艺术家的专辑列表
func (c *Client) GetAlbums(artistID string) ([]models.Album, error) {
	var response struct {
		SubsonicResponse struct {
			Artist struct {
				Album []models.Album `json:"album"`
			} `json:"artist"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"id": artistID,
	}

	if err := c.sendRequest("getArtist.view", params, &response); err != nil {
		return nil, err
	}

	return response.SubsonicResponse.Artist.Album, nil
}

// GetAlbum 获取专辑详情
func (c *Client) GetAlbum(albumID string) (*models.Album, error) {
	var response struct {
		SubsonicResponse struct {
			Album models.Album `json:"album"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"id": albumID,
	}

	if err := c.sendRequest("getAlbum.view", params, &response); err != nil {
		return nil, err
	}

	return &response.SubsonicResponse.Album, nil
}

// GetSongs 获取专辑的歌曲列表
func (c *Client) GetSongs(albumID string) ([]models.Song, error) {
	var response struct {
		SubsonicResponse struct {
			Album struct {
				Song []models.Song `json:"song"`
			} `json:"album"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"id": albumID,
	}

	if err := c.sendRequest("getAlbum.view", params, &response); err != nil {
		return nil, err
	}

	return response.SubsonicResponse.Album.Song, nil
}

// GetSong 获取歌曲详情
func (c *Client) GetSong(songID string) (*models.Song, error) {
	var response struct {
		SubsonicResponse struct {
			Song models.Song `json:"song"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"id": songID,
	}

	if err := c.sendRequest("getSong.view", params, &response); err != nil {
		return nil, err
	}

	return &response.SubsonicResponse.Song, nil
}

// GetSongStreamURL 获取歌曲流URL
func (c *Client) GetSongStreamURL(songID string) (string, error) {
	params := map[string]string{
		"id": songID,
	}
	return c.buildURL("stream.view", params), nil
}

// Search 搜索
func (c *Client) Search(query string, artistCount, albumCount, songCount int) (*models.SearchResult, error) {
	var response struct {
		SubsonicResponse struct {
			SearchResult struct {
				Artist []models.Artist `json:"artist"`
				Album  []models.Album  `json:"album"`
				Song   []models.Song   `json:"song"`
			} `json:"searchResult"`
		} `json:"subsonic-response"`
	}

	params := map[string]string{
		"query":       query,
		"artistCount": fmt.Sprintf("%d", artistCount),
		"albumCount":  fmt.Sprintf("%d", albumCount),
		"songCount":   fmt.Sprintf("%d", songCount),
	}

	if err := c.sendRequest("search2.view", params, &response); err != nil {
		return nil, err
	}

	result := &models.SearchResult{
		Artists: response.SubsonicResponse.SearchResult.Artist,
		Albums:  response.SubsonicResponse.SearchResult.Album,
		Songs:   response.SubsonicResponse.SearchResult.Song,
	}

	return result, nil
}
