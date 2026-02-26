package models

import (
	"time"
)

// MusicFolder 音乐文件夹
type MusicFolder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Artist 艺术家
type Artist struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AlbumCount   int    `json:"albumCount,omitempty"`
	CoverArt     string `json:"coverArt,omitempty"`
	ArtistImageUrl string `json:"artistImageUrl,omitempty"`
}

// Album 专辑
type Album struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	ArtistID     string    `json:"artistId"`
	CoverArt     string    `json:"coverArt,omitempty"`
	SongCount    int       `json:"songCount,omitempty"`
	Duration     int       `json:"duration,omitempty"`
	Year         int       `json:"year,omitempty"`
	Genre        string    `json:"genre,omitempty"`
	Created      time.Time `json:"created,omitempty"`
}

// Song 歌曲
type Song struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Album        string    `json:"album"`
	AlbumID      string    `json:"albumId"`
	Artist       string    `json:"artist"`
	ArtistID     string    `json:"artistId"`
	Track        int       `json:"track,omitempty"`
	DiscNumber   int       `json:"discNumber,omitempty"`
	Year         int       `json:"year,omitempty"`
	Genre        string    `json:"genre,omitempty"`
	Duration     int       `json:"duration"`
	BitRate      int       `json:"bitRate,omitempty"`
	Size         int64     `json:"size,omitempty"`
	Path         string    `json:"path,omitempty"`
	Suffix       string    `json:"suffix,omitempty"`
	ContentType  string    `json:"contentType,omitempty"`
	IsVideo      bool      `json:"isVideo,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	AlbumArtists []string  `json:"albumArtists,omitempty"`
	AlbumArtistIDs []string `json:"albumArtistIds,omitempty"`
}

// Playlist 播放列表
type Playlist struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Comment      string    `json:"comment,omitempty"`
	Owner        string    `json:"owner,omitempty"`
	Public       bool      `json:"public,omitempty"`
	SongCount    int       `json:"songCount,omitempty"`
	Duration     int       `json:"duration,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	Changed      time.Time `json:"changed,omitempty"`
	CoverArt     string    `json:"coverArt,omitempty"`
}

// PlayStatus 播放状态
type PlayStatus struct {
	Playing      bool    `json:"playing"`
	CurrentSong  *Song   `json:"currentSong,omitempty"`
	Position     int     `json:"position"` // 播放位置（秒）
	Duration     int     `json:"duration"` // 歌曲时长（秒）
	Volume       int     `json:"volume"`   // 音量（0-100）
	Repeat       bool    `json:"repeat"`   // 是否循环
	Random       bool    `json:"random"`   // 是否随机播放
}

// SearchResult 搜索结果
type SearchResult struct {
	Artists []Artist `json:"artists"`
	Albums  []Album  `json:"albums"`
	Songs   []Song   `json:"songs"`
}

// APIResponse API 响应基础结构
type APIResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Error   *Error `json:"error,omitempty"`
}

// Error API 错误
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
