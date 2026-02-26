package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/config"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/components"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// ConfigPage 配置页面
type ConfigPage struct {
	Config *config.Config
	Width  int
	Height int
	Status *components.Status
}

// NewConfigPage 创建新的配置页面
func NewConfigPage(cfg *config.Config, width, height int) *ConfigPage {
	return &ConfigPage{
		Config: cfg,
		Width:  width,
		Height: height,
		Status: components.NewStatus(width, 1),
	}
}

// SetSize 设置页面大小
func (p *ConfigPage) SetSize(width, height int) {
	p.Width = width
	p.Height = height
	p.Status.Width = width
}

// Init 初始化页面
func (p *ConfigPage) Init() tea.Cmd {
	return nil
}

// Update 更新页面
func (p *ConfigPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

// View 渲染页面
func (p *ConfigPage) View() string {
	// 构建配置页面内容
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Title.Render("配置"),
		styles.Title.Render("服务器配置"),
		"URL: "+p.Config.Server.URL,
		"用户名: "+p.Config.Server.Username,
		"密码: [隐藏]",
		"",
		styles.Title.Render("播放器配置"),
		"音量: "+fmt.Sprintf("%d", p.Config.Player.Volume),
		"交叉淡入: "+boolToString(p.Config.Player.Crossfade),
		"",
		styles.Title.Render("UI配置"),
		"主题: "+p.Config.UI.Theme,
		"显示专辑封面: "+boolToString(p.Config.UI.Artwork),
		"",
		"按 'c' 返回主页面",
	)

	// 居中显示
	content = lipgloss.Place(
		p.Width,
		p.Height-1,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

	// 添加状态条
	statusBar := p.Status.View()

	// 组合内容和状态条
	return lipgloss.JoinVertical(lipgloss.Top, content, statusBar)
}

// boolToString 将布尔值转换为字符串
func boolToString(b bool) string {
	if b {
		return "是"
	}
	return "否"
}
