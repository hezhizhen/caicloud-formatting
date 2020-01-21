package imports

// UpdateOrder reorders import lines
func UpdateOrder(lines []string) []string {
	var p Packages
	// classify all imported packages
	for _, line := range lines {
		p.Add(line)
	}
	// list all imported packages as expected
	list := p.List()
	return list
}
