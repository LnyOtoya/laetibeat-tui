package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// StatusType 状态类型
type StatusType string

const (
	StatusTypeInfo    StatusType = "info"
	StatusTypeSuccess StatusType = "success"
	StatusTypeWarning StatusType = "warning"
	StatusTypeError   StatusType = "error"
)

// Status 状态组件
type Status struct {
	Type    StatusType
	Message string
	Width   int
	Height  int
	Visible bool
}

// NewStatus 创建新的状态组件
func NewStatus(width, height int) *Status {
	return &Status{
		Type:    StatusTypeInfo,
		Message: "",
		Width:   width,
		Height:  height,
		Visible: false,
	}
}

// ShowInfo 显示信息
func (s *Status) ShowInfo(message string) {
	s.Type = StatusTypeInfo
	s.Message = message
	s.Visible = true
}

// ShowSuccess 显示成功信息
func (s *Status) ShowSuccess(message string) {
	s.Type = StatusTypeSuccess
	s.Message = message
	s.Visible = true
}

// ShowWarning 显示警告信息
func (s *Status) ShowWarning(message string) {
	s.Type = StatusTypeWarning
	s.Message = message
	s.Visible = true
}

// ShowError 显示错误信息
func (s *Status) ShowError(message string) {
	s.Type = StatusTypeError
	s.Message = message
	s.Visible = true
}

// Hide 隐藏状态
func (s *Status) Hide() {
	s.Visible = false
}

// View 渲染状态
func (s *Status) View() string {
	if !s.Visible || s.Message == "" {
		return ""
	}

	// 根据状态类型选择样式
	var statusStyle lipgloss.Style
	switch s.Type {
	case StatusTypeSuccess:
		statusStyle = styles.Success.Copy()
	case StatusTypeWarning:
		statusStyle = lipgloss.NewStyle().Foreground(styles.WarningColor)
	case StatusTypeError:
		statusStyle = styles.Error.Copy()
	default:
		statusStyle = styles.Text.Copy()
	}

	// 格式化消息
	message := s.Message
	if len(message) > s.Width-4 {
		message = message[:s.Width-7] + "..."
	}

	// 居中消息
	padding := (s.Width - len(message)) / 2
	if padding < 0 {
		padding = 0
	}

	paddedMessage := strings.Repeat(" ", padding) + message + strings.Repeat(" ", padding)

	// 确保消息长度正确
	if len(paddedMessage) < s.Width {
		paddedMessage += strings.Repeat(" ", s.Width-len(paddedMessage))
	} else if len(paddedMessage) > s.Width {
		paddedMessage = paddedMessage[:s.Width]
	}

	return statusStyle.Render(paddedMessage)
}

// SetWidth 设置宽度
func (s *Status) SetWidth(width int) {
	s.Width = width
}

// SetHeight 设置高度
func (s *Status) SetHeight(height int) {
	s.Height = height
}
