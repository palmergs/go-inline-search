package main

import (
    "fmt"
    "unicode/utf8"
    "errors"
)

type TokenNode struct {
    nextLetters     map[rune]*TokenNode
    matches         map[string]*TokenMatch
}

type TokenMatch struct {
    ident           string
    category        string
    aspect          string
}

func (node *TokenNode) Insert(value string, index int, ident, category, aspect string) (int, error) {
    if index < len(value) {
        runeValue, width := utf8.DecodeRuneInString(value[index:])
        if width > 0 {
            nextNode := node.nextLetters[runeValue]
            if nextNode == nil {
                nextNode = new(TokenNode)
                nextNode.nextLetters = make(map[rune]*TokenNode)
                nextNode.matches = make(map[string]*TokenMatch)
                node.nextLetters[runeValue] = nextNode
            }
            return nextNode.Insert(value, index + width, ident, category, aspect)
        } 
        return index, errors.New("UTF-8 character width was 0")
    }
    if node.matches[ident] != nil {
      node.matches[ident].category = category
      node.matches[ident].aspect = aspect
    } else {
      node.matches[ident] = &TokenMatch{ident: ident, category: category, aspect: aspect}
    }
    return index, nil
}

func (node *TokenNode) Include(value string, index int) ([]*TokenMatch, error) {
  if index < len(value) {
      runeValue, width := utf8.DecodeRuneInString(value[index:])
      fmt.Printf("Looking for %s found %s\n", value, runeValue)
      if width > 0 {
        nextNode := node.nextLetters[runeValue]
        if nextNode == nil {
            return nil, nil
        }
        return nextNode.Include(value, index + width)
      }
      return nil, errors.New("UTF-8 character width was 0")
  }

  if len(node.matches) > 0 {
      ret := make([]*TokenMatch, 0, len(node.matches))
      for _, value := range node.matches {
          ret = append(ret, value)
      }
      return ret, nil
  } else {
      return nil, nil
  }
}

func main() {
    tn := new(TokenNode)
    fmt.Println("node = øoæ…œ∑áéøπö¥†®˚¬∆Ω≈ç√ñµ≤≥÷≠–ºª", tn)
    arr := []byte("node = øoæ…œ∑áéøπö¥†®˚¬∆Ω≈ç√ñµ≤≥÷≠–ºª")
    for i := range arr {
        fmt.Printf(" %x", arr[i])
    }
    fmt.Println("")
    const nihongo = "日本語"
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }

    str := "日本語"
    for index, runeValue := range str {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

    rootNode := new(TokenNode)
    rootNode.nextLetters = make(map[rune]*TokenNode)
    rootNode.matches = make(map[string]*TokenMatch)
    rootNode.Insert("baltimore", 0, "baltimore", "city", "test")

    found1, _ := rootNode.Include("baltimore", 0)
    fmt.Printf("Root include %s? %s\n", "baltimore", found1)

    found2, _ := rootNode.Include("seattle", 0)
    fmt.Printf("Root include %s? %s\n", "seattle", found2)
}
