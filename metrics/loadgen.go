package metrics

import "os/exec"

func LoadGeneration() {
	cmd := exec.Command("vegeta", "attack", "-duration=10s", "-rate=100", "-targest=targets.txt", "|", "vegeta report", ">>", "reports.txt")
}
