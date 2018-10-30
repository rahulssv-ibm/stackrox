package options

import (
	"sort"

	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/generated/api/v1"
	searchCommon "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/set"
)

// GlobalOptions is exposed for e2e test
var GlobalOptions = []string{
	searchCommon.Cluster.String(),
	searchCommon.Namespace.String(),
	searchCommon.Label.String(),
}

// CategoryToOptionsSet is a map of all option sets by category, with a category for each indexed data type.
var CategoryToOptionsSet map[v1.SearchCategory]set.StringSet

func generateSetFromOptionsMap(maps ...map[searchCommon.FieldLabel]*v1.SearchField) set.StringSet {
	s := set.NewStringSet()
	for _, m := range maps {
		for k, v := range m {
			if !v.GetHidden() {
				s.Add(k.String())
			}
		}
	}
	return s
}

// GetOptions returns the searchable fields for the specified categories
func GetOptions(categories []v1.SearchCategory) []string {
	optionsSet := set.NewStringSet(GlobalOptions...)
	for _, category := range categories {
		optionsSet = optionsSet.Union(CategoryToOptionsSet[category])
	}
	slice := optionsSet.AsSlice()
	sort.Strings(slice)
	return slice
}

func init() {
	CategoryToOptionsSet = make(map[v1.SearchCategory]set.StringSet)
	for category, optionsMap := range globalindex.CategoryToOptionsMap {
		CategoryToOptionsSet[category] = generateSetFromOptionsMap(optionsMap)
	}
}
