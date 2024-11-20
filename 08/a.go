package main

func executeInstr(i instruction, pc, acc int) (int, int) {
	switch i.InstrType {
	case "acc":
		return pc + 1, acc + i.Value
	case "jmp":
		return pc + i.Value, acc
	default:
		return pc + 1, acc
	}
}

func executeProgram(program []instruction) int {
	run := true
	pc := 0
	acc := 0

	for run {
		instr := program[pc]
		if instr.Executed {
			run = false
			continue
		}
		program[pc].Executed = true
		pc, acc = executeInstr(instr, pc, acc)
	}
	return acc
}
