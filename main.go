package main

import (
    "fmt"
    "unicode/utf8"
)

type TokenNode struct {
    nextLetters     map[string]TokenNode
    matches         []TokenMatch
}

type TokenMatch struct {
    ident           int
    category        string
    aspect          string
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
}
