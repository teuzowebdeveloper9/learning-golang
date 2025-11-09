# Structs & Estruturas de Dados em Go

Bem-vindo à pasta `struct` do repositório **learning-go**!  
Aqui estou consolidando meus estudos em Go e, ao mesmo tempo, revisitando estruturas de dados clássicas. A ideia é juntar minhas duas frentes de estudo: aprender Go com mais profundidade **e** entender bem estruturas de dados.

## 🧱 O que está aqui

Nesta pasta estão implementações de estruturas de dados em Go.  
A **primeira estrutura** que escolhi trabalhar foi a **árvore binária de busca (Binary Search Tree, BST)**.  
Depois disso, implementei também um **algoritmo de Sliding Window** — uma técnica muito usada em problemas de subarrays e análise de dados em tempo real.

---

## 🌳 Árvore Binária de Busca (Binary Search Tree)

Uma árvore binária de busca é uma estrutura de dados composta por nós, onde cada nó tem no máximo dois filhos — um “filho à esquerda” e um “filho à direita”. As propriedades principais são:

- Para cada nó, todos os valores no seu sub-árvore da esquerda são menores que o valor do nó.  
- Todos os valores no seu sub-árvore da direita são maiores que o valor do nó.  
- Isso permite buscar, inserir e (em implementações completas) remover valores de forma eficiente — muito melhor que listas simples em muitos casos.

### Visão rápida  
Aqui está uma ilustração para ajudar a visualizar:

![Árvore Binária de Busca](https://pythonhelp.wordpress.com/tag/arvores-binarias-de-busca/)

> *Nota: imagem ilustrativa de uma BST para ajudar no entendimento.*

### Implementação em Go
A implementação em Go inclui:

- `Node` — estrutura para representar cada nó da árvore.  
- `BinarySearchTree` — estrutura que gerencia a árvore inteira (raiz, etc).  
- Métodos principais:
  - `Insert(value int)` — insere valores.  
  - `Search(value int) bool` — busca valores.  
  - `InOrder()`, `PreOrder()`, `PostOrder()` — percorre a árvore.  
  - `Min()` e `Max()` — encontra menor e maior valor.  

---

## ⚙️ Sliding Window

O **algoritmo de Sliding Window (janela deslizante)** é uma técnica usada para otimizar o processamento de intervalos contínuos em arrays — muito comum em problemas como:

- Calcular a soma ou média de uma janela fixa.  
- Encontrar o maior valor em um intervalo móvel.  
- Processar fluxos de dados em tempo real sem recomputar tudo.

### Exemplo visual
![Sliding Window](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTCbSE2ZhYPrkLC9v2z9DhtA9XoeAoGFzOyCQ&s)

> *A janela se move sobre o array, atualizando o resultado de forma incremental.*

### Implementação em Go

No código, implementei duas variações:

- `MaxSumSubarray(nums []int, k int) int` — calcula a maior soma de uma subarray de tamanho `k`.  
- `MovingAverage(nums []int, k int) []float64` — calcula a média móvel de uma janela de tamanho `k`.

Essas funções mostram como a técnica economiza processamento, reaproveitando o resultado anterior e ajustando apenas os elementos que “entram” e “saem” da janela.

---

## 🎯 Por que isso é útil?

- Aprender a estruturar código em Go usando **structs, funções e recursão**.  
- Entender como dados podem ser organizados em **hierarquias (árvores)** ou **intervalos dinâmicos (sliding window)**.  
- Ter base sólida para evoluir para algoritmos mais complexos como **árvores balanceadas, heaps, grafos e algoritmos de streaming**.

---

## 🛠 Próximos passos

- Implementar o método **Delete** na árvore binária.  
- Criar uma versão otimizada do **Sliding Window** usando deque para máximo/mínimo em O(n).  
- Explorar outras estruturas: **filas, pilhas, listas encadeadas e grafos**.  
- Adicionar **testes automatizados** (`testing` package).  
- Documentar e comparar eficiência de cada estrutura.

---

Obrigado por acompanhar este estudo comigo!  
Se quiser comentar, sugerir melhorias ou co-desenvolver, fique à vontade.

Mateus  
**Full-Stack Web Developer / Aprendendo Go & Estruturas de Dados**
