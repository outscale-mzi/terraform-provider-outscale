package outscale

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func dataSourceOutscaleOAPIVMState() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceOutscaleOAPIVMStateRead,
		Schema: getVmStateDataSourceSchema(),
	}
}

func dataSourceOutscaleOAPIVMStateRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	filters, filtersOk := d.GetOk("filter")
	instanceIds, instanceIdsOk := d.GetOk("vm_id")

	if !instanceIdsOk && !filtersOk {
		return errors.New("vm_id or filter must be set")
	}

	params := &fcu.DescribeInstanceStatusInput{}
	if filtersOk {
		params.Filters = buildOutscaleDataSourceFilters(filters.(*schema.Set))
	}
	if instanceIdsOk {
		var ids []*string

		for _, id := range instanceIds.(*schema.Set).List() {
			ids = append(ids, aws.String(id.(string)))
		}

		params.InstanceIds = ids
	}

	params.IncludeAllInstances = aws.Bool(false)

	var resp *fcu.DescribeInstanceStatusOutput
	var err error

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err = conn.VM.DescribeInstanceStatus(params)
		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return resource.NonRetryableError(err)
	})

	if err != nil {
		return err
	}

	filteredStates := resp.InstanceStatuses[:]

	var state *fcu.InstanceStatus
	if len(filteredStates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	}

	if len(filteredStates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	state = filteredStates[0]

	log.Printf("[DEBUG] outscale_oapi_vm_state - Single State found: %s", *state.InstanceId)

	return statusDescriptionOAPIVMStateAttributes(d, state)
}

func statusDescriptionOAPIVMStateAttributes(d *schema.ResourceData, status *fcu.InstanceStatus) error {

	d.SetId(*status.InstanceId)

	d.Set("sub_region_name", status.AvailabilityZone)

	events := eventsSet(status.Events)
	err := d.Set("maintenance_event", events)
	if err != nil {
		return err
	}

	state := flattenedState(status.InstanceState)
	err = d.Set("state", state)
	if err != nil {
		return err
	}

	st := statusSet(status.InstanceStatus)
	err = d.Set("instance_status", st)
	if err != nil {
		return err
	}

	sst := statusSet(status.SystemStatus)
	err = d.Set("state", sst)
	if err != nil {
		return err
	}

	return nil
}

func eventsOAPIVMStateHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["code"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["description"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["not_before"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["not_after"].(string)))
	return hashcode.String(buf.String())
}

func statusSetOAPIVMState(status *fcu.InstanceStatusSummary) *schema.Set {
	s := &schema.Set{
		F: statusHash,
	}

	st := map[string]interface{}{
		"status":  *status.Status,
		"details": detailsSet(status.Details),
	}
	s.Add(st)

	return s
}

func statusHashOAPIVMState(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["status"].(string)))
	return hashcode.String(buf.String())
}

func detailsHashOAPIVMState(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["status"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["details"].(string)))
	return hashcode.String(buf.String())
}

func detailsSetOAPIVMState(details []*fcu.InstanceStatusDetails) *schema.Set {
	s := &schema.Set{
		F: detailsHash,
	}

	for _, v := range details {

		status := map[string]interface{}{
			"name":   *v.Name,
			"status": *v.Status,
		}
		s.Add(status)
	}

	return s
}

func flattenedStateOAPIVMState(state *fcu.InstanceState) map[string]interface{} {
	return map[string]interface{}{
		"code": fmt.Sprintf("%d", *state.Code),
		"name": *state.Name,
	}
}

func eventsSetOAPIVMState(events []*fcu.InstanceStatusEvent) *schema.Set {
	s := &schema.Set{
		//F: eventsHashState,
		F: eventsOAPIVMStateHash,
	}
	for _, v := range events {

		status := map[string]interface{}{
			"code":        *v.Code,
			"description": *v.Description,
			"not_before":  v.NotBefore.Format(time.RFC3339),
			"not_after":   v.NotAfter.Format(time.RFC3339),
		}
		s.Add(status)
	}
	return s
}

func getOAPIVMStateDataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Arguments
		"filter": dataSourceFiltersSchema(),
		// "instance_id": {
		// 	Type:     schema.TypeSet,
		// 	Optional: true,
		// 	Elem:     &schema.Schema{Type: schema.TypeString},
		// },
		// "include_all_instances": {
		// 	Type:     schema.TypeBool,
		// 	Optional: true,
		// },

		// Attributes
		"sub_region_name": { //availability_zone
			Type:     schema.TypeString,
			Computed: true,
		},
		"maintenance_event": { //events_set
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"state_code": { //code
						Type:     schema.TypeString,
						Computed: true,
					},
					"description": { //
						Type:     schema.TypeString,
						Computed: true,
					},
					"not_after": { // not_before
						Type:     schema.TypeString,
						Computed: true,
					},
					"not_before": { // not_after
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},

		"vm_id": { //instance_id
			Type:     schema.TypeBool,
			Optional: true,
		},
		"state": { //instance_state
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"state_code": { // code
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"comment": { // instance_status
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"item": { //details
						Type:     schema.TypeSet,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:     schema.TypeString,
									Computed: true,
								},
								"state": { //status
									Type:     schema.TypeString,
									Computed: true,
								},
							},
						},
					},
					"state": { //state
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		// "comment": { // system_status
		// 	Type:     schema.TypeSet,
		// 	Computed: true,
		// 	Elem: &schema.Resource{
		// 		Schema: map[string]*schema.Schema{
		// 			"item": {
		// 				Type:     schema.TypeSet,
		// 				Computed: true,
		// 				Elem: &schema.Resource{
		// 					Schema: map[string]*schema.Schema{
		// 						"name": { // details
		// 							Type:     schema.TypeString,
		// 							Computed: true,
		// 						},
		// 						"state": { //status
		// 							Type:     schema.TypeString,
		// 							Computed: true,
		// 						},
		// 					},
		// 				},
		// 			},
		// 			"state": { // status
		// 				Type:     schema.TypeString,
		// 				Computed: true,
		// 			},
		// 		},
		// 	},
		// },
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
