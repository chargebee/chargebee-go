### v2.1.1 (2020-04-03)
* * *

* A subresource contract_term has been added in Subscription resource
* New endpoint refund_a_credit_note has been added in Credit notes resource
* New endpoint list_quote_line_groups has been added in Quotes resource
* New endpoint list_contract_terms_for_a_subscription has been added in Subscriptions resource
* New input parameter contract_term_billing_cycle_on_renewal, contract_term[action_at_term_end], contract_term[cancellation_cutoff_period] has been added in create_a_subscription, create_subscription_for_customer, update_a_subscription and reactivate_a_subscription apis
* New input parameter contract_term_cancel_option has been added in cancel_a_subscription api
* New input parameter contract_term[action_at_term_end], contract_term[cancellation_cutoff_period], subscription[contract_term_billing_cycle_on_renewal] has been added in create_subscription_estimate, create_subscription_for_a_customer_estimate, checkout_new_subscription, checkout_existing_subscription apis
* New input parameter business_customer_without_vat_number has been added in create_a_customer and update_billing_info_for_a_customer apis
* New input parameter addons[service_period] replaces the parameter addons[date_from] and addons[date_to] in create_a_quote_for_one-time_charges api
* New input parameter charges[service_period] replaces the parameter charges[date_from] and charges[date_to] in create_a_quote_for_one-time_charges api
* New input parameter consolidated_view has been added in retrieve_quote_as_pdf api
* New input parameter customer[business_customer_without_vat_number] has been added in create_a_subscription, update_a_subscription apis
* New value CHECKOUT_COM has been added in gateway enum
* New value CONTRACT_TERMINATION has been added to eventBasedAddons onEvent enum
* New values IDEAL, GOOGLE_PAY has been added to payment types enum
* New event types CONTRACT_TERM_CREATED, CONTRACT_TERM_RENEWED, CONTRACT_TERM_TERMINATED, CONTRACT_TERM_COMPLETED, CONTRACT_TERM_CANCELLED have been added

### v2.1.0 (2020-02-06)
* * *

* New attribute total_payable and charge_on_aaceptance has been added in Quote.java
* New input parameter cancel_at has been added in cancel_a_subscription api

### v2.0.9 (2020-01-08)
* * * 

* New attribute taxjar_product_code has been added in Addon and Plan resource.
* New input parameter taxjar_product_code has been added in create_an_addon, update_an_addon, create_a_plan and update_a_plan api
* New input parameter taxjar_exemption_category has been added in create_a_customer, update_a_customer , create_a_subscription and import_a_subscription api
* New input parameter no_expiry has been added in gift_subscription_estimate and create_a_gift api
* New attribute no_expiry has been added in Gift resource
* New endpoint update_gift has been added in Gift resource
* New value BUSINESS_CHECKING has been added in AccountType enum
* New value CCD has been added in EcheckType enum
* New event type GIFT_UPDATED has been added
* New enum taxjar_exemption_category has been added

### v2.0.8 (2019-09-13)
* * *

* New endpoints gift_subscription_estimate and create_invoice_estimate have been added in estimate resource.
* New input parameter reference_id in payment_intent has been added in create_a_gift, create_an_invoice, create_using_payment_intent, create_a_subscription, create_subscription_for_customer, update_a_subscription, reactivate_a_subscription, resume_a_subscription, create_a_customer and collect_payment_for_customer api.
* New endpoint record_refund_for_a_payment has been added in transaction resource.

### v2.0.7 (2019-08-27)
* * * 

* New attribute client_profile_id has been added in customer resource.
* New input parameter client_profile_id has been added in create_a_customer, update_a_customer, create_subscription_estimate, create_a_subscription and import_a_subscription api.
* New input parameter payment_intent has been added in create_a_gift, create_an_invoice, reactivate_a_subscription and resume_a_subscription api.
* New input parameter replace_primary_payment_source has been added in create_an_invoice api.
* New attribute currency_code has been added in payment_intent resource.

### v2.0.6 (2019-08-14)
* * *

