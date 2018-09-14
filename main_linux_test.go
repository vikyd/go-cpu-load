// +build linux

package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_allCores(t *testing.T) {
	exeFilePath := "test_tmp/go-cpu-load"

	// build
	if exist, _ := Exists(exeFilePath); exist {
		runCmd("sh remove " + exeFilePath)
	}
	runCmd("go build -o " + exeFilePath)

	// run and get pid
	coresCount := runtime.NumCPU()
	percentage := 33
	pid := startCmd(fmt.Sprintf("%s -c %d -p %d ", exeFilePath, coresCount, percentage))
	// wait for start up
	time.Sleep(1 * time.Second)

	// judge
	var sum float64
	count := 5
	t.Logf("coresCount: %d, percentage: %d", coresCount, percentage)
	usageExpect := float64(coresCount*percentage) / float64(runtime.NumCPU()*100) * 100
	var tolerance float64 = 3

	for i := 0; i < count; i++ {
		usageStr := runBashCmd(fmt.Sprintf("ps -p %d -o %%cpu | tail -n +2", pid))
		usage, err := strconv.ParseFloat(strings.TrimSpace(usageStr), 64)
		if err != nil {
			log.Fatalf("usage return not correct: %s", err.Error())
		}
		usage = usage / float64(coresCount)

		sum += math.Pow(float64(math.Abs(usageExpect-usage)), 2)
		t.Logf("usage: %f", usage)
		t.Logf("usage str: %s", usageStr)
	}
	runCmd(fmt.Sprintf("kill -9 %d", pid))
	t.Logf("usageExpect: %f", usageExpect)

	usageVariance := sum / float64(count)
	t.Logf("usageVariance %f", usageVariance)
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
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error() + "\n" + string(out))
	}
	return string(out)
}

func runBashCmd(cmdStr string) string {
	cmd := exec.Command("bash", "-c", cmdStr)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	stdoutS := string(stdout.Bytes())
	stderrS := string(stderr.Bytes())
	if err != nil {
		log.Fatal(err.Error() + "\n" + "stdout: \n" + stdoutS + "\n" + "stderr:\n" + stderrS)
	}
	return stdoutS
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
