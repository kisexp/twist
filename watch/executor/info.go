package executor

import (
	"runtime"
	"syscall"
)

type Info struct {
	Cmd               string
	Args              []string
	Signal            int
	Timeout           int
	AutoRestart       bool
	PreCmd            string
	PreCmdTimeout     int
	PreCmdIgnoreError bool
}

func NewInfo() Info {
	info := Info{}
	info.Args = make([]string, 0)
	info.Signal = int(syscall.SIGKILL)
	if runtime.GOOS == "window" {
		info.Signal = int(syscall.SIGKILL)
	}
	info.Timeout = 5
	info.AutoRestart = false
	return info
}
