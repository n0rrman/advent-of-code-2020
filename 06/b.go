package main

func sumUnanimousVotes(g []group) int {
	sum := 0
	for _, group := range g {
		voteSet := make(map[vote]bool)
		voteCount := make(map[vote]int)
		for _, p := range group {
			for _, vote := range p {
				oldVal := voteCount[vote]
				voteCount[vote] = oldVal + 1
				if voteCount[vote] == len(group) {
					voteSet[vote] = true
				}
			}
		}
		sum += len(voteSet)
	}
	return sum
}
