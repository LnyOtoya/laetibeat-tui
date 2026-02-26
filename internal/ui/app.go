package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/api"
	"github.com/LnyOtoya/laetibeat-tui/internal/audio/mpv"
	"github.com/LnyOtoya/laetibeat-tui/internal/config"
	"github.com/LnyOtoya/laetibeat-tui/internal/models"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/pages"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// PageType 页面类型
type PageType string

const (
	PageTypeBrowse     PageType = "browse"
	PageTypeNowPlaying PageType = "now_playing"
	PageTypePlaylist   PageType = "playlist"
	PageTypeConfig     PageType = "config"
)

// App 应用模型
type App struct {
	Config         *config.Config
	APIClient      api.APIClient
	MPVClient      *mpv.Client
	BrowsePage     *pages.BrowsePage
	NowPlayingPage *pages.NowPlayingPage
	PlaylistPage   *pages.PlaylistPage
	ConfigPage     *pages.ConfigPage
	CurrentPage    PageType
	Width          int
	Height         int
	IsConnected    bool
}

// NewApp 创建新的应用
func NewApp(cfg *config.Config, apiClient api.APIClient, mpvClient *mpv.Client) *App {
	return &App{
		Config:      cfg,
		APIClient:   apiClient,
		MPVClient:   mpvClient,
		Width:       80,
		Height:      24,
		CurrentPage: PageTypeBrowse,
		IsConnected: false,
	}
}

// Init 初始化应用
func (a *App) Init() tea.Cmd {
	// 初始化页面
	a.BrowsePage = pages.NewBrowsePage(a.APIClient, a.Width, a.Height)
	a.NowPlayingPage = pages.NewNowPlayingPage(a.Width, a.Height)
	a.PlaylistPage = pages.NewPlaylistPage(a.APIClient, a.Width, a.Height)
	a.ConfigPage = pages.NewConfigPage(a.Config, a.Width, a.Height)

	// 连接到服务器
	return a.connectToServer
}

// Update 更新应用
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// 处理窗口大小变化
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		a.Width = msg.Width
		a.Height = msg.Height

		// 更新所有页面大小
		a.BrowsePage.SetSize(msg.Width, msg.Height)
		a.NowPlayingPage.SetSize(msg.Width, msg.Height)
		a.PlaylistPage.SetSize(msg.Width, msg.Height)
		a.ConfigPage.SetSize(msg.Width, msg.Height)
	}

	// 处理键盘事件
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c":
			// 退出应用
			return a, tea.Quit
		case "b":
			// 切换到浏览页面
			a.CurrentPage = PageTypeBrowse
		case "p":
			// 切换到播放列表页面
			a.CurrentPage = PageTypePlaylist
		case "n":
			// 切换到现在播放页面
			a.CurrentPage = PageTypeNowPlaying
		case "c":
			// 切换到配置页面
			a.CurrentPage = PageTypeConfig
		case "space":
			// 播放/暂停
			if a.MPVClient != nil {
				status, err := a.MPVClient.GetStatus()
				if err == nil {
					if status.Playing {
						a.MPVClient.Pause()
					} else {
						a.MPVClient.Resume()
					}
				}
			}
		}
	}

	// 更新当前页面
	switch a.CurrentPage {
	case PageTypeBrowse:
		page, cmd := a.BrowsePage.Update(msg)
		if pageModel, ok := page.(*pages.BrowsePage); ok {
			a.BrowsePage = pageModel
			cmds = append(cmds, cmd)
		}
	case PageTypeNowPlaying:
		page, cmd := a.NowPlayingPage.Update(msg)
		if pageModel, ok := page.(*pages.NowPlayingPage); ok {
			a.NowPlayingPage = pageModel
			cmds = append(cmds, cmd)
		}
	case PageTypePlaylist:
		page, cmd := a.PlaylistPage.Update(msg)
		if pageModel, ok := page.(*pages.PlaylistPage); ok {
			a.PlaylistPage = pageModel
			cmds = append(cmds, cmd)
		}
	case PageTypeConfig:
		page, cmd := a.ConfigPage.Update(msg)
		if pageModel, ok := page.(*pages.ConfigPage); ok {
			a.ConfigPage = pageModel
			cmds = append(cmds, cmd)
		}
	}

	// 处理连接状态更新
	if connectMsg, ok := msg.(ConnectMsg); ok {
		a.IsConnected = connectMsg.Success
		if !connectMsg.Success {
			// 连接失败，显示错误信息
			a.BrowsePage.Status.ShowError(connectMsg.Error)
		}
	}

	// 处理播放状态更新
	if statusMsg, ok := msg.(pages.PlayStatusMsg); ok {
		a.NowPlayingPage.SetStatus(statusMsg.Status)
		a.NowPlayingPage.SetSong(statusMsg.Song)
	}

	// 组合命令
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if len(cmds) > 0 {
		return a, tea.Batch(cmds...)
	}

	return a, nil
}

