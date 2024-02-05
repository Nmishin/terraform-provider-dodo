package sdkv2provider

import (
	"context"
	"fmt"
	"time"
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDodoBrands() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDodoBrandsRead,

		Schema: map[string]*schema.Schema{
			"names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
			},
		},
		Description: heredoc.Doc(fmt.Sprintf(`
			Use this data source to look up [zone](https://api.cloudflare.com/#zone-properties)
			info. This is the singular alternative to %s.
		`, "`cloudflare_zones`")),
	}
}

func dataSourceDodoBrandsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Reading Brands"))
	var diags diag.Diagnostics

	c := meta.(*APIClient)

	brands, _, err := c.client.BrandsApi.ApiV2BrandsGet(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Device Posture Rules: %w", err))
	}

	brandNames := make([]string, 0)
	for _, brand := range brands {
		brandNames = append(brandNames, brand.Name)
	}

	tflog.Debug(ctx, fmt.Sprintf("=====Reading Brands ===== %v", brandNames))
	if err = d.Set("names", brandNames); err != nil {
		return diag.FromErr(fmt.Errorf("error setting matched rules: %w, %v", err, brandNames))
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
