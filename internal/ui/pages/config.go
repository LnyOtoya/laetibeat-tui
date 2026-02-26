package pages

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/config"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/components"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// ConfigPage 配置页面
type ConfigPage struct {
	Config        *config.Config
	Width         int
	Height        int
	Status        *components.Status
	URLInput      textinput.Model
	UsernameInput textinput.Model
	PasswordInput textinput.Model
	VolumeInput   textinput.Model
	FocusIndex    int
}

// NewConfigPage 创建新的配置页面
func NewConfigPage(cfg *config.Config, width, height int) *ConfigPage {
	// 创建URL输入框
	urlInput := textinput.New()
	urlInput.Placeholder = "http://localhost:4040"
	urlInput.SetValue(cfg.Server.URL)
	urlInput.Focus()

	// 创建用户名输入框
	usernameInput := textinput.New()
	usernameInput.Placeholder = "admin"
	usernameInput.SetValue(cfg.Server.Username)

	// 创建密码输入框
	passwordInput := textinput.New()
	passwordInput.Placeholder = "password"
	passwordInput.SetValue(cfg.Server.Password)
	passwordInput.EchoMode = textinput.EchoPassword

	// 创建音量输入框
	volumeInput := textinput.New()
	volumeInput.Placeholder = "80"
	volumeInput.SetValue(fmt.Sprintf("%d", cfg.Player.Volume))

	return &ConfigPage{
		Config:        cfg,
		Width:         width,
		Height:        height,
		Status:        components.NewStatus(width, 1),
		URLInput:      urlInput,
		UsernameInput: usernameInput,
		PasswordInput: passwordInput,
		VolumeInput:   volumeInput,
		FocusIndex:    0,
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
	return textinput.Blink
}

// Update 更新页面
func (p *ConfigPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			// 处理焦点移动
			switch msg.String() {
			case "enter":
				// 保存配置
				p.saveConfig()
				p.Status.ShowSuccess("配置已保存")
			case "up", "shift+tab":
				// 向上移动焦点
				p.FocusIndex--
			case "down", "tab":
				// 向下移动焦点
				p.FocusIndex++
			}

			// 循环处理焦点索引
			if p.FocusIndex < 0 {
				p.FocusIndex = 3
			} else if p.FocusIndex > 3 {
				p.FocusIndex = 0
			}

			// 设置当前焦点
			p.setFocus(p.FocusIndex)
		}
	}

	// 更新输入框
	switch p.FocusIndex {
	case 0:
		p.URLInput, cmd = p.URLInput.Update(msg)
	case 1:
		p.UsernameInput, cmd = p.UsernameInput.Update(msg)
	case 2:
		p.PasswordInput, cmd = p.PasswordInput.Update(msg)
	case 3:
		p.VolumeInput, cmd = p.VolumeInput.Update(msg)
	}

	return p, cmd
}

// View 渲染页面
func (p *ConfigPage) View() string {
	// 构建配置页面内容
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Title.Render("配置"),
		"",
		styles.Title.Render("服务器配置"),
		"URL: "+p.URLInput.View(),
		"用户名: "+p.UsernameInput.View(),
		"密码: "+p.PasswordInput.View(),
		"",
		styles.Title.Render("播放器配置"),
		"音量: "+p.VolumeInput.View(),
		"交叉淡入: "+boolToString(p.Config.Player.Crossfade),
		"",
		styles.Title.Render("UI配置"),
		"主题: "+p.Config.UI.Theme,
		"显示专辑封面: "+boolToString(p.Config.UI.Artwork),
		"",
		styles.Button.Render("按 Enter 保存配置"),
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

// setFocus 设置当前焦点
func (p *ConfigPage) setFocus(index int) {
	// 重置所有输入框的焦点
	p.URLInput.Blur()
	p.UsernameInput.Blur()
	p.PasswordInput.Blur()
	p.VolumeInput.Blur()

	// 设置当前焦点
	switch index {
	case 0:
		p.URLInput.Focus()
	case 1:
		p.UsernameInput.Focus()
	case 2:
		p.PasswordInput.Focus()
	case 3:
		p.VolumeInput.Focus()
	}
}

// saveConfig 保存配置
func (p *ConfigPage) saveConfig() {
	// 更新配置
	p.Config.Server.URL = p.URLInput.Value()
	p.Config.Server.Username = p.UsernameInput.Value()
	p.Config.Server.Password = p.PasswordInput.Value()

	// 解析音量值
	var volume int
	_, err := fmt.Sscanf(p.VolumeInput.Value(), "%d", &volume)
	if err == nil {
		p.Config.Player.Volume = volume
	}
}

// boolToString 将布尔值转换为字符串
func boolToString(b bool) string {
	if b {
		return "是"
	}
	return "否"
}
