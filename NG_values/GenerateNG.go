package NG_values

import (
    "crypto/rand"
    "math/big"
    "runtime"
    "time"
)

const BitSize int = 1023
const timer = 15*time.Second

func GenerateN() (*big.Int) {
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
                return twoQPlusOne
            }
        case <-time.After(timer):
            return nil
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

func GenerateG(p *big.Int) uint8 {
    phi := new(big.Int).Sub(p, big.NewInt(1))
    factors := sympyFactorint(phi)
    
    for g := big.NewInt(2); g.Cmp(p) < 0 && g.Cmp(big.NewInt(100)) < 0; g.Add(g, big.NewInt(1)) {
        found := true
        for _, f := range factors {
            exp := new(big.Int).Div(phi, f)
            if new(big.Int).Exp(g, exp, p).Cmp(big.NewInt(1)) == 0 {
                found = false
                break
            }
        }
        if found {
            return uint8(g.Uint64())
        }
    }
    return 0
}

func sympyFactorint(n *big.Int) []*big.Int {

    factors := []*big.Int{n}
    return factors
}