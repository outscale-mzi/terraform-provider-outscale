package outscale

import (
	"context"
	"fmt"
	"strings"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceOutscaleOAPIPublicIPS() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceOutscalePublicIPSRead,
		Schema: oapiGetPublicIPSDataSourceSchema(),
	}
}

func oapiGetPublicIPSDataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Attributes
		"filter": dataSourceFiltersSchema(),
		"public_ips": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"link_public_ip_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"public_ip_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"vm_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"nic_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"nic_account_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"private_ip": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"public_ip": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"tags": dataSourceTagsSchema(),
				},
			},
		},
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceOutscalePublicIPSRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	req := oscgo.ReadPublicIpsRequest{}

	filters, filtersOk := d.GetOk("filter")

	if filtersOk {
		req.Filters = buildOutscaleOAPIDataSourcePublicIpsFilters(filters.(*schema.Set))
	}

	var resp oscgo.ReadPublicIpsResponse
	err := resource.Retry(60*time.Second, func() *resource.RetryError {
		var err error
		resp, _, err = conn.PublicIpApi.ReadPublicIps(context.Background()).ReadPublicIpsRequest(req).Execute()
		return resource.RetryableError(err)
	})

	if err != nil {
		if e := fmt.Sprint(err); strings.Contains(e, "InvalidAllocationID.NotFound") || strings.Contains(e, "InvalidAddress.NotFound") {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving EIP: %s", err)
	}

	// Verify Outscale returned our EIP
	if len(resp.GetPublicIps()) == 0 {
		return fmt.Errorf("Unable to find EIP: %#v", resp.GetPublicIps())
	}

	addresses := resp.GetPublicIps()

	address := make([]map[string]interface{}, len(addresses))

	for k, v := range addresses {
		add := make(map[string]interface{})

		add["link_public_ip_id"] = v.LinkPublicIpId
		add["public_ip_id"] = v.PublicIpId
		add["vm_id"] = v.VmId
		add["nic_id"] = v.NicId
		add["nic_account_id"] = v.NicAccountId
		add["private_ip"] = v.PrivateIp
		add["public_ip"] = v.PublicIp
		add["tags"] = getOapiTagSet(v.Tags)
		address[k] = add
	}

	d.SetId(resource.UniqueId())

	return d.Set("public_ips", address)
}
