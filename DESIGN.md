# Laetibeat TUI - OpenSubsonic 音乐播放器设计方案

## 1. 项目概述

Laetibeat TUI 是一个基于 Go 语言和 Bubble Tea 框架的终端音乐播放器，支持连接 OpenSubsonic 兼容的音乐服务器。

### 核心特性
- 支持 OpenSubsonic API 认证和交互
- 多种音频格式支持（WAV、FLAC、MP3、OGG 等）
- 基于 mpv 的高质量音频播放
- 美观的终端用户界面
- 响应式设计，适应不同终端大小
- 支持键盘快捷键和鼠标操作

## 2. 技术栈

### 核心框架
- **Go 1.25+**：主要开发语言
- **Bubble Tea** (`github.com/charmbracelet/bubbletea`)：TUI 框架
- **Lip Gloss** (`github.com/charmbracelet/lipgloss`)：终端样式库
- **Bubbles** (`github.com/charmbracelet/bubbles`)：UI 组件库

### 音频后端
- **mpv**：外部媒体播放器，通过 IPC 控制

### 网络与 API
- **标准库**：`net/http`、`encoding/json`、`net/url`
- **OpenSubsonic API**：自定义客户端实现

### 工具库
- **Viper** (`github.com/spf13/viper`)：配置管理
- **UUID** (`github.com/google/uuid`)：生成唯一标识符
- **Errors** (`github.com/pkg/errors`)：增强错误处理

## 3. 目录结构

```
laetibeat-tui/
├── cmd/
│   └── laetibeat/
│       └── main.go           # 应用入口
├── internal/
│   ├── api/                  # OpenSubsonic API 客户端
│   │   ├── client.go         # 客户端主逻辑
│   │   ├── auth.go           # 认证相关
│   │   ├── endpoints.go      # API 端点封装
│   │   └── response.go       # 响应处理
│   ├── audio/                # 音频播放相关
│   │   ├── mpv/              # mpv 控制
│   │   │   ├── client.go     # mpv IPC 客户端
│   │   │   ├── commands.go   # 命令封装
│   │   │   └── events.go     # 事件处理
│   │   ├── player.go         # 播放器接口
│   │   └── stream.go         # 流处理
│   ├── config/               # 配置管理
│   │   ├── config.go         # 配置结构
│   │   └── storage.go        # 配置存储
│   ├── models/               # 数据模型
│   │   ├── artist.go         # 艺术家模型
│   │   ├── album.go          # 专辑模型
│   │   ├── song.go           # 歌曲模型
│   │   ├── playlist.go       # 播放列表模型
│   │   └── response.go       # API 响应模型
│   ├── ui/                   # TUI 界面
│   │   ├── components/       # 可复用组件
│   │   │   ├── list.go       # 增强列表组件
│   │   │   ├── player.go     # 播放器控制组件
│   │   │   ├── search.go     # 搜索组件
│   │   │   └── status.go     # 状态显示组件
│   │   ├── pages/            # 页面
│   │   │   ├── browse.go     # 浏览页面
│   │   │   ├── playlist.go   # 播放列表页面
│   │   │   └── now_playing.go # 现在播放页面
│   │   ├── styles/           # 样式定义
│   │   │   └── styles.go     # 全局样式
│   │   ├── app.go            # 应用状态管理
│   │   └── messages.go       # 消息定义
│   └── utils/                # 工具函数
│       ├── errors.go         # 错误处理
│       ├── http.go           # HTTP 工具
│       └── strings.go        # 字符串工具
├── pkg/                      # 可导出包
│   └── logger/               # 日志工具
├── go.mod                    # Go 模块文件
├── go.sum                    # 依赖校验文件
├── LICENSE                   # 许可证
├── README.md                 # 项目说明
└── DESIGN.md                 # 设计文档
```

## 4. 模块设计与职责

### 4.1 核心模块

#### API 模块 (`internal/api/`)
- **职责**：处理与 OpenSubsonic 服务器的所有通信
- **功能**：
  - 服务器认证（MD5 哈希认证）
  - API 端点调用
  - 响应解析和错误处理
  - 数据缓存管理
- **设计**：
  - 基于接口设计，便于测试和替换
  - 支持多种 API 版本
  - 实现请求重试和错误处理

#### 音频模块 (`internal/audio/`)
- **职责**：处理音频播放和控制
- **功能**：
  - mpv 进程管理
  - IPC 通信
  - 播放控制（播放/暂停、上一曲/下一曲等）
  - 音量控制
  - 播放状态监控
