// maps in go lang

package main

import "fmt"

func main() {
	StudentAge := make(map[string]int)

	StudentAge["Aryya"] = 23
	StudentAge["Saurabh"] = 27
	StudentAge["Prerna"] = 21
	StudentAge["Akrati"] = 19
	StudentAge["Sahiti"] = 42
	StudentAge["Kirti"] = 22

	fmt.Println(StudentAge["Kirti"])
	fmt.Println(StudentAge["Saurabh"])
	fmt.Println(StudentAge)
	fmt.Println(len(StudentAge))

	// nested map

	superhero := map[string]map[string]string{
		"Superman": {
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": {
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
