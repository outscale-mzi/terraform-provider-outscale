package outscale

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceOutscaleLoadBalancerVms() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleLoadBalancerVmsRead,
		Schema: getDataSourceSchemas(map[string]*schema.Schema{
			"load_balancer_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"backend_vm_ids": {
				Type:     schema.TypeList,
				ForceNew: true,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		}),
	}
}

func dataSourceOutscaleLoadBalancerVmsRead(d *schema.ResourceData,
	meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	lb, resp, err := readLbs0(conn, d)
	if err != nil {
		return err
	}

	d.Set("backend_vm_ids", flattenStringList(lb.BackendVmIds))
	d.Set("request_id", resp.ResponseContext.RequestId)
	return nil
}
