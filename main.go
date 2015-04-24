package main

import (
    "fmt"
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
        fmt.Println(arr[i])
    }
}