- **设计**：
  - 基于接口设计，便于替换底层播放器
  - 支持多种播放模式（单曲循环、列表循环等）
  - 实现错误处理和重连机制

#### 配置模块 (`internal/config/`)
- **职责**：管理应用配置
- **功能**：
  - 配置文件读写
  - 默认配置管理
  - 环境变量支持
- **设计**：
  - 使用 Viper 库实现
  - 支持多种配置格式（JSON、YAML 等）
  - 配置验证

#### 模型模块 (`internal/models/`)
- **职责**：定义数据结构
- **功能**：
  - 数据模型定义
  - 模型转换
  - 模型验证
- **设计**：
  - 清晰的结构定义
  - 支持 JSON 序列化/反序列化
  - 包含必要的方法

#### UI 模块 (`internal/ui/`)
- **职责**：处理用户界面和交互
- **功能**：
  - 页面管理
  - 组件渲染
  - 用户输入处理
  - 状态更新
- **设计**：
  - 基于 Bubble Tea 的模型-视图-更新模式
  - 组件化设计，便于复用
  - 响应式布局

### 4.2 接口设计

#### API 客户端接口
```go
type APIClient interface {
    Ping() error
    GetMusicFolders() ([]models.MusicFolder, error)
    GetArtists() ([]models.Artist, error)
    GetAlbums(artistID string) ([]models.Album, error)
    GetSongs(albumID string) ([]models.Song, error)
    GetSongStreamURL(songID string) (string, error)
    // 其他 API 方法...
}
```

#### 播放器接口
```go
type Player interface {
    Play(url string) error
    Pause() error
    Resume() error
    Stop() error
    Next() error
    Previous() error
    SetVolume(volume int) error
    Seek(position int) error
    GetStatus() (models.PlayStatus, error)
    Close() error
}
```

#### UI 组件接口
```go
type Component interface {
    Init() tea.Cmd
    Update(msg tea.Msg) (tea.Model, tea.Cmd)
    View() string
}
```

## 5. 数据流设计

### 5.1 主要数据流

1. **认证流程**
   - 用户输入服务器信息 → 配置存储 → API 客户端初始化 → 认证请求 → 认证结果 → UI 更新

2. **音乐浏览流程**
   - 用户界面操作 → API 调用请求 → API 客户端处理 → 数据模型更新 → 界面渲染

3. **播放控制流程**
   - 用户界面操作 → 播放命令 → 音频模块处理 → mpv 操作 → 播放状态更新 → 界面渲染

4. **搜索流程**
   - 用户输入搜索关键词 → 搜索请求 → API 客户端处理 → 搜索结果 → 界面渲染

### 5.2 状态管理

使用 Bubble Tea 的模型系统管理应用状态：

- **AppModel**：主应用模型，管理全局状态
- **BrowseModel**：浏览页面模型
- **PlayerModel**：播放器模型
- **SearchModel**：搜索模型
- **ConfigModel**：配置模型

状态更新通过消息（Msg）机制实现，确保单向数据流，提高可预测性和可测试性。

## 6. 依赖注入与解耦

### 6.1 依赖注入

使用构造函数注入依赖，避免硬编码依赖关系：

```go
// 示例：创建应用实例
func NewApp(apiClient api.APIClient, player audio.Player, config *config.Config) *AppModel {
    return &AppModel{
        apiClient: apiClient,
        player: player,
        config: config,
        // 其他初始化...
    }
}
```

### 6.2 模块解耦

- **API 模块**：只负责与服务器通信，不依赖 UI 或音频模块
- **音频模块**：只负责音频播放，不依赖 UI 或 API 模块
- **UI 模块**：通过接口与 API 和音频模块交互，不依赖具体实现
- **配置模块**：独立管理配置，被其他模块依赖
- **模型模块**：定义数据结构，被其他模块依赖

### 6.3 事件总线

实现简单的事件总线，用于模块间通信：

```go
type EventBus interface {
    Publish(event Event)
    Subscribe(eventType string, handler EventHandler)
    Unsubscribe(eventType string, handler EventHandler)
}
```

## 7. 配置管理

