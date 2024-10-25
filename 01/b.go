package main

func findSecondReportVal(vals []int) int {
	for _, a := range vals {
		for _, b := range vals {
			target := 2020 - a - b
			for _, tar := range vals {
				if target == tar {
					return tar * a * b
				}
			}
		}
	}
	return 0
}
