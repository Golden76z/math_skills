package main

// Import des packages nécessaires//
import (
	"fmt"
	"os"
)

func main() {

	//*************************************Définition des règles du sudoku**********************************************//

	// Vérifier les erreurs
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	for i, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		} // Vérifie qu'on a 9 arguments et pas plus sur une même ligne

		for j, ch := range arg {
			if ch == '.' {
				sudoku[i][j] = 0

				// Si on lui passe un '.' il l'interprète comme un zero donc une case à remplir

			} else if ch >= '1' && ch <= '9' {
				sudoku[i][j] = int(ch - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	} // Si on met un caractÈre autre qu'entre 1 et 9, on nous retourne Error

	if solution(0, 0) {
		fmt.Println("")
		printSudoku() // Dessiner le tableau
	} else {
		fmt.Println("Error")
	}
}

// ***************************Créer le tableau sudoku à double entrée de 9x9********************************************************//

var sudoku [9][9]int

func printSudoku() {
	for _, x := range sudoku {
		for _, c := range x {
			fmt.Printf("%d ", c)
		}
		// Le %d c'est pour afficher les valeurs entrées dans le go run main.go, pour tester notre code
		fmt.Println()
	}
	fmt.Println()
} // Ici impression du tableau sudoku//

func canInsertNumb(x, y, value int) bool {
	return !insertNumbV(x, value) && // "!" en go c'est pour inverser le résultat de la fonction//
		!insertNumbH(y, value) &&
		!insertNumbC(x, y, value)
} // Vérifier que la valeur insérée respecte les règles définis dans la fonction main//
//Par exemple, 9 ne peut pas se répéter 2 fois sur l'axe x //
// 9 ne peut pas se répéter 2 fois sur l'axe y//
// 9 ne peut pas se répéter 2 fois dans un carré de 3x3//

func insertNumbV(x, value int) bool {
	for i := range [9]int{} {
		if sudoku[i][x] == value {
			return true
		}
	}
	return false
}

//Fonction qui vérifie sur l'axe vertical si la valeur peut être placée//

func insertNumbH(y, value int) bool {
	for i := range [9]int{} {
		if sudoku[y][i] == value {
			return true
		}
	}
	return false
}

//Fonction qui vérifie sur l'axe horizontal si la valeur peut être placée//

func insertNumbC(x, y, value int) bool {
	sx, sy := (x/3)*3, (y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if sudoku[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

//Fonction qui vérifie sur le carré 3x3 si la valeur peut être placée//

//*************************************Définition de la solution du sudoku**********************************************//

func solution(x, y int) bool {
	if y == 9 {
		return true
	} // Vérifie que toutes les lignes du sudoku ont été parcourues et qu'il a trouvé une solution//
	if sudoku[y][x] != 0 {
		return solution(position(x, y))
	} // Vérifie si la case n'est pas remplie et passe à la case suivante en utilisant la récursivité//
	for i := range [9]int{} {
		v := i + 1
		if canInsertNumb(x, y, v) {
			sudoku[y][x] = v
			if solution(position(x, y)) {
				return true
			} // Dans le cas où la case serait vide, le programme va essayer d'insérer une valeur comprise entre 1 et 9//
			// si il a trouvé la bonne valeur, alors il renvoie TRUE, sinon le programme revient en arrière et teste toutes les autres possibilités//
			sudoku[y][x] = 0
		}
	}
	return false
} //Fonction récursive qui résout le sudoku en essayant toutes les valeurs possibles//

func position(x, y int) (int, int) {
	positionX, positionY := (x+1)%9, y
	if positionX == 0 {
		positionY = y + 1
	}
	return positionX, positionY
}

// Permet de calculer la prochaine position d'une valeur dans le tableau sudoku//
