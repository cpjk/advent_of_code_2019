package util

import "os"
import "bufio"
import "strconv"
import "errors"

func InjestIntegerList(fileName string)  ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("could not read file with name ")
	}

	scanner := bufio.NewScanner(file)

	linesStringArray := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		linesStringArray = append(linesStringArray, line)
	}

	numbersArray := make([]int, len(linesStringArray))

	for i, numString := range linesStringArray {
		num, err := strconv.Atoi(numString)
		if err != nil {
			return nil, errors.New("could not parse string ")
		}

		numbersArray[i] = num
	}

	return numbersArray, nil
}
