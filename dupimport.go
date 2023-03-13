package dupimport

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = `found duplicated import tools`

var Analyzer = &analysis.Analyzer{
	Name: `dupimport`,
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		paths := map[string]struct{}{}
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if _, found := paths[path]; found {
				pass.Reportf(ip.Pos(), "WARNING: %s is duplicated import", path)
			}
			paths[path] = struct{}{}
		}
	}
	return nil, nil
}
