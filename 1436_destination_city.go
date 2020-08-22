func destCity(paths [][]string) string {
	routes := make(map[string]string)
	for _, path := range paths {
		routes[path[0]] = path[1]
	}
	src := paths[0][0]
	_, ok := routes[src]
	next := ok
	for next {
		dest, ok := routes[src]
		if ok {
			src = dest
		}
		next = ok
	}
	return src
}