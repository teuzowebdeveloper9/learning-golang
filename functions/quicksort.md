# Functions

## Sobre este repositório

Esta pasta contém implementações de funções em Go, incluindo uma aplicação do algoritmo **Quicksort**. Fiquei com vontade de aplicar o Quicksort aqui para explorar como ele funciona e entender melhor sua lógica. É um algoritmo eficiente e amplamente utilizado para ordenação.

---

## O que é o Quicksort?

O **Quicksort** é um algoritmo de ordenação baseado na técnica de divisão e conquista. Ele funciona da seguinte forma:

1. Escolhe um elemento como **pivô**.
2. Divide o array em duas partes:
   - Elementos menores ou iguais ao pivô.
   - Elementos maiores que o pivô.
3. Recursivamente aplica o Quicksort nas duas partes.
4. Junta as partes ordenadas com o pivô no meio.

### Exemplo visual do Quicksort:
![Quicksort Example](https://upload.wikimedia.org/wikipedia/commons/9/9c/Quicksort-example.gif)

---

## Código

Aqui está a implementação do Quicksort em Go:

```go
package main

import "fmt"

func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := arr[0]
    var left, right []int

    for _, v := range arr[1:] {
        if v <= pivot {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }

    return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
    numbers := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Original slice:", numbers)

    sortedNumbers := quicksort(numbers)
    fmt.Println("Sorted slice:", sortedNumbers)
}
```

---

## Resultado esperado

Se você rodar o código, verá algo como:

```
Original slice: [10 7 8 9 1 5]
Sorted slice: [1 5 7 8 9 10]
```

---

## Referência

Para mais detalhes sobre o algoritmo, veja o exemplo visual acima ou explore o código nesta pasta . 😊
```# Functions

## Sobre este repositório

Este repositório contém implementações de funções em Go, incluindo uma aplicação do algoritmo **Quicksort**. Fiquei com vontade de aplicar o Quicksort aqui para explorar como ele funciona e entender melhor sua lógica. É um algoritmo eficiente e amplamente utilizado para ordenação.

---

## O que é o Quicksort?

O **Quicksort** é um algoritmo de ordenação baseado na técnica de divisão e conquista. Ele funciona da seguinte forma:

1. Escolhe um elemento como **pivô**.
2. Divide o array em duas partes:
   - Elementos menores ou iguais ao pivô.
   - Elementos maiores que o pivô.
3. Recursivamente aplica o Quicksort nas duas partes.
4. Junta as partes ordenadas com o pivô no meio.

### Exemplo visual do Quicksort:
![Quicksort Example](https://upload.wikimedia.org/wikipedia/commons/9/9c/Quicksort-example.gif)

---

## Código

Aqui está a implementação do Quicksort em Go:

```go
package main

import "fmt"

func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := arr[0]
    var left, right []int

    for _, v := range arr[1:] {
        if v <= pivot {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }

    return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
    numbers := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Original slice:", numbers)

    sortedNumbers := quicksort(numbers)
    fmt.Println("Sorted slice:", sortedNumbers)
}
```

---

## Resultado esperado

Se você rodar o código, verá algo como:

```
Original slice: [10 7 8 9 1 5]
Sorted slice: [1 5 7 8 9 10]
```

---

## Referência

Para mais detalhes sobre o algoritmo, veja o exemplo visual acima ou explore o código neste repositório. 😊