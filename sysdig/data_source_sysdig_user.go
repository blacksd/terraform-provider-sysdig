package sysdig

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSysdigUser() *schema.Resource {
	timeout := 30 * time.Second

	return &schema.Resource{
		ReadContext: dataSourceSysdigUserRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(timeout),
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSysdigUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, err := meta.(SysdigClients).sysdigCommonClient()
	if err != nil {
		return diag.FromErr(err)
	}

	u, err := client.GetUserByEmail(ctx, d.Get("email").(string))

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(u.ID))
	d.Set("version", u.Version)
	d.Set("system_role", u.SystemRole)
	d.Set("first_name", u.FirstName)
	d.Set("last_name", u.LastName)

	return nil

}
