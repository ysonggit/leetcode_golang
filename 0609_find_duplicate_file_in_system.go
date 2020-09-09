import (
	"regexp"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	content_map := make(map[string][]string)
	// https://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-go-regular-expressions
	re := regexp.MustCompile(`(?P<fname>[[:word:]]+.txt)\((?P<content>[[:word:]]+)\)`)
	for _, path := range paths {
		tokens := strings.Split(path, " ")
		dirname := tokens[0]
		// fmt.Printf("%#v\n", re.FindAllStringSubmatch("root/a 3.txt(abcd) 2.txt(efg)", -1))
		// >>> [][]string{[]string{"3.txt(abcd)", "3.txt", "abcd"}, []string{"2.txt(efg)", "2.txt", "efg"}}
		file_contents := re.FindAllStringSubmatch(path, -1)
		for _, file_content := range file_contents {
			fname, content := file_content[1], file_content[2]
			content_map[content] = append(content_map[content], dirname+"/"+fname)
		}
	}
	dups := [][]string{}
	for _, file_lists := range content_map {
		if len(file_lists) > 1 {
			dups = append(dups, file_lists)
		}
	}
	return dups
}