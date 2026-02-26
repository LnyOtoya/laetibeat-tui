package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// 颜色定义
var (
	PrimaryColor   = lipgloss.Color("#6366f1") // Indigo
	SecondaryColor = lipgloss.Color("#8b5cf6") // Violet
	AccentColor    = lipgloss.Color("#ec4899") // Pink
	SuccessColor   = lipgloss.Color("#10b981") // Green
	WarningColor   = lipgloss.Color("#f59e0b") // Amber
	ErrorColor     = lipgloss.Color("#ef4444") // Red
	TextColor      = lipgloss.Color("#f3f4f6") // Gray 100
	SubTextColor   = lipgloss.Color("#9ca3af") // Gray 400
	BgColor        = lipgloss.Color("#1f2937") // Gray 800
	CardBgColor    = lipgloss.Color("#374151") // Gray 700
	BorderColor    = lipgloss.Color("#4b5563") // Gray 600
)

// 基础样式
var (
	// 文本样式
	Text = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(BgColor)

	SubText = lipgloss.NewStyle().
		Foreground(SubTextColor).
		Background(BgColor)

	// 标题样式
	Title = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Background(BgColor).
		Bold(true).
		Padding(0, 2)

	// 卡片样式
	Card = lipgloss.NewStyle().
		Background(CardBgColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderColor).
		Padding(1, 2)

	// 按钮样式
	Button = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(PrimaryColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(0, 2).
		Margin(0, 1)

	ButtonHover = Button.Copy().
		Background(SecondaryColor).
		BorderForeground(SecondaryColor)

	// 输入框样式
	Input = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(CardBgColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderColor).
		Padding(0, 2)

	// 进度条样式
	ProgressBar = lipgloss.NewStyle().
		Background(CardBgColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderColor).
		Padding(0, 1)

	ProgressFill = lipgloss.NewStyle().
		Background(PrimaryColor)

	// 列表样式
	List = lipgloss.NewStyle().
		Background(BgColor)

	ListItem = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(BgColor).
		Padding(0, 2)

	ListSelectedItem = ListItem.Copy().
		Background(CardBgColor).
		Bold(true)

	// 状态样式
	StatusBar = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(CardBgColor).
		Padding(0, 2)

	// 错误样式
	Error = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Background(BgColor).
		Bold(true)

	// 成功样式
	Success = lipgloss.NewStyle().
		Foreground(SuccessColor).
		Background(BgColor).
		Bold(true)
)

// 布局样式
var (
	// 主布局
	MainLayout = lipgloss.NewStyle().
		Background(BgColor).
		Padding(1, 2)

	// 内容区域
	ContentArea = lipgloss.NewStyle().
		Background(BgColor)

	// 侧边栏
	Sidebar = lipgloss.NewStyle().
		Background(CardBgColor).
		Border(lipgloss.Border{
			Right: "│",
		},
		).
		BorderForeground(BorderColor).
		Padding(1, 2)

	// 底部栏
	BottomBar = lipgloss.NewStyle().
		Background(CardBgColor).
		Border(lipgloss.Border{
			Top: "─",
		},
		).
		BorderForeground(BorderColor).
		Padding(1, 2)
)

// 响应式布局
func GetMainLayout(width, height int) lipgloss.Style {
	return MainLayout.Copy().
		Width(width).
		Height(height)
}

func GetSidebarLayout(width int) lipgloss.Style {
	return Sidebar.Copy().
		Width(width)
}

func GetContentLayout(width, height int) lipgloss.Style {
	return ContentArea.Copy().
		Width(width).
		Height(height)
}

func GetBottomBarLayout(width int) lipgloss.Style {
	return BottomBar.Copy().
		Width(width)
}
