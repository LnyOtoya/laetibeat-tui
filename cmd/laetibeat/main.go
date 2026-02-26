package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbletea"

	"github.com/LnyOtoya/laetibeat-tui/internal/api"
	"github.com/LnyOtoya/laetibeat-tui/internal/audio/mpv"
	"github.com/LnyOtoya/laetibeat-tui/internal/config"
	"github.com/LnyOtoya/laetibeat-tui/internal/ui"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化API客户端
	apiClient := api.NewClient(
		cfg.Server.URL,
		cfg.Server.Username,
		cfg.Server.Password,
		"laetibeat-tui",
	)

	// 初始化mpv客户端
	mpvClient, err := mpv.NewClient()
	if err != nil {
		log.Fatalf("初始化mpv失败: %v", err)
	}
	defer mpvClient.Close()

	// 设置mpv音量
	mpvClient.SetVolume(cfg.Player.Volume)

	// 初始化应用
	app := ui.NewApp(cfg, apiClient, mpvClient)

	// 启动应用
	p := tea.NewProgram(app, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatalf("启动应用失败: %v", err)
	}

	fmt.Println("应用已退出")
}
