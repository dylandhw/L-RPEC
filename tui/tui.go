package main

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Run Vegeta stress testing"},

		selected: make(map[int]struct{}),
	}
}
