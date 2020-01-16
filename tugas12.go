package tugas12

import (
	"fmt"
	"regexp"
)

func Tugas12(text string) {
	regex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		return
	}

	res1 := regex.FindAllString(text, -1)
	fmt.Println(res1)

	res2 := regex.MatchString(text)
	fmt.Println(res2)
}
