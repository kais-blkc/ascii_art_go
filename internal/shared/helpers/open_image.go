package helpers

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	OSLinux   = "linux"
	OSWindows = "windows"
	OSDarwin  = "darwin"
)

func OpenImage(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case OSLinux:
		cmd = exec.Command("xdg-open", path)
	case OSWindows:
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
	case OSDarwin:
		cmd = exec.Command("open", path)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
