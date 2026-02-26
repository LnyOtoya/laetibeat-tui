package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/api"
	"github.com/LnyOtoya/laetibeat-tui/internal/audio/mpv"
	"github.com/LnyOtoya/laetibeat-tui/internal/config"
	"github.com/LnyOtoya/laetibeat-tui/internal/models"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/messages"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/pages"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// PageType 页面类型
type PageType string

const (
	PageTypeConfig PageType = "config"
)

// App 应用模型
type App struct {
	Config      *config.Config
	APIClient   api.APIClient
	MPVClient   *mpv.Client
	ConfigPage  *pages.ConfigPage
	CurrentPage PageType
	Width       int
	Height      int
	IsConnected bool
}

// NewApp 创建新的应用
func NewApp(cfg *config.Config, apiClient api.APIClient, mpvClient *mpv.Client) *App {
	return &App{
		Config:      cfg,
		APIClient:   apiClient,
		MPVClient:   mpvClient,
		Width:       80,
		Height:      24,
		CurrentPage: PageTypeConfig,
		IsConnected: false,
	}
}

// Init 初始化应用
func (a *App) Init() tea.Cmd {
	// 初始化配置页面
	a.ConfigPage = pages.NewConfigPage(a.Config, a.Width, a.Height)

	// 连接到服务器
	connectCmd := a.connectToServer

	return connectCmd
}

// Update 更新应用
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// 处理窗口大小变化
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		a.Width = msg.Width
		a.Height = msg.Height

		// 更新配置页面大小
		a.ConfigPage.SetSize(msg.Width, msg.Height)
	}

	// 处理键盘事件
	if msg, ok := msg.(tea.KeyMsg); ok {
		// 只处理ctrl+c退出
		if msg.String() == "ctrl+c" {
			return a, tea.Quit
		}
	}

	// 处理切换页面消息
	if switchMsg, ok := msg.(messages.SwitchPageMsg); ok {
		// 只处理配置页面
		if switchMsg.Page == "config" {
			a.CurrentPage = PageTypeConfig
		}
		return a, nil
	}

	// 更新当前页面
	if a.CurrentPage == PageTypeConfig {
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
			a.ConfigPage.Status.ShowError(connectMsg.Error)
		}
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
	if a.CurrentPage == PageTypeConfig {
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
	if a.CurrentPage == PageTypeConfig {
		pageText = "配置"
	}
	statusText += " | 页面: " + pageText

	// 添加快捷键提示
	shortcuts := "C:配置 ESC:返回 Ctrl+C:退出"

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

	return nil
}
