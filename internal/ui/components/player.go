package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/models"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// Player 播放器组件
type Player struct {
	Status   models.PlayStatus
	Progress progress.Model
	Volume   progress.Model
	Width    int
	Height   int
}

// NewPlayer 创建新的播放器组件
func NewPlayer(width, height int) *Player {
	// 初始化进度条
	progressBar := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(width-4),
	)

	// 初始化音量条
	volumeBar := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(20),
	)

	return &Player{
		Status: models.PlayStatus{
			Playing: false,
			Volume:  80,
		},
		Progress: progressBar,
		Volume:   volumeBar,
		Width:    width,
		Height:   height,
	}
}

// SetStatus 设置播放状态
func (p *Player) SetStatus(status models.PlayStatus) {
	p.Status = status

	// 更新进度条
	if status.Duration > 0 {
		progress := float64(status.Position) / float64(status.Duration)
		p.Progress.SetPercent(progress)
	}

	// 更新音量条
	volume := float64(status.Volume) / 100.0
	p.Volume.SetPercent(volume)
}

// View 渲染播放器
func (p *Player) View() string {
	var builder strings.Builder

	// 歌曲信息
	if p.Status.CurrentSong != nil {
		title := p.Status.CurrentSong.Title
		artist := p.Status.CurrentSong.Artist
		album := p.Status.CurrentSong.Album

		// 截断过长的文本
		maxWidth := p.Width - 4
		if len(title) > maxWidth {
			title = title[:maxWidth-3] + "..."
		}

		songInfo := fmt.Sprintf("%s - %s", title, artist)
		if len(album) > 0 {
			songInfo += fmt.Sprintf(" (\x1b[2m%s\x1b[0m)", album)
		}

		builder.WriteString(styles.Title.Render(songInfo) + "\n")
	} else {
		builder.WriteString(styles.Title.Render("未播放任何歌曲") + "\n")
	}

	// 进度条
	progressText := p.formatTime(p.Status.Position) + " / " + p.formatTime(p.Status.Duration)

	// 计算进度文本和进度条的布局
	progressWidth := p.Width - 4
	textWidth := len(progressText)
	barWidth := progressWidth - textWidth - 2

	if barWidth > 0 {
		p.Progress.Width = barWidth
		progressView := lipgloss.JoinHorizontal(lipgloss.Center,
			progressText,
			p.Progress.View(),
		)
		builder.WriteString(styles.ProgressBar.Render(progressView) + "\n")
	}

	// 控制按钮和音量
	playButton := "▶"
	if p.Status.Playing {
		playButton = "⏸"
	}

	controlButtons := fmt.Sprintf("%s ⏮ ⏭ 🔊", playButton)
	volumeText := strconv.Itoa(p.Status.Volume) + "%"

	// 计算控制区域布局
	controlWidth := p.Width - 4
	buttonsWidth := len(controlButtons)
	volumeWidth := len(volumeText) + 22 // 音量条宽度

	if buttonsWidth+volumeWidth <= controlWidth {
		spacerWidth := controlWidth - buttonsWidth - volumeWidth
		spacer := strings.Repeat(" ", spacerWidth)

		controlView := lipgloss.JoinHorizontal(lipgloss.Center,
			controlButtons,
			spacer,
			p.Volume.View(),
			" "+volumeText,
		)
		builder.WriteString(styles.StatusBar.Render(controlView) + "\n")
	}

	// 状态信息
	statusInfo := fmt.Sprintf("循环: %v | 随机: %v", p.Status.Repeat, p.Status.Random)
	builder.WriteString(styles.SubText.Render(statusInfo) + "\n")

	return builder.String()
}

// formatTime 格式化时间（秒）为 MM:SS 格式
func (p *Player) formatTime(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

// SetWidth 设置宽度
func (p *Player) SetWidth(width int) {
	p.Width = width
	p.Progress.Width = width - 4
}

// SetHeight 设置高度
func (p *Player) SetHeight(height int) {
	p.Height = height
}
