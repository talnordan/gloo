package translator

import (
	"fmt"
	"sort"

	envoycache "github.com/envoyproxy/go-control-plane/pkg/cache"

	"github.com/solo-io/glue/discovery"
	"github.com/solo-io/glue/pkg/api/types/v1"
	"github.com/solo-io/glue/secrets"
	"github.com/solo-io/glue/translator/plugin"
)

type Translator struct {
	plugins []plugin.Plugin
}

func NewTranslator() *Translator {
	return &Translator{}
}

func (t *Translator) Translate(cfg v1.Config,
	clusters discovery.Clusters,
	secretMap secrets.SecretMap) (envoycache.Snapshot, error) {

	// runTranslation

	// combine with cluster + endpoints
	// stable sort

	// computer snapshort version
	return envoycache.Snapshot{}, fmt.Errorf("not implemented")
}

func (t *Translator) getAllDependencies() []string {

}

func (t *Translator) runValidation() []error {

}

func (t *Translator) runTranslation() envoycache.Snapshot {
	// compute virtual VirtualHosts
	// compute Routes
	// ...

	// do a stable sort

}

func sortFilters(filters []plugin.FilterWrapper) []plugin.FilterWrapper {
	// sort them accoirding to stage and then according to the name.
	less := func(i, j int) bool {
		filteri := filters[i]
		filterj := filters[j]
		if filteri.Stage != filterj.Stage {
			return filteri.Stage < filterj.Stage
		}
		return filteri.Filter.Name < filterj.Filter.Name
	}
	sort.Slice(filters, less)
	return filters
}
