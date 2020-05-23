package main

import "fmt"

func sma(n int) func(float64) float64 {
    s := make([]float64, 0, n)
    i, sum, rn := 0, 0., 1/float64(n)
    return func(x float64) float64 {
        if len(s) < n {
            sum += x
            s = append(s, x)
            return sum / float64(len(s))
        }
        s[i] = x
        i++
        if i == n {
            i = 0
        }
        sum = 0
        for _, x = range s {
            sum += x
        }
        return sum * rn
    }
}

func main() {
    sma3 := sma(3)
    sma5 := sma(5)
    fmt.Println("x       sma3   sma5")
    for _, x := range []float64{10 , 15 , 12 , 11 , 16, 12, 14 ,13 , 13} {
        fmt.Printf("%5.3f  %5.3f  %5.3f\n", x, sma3(x), sma5(x))
    }
}
