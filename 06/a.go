package main

func sumUniqueVotes(g []group) int {
	sum := 0
	for _, group := range g {
		voteSet := make(map[vote]bool)
		for _, p := range group {
			for _, vote := range p {
				voteSet[vote] = true
			}
		}
		sum += len(voteSet)
	}
	return sum
}
