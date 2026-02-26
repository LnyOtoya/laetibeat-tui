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

// BrowsePage 浏览页面
type BrowsePage struct {
	API       api.APIClient
	List      *components.List
	Status    *components.Status
	Search    *components.Search
	Width     int
	Height    int
	ViewMode  ViewMode
	CurrentID string
}

// ViewMode 浏览模式
type ViewMode string

const (
	ViewModeArtists ViewMode = "artists"
	ViewModeAlbums  ViewMode = "albums"
	ViewModeSongs   ViewMode = "songs"
)

// NewBrowsePage 创建新的浏览页面
func NewBrowsePage(apiClient api.APIClient, width, height int) *BrowsePage {
	// 创建状态组件
	status := components.NewStatus(width, 1)

	// 创建搜索组件
	search := components.NewSearch(width-4, 1, func(query string) {
		// 搜索逻辑将在Update中处理
	})

	// 创建空列表
	list := components.NewList("艺术家", []list.Item{}, width-4, height-10)

	return &BrowsePage{
		API:       apiClient,
		List:      list,
		Status:    status,
		Search:    search,
		Width:     width,
		Height:    height,
		ViewMode:  ViewModeArtists,
		CurrentID: "",
	}
}

// Init 初始化页面
func (p *BrowsePage) Init() tea.Cmd {
	// 加载艺术家列表
	return p.loadArtists
}

