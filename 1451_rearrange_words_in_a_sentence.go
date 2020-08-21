import (
    "sort"
    "strings"
)

type byLength [] string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func arrangeWords(text string) string {
    words := strings.Fields(strings.ToLower(text))
    sort.Stable(byLength(words)) // must use stable sort to keep original order of equal element
    res := strings.Join(words[:]," ")
    return strings.ToUpper(res[0:1]) + res[1:]   
}