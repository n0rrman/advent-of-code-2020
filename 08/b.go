package main

func executeModifiedPrograms(program []instruction) int {
	for i, instr := range program {
		newProg := make([]instruction, len(program))
		copy(newProg, program)
		resetProgram(&newProg)

		if instr.InstrType == "nop" {
			newProg[i].InstrType = "jmp"
		} else if instr.InstrType == "jmp" {
			newProg[i].InstrType = "nop"
		} else {
			continue
		}
		acc, executed := executeProgramWithErrCode(newProg)
		if executed {
			return acc
		}
	}
	return -2
}

func resetProgram(program *[]instruction) {
	for i := range *program {
		(*program)[i].Executed = false
	}
}

func executeProgramWithErrCode(program []instruction) (int, bool) {
	run := true
	pc := 0
	acc := 0
	executed := false

	for run {
		if pc == len(program) {
			run = false
			executed = true
			continue
		}
		instr := program[pc]
		if instr.Executed {
			run = false
			continue
		}
		program[pc].Executed = true
		pc, acc = executeInstr(instr, pc, acc)
	}

	return acc, executed
}
