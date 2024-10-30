package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//https://yourbasic.org/golang/round-float-to-int/
//https://www.delftstack.com/fr/howto/go/how-to-read-a-file-line-by-line-in-go/
//https://devopssec.fr/article/variables-golang

func Average(tab []float64) float64 {
	var average float64
	var compteur int
	for i := 0; i < len(tab); i++ {
		average += tab[i]
		compteur++
	}
	average = average / float64(compteur)
	return average
}

func Median(tab []float64) float64 {
	var median float64
	var compteur float64
	for i := 0; i < len(tab)-1; i++ {
		for j := 0; j < len(tab)-1-i; j++ {
			if tab[j] > tab[j+1] {
				tab[j], tab[j+1] = tab[j+1], tab[j]
			}
		}
	}
	if len(tab)%2 != 0 {
		for i := 0; i < len(tab); i++ {
			compteur++
			if compteur == float64(len(tab))/2 {
				median = tab[i]
			}
		}
	} else if len(tab)%2 == 0 {
		for i := 0; i < len(tab); i++ {
			compteur++
			if compteur == float64(len(tab))/2 {
				median = (tab[i] + tab[i+1]) / 2
			}
		}
	}
	return median
}

func Variance(tab []float64) float64 {
	var variance float64
	for i := 0; i < len(tab); i++ {
		variance = variance + (tab[i] * tab[i])
	}
	variance = variance / float64(len(tab))
	variance -= Average(tab) * Average(tab)
	return variance
}

func StandardDeviation(tab []float64) float64 {
	return math.Sqrt(float64(Variance(tab)))
}

func main() {
	//On vérifie la longueur des arguments
	if len(os.Args) != 2 {
		fmt.Println("Nombre d'arguments non valide")
		os.Exit(0)
	}
	//On ouvre le fichier et on le stocke dans une variable
	file, err := os.OpenFile(os.Args[1], os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Création/Ouverture du fichier sample.txt impossible")
		os.Exit(1)
	}
	//On crée un tableau de string vide qui va stocker et séparer mes nombres
	finaltab := []string{}
	//Buffio lit mon fichier lignes par lignes
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		//On append toutes nos valeurs une à une
		finaltab = append(finaltab, fileScanner.Text())
	}
	//On crée un tableau de float64 pour pouvoir utiliser nos fonctions dessus
	tabfloat := []float64{}
	//On fait une boucle pour convertir tout notre tableau de string en float64
	for _, word := range finaltab {
		//Convertion string en int
		tempword, error := strconv.Atoi(word)
		if error != nil {
			fmt.Println("Erreur: donnée non décimale")
		}
		//On rajoute au tableau de float64 notre int qu'on converti en même temps
		tabfloat = append(tabfloat, float64(tempword))
	}
	//Print de l'application de toutes mes fonctions
	fmt.Println("Average:", int(math.Round(Average(tabfloat))))
	fmt.Println("Median:", int(math.Round(Median(tabfloat))))
	fmt.Println("Variance:", int(math.Round(Variance(tabfloat))))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDeviation(tabfloat))))
}
