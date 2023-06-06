package xenv

import (
	"os"
	"runtime"
	"strings"
)

func IsWin() bool {
	return runtime.GOOS == "windows"
}

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsWSL() bool {
	b := make([]byte, 1024)
	f, err := os.Open("/proc/version")
	if err == nil {
		_, _ = f.Read(b)
		_ = f.Close()
		return strings.Contains(string(b), "Microsoft")
	}
	return false
}
