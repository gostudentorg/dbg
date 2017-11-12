package dbg

import (
	"os"
	"os/exec"
	"runtime/pprof"
)

var (
	fileCPU string
)

func CPUProfileStart(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = pprof.StartCPUProfile(f)
	if err == nil {
		fileCPU = fileName
	}
	return err
}

// CPUProfileOpenBrowser creates a go-torch cpu flamegraph and opens it in the browser
// Needs gotorch to be installed with flamegraph scripts in PATH
// TODO: make open cmd plattform independent
func CPUProfileOpenBrowser() error {
	pprof.StopCPUProfile()
	cmd := exec.Command("go-torch", fileCPU)
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("open", "torch.svg")
	return cmd.Run()
}
