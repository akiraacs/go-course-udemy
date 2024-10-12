package main

import "fmt"

func main() {

	user := map[string]string{
		"name": "Adrian",
		"age":  "23",
	}

	userInfos := map[string]map[string]string{
		"adress": {
			"street":  "Rua Colorado",
			"number":  "283",
			"zipcode": "29102841",
		},
		"college": {
			"course": "Technology and Information Systems",
		},
	}

	user["phone"] = "2799999999"
	userInfos["college"]["type"] = "Bachelor"

	fmt.Println(userInfos)
	fmt.Println(user)
}
