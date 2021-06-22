package taplytics_tf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testProviders = map[string]*schema.Provider{
	"taplytics-tf": Provider(),
}

func TestProvider(t *testing.T) {
	for name, provider := range testProviders {
		if err := provider.InternalValidate(); err != nil {
			t.Fatalf("%s: Error: %s",name, err)
		}
	}
}
