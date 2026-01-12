package chargebee

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeListParams(t *testing.T) {
	tests := []struct {
		Name   string
		Input  interface{}
		Output *url.Values
	}{
		{
			Name: "Subscription list request",
			Input: &SubscriptionListRequest{
				Limit: Int32(2),
				ItemId: &StringFilter{
					In: []string{"9_99_plan", "cbdemo_advanced"},
					Is: "cbdemo_advanced",
				},
			},
			Output: &url.Values{
				"limit":       []string{"2"},
				"item_id[in]": []string{"[\"9_99_plan\",\"cbdemo_advanced\"]"},
				"item_id[is]": []string{"cbdemo_advanced"},
			},
		},
		{
			Name: "Customer List request",
			Input: &CustomerListRequest{
				FirstName: &StringFilter{
					Is: "John",
				},
				LastName: &StringFilter{
					Is: "Doe",
				},
				Email: &StringFilter{
					Is: "john@test.com",
				},
			},
			Output: &url.Values{
				"first_name[is]": []string{"John"},
				"last_name[is]":  []string{"Doe"},
				"email[is]":      []string{"john@test.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := SerializeListParams(tt.Input)
			assert.True(t, reflect.DeepEqual(out, tt.Output), fmt.Sprintf("%s Serialized input and output didn't match", tt.Name))
		})
	}
}

func TestSerializeParams(t *testing.T) {
	tests := []struct {
		Name   string
		Input  interface{}
		Output *url.Values
	}{
		{
			Name: "Export Customer Request",
			Input: &ExportCustomersRequest{
				Customer: &ExportCustomersCustomer{
					FirstName: &StringFilter{
						Is: "John",
					},
					LastName: &StringFilter{
						Is: "Doe",
					},
				},
			},
			Output: &url.Values{
				"customer[first_name][is]": []string{"John"},
				"customer[last_name][is]":  []string{"Doe"},
			},
		},
		{
			Name: "Coupon: Create For Item Request Params",
			Input: &CouponCreateForItemsRequest{
				Id:                 "summer_offer",
				Name:               "Summer Offer",
				DiscountPercentage: Float64(10.0),
				DiscountType:       CouponDiscountTypePercentage,
				DurationType:       CouponDurationTypeForever,
				ApplyOn:            CouponApplyOnEachSpecifiedItem,
				ItemConstraints: []*CouponCreateForItemsItemConstraint{
					{
						Constraint: CouponItemConstraintConstraintSpecific,
						ItemType:   CouponItemConstraintItemTypePlan,
					},
				},
			},
			Output: &url.Values{
				"id":                              []string{"summer_offer"},
				"name":                            []string{"Summer Offer"},
				"discount_percentage":             []string{"10"},
				"discount_type":                   []string{"percentage"},
				"duration_type":                   []string{"forever"},
				"apply_on":                        []string{"each_specified_item"},
				"item_constraints[constraint][0]": []string{"specific"},
				"item_constraints[item_type][0]":  []string{"plan"},
			},
		},
		{
			Name: "Purchase: Create Request Params",
			Input: &PurchaseCreateRequest{
				CustomerId: "__dev__1234",
				PurchaseItems: []*PurchaseCreatePurchaseItem{
					{
						Index:       Int32(2),
						ItemPriceId: "cbdemo_basic-USD-weekly",
						Quantity:    Int32(5),
					},
					{
						Index:       Int32(0),
						ItemPriceId: "cbdemo_basic-USD-monthly",
						Quantity:    Int32(10),
					},
				},
				SubscriptionInfo: []*PurchaseCreateSubscriptionInfo{
					{
						Index: Int32(0),
						MetaData: map[string]interface{}{
							"features": map[string]interface{}{
								"usage-limit":        "5GB",
								"speed-within-quota": "2MBbps",
								"post-usage-quota":   "512kbps",
							},
						},
					},
				},
			},
			Output: &url.Values{
				"customer_id":                      []string{"__dev__1234"},
				"purchase_items[index][0]":         []string{"2"},
				"purchase_items[item_price_id][0]": []string{"cbdemo_basic-USD-weekly"},
				"purchase_items[quantity][0]":      []string{"5"},
				"purchase_items[index][1]":         []string{"0"},
				"purchase_items[item_price_id][1]": []string{"cbdemo_basic-USD-monthly"},
				"purchase_items[quantity][1]":      []string{"10"},
				"subscription_info[index][0]":      []string{"0"},
				"subscription_info[meta_data][0]":  []string{"{\"features\":{\"post-usage-quota\":\"512kbps\",\"speed-within-quota\":\"2MBbps\",\"usage-limit\":\"5GB\"}}"},
			},
		},
		{
			Name: "Customer: Create Request Params",
			Input: &CustomerCreateRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@test.com",
				Locale:    "fr-CA",
				BillingAddress: &CustomerCreateBillingAddress{
					FirstName: "John",
					LastName:  "Doe",
				},
				Taxability: TaxabilityExempt,
				ExemptionDetails: []map[string]interface{}{
					{
						"loc": map[string]interface{}{
							"addr": "sdfsd,dsf",
							"st":   "TX",
							"ctry": "USA",
							"zip":  "87799",
						},
						"cat":  0,
						"exnb": true,
						"frc":  true,
						"scp":  1920,
					},
					{
						"loc": map[string]interface{}{
							"addr": "john.doe street",
							"st":   "FR",
							"ctry": "EUR",
							"zip":  "1234",
						},
						"cat":  0,
						"exnb": true,
						"frc":  true,
						"scp":  1920,
					},
				},
			},
			Output: &url.Values{
				"exemption_details":           []string{"[{\"cat\":0,\"exnb\":true,\"frc\":true,\"loc\":{\"addr\":\"sdfsd,dsf\",\"ctry\":\"USA\",\"st\":\"TX\",\"zip\":\"87799\"},\"scp\":1920},{\"cat\":0,\"exnb\":true,\"frc\":true,\"loc\":{\"addr\":\"john.doe street\",\"ctry\":\"EUR\",\"st\":\"FR\",\"zip\":\"1234\"},\"scp\":1920}]"},
				"last_name":                   []string{"Doe"},
				"first_name":                  []string{"John"},
				"email":                       []string{"john@test.com"},
				"locale":                      []string{"fr-CA"},
				"billing_address[last_name]":  []string{"Doe"},
				"billing_address[first_name]": []string{"John"},
				"taxability":                  []string{"exempt"},
			},
		},
		{
			Name: "Subscription: Create With Items",
			Input: &SubscriptionCreateWithItemsRequest{
				SubscriptionItems: []*SubscriptionCreateWithItemsSubscriptionItem{
					{
						ItemPriceId:   "9_99_plan-USD-Monthly",
						BillingCycles: Int32(2),
						Quantity:      Int32(1),
						UnitPrice:     Int64(1000),
					},
				}, CouponIds: []string{"summer_offer1", "summer_offer1"},
			},
			Output: &url.Values{
				"coupon_ids[0]":                         []string{"summer_offer1"},
				"coupon_ids[1]":                         []string{"summer_offer1"},
				"subscription_items[item_price_id][0]":  []string{"9_99_plan-USD-Monthly"},
				"subscription_items[billing_cycles][0]": []string{"2"},
				"subscription_items[quantity][0]":       []string{"1"},
				"subscription_items[unit_price][0]":     []string{"1000"},
			},
		},
		{
			Name: "Subscription: Create Request Params",
			Input: &SubscriptionCreateRequest{
				PlanId: "cbdemo_grow",
				Customer: &SubscriptionCreateCustomer{
					Email:          "john@user.com",
					FirstName:      "John",
					LastName:       "Doe",
					Locale:         "frCA",
					Phone:          "+19499999999",
					AutoCollection: AutoCollectionOn,
				},
				Addons: []*SubscriptionCreateAddon{
					{
						Id: "cbdemo_conciergesupport",
					},
					{
						Id:       "cbdemo_additionaluser",
						Quantity: Int32(2),
					},
				},
				CouponIds: []string{"cbdemo_earlybird", "cb_test"},
			},
			Output: &url.Values{
				"coupon_ids[0]":             []string{"cbdemo_earlybird"},
				"coupon_ids[1]":             []string{"cb_test"},
				"customer[phone]":           []string{"+19499999999"},
				"plan_id":                   []string{"cbdemo_grow"},
				"customer[last_name]":       []string{"Doe"},
				"customer[locale]":          []string{"frCA"},
				"addons[id][0]":             []string{"cbdemo_conciergesupport"},
				"addons[id][1]":             []string{"cbdemo_additionaluser"},
				"addons[quantity][1]":       []string{"2"},
				"customer[email]":           []string{"john@user.com"},
				"customer[auto_collection]": []string{"on"},
				"customer[first_name]":      []string{"John"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := SerializeParams(tt.Input)
			assert.True(t, reflect.DeepEqual(out, tt.Output), fmt.Sprintf("%s Serialized input and output didn't match", tt.Name))
			if t.Failed() {
				t.Log(out)
				t.Log(tt.Output)
			}
		})
	}
}

func TestUtil(t *testing.T) {

	before := &SubscriptionCreateRequest{
		PlanId: "cbdemo_grow",
		Customer: &SubscriptionCreateCustomer{
			Email:          "john@user.com",
			FirstName:      "John",
			LastName:       "Doe",
			Locale:         "frCA",
			Phone:          "+19499999999",
			AutoCollection: AutoCollectionOn,
		},
		Addons: []*SubscriptionCreateAddon{
			{
				Id: "cbdemo_conciergesupport",
			},
			{
				Id:       "cbdemo_additionaluser",
				Quantity: Int32(2),
			},
		},
		CouponIds: []string{"cbdemo_earlybird"},
	}
	after := &url.Values{
		"coupon_ids[0]":             []string{"cbdemo_earlybird"},
		"customer[phone]":           []string{"+19499999999"},
		"plan_id":                   []string{"cbdemo_grow"},
		"customer[last_name]":       []string{"Doe"},
		"customer[locale]":          []string{"frCA"},
		"addons[id][0]":             []string{"cbdemo_conciergesupport"},
		"addons[id][1]":             []string{"cbdemo_additionaluser"},
		"addons[quantity][1]":       []string{"2"},
		"customer[email]":           []string{"john@user.com"},
		"customer[auto_collection]": []string{"on"},
		"customer[first_name]":      []string{"John"},
	}
	assert.Equal(t, SerializeParams(before), after)
}
