package main

const (
	NULL  = 0
	FALSE = -1
	TRUE  = 1
)

func countBagColours(start string, bags map[string]bag) int {
	memo := map[string]int{}
	for _, bag := range bags {
		memo[bag.colour] = traverseBag(bag.colour, start, bags, &memo)
	}

	count := 0
	for _, bag := range memo {
		if bag == TRUE {
			count++
		}
	}
	return count
}

func traverseBag(start, end string, b map[string]bag, m *map[string]int) int {
	if (*m)[start] != NULL {
		return (*m)[start]
	}

	contained := b[start].bags
	for _, c := range contained {
		if c == end || traverseBag(c, end, b, m) == TRUE {
			(*m)[start] = TRUE
			return TRUE
		}
	}
	(*m)[start] = FALSE
	return FALSE
}
