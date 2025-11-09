# Structs & Estruturas de Dados em Go

Bem-vindo à pasta `struct` do repositório **learning-go**!  
Aqui estou consolidando meus estudos em Go e, ao mesmo tempo, revisitando estruturas de dados clássicas. A ideia é juntar minhas duas frentes de estudo: aprender Go com mais profundidade **e** entender bem estruturas de dados.

## 🧱 O que está aqui

Nesta pasta estão implementações de estruturas de dados em Go.  
A **primeira estrutura** que escolhi trabalhar é a **árvore binária de busca (Binary Search Tree, BST)**.

## 🌳 O que é uma Árvore Binária de Busca?

Uma árvore binária de busca é uma estrutura de dados composta por nós, onde cada nó tem no máximo dois filhos — um “filho à esquerda” e um “filho à direita”. As propriedades principais são:

- Para cada nó, todos os valores no seu sub-árvore da esquerda são menores que o valor do nó.  
- Todos os valores no seu sub-árvore da direita são maiores que o valor do nó.  
- Isso permite buscar, inserir e (em implementações completas) remover valores de forma eficiente — muito melhor que listas simples em muitos casos.

### Visão rápida  
Aqui está uma ilustração para ajudar a visualizar:

![Árvore Binária de Busca](https://pythonhelp.wordpress.com/tag/arvores-binarias-de-busca/)

> *Nota: imagem ilustrativa de uma BST para ajudar no entendimento.*

## 🔍 O que eu implementei

Na implementação em Go você encontrará:

- `Node` — estrutura para representar cada nó da árvore.  
- `BinarySearchTree` — estrutura para gerenciar a árvore inteira (raiz, etc).  
- Métodos básicos:
  - `Insert(value int)` — para inserir valores.  
  - `Search(value int) bool` — para buscar a existência de um valor.  
  - `InOrder()`, `PreOrder()`, `PostOrder()` — para percursos da árvore.  
  - `Min()` e `Max()` — para encontrar o menor e o maior valor da árvore.  
- Um exemplo no `main.go` para demonstrar a árvore em funcionamento.

## 🎯 Por que isso é útil?

- Aprender a estruturar código em Go usando structs, métodos e recursão.  
- Entender como dados podem ser organizados em hierarquias e como isso melhora certas operações (busca, ordenação, etc).  
- Ter uma base para evoluir: depois dessa árvore, posso implementar **remoção de nós**, **árvores balanceadas**, **heaps**, **grafos**, e muito mais.

## 🛠 Próximos passos

- Implementar a operação **Delete** para remover nós da árvore.  
- Explorar árvores balanceadas (por exemplo, AVL ou Red-Black Tree).  
- Ampliar para outras estruturas: filas, pilhas, listas encadeadas, grafos.  
- Integrar testes automatizados em Go (`testing` package).  
- Documentar cada estrutura e uso prático.

---

Obrigado por acompanhar este estudo comigo!  
Se quiser comentar, sugerir melhorias ou co-desenvolver, fique à vontade.

Matheus  
Full-Stack Web Developer / Aprendendo Go & Estruturas de Dados  
