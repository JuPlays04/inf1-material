package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion EvenPrefix.
 * ERREICHBARE PUNKTE: 10
 */

// EvenPrefix erwartet eine Liste von Zahlen und liefert
// die längste Anfangs-Folge, die nur gerade Zahlen enthält.
// D.h. eine Liste mit allen geraden Zahlen vor der ersten ungeraden.
func EvenPrefix(nums []int) []int {
	peter := []int{}
	for _, n := range nums {

		if n%2 != 0 {
			break
		}
		peter = append(peter, n)
	}

	return peter
}
