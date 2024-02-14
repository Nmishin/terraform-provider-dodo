package sdkv2provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	api "github.com/Nmishin/dodo-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		if s.Computed {
			desc += " Generated."
		}
		return strings.TrimSpace(desc)
	}
}

func Provider(version string, testing bool) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: fmt.Sprintf("The API Token for operations."),
				DefaultFunc: schema.EnvDefaultFunc("DODO_BASE_URL", "https://globalapi.dodopizza.com"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dodo_brand": dataSourceDodoBrands(),
		},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure(version, testing),
	}
}

type APIClient struct {
	client *api.APIClient
}

func providerConfigure(version string, testing bool) schema.ConfigureContextFunc {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		apiURL := d.Get("url").(string)

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics

		dodoURL, err := url.Parse(apiURL)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		config := api.NewConfiguration()
		config.Host = dodoURL.Host
		config.Scheme = dodoURL.Scheme

		apiClient := api.NewAPIClient(config)

		return &APIClient{
			client: apiClient,
		}, diags
	}
}