### 7.1 配置项

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `server.url` | string | "http://localhost:4040" | 服务器 URL |
| `server.username` | string | "admin" | 用户名 |
| `server.password` | string | "" | 密码（加密存储） |
| `player.volume` | int | 80 | 音量（0-100） |
| `player.crossfade` | bool | false | 交叉淡入淡出 |
| `ui.theme` | string | "default" | UI 主题 |
| `ui.artwork` | bool | true | 显示专辑封面 |
| `cache.enabled` | bool | true | 启用缓存 |
| `cache.ttl` | int | 3600 | 缓存有效期（秒） |

### 7.2 配置存储

- **配置文件**：`~/.config/laetibeat/config.yaml`
- **环境变量**：支持 `LAETIBEAT_` 前缀的环境变量
- **命令行参数**：支持覆盖配置项

## 8. 错误处理

### 8.1 错误分类

- **API 错误**：服务器返回的错误
- **网络错误**：网络连接问题
- **播放器错误**：mpv 相关错误
- **配置错误**：配置文件问题
- **UI 错误**：界面渲染问题

### 8.2 错误处理策略

- **错误包装**：使用 `github.com/pkg/errors` 包装错误，保留堆栈信息
- **错误传播**：通过返回值传播错误，避免 panic
- **错误展示**：在 UI 中显示用户友好的错误信息
- **错误恢复**：实现关键操作的错误恢复机制
- **错误日志**：详细记录错误信息到日志

## 9. 性能优化

### 9.1 网络优化

- **请求缓存**：缓存常用 API 响应
- **批量请求**：合并多个 API 请求
- **延迟加载**：按需加载数据，避免一次性加载过多
- **连接池**：复用 HTTP 连接

### 9.2 渲染优化

- **增量渲染**：只渲染变化的部分
- **虚拟列表**：处理大型列表时只渲染可见部分
- **异步渲染**：耗时操作放在后台线程

### 9.3 音频优化

- **缓冲管理**：合理设置 mpv 缓冲大小
- **格式处理**：根据音频格式选择最佳播放参数
- **网络流优化**：针对网络流调整播放策略

## 10. 测试策略

### 10.1 单元测试

- **API 模块**：测试认证和 API 调用
- **音频模块**：测试播放器控制
- **配置模块**：测试配置读写
- **模型模块**：测试数据转换
- **工具模块**：测试工具函数

### 10.2 集成测试

- **API 集成**：测试与真实服务器的交互
- **音频集成**：测试与 mpv 的集成
- **UI 集成**：测试页面切换和交互

### 10.3 端到端测试

- **完整流程测试**：测试从启动到播放的完整流程
- **错误场景测试**：测试各种错误场景的处理
- **性能测试**：测试在大型库中的性能

## 11. 部署与分发

### 11.1 依赖管理

- **Go 模块**：使用 `go.mod` 和 `go.sum` 管理依赖
- **mpv**：用户需要预先安装 mpv

### 11.2 构建

```bash
# 构建可执行文件
go build -o laetibeat ./cmd/laetibeat

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o laetibeat-linux ./cmd/laetibeat
GOOS=darwin GOARCH=amd64 go build -o laetibeat-macos ./cmd/laetibeat
GOOS=windows GOARCH=amd64 go build -o laetibeat.exe ./cmd/laetibeat
```

### 11.3 安装

- **手动安装**：复制可执行文件到 PATH 目录
- **包管理器**：支持通过 Homebrew、Scoop 等安装

## 12. 开发计划

### 12.1 阶段一：基础架构

- [x] 项目初始化和依赖管理
- [x] API 客户端实现
- [x] 音频模块实现
- [x] 配置模块实现
- [x] 模型定义

### 12.2 阶段二：核心功能

- [ ] UI 框架搭建
- [ ] 认证和配置页面
- [ ] 音乐浏览功能
- [ ] 播放控制功能
- [ ] 搜索功能

### 12.3 阶段三：完善与优化

- [ ] 播放列表管理
- [ ] 现在播放页面
- [ ] 错误处理和日志
- [ ] 性能优化
- [ ] 测试和修复

### 12.4 阶段四：发布

- [ ] 文档编写
- [ ] 构建和分发
- [ ] 社区反馈和迭代

## 13. 总结

Laetibeat TUI 采用模块化、解耦的设计，使用现代 Go 语言实践和 Bubble Tea 框架，为用户提供一个功能完整、界面美观的终端音乐播放器。

通过清晰的目录结构、接口设计和依赖管理，确保了代码的可维护性和可扩展性。同时，通过性能优化和错误处理，确保了应用的稳定性和响应速度。

该设计方案为开发团队提供了清晰的指导，有助于快速构建一个高质量的 OpenSubsonic 兼容音乐播放器。