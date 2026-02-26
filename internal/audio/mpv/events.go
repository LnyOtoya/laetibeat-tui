package mpv

import "encoding/json"

// Event mpv事件结构
type Event struct {
	Event        string          `json:"event"`
	ID           int             `json:"id,omitempty"`
	Error        *EventError     `json:"error,omitempty"`
	Data         json.RawMessage `json:"data,omitempty"`
	Paused       bool            `json:"paused,omitempty"`
	PlaybackTime float64         `json:"playback-time,omitempty"`
	Duration     float64         `json:"duration,omitempty"`
	Reason       string          `json:"reason,omitempty"`
	File         string          `json:"file,omitempty"`
}

// EventError 事件错误
type EventError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// EventTypes mpv事件类型
const (
	EventStartFile       = "start-file"
	EventEndFile         = "end-file"
	EventFileLoaded      = "file-loaded"
	EventPause           = "pause"
	EventUnpause         = "unpause"
	EventPlaybackRestart = "playback-restart"
	EventSeek            = "seek"
	EventPropertyChange  = "property-change"
	EventAudioReconfig   = "audio-reconfig"
	EventClientMessage   = "client-message"
	EventLogMessage      = "log-message"
	EventIdle            = "idle"
	EventShutdown        = "shutdown"
)

// EndFileReason end-file事件原因
const (
	EndFileReasonEOF      = "eof"
	EndFileReasonStop     = "stop"
	EndFileReasonQuit     = "quit"
	EndFileReasonError    = "error"
	EndFileReasonRedirect = "redirect"
	EndFileReasonUnknown  = "unknown"
)
