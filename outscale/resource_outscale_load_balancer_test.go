package outscale

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleOAPILBUBasic(t *testing.T) {
	var conf oscgo.LoadBalancer

	r := acctest.RandIntRange(0, 50)
	region := os.Getenv("OUTSCALE_REGION")
	zone := fmt.Sprintf("%sa", region)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: "outscale_load_balancer.bar",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOutscaleOAPILBUDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleOAPILBUConfig(r),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPILBUExists("outscale_load_balancer.bar", &conf),
					resource.TestCheckResourceAttr(
						"outscale_load_balancer.bar", "subregion_names.#", "1"),
					resource.TestCheckResourceAttr(
						"outscale_load_balancer.bar", "subregion_names.0", zone),
					resource.TestCheckResourceAttr(
						"outscale_load_balancer.bar",
						"listeners.#", "1"),
				)},
		},
	})
}

func testAccCheckOutscaleOAPILBUDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_load_balancer" {
			continue
		}

		var err error
		var resp oscgo.ReadLoadBalancersResponse
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			filter := &oscgo.FiltersLoadBalancer{
				LoadBalancerNames: &[]string{rs.Primary.ID},
			}

			req := &oscgo.ReadLoadBalancersRequest{
				Filters: filter,
			}

			resp, _, err = conn.LoadBalancerApi.ReadLoadBalancers(
				context.Background()).ReadLoadBalancersRequest(*req).Execute()

			if err != nil {
				if strings.Contains(fmt.Sprint(err), "Throttling") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})

		if err == nil {
			if len(*resp.LoadBalancers) != 0 &&
				*(*resp.LoadBalancers)[0].LoadBalancerName ==
					rs.Primary.ID {
				return fmt.Errorf("LBU still exists")
			}
		}

		if strings.Contains(fmt.Sprint(err), "LoadBalancerNotFound") {
			return nil
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckOutscaleOAPILBUExists(n string, res *oscgo.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LBU ID is set")
		}

		conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

		var err error
		var resp oscgo.ReadLoadBalancersResponse
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			filter := &oscgo.FiltersLoadBalancer{
				LoadBalancerNames: &[]string{rs.Primary.ID},
			}

			req := &oscgo.ReadLoadBalancersRequest{
				Filters: filter,
			}

			resp, _, err = conn.LoadBalancerApi.ReadLoadBalancers(
				context.Background()).ReadLoadBalancersRequest(*req).Execute()

			if err != nil {
				if strings.Contains(fmt.Sprint(err), "Throttling") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})

		if err != nil {
			return err
		}

		if len(*resp.LoadBalancers) != 1 ||
			*(*resp.LoadBalancers)[0].LoadBalancerName != rs.Primary.ID {
			return fmt.Errorf("LBU not found")
		}

		res = &(*resp.LoadBalancers)[0]

		if res.NetId != nil {
			sgid := rs.Primary.Attributes["source_security_group_id"]
			if sgid == "" {
				return fmt.Errorf("Expected to find source_security_group_id for LBU, but was empty")
			}
		}

		return nil
	}
}

func testAccOutscaleOAPILBUConfig(r int) string {
	return fmt.Sprintf(`
resource "outscale_load_balancer" "bar" {
  subregion_names = ["%sa"]
  load_balancer_name               = "foobar-terraform-elb-%d"
  listeners {
    backend_port = 8000
    backend_protocol = "HTTP"
    load_balancer_port = 80
    load_balancer_protocol = "HTTP"
  }

	tags {
		key = "name"
		value = "baz"
	}

}
`, os.Getenv("OUTSCALE_REGION"), r)
}
