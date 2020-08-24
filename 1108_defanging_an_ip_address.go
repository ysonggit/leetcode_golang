import "regexp"

func defangIPaddr(address string) string {
	re := regexp.MustCompile(`(\.)`)
	return re.ReplaceAllString(address, "[.]")
}

import "strings"
func defangIPaddr_2(address string) string {
    tokens := strings.Split(address, ".")
    return strings.Join(tokens[:], "[.]")
}