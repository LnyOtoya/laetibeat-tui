package mpv

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/Microsoft/go-winio"
	"github.com/pkg/errors"

	"github.com/LnyOtoya/laetibeat-tui/internal/models"
)

// Client mpv IPC客户端
type Client struct {
	SocketPath    string
	MPVProcess    *exec.Cmd
	Conn          net.Conn
	Connected     bool
	Mutex         sync.Mutex
	EventHandlers map[string][]EventHandler
}

// EventHandler 事件处理器类型
type EventHandler func(event *Event)

// NewClient 创建新的mpv客户端
func NewClient() (*Client, error) {
	// 检查mpv是否安装
	if _, err := exec.LookPath("mpv"); err != nil {
		return nil, errors.Wrap(err, "mpv not found in PATH. Please install mpv first.")
	}

	// 在Windows上，使用唯一的命名管道格式
	pipeName := fmt.Sprintf(`\\.\pipe\mpv-%d`, time.Now().UnixNano())

	// 启动mpv进程
	cmd := exec.Command(
		"mpv",
		"--no-video",
		"--idle=yes",
		"--input-ipc-server="+pipeName,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start mpv")
	}

	client := &Client{
		SocketPath:    pipeName,
		MPVProcess:    cmd,
		Connected:     false,
		EventHandlers: make(map[string][]EventHandler),
	}

	// 等待mpv启动并创建命名管道
	time.Sleep(1000 * time.Millisecond)

	// 连接到mpv
	if err := client.Connect(); err != nil {
		// 清理进程
		if client.MPVProcess != nil && client.MPVProcess.Process != nil {
			client.MPVProcess.Process.Kill()
		}
		return nil, err
	}

	// 启动事件监听
	go client.listenEvents()

	return client, nil
}

// Connect 连接到mpv IPC服务器
func (c *Client) Connect() error {
	// 在Windows上，我们需要使用winio包来连接到命名管道
	conn, err := winio.DialPipe(c.SocketPath, nil)
	if err != nil {
		return errors.Wrap(err, "failed to connect to mpv named pipe")
	}

	c.Conn = conn
	c.Connected = true

	return nil
}

// Close 关闭客户端
func (c *Client) Close() error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 断开连接
	if c.Conn != nil {
		c.Conn.Close()
		c.Conn = nil
	}

	// 终止mpv进程
	if c.MPVProcess != nil && c.MPVProcess.Process != nil {
		c.MPVProcess.Process.Kill()
		c.MPVProcess.Wait()
		c.MPVProcess = nil
	}

	c.Connected = false
	return nil
}

// SendCommand 发送命令到mpv
func (c *Client) SendCommand(command []interface{}) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if !c.Connected || c.Conn == nil {
		return errors.New("not connected to mpv")
	}

	// 构建命令对象
	cmdObj := map[string]interface{}{
		"command": command,
	}

	// 序列化为JSON
	cmdJSON, err := json.Marshal(cmdObj)
	if err != nil {
		return errors.Wrap(err, "failed to marshal command")
	}

	// 添加换行符
	cmdJSON = append(cmdJSON, '\n')

	// 发送命令
	_, err = c.Conn.Write(cmdJSON)
	if err != nil {
		c.Connected = false
		return errors.Wrap(err, "failed to send command")
	}

	// 读取响应
	response, err := c.readResponse()
	if err != nil {
		return err
	}

	// 解析响应
	var respObj map[string]interface{}
	if err := json.Unmarshal(response, &respObj); err != nil {
		return errors.Wrap(err, "failed to unmarshal response")
	}

	// 检查是否有错误
	if errorObj, ok := respObj["error"].(map[string]interface{}); ok {
		return errors.Errorf("mpv error: %s", errorObj["message"])
	}

	return nil
}

// SetProperty 设置属性
func (c *Client) SetProperty(name string, value interface{}) error {
	return c.SendCommand([]interface{}{"set_property", name, value})
}

// GetProperty 获取属性
func (c *Client) GetProperty(name string) (interface{}, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if !c.Connected || c.Conn == nil {
		return nil, errors.New("not connected to mpv")
	}

	// 构建命令对象
	cmdObj := map[string]interface{}{
		"command": []interface{}{"get_property", name},
	}

	// 序列化为JSON
	cmdJSON, err := json.Marshal(cmdObj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal command")
	}

	// 添加换行符
	cmdJSON = append(cmdJSON, '\n')

	// 发送命令
	_, err = c.Conn.Write(cmdJSON)
	if err != nil {
		c.Connected = false
		return nil, errors.Wrap(err, "failed to send command")
	}

	// 读取响应
	response, err := c.readResponse()
	if err != nil {
		return nil, err
	}

	// 解析响应
	var respObj map[string]interface{}
	if err := json.Unmarshal(response, &respObj); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	// 检查是否有错误
	if errorObj, ok := respObj["error"].(map[string]interface{}); ok {
		return nil, errors.Errorf("mpv error: %s", errorObj["message"])
	}

	// 返回结果
	return respObj["data"], nil
}

// readResponse 读取响应
func (c *Client) readResponse() ([]byte, error) {
	var buffer []byte
	var buf [1024]byte

	for {
		n, err := c.Conn.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				c.Connected = false
			}
			return nil, err
		}

		buffer = append(buffer, buf[:n]...)

		// 检查是否包含完整的JSON对象（以换行符结束）
		if strings.Contains(string(buffer), "\n") {
			break
		}
	}

	return buffer, nil
}

// listenEvents 监听事件
func (c *Client) listenEvents() {
	// 启用事件监听
	c.SendCommand([]interface{}{"enable_event", "all"})

	for c.Connected {
		// 读取事件
		eventData, err := c.readResponse()
		if err != nil {
			break
		}

		// 解析事件
		var event Event
		if err := json.Unmarshal(eventData, &event); err != nil {
			continue
		}

		// 处理事件
		c.handleEvent(&event)
	}
}

// handleEvent 处理事件
func (c *Client) handleEvent(event *Event) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 调用注册的事件处理器
	if handlers, ok := c.EventHandlers[event.Event]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

// On 注册事件处理器
func (c *Client) On(eventType string, handler EventHandler) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.EventHandlers[eventType] = append(c.EventHandlers[eventType], handler)
}

// Off 移除事件处理器
func (c *Client) Off(eventType string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 移除指定事件类型的所有处理器
	delete(c.EventHandlers, eventType)
}

// GetStatus 获取播放状态
func (c *Client) GetStatus() (models.PlayStatus, error) {
	status := models.PlayStatus{
		Playing: false,
		Volume:  80,
		Repeat:  false,
		Random:  false,
	}

	return status, nil
}
