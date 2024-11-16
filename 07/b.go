package main

func countContainedBags(start string, bags map[string]bag) int {
	return recCount(start, bags) - 1
}

func recCount(start string, bags map[string]bag) int {
	count := 0
	for _, bag := range bags[start].bags {
		count += recCount(bag, bags)
	}

	return count + 1
}