// Update 更新页面
func (p *BrowsePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// 处理搜索消息
	if searchMsg, ok := msg.(tea.KeyMsg); ok {
		if searchMsg.Type == tea.KeyEnter {
			query := p.Search.Model().Value()
			if query != "" {
				p.Status.ShowInfo("正在搜索: " + query)
				cmds = append(cmds, p.search(query))
			}
		}
	}

	// 更新搜索组件
	searchModel, err := p.Search.Update(msg)
	if err == nil {
		p.Search = &components.Search{
			Input:    searchModel,
			Width:    p.Search.Width,
			Height:   p.Search.Height,
			OnSearch: p.Search.OnSearch,
		}
	}

	// 更新列表组件
	// 注意：我们不能直接修改List的内部状态，因为list字段是未导出的
	// 这里需要通过其他方式更新List，或者修改List组件的设计

	// 处理列表选择
	// 注意：在新版本的bubbles/list中，SelectMsg可能已经不存在
	// 我们需要通过其他方式处理列表选择，例如通过键盘事件

	// 处理加载完成消息
	switch msg := msg.(type) {
	case artistsLoadedMsg:
		p.handleArtistsLoaded(msg)
	case albumsLoadedMsg:
		p.handleAlbumsLoaded(msg)
	case songsLoadedMsg:
		p.handleSongsLoaded(msg)
	case searchResultMsg:
		p.handleSearchResult(msg)
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
func (p *BrowsePage) View() string {
	// 构建页面布局
	var content strings.Builder

	// 标题
	title := "浏览音乐库"
	content.WriteString(styles.Title.Render(title) + "\n")

	// 搜索栏
	content.WriteString(p.Search.View() + "\n")

	// 状态信息
	statusView := p.Status.View()
	if statusView != "" {
		content.WriteString(statusView + "\n")
	}

	// 列表
	content.WriteString(p.List.View())

	// 帮助信息
	helpText := "↑↓: 选择 ↵: 进入/播放 Ctrl+C: 退出"
	content.WriteString("\n" + styles.SubText.Render(helpText))

	// 应用样式
	return styles.ContentArea.Render(content.String())
}

// SetSize 设置页面大小
func (p *BrowsePage) SetSize(width, height int) {
	p.Width = width
	p.Height = height

	// 更新组件大小
	p.Search.SetWidth(width - 4)
	p.List = components.NewList(p.List.Model().Title, p.List.Model().Items(), width-4, height-10)
	p.Status.SetWidth(width - 4)
}

// handleListSelect 处理列表选择
// 注意：在新版本的bubbles/list中，SelectMsg可能已经不存在
// 这个方法暂时保留，但需要重新实现
func (p *BrowsePage) handleListSelect() {
	// 获取当前选中的项
	item, ok := p.List.SelectedItem()
	if !ok {
		return
	}

	switch p.ViewMode {
	case ViewModeArtists:
		// 选择艺术家，加载其专辑
		artist, ok := item.(artistItem)
		if ok {
			p.CurrentID = artist.ID
			p.ViewMode = ViewModeAlbums
			p.Status.ShowInfo("正在加载专辑...")
			p.loadAlbums(artist.ID)
		}
	case ViewModeAlbums:
		// 选择专辑，加载其歌曲
		album, ok := item.(albumItem)
		if ok {
			p.CurrentID = album.ID
			p.ViewMode = ViewModeSongs
			p.Status.ShowInfo("正在加载歌曲...")
			p.loadSongs(album.ID)
		}
	case ViewModeSongs:
		// 选择歌曲，播放
		song, ok := item.(songItem)
		if ok {
			p.Status.ShowInfo("正在播放: " + song.TitleName)
			// 播放逻辑将在主应用中处理
		}
	}
}

// 消息类型

// artistsLoadedMsg 艺术家加载完成消息
type artistsLoadedMsg struct {
	artists []models.Artist
}

// albumsLoadedMsg 专辑加载完成消息
type albumsLoadedMsg struct {
	albums []models.Album
}

// songsLoadedMsg 歌曲加载完成消息
type songsLoadedMsg struct {
	songs []models.Song
}

// searchResultMsg 搜索结果消息
type searchResultMsg struct {
	result *models.SearchResult
}

// errorMsg 错误消息
type errorMsg struct {
	error error
}

// 列表项类型

// artistItem 艺术家列表项
type artistItem struct {
	ID   string
	Name string
}

// Title 获取标题
func (i artistItem) Title() string {
	return i.Name
}

// Description 获取描述
func (i artistItem) Description() string {
	return ""
}

// FilterValue 获取过滤值
func (i artistItem) FilterValue() string {
	return i.Name
}

// albumItem 专辑列表项
type albumItem struct {
	ID        string
	TitleName string
	Artist    string
	Year      int
}

// Title 获取标题
func (i albumItem) Title() string {
	return i.TitleName
}

// Description 获取描述
func (i albumItem) Description() string {
	if i.Year > 0 {
		return i.Artist + " (" + fmt.Sprintf("%d", i.Year) + ")"
	}
	return i.Artist
}

// FilterValue 获取过滤值
func (i albumItem) FilterValue() string {
	return i.TitleName + " " + i.Artist
}

// songItem 歌曲列表项
type songItem struct {
	ID        string
	TitleName string
	Artist    string
	Album     string
	Track     int
	Duration  int
}

// Title 获取标题
func (i songItem) Title() string {
	if i.Track > 0 {
		return fmt.Sprintf("%d. %s", i.Track, i.TitleName)
	}
	return i.TitleName
}

// Description 获取描述
func (i songItem) Description() string {
	duration := i.formatDuration(i.Duration)
	return i.Artist + " - " + i.Album + " (" + duration + ")"
}

// FilterValue 获取过滤值
func (i songItem) FilterValue() string {
	return i.TitleName + " " + i.Artist + " " + i.Album
}

// formatDuration 格式化时长
func (i songItem) formatDuration(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

// 加载数据的命令

// loadArtists 加载艺术家列表
func (p *BrowsePage) loadArtists() tea.Msg {
	artists, err := p.API.GetArtists()
	if err != nil {
		return errorMsg{error: err}
	}
	return artistsLoadedMsg{artists: artists}
}

// loadAlbums 加载专辑列表
func (p *BrowsePage) loadAlbums(artistID string) tea.Cmd {
	return func() tea.Msg {
		albums, err := p.API.GetAlbums(artistID)
		if err != nil {
			return errorMsg{error: err}
		}
		return albumsLoadedMsg{albums: albums}
	}
}

// loadSongs 加载歌曲列表
func (p *BrowsePage) loadSongs(albumID string) tea.Cmd {
	return func() tea.Msg {
		songs, err := p.API.GetSongs(albumID)
		if err != nil {
			return errorMsg{error: err}
		}
		return songsLoadedMsg{songs: songs}
	}
}

// search 搜索
func (p *BrowsePage) search(query string) tea.Cmd {
	return func() tea.Msg {
		result, err := p.API.Search(query, 20, 20, 20)
		if err != nil {
			return errorMsg{error: err}
		}
		return searchResultMsg{result: result}
	}
}

// 处理加载完成的方法

// handleArtistsLoaded 处理艺术家加载完成
func (p *BrowsePage) handleArtistsLoaded(msg artistsLoadedMsg) {
	// 转换为列表项
	var items []list.Item
	for _, artist := range msg.artists {
		items = append(items, artistItem{
			ID:   artist.ID,
			Name: artist.Name,
		})
	}

	// 更新列表
	p.List = components.NewList("艺术家", items, p.Width-4, p.Height-10)
	p.Status.Hide()
}

// handleAlbumsLoaded 处理专辑加载完成
func (p *BrowsePage) handleAlbumsLoaded(msg albumsLoadedMsg) {
	// 转换为列表项
	var items []list.Item
	for _, album := range msg.albums {
		items = append(items, albumItem{
			ID:        album.ID,
			TitleName: album.Title,
			Artist:    album.Artist,
			Year:      album.Year,
		})
	}

	// 更新列表
	p.List = components.NewList("专辑", items, p.Width-4, p.Height-10)
	p.Status.Hide()
}

// handleSongsLoaded 处理歌曲加载完成
func (p *BrowsePage) handleSongsLoaded(msg songsLoadedMsg) {
	// 转换为列表项
	var items []list.Item
	for _, song := range msg.songs {
		items = append(items, songItem{
			ID:        song.ID,
			TitleName: song.Title,
			Artist:    song.Artist,
			Album:     song.Album,
			Track:     song.Track,
			Duration:  song.Duration,
		})
	}

	// 更新列表
	p.List = components.NewList("歌曲", items, p.Width-4, p.Height-10)
	p.Status.Hide()
}

// handleSearchResult 处理搜索结果
func (p *BrowsePage) handleSearchResult(msg searchResultMsg) {
	// 转换为列表项
	var items []list.Item

	// 添加艺术家
	for _, artist := range msg.result.Artists {
		items = append(items, artistItem{
			ID:   artist.ID,
			Name: "艺术家: " + artist.Name,
		})
	}

	// 添加专辑
	for _, album := range msg.result.Albums {
		items = append(items, albumItem{
			ID:        album.ID,
			TitleName: "专辑: " + album.Title,
			Artist:    album.Artist,
			Year:      album.Year,
		})
	}

	// 添加歌曲
	for _, song := range msg.result.Songs {
		items = append(items, songItem{
			ID:        song.ID,
			TitleName: "歌曲: " + song.Title,
			Artist:    song.Artist,
			Album:     song.Album,
			Track:     song.Track,
			Duration:  song.Duration,
		})
	}

	// 更新列表
	p.List = components.NewList("搜索结果", items, p.Width-4, p.Height-10)
	p.Status.Hide()
}
