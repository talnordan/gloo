package translator

import (
	"fmt"

	envoycache "github.com/envoyproxy/go-control-plane/pkg/cache"

	"github.com/solo-io/glue/pkg/api/types/v1"
	"github.com/solo-io/glue/discovery"
	"github.com/solo-io/glue/secrets"
)

type Translator struct{
	plugins []Plugins
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

func (t *Translator) runTranslation() envoycache.Snapshot{} {
	// compute virtual VirtualHosts
	// compute Routes
	// ...

	// do a stable sort

}
