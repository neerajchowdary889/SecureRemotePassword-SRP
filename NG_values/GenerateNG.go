package NG_values

import (
    "crypto/rand"
    "math/big"
    "runtime"
    "time"
)
const G uint8 = 16
const BitSize int = 2047
const timer = 15*time.Second

func GenerateNG() (*big.Int, uint8) {
    ch := make(chan *big.Int)
    for i := 0; i < runtime.NumCPU(); i++ {
        go searchSophieGermainPrime(ch)
    }
    for {
        select {
        case Q := <-ch:
            if Q != nil {
                twoQPlusOne := new(big.Int).Mul(Q, big.NewInt(2))
                twoQPlusOne.Add(twoQPlusOne, big.NewInt(1))
                return twoQPlusOne, G
            }
        case <-time.After(timer):
            return nil, 0
        }
    }
}

func searchSophieGermainPrime(ch chan<- *big.Int) {
    for {
        Q, err := rand.Prime(rand.Reader, BitSize)
        if err != nil {
            ch <- nil
            return
        }
        if validSophieGermain(Q) {
            ch <- Q
            return
        }
    }
}

func validSophieGermain(Q *big.Int) bool {
    twoPPlusOne := new(big.Int).Mul(Q, big.NewInt(2))
    twoPPlusOne.Add(twoPPlusOne, big.NewInt(1))
    return twoPPlusOne.ProbablyPrime(20)
}