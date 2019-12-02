package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type memory []int // Type alias for an int array represnting program memory

func (mem memory) ValueAtAddress(address int) (int, error) {
	if address >= len(mem) {
		return 0, fmt.Errorf("Attempted to read memory out of bounds at %d, memory size: %d", address, len(mem))
	}
	return mem[address], nil
}

func (mem memory) ValueAtPtr(address int) (int, error) {
	if address >= len(mem) {
		return 0, fmt.Errorf("Attempted to read memory out of bounds at %d, memory size: %d", address, len(mem))
	}
	ptrAddress := mem[address]
	if ptrAddress >= len(mem) {
		return 0, fmt.Errorf("Attempted to read memory out of bounds at %d, memory size: %d", ptrAddress, len(mem))
	}

	return mem[ptrAddress], nil
}

func (mem *memory) WriteToAddress(address, value int) error {
	if address >= len(*mem) {
		return fmt.Errorf("Attempted to write memory out of bounds at %d, memory size: %d", address, len(*mem))
	}
	(*mem)[address] = value
	return nil
}

func runProgram(mem memory) (int, error) {
	// Initalize program counter
	pc := 0
	for {
		opcode, err := mem.ValueAtAddress(pc)
		if err != nil {
			log.Fatal("Error at PC=", pc, ": ", err)
		}

		if opcode == 99 {
			// Print result
			result, _ := mem.ValueAtAddress(0)
			return result, nil
		}

		pc++
		arg1, err := mem.ValueAtPtr(pc)
		if err != nil {
			return 0, fmt.Errorf("Error at PC=", pc, ": ", err)
		}

		pc++
		arg2, err := mem.ValueAtPtr(pc)
		if err != nil {
			return 0, fmt.Errorf("Error at PC=", pc, ": ", err)
		}

		result, err := calculateOpResult(opcode, arg1, arg2)
		if err != nil {
			return 0, fmt.Errorf("Error at PC=", pc, ": ", err)
		}

		pc++
		destAddress, err := mem.ValueAtAddress(pc)
		if err != nil {
			return 0, fmt.Errorf("Error at PC=", pc, ": ", err)
		}

		mem.WriteToAddress(destAddress, result)
		pc++
	}
}

// Read input file and return an array representing the computer's "program memory"
func readInput(path string) (memory, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	mem := make(memory, 0)
	scanner := bufio.NewReader(file)
	for {
		tokenStr, pErr := scanner.ReadSlice(',')
		if pErr != nil && pErr != io.EOF {

			return nil, err
		}
		token, err := strconv.Atoi(string(tokenStr[:len(tokenStr)-1]))
		if err != nil {
			return nil, err
		}
		mem = append(mem, token)

		if pErr == io.EOF {
			break
		}
	}
	return mem, nil
}

func calculateOpResult(opcode, arg1, arg2 int) (int, error) {
	switch opcode {
	case 1:
		return arg1 + arg2, nil
	case 2:
		return arg1 * arg2, nil
	}

	return 0, fmt.Errorf("Unknown opcode %d", opcode)
}

func main() {
	var args = os.Args

	initalMem, err := readInput(args[1])
	if err != nil {
		log.Fatal("Failed to read task input: ", err)
	}

	for op1 := 0; op1 < 100; op1++ {
		for op2 := 0; op2 < 100; op2++ {
			mem := make(memory, len(initalMem))
			copy(mem, initalMem)

			mem.WriteToAddress(1, op1)
			mem.WriteToAddress(2, op2)

			result, err := runProgram(mem)
			if err != nil {
				log.Fatal("Failed to run program with inital conditions ", op1, " ", op2, ": ", err)
			}
			if result == 19690720 {
				fmt.Printf("Solution found with op1: %d op2: %d\n", op1, op2)
				return
			}
		}
	}
	log.Fatal("No solution found :(")
}