// View 渲染应用
func (a *App) View() string {
	// 构建页面内容
	var pageView string
	switch a.CurrentPage {
	case PageTypeBrowse:
		pageView = a.BrowsePage.View()
	case PageTypeNowPlaying:
		pageView = a.NowPlayingPage.View()
	case PageTypePlaylist:
		pageView = a.PlaylistPage.View()
	case PageTypeConfig:
		pageView = a.ConfigPage.View()
	}

	// 构建状态栏
	statusBar := a.buildStatusBar()

	// 组合页面和状态栏
	content := lipgloss.JoinVertical(lipgloss.Top, pageView, statusBar)

	// 应用样式
	return styles.MainLayout.Render(content)
}

// buildStatusBar 构建状态栏
func (a *App) buildStatusBar() string {
	// 构建状态栏内容
	statusText := "Laetibeat TUI"

	// 添加连接状态
	if a.IsConnected {
		statusText += " | 已连接"
	} else {
		statusText += " | 未连接"
	}

	// 添加当前页面
	pageText := ""
	switch a.CurrentPage {
	case PageTypeBrowse:
		pageText = "浏览"
	case PageTypeNowPlaying:
		pageText = "现在播放"
	case PageTypePlaylist:
		pageText = "播放列表"
	case PageTypeConfig:
		pageText = "配置"
	}
	statusText += " | 页面: " + pageText

	// 添加快捷键提示
	shortcuts := "B:浏览 N:现在播放 P:播放列表 C:配置 空格:播放/暂停"

	// 布局状态栏
	statusStyle := styles.StatusBar.Copy().Width(a.Width)
	left := statusStyle.Copy().Align(lipgloss.Left).Render(statusText)
	right := statusStyle.Copy().Align(lipgloss.Right).Render(shortcuts)

	return lipgloss.JoinHorizontal(lipgloss.Bottom, left, right)
}

// connectToServer 连接到服务器
func (a *App) connectToServer() tea.Msg {
	// 测试连接
	err := a.APIClient.Ping()
	if err != nil {
		return ConnectMsg{
			Success: false,
			Error:   "连接到服务器失败: " + err.Error(),
		}
	}

	return ConnectMsg{
		Success: true,
		Error:   "",
	}
}

// ConnectMsg 连接状态消息
type ConnectMsg struct {
	Success bool
	Error   string
}

// PlaySong 播放歌曲
func (a *App) PlaySong(song *models.Song) error {
	// 获取歌曲流URL
	streamURL, err := a.APIClient.GetSongStreamURL(song.ID)
	if err != nil {
		return err
	}

	// 播放歌曲
	if err := a.MPVClient.Play(streamURL); err != nil {
		return err
	}

	// 更新现在播放页面
	a.NowPlayingPage.SetSong(song)

	// 切换到现在播放页面
	a.CurrentPage = PageTypeNowPlaying

	return nil
}
