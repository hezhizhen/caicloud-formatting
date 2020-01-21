package imports

// Packages holds all imported packages for one file
// and separates them into 4 categories.
type Packages struct {
	Std     []string // standard libraries for Go
	Local   []string // packages defined in the current repository
	Company []string // packages from other repositories of caicloud
	Others  []string // packages from somewhere else
}

// Add adds one imported package to its category.
func (p *Packages) Add(line string) {
	pkg, origin := extractPackage(line)
	if pkg == "" {
		return
	}
	// classify
	category := classify(pkg)
	switch category {
	case "std":
		p.Std = append(p.Std, origin)
	case "local":
		p.Local = append(p.Local, origin)
	case "company":
		p.Company = append(p.Company, origin)
	case "others":
		p.Others = append(p.Others, origin)
	}
}

// List lists all packages in a specific order:
//     std
//     local
//     company
//     others
// Add a blank line for separating different categories.
// If the convention is changed, update this function then.
func (p *Packages) List() []string {
	var ret []string
	for i, l := range p.Std {
		ret = appendPackage(ret, l, i == 0)
	}
	for i, l := range p.Local {
		ret = appendPackage(ret, l, i == 0)
	}
	for i, l := range p.Company {
		ret = appendPackage(ret, l, i == 0)
	}
	for i, l := range p.Others {
		ret = appendPackage(ret, l, i == 0)
	}
	return ret
}
