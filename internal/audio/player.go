package audio

import (
	"github.com/LnyOtoya/laetibeat-tui/internal/models"
)

// Player 播放器接口
type Player interface {
	Play(url string) error
	Pause() error
	Resume() error
	Stop() error
	Next() error
	Previous() error
	SetVolume(volume int) error
	Seek(position int) error
	GetStatus() (models.PlayStatus, error)
	Close() error
}

// PlaybackOptions 播放选项
type PlaybackOptions struct {
	Volume    int
	Crossfade bool
	Repeat    bool
	Random    bool
}

// NewPlaybackOptions 创建默认播放选项
func NewPlaybackOptions() *PlaybackOptions {
	return &PlaybackOptions{
		Volume:    80,
		Crossfade: false,
		Repeat:    false,
		Random:    false,
	}
}
