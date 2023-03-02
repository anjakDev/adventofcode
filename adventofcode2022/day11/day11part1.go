package main

import (
	"adventofcode/util"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const MONKEY_LABEL = "M"
const STARTING_LABEL = "Starting items"
const OPERATION_LABEL = "Operation"
const TEST_LABEL = "Test"
const TRUE_LABEL = "If true"
const FALSE_LABEL = "If false"

type MonkeyBusiness struct {
	itemsToInspect      []int
	operator            string
	opValue             string
	testDivisibleNumber int
	ifTrueRecipientIdx  int
	ifFalseRecipientIdx int
	inspectedItemCount  int
}

func partone() int {
	fileScanner, closer := util.NewScanner("input-day11.txt")
	defer closer.Close()

	// Define a regular expression to match numbers in the string
	re := regexp.MustCompile(`\d+`)

	monkeys := make([]*MonkeyBusiness, 0)

	// Parse the monkeys from the file and create a list of monkeys with their attributes.
	monkeyIdx := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Next monkey is coming up.
		if len(line) == 0 {
			monkeyIdx++
			continue
		}

		// Create an empty monkey to be filled by the consecutive lines.
		trimmedLine := strings.TrimSpace(line)
		if string(trimmedLine[0]) == MONKEY_LABEL {
			monkeys = append(monkeys, new(MonkeyBusiness))
			monkeys[monkeyIdx].inspectedItemCount = 0
			continue
		}

		// Split attribute into type and value
		parts := strings.Split(trimmedLine, ": ")
		if len(parts) < 2 {
			panic("invalid attribute string")
		}

		// Check which monkey attribute is contained in the line read
		switch parts[0] {
		case STARTING_LABEL:
			// Find all matches of the regular expression in the string
			matches := re.FindAllString(parts[1], -1)

			itemsToInspect := make([]int, len(matches))

			for i, match := range matches {
				num, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				itemsToInspect[i] = num
			}
			monkeys[monkeyIdx].itemsToInspect = itemsToInspect
		case OPERATION_LABEL:
			opParts := strings.Split(parts[1], " ")
			if len(opParts) < 5 {
				panic("invalid operation string")
			}
			monkeys[monkeyIdx].operator = opParts[3]
			monkeys[monkeyIdx].opValue = opParts[4]
		case TEST_LABEL:
			matches := re.FindAllString(parts[1], -1)
			if len(matches) != 1 {
				panic("invalid test string")
			}
			divisibleNumber, err := strconv.Atoi(matches[0])
			if err != nil {
				panic(err)
			}

			monkeys[monkeyIdx].testDivisibleNumber = divisibleNumber
		case TRUE_LABEL:
			matches := re.FindAllString(parts[1], -1)
			if len(matches) != 1 {
				panic("invalid test string")
			}
			num, err := strconv.Atoi(matches[0])
			if err != nil {
				panic(err)
			}
			monkeys[monkeyIdx].ifTrueRecipientIdx = num
		case FALSE_LABEL:
			matches := re.FindAllString(parts[1], -1)
			if len(matches) != 1 {
				panic("invalid test string")
			}
			num, err := strconv.Atoi(matches[0])
			if err != nil {
				panic(err)
			}
			monkeys[monkeyIdx].ifFalseRecipientIdx = num
		}
	}

	// Let the monkeys inspect the items for 20 rounds.
	for rounds := 0; rounds < 20; rounds++ {
		for _, monkey := range monkeys {
			for _, itemWorryLevel := range monkey.itemsToInspect {
				monkey.inspectedItemCount++

				var constant int
				if monkey.opValue == "old" {
					constant = itemWorryLevel
				} else {
					constant, _ = strconv.Atoi(monkey.opValue)
				}

				var newItemWorryLevel int
				switch monkey.operator {
				case "+":
					newItemWorryLevel = itemWorryLevel + constant
				case "*":
					newItemWorryLevel = itemWorryLevel * constant
				}

				newItemWorryLevel = newItemWorryLevel / 3

				var recipientMonkeyIdx int
				if math.Mod(float64(newItemWorryLevel), float64(monkey.testDivisibleNumber)) != 0 {
					recipientMonkeyIdx = monkey.ifFalseRecipientIdx
				} else {
					recipientMonkeyIdx = monkey.ifTrueRecipientIdx
				}
				monkeys[recipientMonkeyIdx].itemsToInspect = append(monkeys[recipientMonkeyIdx].itemsToInspect, newItemWorryLevel)
			}
			monkey.itemsToInspect = nil
		}
	}

	// Sort monkeys by the amount of items they each inspected, in descending order.
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItemCount > monkeys[j].inspectedItemCount
	})

	return monkeys[0].inspectedItemCount * monkeys[1].inspectedItemCount
}
