package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/LnyOtoya/laetibeat-tui/internal/api"
	"github.com/LnyOtoya/laetibeat-tui/internal/models"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/components"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// PlaylistPage 播放列表页面
type PlaylistPage struct {
	API    api.APIClient
	List   *components.List
	Status *components.Status
	Width  int
	Height int
}

// NewPlaylistPage 创建新的播放列表页面
func NewPlaylistPage(apiClient api.APIClient, width, height int) *PlaylistPage {
	// 创建状态组件
	status := components.NewStatus(width, 1)

	// 创建空列表
	list := components.NewList("播放列表", []list.Item{}, width-4, height-8)

	return &PlaylistPage{
		API:    apiClient,
		List:   list,
		Status: status,
		Width:  width,
		Height: height,
	}
}

// Init 初始化页面
func (p *PlaylistPage) Init() tea.Cmd {
	// 加载播放列表
	return p.loadPlaylists
}

// Update 更新页面
func (p *PlaylistPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// 更新列表组件
	// 注意：我们不能直接修改List的内部状态，因为list字段是未导出的
	// 这里需要通过其他方式更新List，或者修改List组件的设计

	// 处理列表选择
	// 注意：在新版本的bubbles/list中，SelectMsg可能已经不存在
	// 我们需要通过其他方式处理列表选择，例如通过键盘事件

	// 处理加载完成消息
	switch msg := msg.(type) {
	case playlistsLoadedMsg:
		p.handlePlaylistsLoaded(msg)
	case playlistSongsLoadedMsg:
		p.handlePlaylistSongsLoaded(msg)
	case errorMsg:
		p.Status.ShowError(msg.error.Error())
	}

	// 组合命令
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if len(cmds) > 0 {
		return p, tea.Batch(cmds...)
	}

	return p, nil
}

// View 渲染页面
func (p *PlaylistPage) View() string {
	// 构建页面布局
	var content strings.Builder

	// 标题
	title := "播放列表"
	content.WriteString(styles.Title.Render(title) + "\n")

	// 状态信息
	statusView := p.Status.View()
	if statusView != "" {
		content.WriteString(statusView + "\n")
	}

	// 列表
	content.WriteString(p.List.View())

	// 帮助信息
	helpText := "↑↓: 选择 ↵: 查看/播放 Ctrl+C: 退出"
	content.WriteString("\n" + styles.SubText.Render(helpText))

	// 应用样式
	return styles.ContentArea.Render(content.String())
}

// SetSize 设置页面大小
func (p *PlaylistPage) SetSize(width, height int) {
	p.Width = width
	p.Height = height

	// 更新组件大小
	p.List = components.NewList(p.List.Model().Title, p.List.Model().Items(), width-4, height-8)
	p.Status.SetWidth(width - 4)
}

// 消息类型

// playlistsLoadedMsg 播放列表加载完成消息
type playlistsLoadedMsg struct {
	playlists []models.Playlist
}

// playlistSongsLoadedMsg 播放列表歌曲加载完成消息
type playlistSongsLoadedMsg struct {
	playlistID string
	songs      []models.Song
}

// 列表项类型

// playlistItem 播放列表列表项
type playlistItem struct {
	ID        string
	Name      string
	SongCount int
	Duration  int
}

// Title 获取标题
func (i playlistItem) Title() string {
	return i.Name
}

// Description 获取描述
func (i playlistItem) Description() string {
	duration := i.formatDuration(i.Duration)
	return fmt.Sprintf("%d 首歌曲 · %s", i.SongCount, duration)
}

// FilterValue 获取过滤值
func (i playlistItem) FilterValue() string {
	return i.Name
}

// formatDuration 格式化时长
func (i playlistItem) formatDuration(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

// 加载数据的命令

// loadPlaylists 加载播放列表
func (p *PlaylistPage) loadPlaylists() tea.Msg {
	// 这里应该调用API获取播放列表
	// 暂时返回空列表
	return playlistsLoadedMsg{playlists: []models.Playlist{}}
}

// loadPlaylistSongs 加载播放列表歌曲
func (p *PlaylistPage) loadPlaylistSongs(playlistID string) tea.Cmd {
	return func() tea.Msg {
		// 这里应该调用API获取播放列表歌曲
		// 暂时返回空列表
		return playlistSongsLoadedMsg{
			playlistID: playlistID,
			songs:      []models.Song{},
		}
	}
}

// 处理加载完成的方法

// handlePlaylistsLoaded 处理播放列表加载完成
func (p *PlaylistPage) handlePlaylistsLoaded(msg playlistsLoadedMsg) {
	// 转换为列表项
	var items []list.Item
	for _, playlist := range msg.playlists {
		items = append(items, playlistItem{
			ID:        playlist.ID,
			Name:      playlist.Name,
			SongCount: playlist.SongCount,
			Duration:  playlist.Duration,
		})
	}

	// 更新列表
	p.List = components.NewList("播放列表", items, p.Width-4, p.Height-8)
	p.Status.Hide()
}

// handlePlaylistSongsLoaded 处理播放列表歌曲加载完成
func (p *PlaylistPage) handlePlaylistSongsLoaded(msg playlistSongsLoadedMsg) {
	// 这里应该显示播放列表歌曲
	// 暂时只显示状态
	p.Status.ShowSuccess("播放列表加载完成")
}
