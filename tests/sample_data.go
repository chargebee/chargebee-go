package tests

func simpleSubscription() []byte {
	return []byte(`{
		"subscription":{
			"id":"simple_subscription",
			"plan_id":"basic",
			"cf_gender":"Female"
		},
		"customer":{
			"first_name":"simple",
			"last_name":"subscription"
		}
	}`)
}
func nestedSubscription() []byte {
	return []byte(`{
		"subscription":{
			"id":"nested_subscription",
			"plan_id":"basic",
			"addons":[
				{
					"id":"monitor",
					"quantity":"10"
				},
				{
					"id":"ssl"
				}
			]
	 	}
	}`)
}
func eventData() string {
	return string(`{ 
		"content": {
			 "card": { 
				 "card_type": "visa", 
				 "customer_id": "8avPlNabxST86", 
				 "expiry_month": 10, 
				 "expiry_year": 2012, 
				 "gateway": "chargebee", 
				 "iin": "411111", 
				 "last4": "1111", 
				 "masked_number": "411111******1111",
				"object": "credit_card", 
				"status": "valid" 
			}, 
			"customer": { 
				"card_status": "valid", 
				"created_at": 1339951248, 
				"email": "rr@chargebee.com", 
				"id": "8avPlNabxST86", 
				"object": "customer" 
			}, 
			"invoice": { 
				"amount": 900, 
				"end_date": 1339951248, 
				"id": "3", 
				"line_items": [ { 
					"amount": 900, 
					"date_from": 1339951248, 
					"date_to": 1342543248, 
					"description": "Plan (1 x No Trial) Charges for term (17Jun2012  17Jul2012)", 
					"object": "line_item", 
					"quantity": 1, 
					"unit_amount": 900 
				} ], 
				"object": "invoice", 
				"paid_on": 1339951249, 
				"recurring": true, 
				"start_date": 1339951248, 
				"status": "paid", 
				"sub_total": 900, 
				"subscription_id": "8avPlNabxST86" 
			}, 
			"subscription": {
				"activated_at": 1339951248, 
				"created_at": 1339951248, 
				"current_term_end": 1342543248, 
				"current_term_start": 1339951248, 
				"due_invoices_count": 0, 
				"id": "8avPlNabxST86", 
				"object": "subscription", 
				"plan_id": "no_trial", 
				"plan_quantity": 1, 
				"status": "active" 
			}, 
			"transaction": { 
				"amount": 900, 
				"date": 1339951249, 
				"gateway": "chargebee", 
				"id": "txn_HoR7OrcNacEsxKR",
				"id_at_gateway": "cb_HoR7OrcNacEsxMS", 
				"invoice_id": "3", 
				"masked_card_number": "411111******1111", 
				"object": "transaction", 
				"status": "success", 
				"subscription_id": "8avPlNabxST86", 
				"type": "payment" 
			} 
		}, 
		"event_type": "payment_succeeded", 
		"id": "ev_HoR7OrcNacEsxgT", 
		"object": "event", 
		"occurred_at": 1339951249, 
		"webhook_status": "scheduled" 
	}`)
}
func listSubscription() []byte {
	return []byte(`{
		"list":[
			{
				"subscription":{
					"id":"simple_subscription",
					"plan_id":"basic"
				},
				"customer":{
					"first_name":"simple",
					"last_name":"subscription"
				}
			},
			{
				"subscription":{
					"id":"simple_subscription",
					"plan_id":"basic"
				},
				"customer":{
					"first_name":"simple",
					"last_name":"subscription"
				}
			}
		]
	}`)
}
func subscriptionWithCustomfield() []byte {
	return []byte(`{
			"subscription":{
			"id":"simple_subscription",
			"plan_id":"basic",
			"cf_gender":"Female"
		},
		"customer":{
			"first_name":"simple",
			"last_name":"subscription",
			"cf_DOB":"12-10-1994"
		}
}`)
}
