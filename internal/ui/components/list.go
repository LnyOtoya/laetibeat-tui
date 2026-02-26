package components

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/LnyOtoya/laetibeat-tui/internal/ui/styles"
)

// List 自定义列表组件
type List struct {
	list list.Model
}

// NewList 创建新的列表组件
func NewList(title string, items []list.Item, width, height int) *List {
	// 创建列表项
	delegate := list.NewDefaultDelegate()

	// 自定义列表项样式
	delegate.Styles.NormalTitle = styles.ListItem.Copy()
	delegate.Styles.NormalDesc = styles.SubText.Copy()
	delegate.Styles.SelectedTitle = styles.ListSelectedItem.Copy()
	delegate.Styles.SelectedDesc = styles.SubText.Copy().
		Background(styles.CardBgColor)

	// 创建列表
	listModel := list.New(items, delegate, width, height)
	listModel.Title = title
	listModel.Styles.Title = styles.Title.Copy()
	listModel.Styles.PaginationStyle = styles.SubText.Copy()
	listModel.Styles.HelpStyle = styles.SubText.Copy()

	// 设置过滤
	listModel.SetFilteringEnabled(true)

	return &List{
		list: listModel,
	}
}

// Model 获取底层的list.Model
func (l *List) Model() list.Model {
	return l.list
}

// Update 更新列表
func (l *List) Update(msg interface{}) (*List, tea.Cmd) {
	updatedListModel, cmd := l.list.Update(msg)
	return &List{list: updatedListModel}, cmd
}

// View 渲染列表
func (l *List) View() string {
	return l.list.View()
}

// SetItems 设置列表项
func (l *List) SetItems(items []list.Item) {
	l.list.SetItems(items)
}

// AddItem 添加列表项
func (l *List) AddItem(item list.Item) {
	items := l.list.Items()
	items = append(items, item)
	l.list.SetItems(items)
}

// RemoveItem 移除列表项
func (l *List) RemoveItem(index int) {
	items := l.list.Items()
	if index >= 0 && index < len(items) {
		items = append(items[:index], items[index+1:]...)
		l.list.SetItems(items)
	}
}

// SelectedItem 获取选中的列表项
func (l *List) SelectedItem() (list.Item, bool) {
	if len(l.list.Items()) == 0 {
		return nil, false
	}
	index := l.list.Index()
	if index < 0 || index >= len(l.list.Items()) {
		return nil, false
	}
	return l.list.Items()[index], true
}

// SelectedIndex 获取选中的列表项索引
func (l *List) SelectedIndex() int {
	return l.list.Index()
}

// SetWidth 设置列表宽度
func (l *List) SetWidth(width int) {
	l.list.SetWidth(width)
}

// SetHeight 设置列表高度
func (l *List) SetHeight(height int) {
	l.list.SetHeight(height)
}

// SetTitle 设置列表标题
func (l *List) SetTitle(title string) {
	l.list.Title = title
}

// Filter 设置过滤文本
func (l *List) Filter(filter string) {
	// 过滤功能由列表组件自动处理
}

// ClearFilter 清除过滤
func (l *List) ClearFilter() {
	l.list.ResetFilter()
}

// SetShowHelp 设置是否显示帮助
func (l *List) SetShowHelp(show bool) {
	l.list.SetShowHelp(show)
}

// SetShowPagination 设置是否显示分页
func (l *List) SetShowPagination(show bool) {
	l.list.SetShowPagination(show)
}

// SetShowStatusBar 设置是否显示状态栏
func (l *List) SetShowStatusBar(show bool) {
	l.list.SetShowStatusBar(show)
}

// Item 列表项接口
type Item interface {
	Title() string
	Description() string
	FilterValue() string
}

// DefaultItem 默认列表项
type DefaultItem struct {
	title       string
	description string
	filterValue string
}

// NewDefaultItem 创建默认列表项
func NewDefaultItem(title, description string) DefaultItem {
	return DefaultItem{
		title:       title,
		description: description,
		filterValue: title + " " + description,
	}
}

// Title 获取标题
func (i DefaultItem) Title() string {
	return i.title
}

// Description 获取描述
func (i DefaultItem) Description() string {
	return i.description
}

// FilterValue 获取过滤值
func (i DefaultItem) FilterValue() string {
	return i.filterValue
}
