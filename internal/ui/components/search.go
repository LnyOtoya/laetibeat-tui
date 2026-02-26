package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// Search 搜索组件
type Search struct {
	Input    textinput.Model
	Width    int
	Height   int
	OnSearch func(query string)
}

// NewSearch 创建新的搜索组件
func NewSearch(width, height int, onSearch func(query string)) *Search {
	// 初始化文本输入
	input := textinput.New()
	input.Placeholder = "搜索艺术家、专辑或歌曲..."
	input.Prompt = "🔍 "
	input.Focus()

	// 样式设置在新版本中可能不同，暂时移除

	return &Search{
		Input:    input,
		Width:    width,
		Height:   height,
		OnSearch: onSearch,
	}
}

// Update 更新搜索组件
func (s *Search) Update(msg interface{}) (textinput.Model, tea.Cmd) {
	return s.Input.Update(msg)
}

// View 渲染搜索组件
func (s *Search) View() string {
	return styles.Input.Render(s.Input.View())
}

// Model 获取底层的textinput.Model
func (s *Search) Model() textinput.Model {
	return s.Input
}

// SetWidth 设置宽度
func (s *Search) SetWidth(width int) {
	s.Width = width
	s.Input.Width = width - 8 // 减去边距和提示
}

// SetHeight 设置高度
func (s *Search) SetHeight(height int) {
	s.Height = height
}

// Submit 提交搜索
func (s *Search) Submit() {
	if s.OnSearch != nil {
		s.OnSearch(s.Input.Value())
	}
}

// Clear 清除搜索
func (s *Search) Clear() {
	s.Input.SetValue("")
	s.Input.Focus()
}
