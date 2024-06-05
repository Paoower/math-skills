<h1> Math-skills </h1>

## Arborescence

- `main.go`
- `calculation/`
  - `average.go`
  - `median.go`
  - `variance.go`
  - `deviation.go`

## main.go

Le fichier `main.go` gère la lecture des données depuis un fichier texte et appelle les fonctions de calcul des statistiques. Voici le fonctionnement détaillé :

```go
package main

import (
	"calculation/calculation"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Nom de fichier manquant")
		return
	} else if len(os.Args) != 2 {
		fmt.Println("Mauvaise utilisation de l'outil")
		return
	}
	inputfile := os.Args[1]
	input, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	var data []float64
	for _, line := range lines {
		if line == "" {
			continue
		}
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Erreur de conversion de la ligne '%s' en float : %v\n", line, err)
			return
		}
		data = append(data, value)
	}
	Statistic(data)
}

func Statistic(data []float64) {
	fmt.Println("Moyenne :", int(math.Round(calculation.Average(data))))
	fmt.Println("Médiane :", int(math.Round(calculation.Median(data))))
	fmt.Println("Variance :", int(math.Round(calculation.Variance(data))))
	fmt.Println("Écart Type :", int(math.Round(calculation.Deviation(data))))
}
```

## Average.go

- **Moyenne** : La fonction `Average` parcourt chaque élément de la tranche de nombres, les additionne et divise ensuite la somme par le nombre total d'éléments pour obtenir la moyenne.

```go
package calculation

func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
```

## Median.go

- **Médiane** : La fonction `Median` trie d'abord la tranche de nombres en utilisant `sort.Float64s`. Une fois les nombres triés, elle trouve la médiane de la manière suivante :
    - Si le nombre d'éléments est impair, la médiane est l'élément au centre de la tranche triée, c'est-à-dire `data[n/2]`.
    - Si le nombre d'éléments est pair, la médiane est la moyenne des deux éléments centraux. Dans ce cas, la fonction identifie les deux valeurs du milieu : `data[n/2 - 1]` et `data[n/2]`, puis calcule la moyenne de ces deux valeurs en utilisant la formule `(data[n/2 - 1] + data[n/2]) / 2`.
 
```go
package calculation

import "sort"

func Median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return data[(n+1)/2]
	}
	return data[n/2]
}
```

## Variance.go

- **Variance** : La fonction `Variance` commence par calculer la moyenne des nombres à l'aide de la fonction `Average`. Ensuite, elle calcule la somme des carrés des écarts entre chaque nombre et la moyenne, puis divise cette somme par le nombre total d'éléments pour obtenir la variance.

```go
package calculation

func Variance(data []float64) float64 {
	avg := Average(data)
	sum := 0.0
	for _, value := range data {
		sum += (value - avg) * (value - avg)
	}
	return sum / float64(len(data))
}
```

## Deviation.go

- **Écart Type** : La fonction `Deviation` utilise la fonction `Variance` pour obtenir la variance, puis calcule la racine carrée de la variance pour obtenir l'écart type.

```go
package calculation

import "math"

func Deviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
```

## Utilisation

Pour utiliser le programme :

Clonez ce dêpot et exécutez le programme en utilisant la commande suivante :

```bash
git clone https://zone01normandie.org/git/mtrebert/math-skills.git
cd math-skills/
go run . <inputfile>
```





