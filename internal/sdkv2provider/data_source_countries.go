package sdkv2provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDodoCountries() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDodoCountriesRead,

		Schema: map[string]*schema.Schema{
			"brand": {
				Type:     schema.TypeString,
				Required: true,
			},
			"countries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of countries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDodoCountriesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Reading countries"))

	c := meta.(*APIClient)

	brand := d.Get("brand").(string)

	countries, _, err := c.client.CountriesApi.ApiV2BrandCountriesGet(ctx, brand)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing countries: %w", err))
	}

	countriesIds := make([]string, 0)
	countriesDetails := make([]interface{}, 0)
	for _, c := range countries {
		countriesDetails = append(countriesDetails, map[string]interface{}{
			"id":       c.Id,
			"code":     c.Code,
			"name":     c.Name,
			"currency": c.Currency,
		})
		countriesIds = append(countriesIds, strconv.Itoa(int(c.Id)))
	}

	if err = d.Set("countries", countriesDetails); err != nil {
		return diag.Errorf("error setting countries: %w", err)
	}

	d.SetId(stringListChecksum(countriesIds))
	return nil
}
