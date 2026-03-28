package metrics

import (
	"fmt"
	"os/exec"
)

func LoadGeneration() {
	cmd := exec.Command("vegeta", "attack", "-duration=10s", "-rate=100", "-targest=targets.txt", "|", "vegeta report", ">>", "reports.txt")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("err")
	}
	fmt.Println(string(output))
}
