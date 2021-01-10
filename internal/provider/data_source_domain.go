package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a domain.",

		ReadContext: dataSourceDomainRead,

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Description: "ID of the domain.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"domain": {
				Description: "Name of the domain.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expiry_date": {
				Description: "Expiry date.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"registered_date": {
				Description: "Registered date.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"renew": {
				Description: "Whether the domain should be renewed.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"registrant": {
				Description: "Name of the registrant.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Domain status.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"nameservers": {
				Description: "List of nameservers.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"services": {
				Description: "Domain services.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"registrar": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dns": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"webhotel": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client
	auth := meta.(*apiClient).auth

	domainID := int32(d.Get("domain_id").(int))

	resp, _, err := client.DomainsApi.GetDomain(auth, domainID).Execute()
	if err.Error() != "" {
		return diag.Errorf("error getting domain (ID: %d): %s %s", domainID, err.Error(), err.Body())
	}

	d.Set("domain", resp.Domain)
	d.Set("expiry_date", resp.ExpiryDate)
	d.Set("registered_date", resp.RegisteredDate)
	d.Set("renew", resp.Renew)
	d.Set("registrant", resp.Registrant)
	d.Set("status", resp.Status)
	d.Set("nameservers", resp.Nameservers)

	var services []interface{}
	if resp.Services != nil {
		services = append(services, map[string]interface{}{
			"registrar": resp.Services.Registrar,
			"dns":       resp.Services.Dns,
			"email":     resp.Services.Email,
			"webhotel":  resp.Services.Webhotel,
		})
	}
	d.Set("services", services)

	d.SetId(strconv.Itoa(int(resp.GetId())))

	return nil
}
