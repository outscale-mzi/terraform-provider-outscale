package outscale

// func TestAccOutscaleOutboundRule_importBasic(t *testing.T) {
// 	resourceName := "outscale_outbound_rule.egress_1"

// 	rInt := acctest.RandInt()

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckOutscaleSGRuleDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccOutscaleSecurityGroupRuleEgressConfig(rInt),
// 			},

// 			resource.TestStep{
// 				ResourceName:            resourceName,
// 				ImportState:             true,
// 				ImportStateVerify:       true,
// 				ImportStateVerifyIgnore: []string{"associate_public_ip_address", "user_data", "security_group"},
// 			},
// 		},
// 	})
// }
