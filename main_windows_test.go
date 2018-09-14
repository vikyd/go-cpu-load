// +build windows

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func Test_allCores(t *testing.T) {
	exeFilePath := "test_tmp\\go-cpu-load.exe"

	// build
	if exist, _ := Exists(exeFilePath); exist {
		runCmd("cmd /C del " + exeFilePath)
	}
	runCmd("go build -o " + exeFilePath)

	// run and get pid
	coresCount := runtime.NumCPU() 
	percentage := 33
	pid := startCmd(fmt.Sprintf("%s -c %d -p %d ", exeFilePath, coresCount, percentage))

	// judge
	var sum float64
	count := 5
	t.Logf("coresCount: %d, percentage: %d", coresCount, percentage)
	usageExpect := float64(coresCount*percentage) / float64(runtime.NumCPU()*100) * 100
	var tolerance float64 = 2
	for i := 0; i < count; i++ {
		usageStr := runCmd(fmt.Sprintf("test_bin\\cpu_usage\\cpu_usage_win.exe %d", pid))
		usage, _ := strconv.Atoi(string(usageStr))
		usageF := float64(usage)
		sum += math.Pow(float64(math.Abs(usageExpect-usageF)), 2)
		t.Logf("usage: %s", usageStr)
	}
	runCmd(fmt.Sprintf("taskkill /f /pid %d", pid))
	t.Logf("usageExpect: %f", usageExpect)

	usageVariance := sum / float64(count)
	t.Logf("usageDelta %f", usageVariance)
	if usageVariance > tolerance {
		t.Errorf("usageVariance %f > tolerance %f", usageVariance, tolerance)
	}
}

// ---- tools func ----------------------------
// starts the specified command and wait for it to complete
func runCmd(cmdStr string) string {
	parts := strings.Split(cmdStr, " ")
	cmdCmd := parts[0]
	cmdArgs := parts[1:]
	cmd := exec.Command(cmdCmd, cmdArgs...)
	out, err := cmd.Output()
	// err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

// starts the specified command but does not wait for it to complete
func startCmd(cmdStr string) int {
	parts := strings.Split(cmdStr, " ")
	cmdCmd := parts[0]
	cmdArgs := parts[1:]
	cmd := exec.Command(cmdCmd, cmdArgs...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return cmd.Process.Pid
}

func calcCPULoadTargetPercentage(coresCount int, percentage int) int {
	return (coresCount * percentage) / (runtime.NumCPU() * 100)
}

// Exists reports whether the named file or directory exists.
// https://stackoverflow.com/a/22467409/2752670
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}
