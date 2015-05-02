package tokensearch

import (
	"bufio"
	"fmt"
	"strings"
)

type TokenMatches struct {
	match 		*TokenMatch
	locations	map[int]int
}

func MatchesInString(input string) (map[string]*TokenMatches, error) {

	matches := make(map[string]*TokenMatches)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return matches, err
	}
	fmt.Printf("Found %d runes\n", count)
	return matches, nil
}