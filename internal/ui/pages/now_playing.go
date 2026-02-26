package pages

import (
	"fmt"
	"strings"

	"github.com/LnyOtoya/laetibeat-tui/internal/models"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/components"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// NowPlayingPage 现在播放页面
type NowPlayingPage struct {
	Player *components.Player
	Status *components.Status
	Song   *models.Song
	Width  int
	Height int
}

// NewNowPlayingPage 创建新的现在播放页面
func NewNowPlayingPage(width, height int) *NowPlayingPage {
	// 创建播放器组件
	player := components.NewPlayer(width-4, 10)

	// 创建状态组件
	status := components.NewStatus(width-4, 1)

	return &NowPlayingPage{
		Player: player,
		Status: status,
		Width:  width,
		Height: height,
	}
}

// Init 初始化页面
func (p *NowPlayingPage) Init() tea.Cmd {
	return nil
}

// Update 更新页面
func (p *NowPlayingPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 处理播放状态更新
	if statusMsg, ok := msg.(PlayStatusMsg); ok {
		p.Player.SetStatus(statusMsg.Status)
		p.Song = statusMsg.Song
	}

	return p, nil
}

// View 渲染页面
func (p *NowPlayingPage) View() string {
	// 构建页面布局
	var content strings.Builder

	// 标题
	title := "现在播放"
	content.WriteString(styles.Title.Render(title) + "\n")

	// 歌曲信息
	if p.Song != nil {
		// 歌曲标题
		songTitle := p.Song.Title
		if len(songTitle) > p.Width-8 {
			songTitle = songTitle[:p.Width-11] + "..."
		}
		content.WriteString(styles.Text.Render("歌曲: "+songTitle) + "\n")

		// 艺术家
		artist := p.Song.Artist
		if len(artist) > p.Width-8 {
			artist = artist[:p.Width-11] + "..."
		}
		content.WriteString(styles.Text.Render("艺术家: "+artist) + "\n")

		// 专辑
		album := p.Song.Album
		if len(album) > p.Width-8 {
			album = album[:p.Width-11] + "..."
		}
		content.WriteString(styles.Text.Render("专辑: "+album) + "\n")

		// 年份和流派
		yearGenre := ""
		if p.Song.Year > 0 {
			yearGenre += fmt.Sprintf("%d", p.Song.Year)
		}
		if p.Song.Genre != "" {
			if yearGenre != "" {
				yearGenre += " · "
			}
			yearGenre += p.Song.Genre
		}
		if yearGenre != "" {
			content.WriteString(styles.SubText.Render(yearGenre) + "\n")
		}
	} else {
		content.WriteString(styles.Text.Render("未播放任何歌曲") + "\n")
	}

	// 空行
	content.WriteString("\n")

	// 播放器组件
	content.WriteString(p.Player.View() + "\n")

	// 状态信息
	statusView := p.Status.View()
	if statusView != "" {
		content.WriteString(statusView + "\n")
	}

	// 帮助信息
	helpText := "←: 返回 ↑↓: 调整 →: 播放/暂停 空格: 播放/暂停"
	content.WriteString("\n" + styles.SubText.Render(helpText))

	// 应用样式
	return styles.ContentArea.Render(content.String())
}

// SetSize 设置页面大小
func (p *NowPlayingPage) SetSize(width, height int) {
	p.Width = width
	p.Height = height

	// 更新组件大小
	p.Player.SetWidth(width - 4)
	p.Status.SetWidth(width - 4)
}

// SetSong 设置当前歌曲
func (p *NowPlayingPage) SetSong(song *models.Song) {
	p.Song = song
}

// SetStatus 设置播放状态
func (p *NowPlayingPage) SetStatus(status models.PlayStatus) {
	p.Player.SetStatus(status)
}

// PlayStatusMsg 播放状态更新消息
type PlayStatusMsg struct {
	Status models.PlayStatus
	Song   *models.Song
}
