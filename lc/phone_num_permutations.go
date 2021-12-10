package lc

import "fmt"

func LetterCombinationsMain() {
	fmt.Println(letterCombinations("472235")) // grep isabel c:
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letters := map[rune][]rune{
		'2': []rune("abc"),
		'3': []rune("def"),
		'4': []rune("ghi"),
		'5': []rune("jkl"),
		'6': []rune("mno"),
		'7': []rune("pqrs"),
		'8': []rune("tuv"),
		'9': []rune("wxyz"),
	}

	perms := 1
	for _, c := range digits {
		perms *= len(letters[c])
	}

	res := make([][]rune, perms)
	for i := 0; i < len(res); i++ {
		res[i] = make([]rune, len(digits))
	}

	step := perms
	for i, d := range digits {
		step /= len(letters[d])

		for j := 0; j < len(res); j++ {
			res[j][i] = letters[d][(j/step)%len(letters[d])]
		}
	}

	ret := make([]string, perms)
	for i, rr := range res {
		ret[i] = string(rr)
	}

	return ret
}
