package mpv

import (
	"github.com/pkg/errors"
)

// Play 播放音频
func (c *Client) Play(url string) error {
	return c.SendCommand([]interface{}{"loadfile", url, "replace"})
}

// Pause 暂停播放
func (c *Client) Pause() error {
	return c.SetProperty("pause", true)
}

// Resume 恢复播放
func (c *Client) Resume() error {
	return c.SetProperty("pause", false)
}

// Stop 停止播放
func (c *Client) Stop() error {
	return c.SendCommand([]interface{}{"stop"})
}

// Next 下一曲
func (c *Client) Next() error {
	return c.SendCommand([]interface{}{"playlist-next"})
}

// Previous 上一曲
func (c *Client) Previous() error {
	return c.SendCommand([]interface{}{"playlist-prev"})
}

// SetVolume 设置音量（0-100）
func (c *Client) SetVolume(volume int) error {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	return c.SetProperty("volume", volume)
}

// Seek 调整播放进度（秒）
func (c *Client) Seek(position int) error {
	return c.SendCommand([]interface{}{"seek", position, "absolute"})
}

// SeekRelative 相对调整播放进度（秒）
func (c *Client) SeekRelative(seconds int) error {
	return c.SendCommand([]interface{}{"seek", seconds, "relative"})
}

// SetRepeat 设置循环模式
// mode: "no", "inf", "once"
func (c *Client) SetRepeat(mode string) error {
	validModes := map[string]bool{
		"no":   true,
		"inf":  true,
		"once": true,
	}
	
	if !validModes[mode] {
		return errors.New("invalid repeat mode")
	}
	
	return c.SetProperty("loop", mode)
}

// SetRandom 设置随机播放
func (c *Client) SetRandom(random bool) error {
	return c.SetProperty("shuffle", random)
}

// AddToPlaylist 添加到播放列表
func (c *Client) AddToPlaylist(url string) error {
	return c.SendCommand([]interface{}{"loadfile", url, "append"})
}

// ClearPlaylist 清空播放列表
func (c *Client) ClearPlaylist() error {
	return c.SendCommand([]interface{}{"playlist-clear"})
}

// RemoveFromPlaylist 从播放列表移除项目
func (c *Client) RemoveFromPlaylist(index int) error {
	return c.SendCommand([]interface{}{"playlist-remove", index})
}

// SetCrossfade 设置交叉淡入淡出（秒）
func (c *Client) SetCrossfade(seconds float64) error {
	return c.SetProperty("audio-fade", seconds)
}
