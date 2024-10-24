package main

func findReportVal(vals []int) int {
	for _, val := range vals {
		target := 2020 - val
		for _, tar := range vals {
			if target == tar {
				return val * tar
			}
		}
	}
	return 0
}
