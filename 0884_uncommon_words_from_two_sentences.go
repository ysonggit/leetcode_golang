import "strings"

func uncommonFromSentences(A string, B string) []string {
	words_a, words_b := strings.Fields(A), strings.Fields(B)
	dict_a, dict_b := createFrequencyMap(words_a), createFrequencyMap(words_b)
	unique_a, unique_b := getUniqueWords(dict_a), getUniqueWords(dict_b)
	uncomm := make([]string, 0)
	for _, w := range unique_a {
		if _, ok := dict_b[w]; !ok {
			uncomm = append(uncomm, w)
		}
	}
	for _, w := range unique_b {
		if _, ok := dict_a[w]; !ok {
			uncomm = append(uncomm, w)
		}
	}
	return uncomm
}

func createFrequencyMap(arr []string) map[string]int {
	m := make(map[string]int)
	for _, i := range arr {
		m[i] += 1
	}
	return m
}

func getUniqueWords(dict map[string]int) []string {
	res := make([]string, 0)
	for key, num := range dict {
		if num == 1 {
			res = append(res, key)
		}
	}
	return res
}