package metrics

import (
	"fmt"
)

func Tests() {
	fmt.Print("===RUNNING VEGETA STRESS TESTING===")
	LoadGeneration()
}
