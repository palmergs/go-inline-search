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

func MatchesInString(root *TokenNode, input string) (map[string]*TokenMatches, error) {

	matches := make(map[string]*TokenMatches)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Printf("%v", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return matches, err
	}
	fmt.Printf("Found %d runes\n", count)
	return matches, nil
}