# AC1ESTD
package main

import "fmt"

func calculaMedia(valores ...float64) float64 {
    if len(valores) == 0 {
        return 0 
    }
    
    soma := 0.0
    for i := 0; i < len(valores); i++ {
        soma += valores[i]
    }
    
    return soma / float64(len(valores))
}

func verificaParidade(numero int) string {
    if numero%2 == 0 {
        return "par"
    }
    return "ímpar"
}

func minhaPotencia(base, exp int) int {
    if exp == 0 {
        return 1
    }
    resultado := 1
    for i := 0; i < exp; i++ {
        resultado = resultado * base
    }
    return resultado
}

func converteCelsiusParaFahrenheit(C float64) float64 {
    return (C * 9 / 5) + 32
}
