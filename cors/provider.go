package cors

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/terraform-provider-cors/cors/config/config"
	//"github.com/chnsz/terraform-provider-cors/cors/services/cbr"
)

// Provider returns a schema.Provider for CORS.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CORS_REGION", nil),
			},

			"domain_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CORS_DOMAIN_ID", nil),
			},

			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CORS_PROJECT_ID", nil),
			},

			"app_code": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CORS_APP_CODE", nil),
			},

			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CORS_MAX_RETRIES", 5),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			//"huaweicloud_cbr_vaults":  cbr.DataSourceCbrVaults(),
		},

		ResourcesMap: map[string]*schema.Resource{
			//"huaweicloud_cbr_vault":  cbr.ResourceCbrVault(),
		},
		ConfigureContextFunc = configProvider
	}

	return provider
}

func configProvider(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := config.Config{
		Region:     d.Get("region").(string),
		DomainID:   d.Get("domain_id").(string),
		ProjectID:  d.Get("project_id").(string),
		AppCode:    d.Get("app_code").(string),
		MaxRetries: d.Get("max_retries").(int),
	}

	return &config, nil
}