* New resource payment_intent has been added.
* New input parameter 'id' in payment_intent sub-param has been added in create_a_customer, collect_payment_for_customer, create_using_payment_intent, create_a_subscription, create_subscription_for_customer and update_a_subscription .
* New event_types PAYMENT_INTENT_CREATED and PAYMENT_INTENT_UPDATED have been added.

### v2.0.5 (2019-08-07)
* * *

* New attributes tax_amount_in_local_currency and local_currency_code have been added in line_item_tax sub-resource of credit_note, credit_note_estimate, invoice, invoice_estimate, order and quote resources.
* New attributes sub_total_in_local_currency, total_in_local_currency and local_currency_code have been added in credit_note and invoice resources.
* New input parameter gw_payment_method_id has been added in create_a_customer, collect_payment_for_customer, create_using_payment_intent, create_a_subscription, create_subscription_for_customer and update_a_subscription api.
* New attributes name, invoice_id and notes have been added in quote resource.
* New input parameters name, notes and expires_at have been added in create_a_quote_for_a_new_subscription, create_a_quote_for_update_subscription and create_a_quote_for_one-time_charges api.
* New input parameters id, auto_collection and po_number have been added in convert_a_quote api.

### v2.0.4 (2019-07-22)
* * *

* The attribute ref_tx_id has been added in card resource.
* The attribute custmer_id has been added in credit_note_estimate and invoice_estimate resource.
* The input parameters payment_intent with gateway_account_id and gw_token have been added in create a customer, collect payment for customer, create subscription and create subscription for customer API.
* New enum dunning_type has been added.
* New value vantiv has been added in the gateway enum.
* New value sepa_credit has been added in payment method enum. 
* The attribute dunning_attempts has been added in invoice resource.
* New enpoint create_using_payment_intent has been added in payment_source resouce.
* The input parameter reference_transaction has been added in update card payment source API.
* New endpoint list quote has been added in qoutes resource.
* The attributes initiator_type and three_d_secure have been added in transaction resource.
* New attribute scheme has been added in virtual_bank_account resource.
* New input parameter scheme has been added in create a virtual bank account using permanent token and create a virtual bank account API.

### v2.0.3 (2019-07-08)
* * *

* The resources hierarchy and token are added.
* Value ‘day’ is added in period_unit and shipping_frequency_period_unit .
* The parameters fractional_correction, comment, date_from and date_to are added in Create credit note API.
* The attribute customer_id is added to line_item sub resource of Credit note, Credit note estimate, Invoice, Invoice estimate and Quote.
* Endpoints Link a customer, Delink a customer, Get hierarchy, CreateUsingToken are added.
* The attributes business_customer_without_vat_number and relationship are added to customer resource.
* The filter parameters parent_id, payment_owner_id, invoice_owner_id are added in List customers, Export revenue recognition reports, Export Deferred Revenue Reports and Export customers APIs.
* The parameter token_id is added in Collect payment API.
* Event types hierarchy_created, hierarchy_deleted, token_created, token_consumed and token_expired are added.
* The parameter service_period_in_days is added in Create subscription estimate, Create subscription for customer estimate, Update subscription estimate, Checkout new, Checkout existing, Create subscription for customer quote, Update subscription quote, Create subscription, Create subscription for customer, Update subscription, Import subscription and Import subscription for customer.
* The filter parameter payment_owner is added in Export revenue recognition reports,Export Deferred Revenue Reports, Export Invoice and List Invoice APIs.
* The attribute payment_owner is added to invoice.
* The attributes date_from and date_to are added to Create an invoice, create a invoice for add-on and create a invoice for charge,Create a quote for one-time charges, Add charge at term end and Add add-on at term end.
* The parameter comment is added to Stop dunning, Apply payments, Apply credits, Add charge, Add add-on charge, Update details and Close Invoice APIs.
* The parameter ‘claim_credits’ is added to delete invoice API.
* The attribute override_relationship is added to subscription resource.
* The parameter token_id is added to Create subscription API.
* The parameter override_relationship is added to Create subscription for customer and Update subscription APIs.
* The attribute service_period_in_days is added to event add ons sub resource.

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
