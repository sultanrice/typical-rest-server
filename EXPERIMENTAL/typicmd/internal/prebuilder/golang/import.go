package golang

// Imports is slice of import
type Imports []Import

// Import is plain old go object for import
type Import struct {
	Alias       string
	PackageName string
}

// WithAlias to add new import with alias
func (i *Imports) WithAlias(alias, pkg string) {
	*i = append(*i, Import{
		Alias:       alias,
		PackageName: pkg,
	})
}

// AddImport to add new import
func (i *Imports) AddImport(pkg string) {
	i.WithAlias("", pkg)
}

// BlankImport to add new blank import
func (i *Imports) BlankImport(pkg string) {
	i.WithAlias("_", pkg)
}
