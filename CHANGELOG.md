### v2.0.2 (2019-04-11)
* * *

* The attributes avalara_sale_type, avalara_transaction_type and avalara_service_type are added in Addon and plan resource.
* The input parameters avalara_sale_type, avalara_transaction_type , avalara_service_type are added in create addon , update addon  ,create plan , update  plan, create invoice  , create invoice for charge, add_charge , add_charge_at_term_end and create_for_onetime_charges api endpoints.
* The attributes is_partial_tax_applied, is_non_compliance_tax and taxable_amount are added in line_item_taxes of credit_note ,credit_note_estimate , invoice, invoice_estimate , quote and order resources.
* The attributes exemption_details and customer_type are added in customer resource .
* The input parameters exemption_details and customer_type are added in create customer, update customer , create subscription estimate, create subscription and import subscription api endpoints.
* The enum values federal and unincorporated are added in tax_juris_type.
* The enum value export is added in tax_override_reason .
* The input parameter cancelled_at is added in cancel order api endpoint.
* New endpoint delete_local is added in payment_source and virtual_bank_account resources.

### v2.0.1 (2019-03-07)
* * *

* The attributes created_at, resource_version and updated_at are added in card, payment_source and virtual_bank_account resources.
* The input filter parameter sort_by with updated_at attribute is added in list customer and list subscription api endpoints.
* New endpoint export orders is added in export resource.
* New endpoint accept quote is added in hosted_pages resource.
* The input filter parameters updated_at and created_at are added in list payment_source api endpoint and list virtual_bank_accounts endpoint .
* New endpoint delete an offline transaction is added in transaction resource.

### v2.0.0 (2019-01-10)
* * *

Chargebee [API V2](https://apidocs.chargebee.com/docs/api#versions) for golang is now live!
