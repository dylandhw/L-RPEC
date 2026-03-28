package metrics

import (
	"fmt"
	"os/exec"
)

func LoadGeneration() {
	cmd := exec.Command("bash", "-c", "vegeta attack -duration=1s -rate=100000 -targets=/home/dw/Code/L-RPEC/metrics/targets.txt | vegeta report >> /home/dw/Code/L-RPEC/metrics/reports.txt")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("error:", err)
	}
}

// vegeta attack -duration=10s -rate=100 -targets=targets.txt | vegeta report >> reports.txt
