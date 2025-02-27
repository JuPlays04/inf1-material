package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	f := -1
	l := -1
	sol := []string{}
	for i := 0; i < len(list); i++ {
		if list[i] == first {
			f = i
		}
		if list[i] == last {
			l = i
		}

	}

	if f == -1 || l == -1 || f > l {
		return sol
	}
	sol = append(sol, list[:f]...)
	sol = append(sol, list[l+1:]...)

	return sol
}
