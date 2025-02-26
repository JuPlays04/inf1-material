package aufgabe1

import "strings"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	shortest := ""

	for _, peter := range list {
		if strings.HasPrefix(peter, "abc") {
			if shortest == "" || len(peter) < len(shortest) {
				shortest = peter
			}

		}
	}
	return shortest
}
