🔴 **Attention**: The support for v2 will eventually be discontinued on **December 31st 2023** and will no longer receive any further updates. We strongly recommend [upgrading to v3](https://github.com/chargebee/chargebee-go/wiki/Migration-guide-for-v3) as soon as possible.

* * * 

### v2.27.0 (2023-11-30)
* * *

#### New Enum values:
* swiss_reference has been added to PaymentReferenceNumber#TypeEnum.

#### New Input parameters:
* subscription[po_number] has been added to Quote#EditCreateSubForCustomerQuoteRequest in Quote resource. 
* subscription[po_number] has been added to Quote#EditCreateSubCustomerQuoteForItemsRequest in Quote resource.

* * * 

### v2.26.0 (2023-10-31)
* * *

#### New attributes:
* statement_descriptor has been added to invoice resource.
* source has been added to the Hosted_Pages resource.

#### New Enum values:
* pay_com has been added to Gateway.
* pay_to, faster_payments, sepa_instant_transfer has been added to Customer#TypeEnum.
* pay_to, faster_payments, sepa_instant_transfer has been added to PaymentMethodEnum.
* pay_to, faster_payments, sepa_instant_transfer has been added to PaymentMethodTypeEnum.
* pay_to, faster_payments, sepa_instant_transfer has been added to TypeEnum.
* pay_to, faster_payments, sepa_instant_transfer, amazon_payments has been added to PaymentIntent#PaymentMethodTypeEnum.

#### New Input parameters:
* source has been added to HostedPage#CheckoutOneTimeForItemsRequest.
* source has been added to HostedPage#CheckoutNewForItemsRequest.
* source has been added to HostedPage#CheckoutExistingForItemsRequest.
* statement_descriptor_descriptor has been added to Invoice#CreateRequest.
* statement_descriptor_descriptor has been added to Invoice#CreateForChargeItemsAndChargesRequest.
* statement_descriptor_descriptor has been added to Invoice#UpdateDetailsRequest.
* statement_descriptor_additional_info has been added to Invoice#CreateRequest.
* statement_descriptor_additional_info has been added to Invoice#CreateForChargeItemsAndChargesRequest.
* statement_descriptor_additional_info has been added to Invoice#UpdateDetailsRequest.
* statement_descriptor_descriptor has been added to Subscription#CreateRequest.
* statement_descriptor_descriptor has been added to Subscription#CreateForCustomerRequest.
* statement_descriptor_descriptor has been added to Subscription#CreateWithItemsRequest.
* statement_descriptor_descriptor has been added to Subscription#UpdateRequest.
* statement_descriptor_descriptor has been added to Subscription#UpdateForItemsRequest.
* statement_descriptor_descriptor has been added to Subscription#ReactivateRequest.
* statement_descriptor_additional_info has been added to Subscription#CreateRequest.
* statement_descriptor_additional_info has been added to Subscription#CreateForCustomerRequest.
* statement_descriptor_additional_info has been added to Subscription#CreateWithItemsRequest.
* statement_descriptor_additional_info has been added to Subscription#UpdateRequest.
* statement_descriptor_additional_info has been added to Subscription#UpdateForItemsRequest.
* statement_descriptor_additional_info has been added to Subscription#ReactivateRequest.

#### Updates to old parameters:
* duration_type, discount_type in coupons has been made optional.
* proration_type in addons is now supported.
* csv_tax_rule has been removed.
* Tax1JurisTypeEnum has been removed.
* Tax2JurisTypeEnum has been removed.
* Tax3JurisTypeEnum has been removed.
* Tax4JurisTypeEnum has been removed.

### v2.25.0 (2023-09-26)
* * *

#### New attributes:
* venmo has been added to the PaymentSource resource.
* proration_type has been added to the QuotedCharge, QuotedSubscription and Subscription resource.

#### New Enum values:
* ebanx has been added to Gateway.
* venmo has been added to Customer#TypeEnum.
* venmo has been added to PaymentMethodEnum.
* venmo has been added to PaymentMethodTypeEnum.
* venmo has been added to TypeEnum.
* venmo has been added to PaymentIntent#PaymentMethodTypeEnum.

#### New Input parameters:
* einvoicing_method has been added to HostedPage#CheckoutOneTimeForItemsRequest.
* einvoicing_method has been added to HostedPage#CheckoutNewForItemsRequest.
* additional_information has been added to PaymentSource#CardAdditionalInformation.

### v2.24.1 (2023-09-15)
* * *

#### Fixes:
* Fixed the issue of nested field serialization.


### v2.24.0 (2023-09-05)
* * *

#### New Resource:
* CsvTaxRule has been added.

#### New attributes:
* direct_debit_scheme has been added to the PaymentSource#PaymentSourceBankAccount resource. 

#### New Enum :
* DirectDebitSchemeEnum has been added.
* Tax1JurisTypeEnum has been added.
* Tax2JurisTypeEnum has been added.
* Tax3JurisTypeEnum has been added.
* Tax4JurisTypeEnum has been added.

#### New Input parameters:
* avalara_tax_code, hsn_code, taxjar_product_code has been added to Invoice#AddChargeRequest. 


### v2.23.0 (2023-07-31)
* * *

#### New Attributes:
* tax_category has been added to the CreditNote, Quote and Invoice resource. 
* proration_type has been added in Addon resource.

#### New Enum values:
* tax has been added to EntityType enum in Invoice resource.
* payment_source_locally_deleted has been added to EventType.

#### New Input parameters:

* CouponId and CouponApplyTill has been added to Subscritpion#CreateRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#CreateForCustomerRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#CreateWithItemsRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#UpdateRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#UpdateForItemsRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#ImportSubscriptionRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#ImportForCustomerRequest in Subscritpion resource. 
* CouponId and CouponApplyTill has been added to Subscritpion#ImportForItemsRequest in Subscritpion resource. 
* cancel_reason_code has been added to Subscritpion#ImportForItemsRequest in Subscritpion resource.
* proration_type has been added in addon#createRequest and addon#UpdateRequest in Addon resource.
* addons[proration_type] has been added in Estimate#UpdateSubscriptionRequest in Estimate resource.
* addons[proration_type]  has been added in Subscription#UpdateRequest in Subscritpion resource.

#### New Enum Class:
* ProrationType enum has been added to addon resource.
* ProrationType enum has been added.


### v2.22.0 (2023-06-30)
* * *

#### New endpoints:
* ViewVoucher#ViewVoucherRequest has been added to the HostedPage resource.
* InvoiceListPaymentReferenceNumbers#InvoiceListPaymentReferenceNumbersRequest has been added to the Invoice resource.

#### New Resource:
* PaymentReferenceNumber has been added.

#### New attributes:
* local_currency_exchange_rate has been added to the CreditNote and Invoice resource. 

#### New Enum values:
* view_voucher has been added to Type enum in HostedPage resource.
* paused has been added in StoreStatus enum in InAppSubscription resource.
* metrics_global and windcave has been added in Gateway enum. 


#### New Input parameters:
* payment_initiator has been added to Customers#CollectPaymentRequest. 
* payment_initiator has been added to Invoice#CreateRequest.
* payment_initiator has been added to Invoice#ChargeRequest.
* payment_initiator has been added to Invoice#ChargeAddonRequest.
* payment_initiator has been added to Invoice#CollectPaymentRequest.
* payment_initiator has been added to Subscription#CreateRequest.
* payment_initiator has been added to Invoice#CreateForChargeItemRequest.
* payment_initiator has been added to Subscription#CreateForCustomerRequest.
* payment_initiator has been added to Subscription#CreateWithItemsRequest.
* payment_initiator has been added to Subscription#UpdateForItemsRequest.
* payment_initiator has been added to Subscription#ReactivateRequest.
* payment_initiator has been added to Subscription#ResumeRequest.
* payment_reference_numbers[id] has been added to Invoice#ImportInvoiceRequest.
* payment_reference_numbers[type] has been added to Invoice#ImportInvoiceRequest.
* payment_reference_numbers[number] has been added to Invoice#ImportInvoiceRequest.


### v2.21.0 (2023-05-31)
* * *

#### New endpoints:
* PaymentSource#CreateVoucherPaymentSource has been added to the PaymentSource resource.
* EventsRequest#HostedPage has been added to the PaymentSource resource.

#### New Resource:
* PaymentVoucher has been added.

#### New attributes:
* boleto and billing_address has been added to the PaymentSource resource.
* product_id has been added to the ItemPrice resource.

#### New Enum Class:
* EventNameEnum has been added.
* PaymentVoucherTypeEnum has been added.
* VoucherTypeEnum has been added.

#### New Enum values:
* product and variant has been added to EntityType enum.
* product_created, product_updated, product_deleted, variant_created, variant_updated and variant_deleted enums have been added in EventType enum.
* voucher_created, voucher_expired and voucher_create_failed have been added in EventType enum.
* boleto has been added in PaymentMethod and OfflinePaymentMethod  and PaymentMethodTypeEnum#PaymentIntent.


### v2.20.0 (2023-05-16)
* * *

#### New Feature:
* Added SetIdempotencyKey("UUID") utility to pass **Idempotency key** along with request headers to allow a safe retry of POST requests.
* Added IsIdempotencyReplayed() utility to differentiate between original and replayed requests.
* Added GetResponseHeaders() utility to fetch the response headers.


### v2.19.0 (2023-04-28)
* * *

#### Fixes:
* SubscriptionId attribute has been maid as required in InAppSubscription resource.

#### New Attributes: 
* TotalContractValueBeforeTax has been added to the ContractTerm resource.
* TotalContractValueBeforeTax#SubscriptionContractTerm has been added to the Subscription resource.
* TotalContractValueBeforeTax#SubscriptionEstimateContractTerm has been addded to the SubscriptionEstimate resource.
* CouponConstraints has been added to the Coupon resource.

#### Added Input Parameters:
* contract_term[total_amount_raised_before_tax]#ImportSubscriptionRequest, contract_term[total_amount_raised_before_tax]#ImportForItemsRequest, contract_term[total_amount_raised_before_tax]#ImportContractTermRequest and contract_term[total_amount_raised_before_tax]#ImportForCustomerRequest parameter has been added to Subscription resource.
* contract_term[total_contract_value_before_tax]#ImportContractTermRequest parameter has been added to Subscription resource.
* coupon_constraints[entity_type]#CreateForItemsRequestParams, coupon_constraints[type]#CreateForItemsRequestParams and coupon_constraints[value]#CreateForItemsRequestParams parameter has been added to the Coupon resource.
* coupon_constraints[entity_type]#UpdateForItemsRequestParams, coupon_constraints[type]#UpdateForItemsRequestParams and coupon_constraints[value]#UpdateForItemsRequestParams parameter has been added to the Coupon resource.
* export_type#CustomersRequest and export_type#SubscriptionsRequest parameter has been added to the Export resource.

#### New Enum Class:
* ExportType has been added.

#### New Enum values:
* pending_authorization has been added to StatusEnum#PaymentIntentPaymentAttempt to the PaymentIntent resource.


### v2.18.0 (2023-03-24)
* * *

#### New Utility Method:
* Added functionality to set custom context in HTTP request.

#### Fixes:
* Fixed list_discounts subscriptions API error.

#### New Attributes: 
* Einvoice#ReferenceNumber has been added to the CreditNote resource.
* Einvoice#ReferenceNumber has been added to the Invoice resource.
* einvoicing_method has been added to the Customer resource.
* StoreStatus and InvoiceId have been addded to the in_app_subscriptions resource.

#### Added Input Parameters:
* UpdateForItemsRequestParams#discount_quantity and UpdateRequestParams#discount_quantity parameter has been added to the coupon resource.
* UpdateBillingInfoRequestParams#einvoicing_method and CreateRequestParams#einvoicing_method parameter has been added to the customer resource.
* CreateCustomerParams#einvoicing_method , UpdateCustomerParams#einvoicing_metho and  UpdateForItemsCustomerParams#einvoicing_method parameter has been added to the Subscription resource.
* CreateSubscriptionInfoParams#meta_data parameter has been added to the purchase resource.

#### New Endpoints:
* credit_note#SendEinvoice has been added to the CreditNote resource. 
* invoice#SendEinvoice has been added to the Invoice resource.
* in_app_subscriptions#ImportSubscription and in_app_subscriptions#RetrieveStoreSubs requests have been added to the in_app_subscriptions resource.

#### New Enum Class:
* EinvoicingMethod has been added.
* StoreStatus has been added in in_app_subscriptions resource.


### v2.17.0 (2023-02-17)
* * *

#### Fixes:
* Fixed custom http client request creation issue.

#### New Attributes:
* resource_version and updated_at parameter has been added to the tax_withheld resource.

#### Added input parameters:
* ListRequestParams#updated_at parameter has been added to attached_item resource.
* AttachedItemsAttachedItemParams#updated_at parameter has been added to export resource.
* CheckoutGiftRequestParams#coupon_ids parameter has been added to hosted_page resource.
* ProcessReceiptCustomerParams#email, ProcessReceiptCustomerParams#first_name and  ProcessReceiptCustomerParams#last_name parameters has been added to in_app_subscription resource.

#### New endpoints:
* invoice#RecordTaxWithheld and invoice#RemoveTaxWithheld has been added to the invoice resource.
* credit_note#RemoveTaxWithheldRefund has been added to the credit_note resource.

#### New Enum values:
* custom has been added to payment_method Enum.
* ecentric has been added to gateway Enum.

#### Removed input parameters:
* EstimateRequestParams#EstimateInvoiceInfoParams parameter has been removed from purchase resource.

#### Deprecated input parameters:
* CheckoutGiftRequest#coupon parameter have been deprecated in hosted_page resource.


### v2.16.0 (2023-01-13)
* * *

#### New Attributes:
* ShippingAddress and BillingAddress parameter has been added to the credit_note.
* is_advance_charge has been added to the unbilled_charge.

#### Added input parameters:
* CreateSubForCustomerQuoteSubscriptionParams#po_number and CreateSubItemsForCustomerQuoteSubscriptionParams#po_number parameter has been added to Quote resource.
* ImportUnbilledChargesUnbilledChargeParams#is_advance_charge parameter has been added to Subscription resource.
* ImportInvoiceRequestParams#has_advance_charges parameter has been added to Invoice resource.

#### New endpoints:
* Invoice#DeleteLineItems has been added to the Invoice resource.

#### New Enum values:
* subscription_trial_extended has been added to event_type enum.
* zero_value_item has been added to tax_exempt_reason enum.
* bank_of_america has been added to gateway Enum.


### v2.15.0 (2022-11-21)
* * *

#### Fixes:
* Fixed jsonArray Serialisation issue.

#### New Attributes:
* business_entity_id filter parameter has been added to the Customer.
* is_written_off, write_off_amount, write_off_date and credit_note parameter have been added to the Invoice.

#### Added input parameters:
* business_entity_id filter parameter has been added to the  Export#RevenueRecognitionRequestParams, Export#DeferredRevenueRequestParams, Export#CustomersRequestParams and Subscription#ListRequestParams API.
* skip_billing_cycles parameter has been added to the Estimate#PauseSubscriptionSubscriptionParams API.
* skip_billing_cycles parameter has been added to the  Subscription#PauseRequestParams.
* id have been added to the Invoices#ImportInvoiceCreditNoteParams API.

#### New Enum values:
* billing_cycles has been added to pause_option.


### v2.14.1 (2022-11-16)
* * *

### Fixes:
* Fixed the broken import package.


### v2.14.0 (2022-11-10)
* * *

#### New endpoints:
* credit_notes#import_credit_note has been added to the credit_notes resource.
* subscriptions#import_unbilled_charges has been added to the subscriptions resource.

#### New attributes:
* shipped_at parameter has been added to the orders#list_orders API.

#### Added input parameters:
* voided_at and void_reason_code have been added to the invoices#import_invoice API.

#### New attributes:
* business_entity_id has been added to the orders resource.

#### New Enum values:
* registered has been added to the status enum in credit_note_einvoice and invoice_einvoice subresources.
* type enum has been added to the subscription_discount subresource.
* subscription_entitlements_created has been added to the event_type enum.


### v2.13.0 (2022-09-20)
* * *

### Fixes:
* Fixed race conditions while making concurrent requests.

#### New endpoints:
* subscriptions#list_discounts has been added to the subscriptions resource.

#### New Resource:
* discount has been added.

#### New attributes:
* billing_month has been added to the customer resource.

#### Added input parameters:
* billing_month has been added to the customers#change_billing_date API.
* line_items[subscription_id] has been added to the invoices#import_invoice API.
* layout has been added to hosted_pages#checkout_onetime_for_items, hosted_pages#checkout_new_for_items and hosted_pages#checkout_existing_for_items APIs.
* discounts[apply_on], discounts[percentage], discounts[amount] and discounts[item_price_id] have been added to estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#update_subscription_for_items, hosted_pages#checkout_onetime_for_items, hosted_pages#checkout_new_for_items, hosted_pages#checkout_existing_for_items, invoices#create_for_charge_items_and_charges, quotes#create_subscription_for_items, quotes#edit_create_subscription_quote_for_items, quotes#update_subscription_quote_for_items, quotes#edit_update_subscription_quote_for_items, quotes#create_for_charge_items_and_charges, quotes#edit_for_charge_items_and_charges, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items and subscriptions#import_subscription_for_items APIs.
* discounts[duration_type], discounts[period], discounts[period_unit] and discounts[included_in_mrr] have been added to estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#update_subscription_for_items, hosted_pages#checkout_new_for_items, hosted_pages#checkout_existing_for_items, quotes#create_subscription_for_items, quotes#edit_create_subscription_quote_for_items, quotes#update_subscription_quote_for_items, quotes#edit_update_subscription_quote_for_items, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items and subscriptions#import_subscription_for_items APIs.
* discounts[operation_type] and discounts[id] have been added to estimates#update_subscription_for_items, quotes#update_subscription_quote_for_items, hosted_pages#checkout_existing_for_items, quotes#edit_update_subscription_quote_for_items, subscriptions#update_subscription_for_items and subscriptions#import_subscription_for_items APIs.

#### New Enum values:
* global_payments has been added to gateway enum.
* layout, apply_on, duration_type and operation_type enumshave been added.

### v2.12.0 (2022-08-22)
* * *

#### New endpoints:
* Purchase#Retrieve has been added to the Purchase resource.

#### New attributes:
* ResourceVersion has been added to the Token resource.
* UpdatedAt has been added to the Token and UnbilledCharge resources.
* ReferenceLineItemId has been added to the CreditNote#LineItem, CreditNoteEstimate#LineItem, Quote#LineItem, QuoteLineGroup#LineItem, InvoiceEstimate#LineItem and Invoice#LineItem subResources.
* Index has been added to the Order#ShippingAddress, Invoice#ShippingAddress, Quote#ShippingAddress, QuotedCharge#ItemTier, QuotedSubscription#ItemTier, Subscription#ItemTier, Subscription#ShippingAddress and SubscriptionEstimate#ShippingAddress subResources.
* VoidWithCreditNote has been added to the Invoice#VoidInvoiceRequest subresources.
* PaymentMethodDetails has been added to the Transaction resource.

#### New Resource:
* InAppSubscription have been added.

#### Removed input parameters:
* BusinessEntityId has been removed from Purchase#CreateRequest and Purchase#EstimateRequest.


### v2.11.0 (2022-07-08)
* * *

### Fixes:
* Fixed Status Enum related issues
* Fixed EmbeddedResource related issues 

#### New endpoints:
* hostedpage#PreCancel has been added to the hostedpage resource.

#### New attributes:
* business_entity_id have been added to the CreditNote, Customer,HostedPage, Invoice, PaymentIntent, Quote, Subscription and Transaction resources.
* coupon_set_code have been added to the CreditNote, CreditNoteEstimate, Invoice, InvoiceEstimate, Quote and QuoteLineGroup resources.
* List of SubscriptionEstimate have been added to Estimate resource.

#### New Resource:
* ImpactedSubscription, ImpactedItem, Purchase have been added.

#### New Enum values:
* direct_debit has been added to payment_method_type enum.
* bancontact, not_applicable added to payment_source_card_brand enum.
* business_entity added to entity_type Enum
* business_entity_created, business_entity_updated, business_entity_deleted, purchase_created added to event_type Enum
* chargebee_payments added to gateway Enum


### v2.10.0 (2022-05-23)
* * *

#### New endpoints:
* invoice#sync_usages and invoice#resend_einvoice have been added to the invoice resource.
* credit_notes#resend_einvoice has been added to the credit_notes resource.
* features#list_features, features#create_a_feature, features#update_a_feature, features#retrieve_a_feature, features#delete_a_feature, features#activate_a_feature, features#archive_a_feature and features#reactivate_a_feature have  been added to the features resource.
* subscription_entitlements#subscription_entitlements_for_subscription and subscription_entitlements#set_subscription_entitlement_availability have been added to the subscription_entitlements resource.
* item_entitlements#item_entitlements_for_an_item, item_entitlements#item_entitlements_for_a_feature, item_entitlements#add_an_item_entitlements and item_entitlements#upsert_or_remove_an_item_entitlements_for_item have been added to the item_entitlements resource.
* entitlement_overrides#add_entitlement_override_for_a_subscription and entitlement_overrides#list_entitlement_override_for_a_subscription have been added to the entitlement_overrides resource.

#### New Resource:
* features, subscription_entitlements, item_entitlements and entitlement_overrides have been added.

#### New filter parameters:
* einvoice[status] filter parameter has been added in credit_notes#list_credit_notes api.

#### New Enum values:
* paypal_express_checkout has been added to payment_method_type enum.
* feature_created, feature_updated, feature_deleted, feature_activated, feature_reactivated, feature_archived, item_entitlements_updated, entitlement_overrides_updated, entitlement_overrides_removed, item_entitlements_removed and entitlement_overrides_auto_removed have been added to the event_type enum.
* action enum has been added.

### v2.9.0 (2022-04-25)
* * *

#### New endpoints:
* UnbilledCharge#CreateUnbilledCharge has been added to UnbilledCharge resource. Applicable only for PC1.0.

#### New attributes:
* channel have been added to the Addon, AttachedItem, CreditNote, Customer, Invoice, ItemFamily, ItemPrice, Plan and Subscription resources.
* external_name and channel have been added to the Item resource.

#### New input parameters:
* external_name have been added to Item#CreateRequest and Item#UpdateRequest.
* channel have been added to Addon#AddonListRequest.
* channel have been added to CreditNote#CreditNoteListRequest.
* channel have been added to Customer#CustomerListRequest.
* invoice[channel], subscription[channel] and customer[channel] have been added to Export#RevenueRecognitionRequest and Export#DeferredRevenueRequest.
* plan[channel] have been added to Export#PlansRequest.
* addon[channel] have been added to Export#AddonsRequest.
* customer[channel] have been added to Export#CustomersRequest.
* subscription[channel] have been added to Export#SubscriptionsRequest.
* invoice[channel] have been added to Export#InvoicesRequest. 
* credit_note[channel] have been added to Export#CreditNotesRequest.
* item[channel] have been added to Export#ItemsRequest.
* item_price[channel] have been added to Export#ItemPricesRequest.
* channel have been added to Invoice#InvoiceListRequest.
* channel have been added to Item#ItemListRequest.
* channel have been added to ItemPrice#ItemPriceListRequest.
* channel have been added to Plan#PlanListRequest.
* channel have been added to Subscription#SubscriptionListRequest.

### Existing input parameter changes:
* item_family_id in Item#CreateRequest has been made as required field.

### Removed enum values:
* coupon_expired has been removed from EventTypeEnum. 

### v2.8.0 (2022-03-14)
* * *

#### New Input parameters:
* net_term_days has been added to subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items, subscriptions#import_subscription_for_items endpoints.


### v2.7.0 (2022-03-08)
* * *

### Fixes:
* Escaping "id" field with PathEscape now onwards there will not be any panic when we have id with any special character in it

#### New attributes:
* upi, mandates and their subresources have been added to the payment_source resource.

#### New Input parameters:
* bank_account[phone] have been added to customers#create_bank_account
* payment_intent[payment_method_type] have been added to customers#create_payment_intent, customers#collect_payment_intent
* payment_intent[payment_method_type] have been added to estimate#gift_subscription and estimate#gift_subscription_for_items.
* payment_intent[payment_method_type] have been added to gift#create_payment_intent and gift#create_gift_for_items_payment_intent.
* customer[is_einvoice_enabled], customer[entity_identifier_scheme], customer[entity_identifier_standard], entity_identifiers[id], entity_identifiers[scheme], entity_identifiers[value], entity_identifiers[operation], entity_identifiers[standard] have been added to hosted_page#checkout_onetime_for_items, hosted_page#checkout_new_for_items and hosted_page#checkout_existing_for_items.
* bank_account[phone] have been added to invoice#create_bank_account, invoice#create_bank_account_for_chargeitems_and_charges
* payment_intent[payment_method_type] have been added to invoice#create_payment_intent, invoice#create_payment_intent_for_chargeitems_and_charges
* bank_account[phone] have been added to payment_source#create_bank_account
* payment_intent[payment_method_type] have been added to payment_source#create_using_payment_intent
* subscription_id have been added to payment_source#list_payment_source
* bank_account[phone] have been added to subscription#create_bank_account
* payment_intent[payment_method_type] have been added to subscription#create_payment_intent, subscription#create_payment_intent_for_customer, subscription#create_payment_intent_with_items, subscription#update_payment_intent, subscription#update_payment_intent_for_items, subscription#reactivate_payment_intent, subscription#resume_payment_intent
* activated_at have been added to subscription#import_subscription, subscription#import_subscription_for_customer, subscription#import_subscription_for_items

#### New Enum values:
* upi, netbanking_emandates enum has been added in customer_payment_method_type subresource of customer resource
* current enum has been added in account_type.
* quickbooks, razorpay enum has been added in gateway.
* upi, netbanking_emandates enum has been added in payment_method,payment_method_type, type
* upi, netbanking_emandates, apple_pay enum has been added in payment_method_type subresource of payment_intent resource
* current enum has been added in bank_account_type subresource of payment_source.

#### Deprecated attributes:
* coupon attribute have been added and deprecated in hosted_page resource.

#### Deprecated enums:
* NetdPaymentDueReminder have been deprecated in event_type.


### v2.6.0 (2022-01-21)
* * *

### Fixes:
* Fixed issue for not using existing httpClient when sending requests.

#### New endpoints:
* credit_notes#download_einvoice has been added to the credit_notes resource.
* invoice#download_einvoice has been added to the invoice resource.

#### New attributes:
* is_einvoice_enabled, entity_identifier_scheme, entity_identifier_standard and entity_identifiers[] have been added to the customer resource.
* einvoice has been added to the invoice resource.
* einvoice has been added to the credit_notes resource.
* mime_type has been added to the download resource.

#### New Input parameters:
* entity_identifier_scheme, entity_identifier_standard, is_einvoice_enabled, entity_identifiers[id][0..N], entity_identifiers[scheme][0..N], entity_identifiers[value][0..N], entity_identifiers[standard][0..N] have been added to customers#create_a_customer, customers#update_billing_info_for_a_customer apis.
* customer[entity_identifier_scheme], customer[entity_identifier_standard], customer[is_einvoice_enabled], entity_identifiers[id][0..N], entity_identifiers[scheme][0..N], entity_identifiers[value][0..N], entity_identifiers[standard][0..N] have been added to the subscriptions#create_a_subscription api.
* customer[entity_identifier_scheme], customer[entity_identifier_standard], customer[is_einvoice_enabled] have been added to subscriptions#update_a_subscription and subscriptions#update_subscription_for_items apis.

#### New Enum values:
* operation enum has been added.
* status enum has been added in credit_notes_einvoice subresource of credit_notes resource.

#### Deprecated attributes:
* user, type, payment_method and exchange_rate have been deprecated from TaxWithHeld resource.

#### Deprecated enums:
* type and payment_method have been deprecated in TaxWithHeld resource.

#### Updated parameters:
* hierarchy_operation_type has been made mandatory in customers#get_hierarchy api.

#### Removed Filter parameters:
* create_pending_invoices has been removed from subscriptions#list_subscriptions api.

### v2.5.1 (2022-01-05)
* * *

### Fixes:
* Improved error message for Invalid JSON response.
* Fixed issue while coupon retrieval having '%' in coupon ID.
* Remove .DS_store files

### v2.5.0 (2021-12-08)
* * *

### Fixes:
* Exposed function chargebee#UpdateTotalHTTPTimeout to set custom time out.
* Updated default http timeout of request to 80000ms. 

#### New endpoints:
* payment_sources#update_bank_account have been added in payment_sources resource.
* item_price#item_price_find_applicable_items and item_price#item_price_find_applicable_item_prices have been added in item_price resource.

#### New Attributes:
* hsn_code have been added to the resource addon, item_price and plan.
* first_name, last_name and email have been added to the resource payment_sources.

#### New Resource:
* TaxWithheld has been added.Applicable only for API V2. 

#### New Input parameters:
* hsn_code have been added to addons#create_an_addon, addons#update_an_addon, plan#create_an_plan and plan#update_an_plan  apis.
* bank_account[first_name],bank_account[last_name] and bank_account[email] have been added to payment_sources#update_bank_account api.
* charges[hsn_code] have been added in estimate#Create_Invoice, estimate#Create_Invoice_For_Items, hosted_pages#Checkout_One_Time, hosted_pages#Checkout_One_Time_For_Items, invoice#create_an_invoice, invoice#Create_For_Charge_Items_And_Charges and unbilledCharge#create_an_unbilledCharge apis.
* tax_detail[hsn_code] have been added in item_price#create_an_itemPrice and item_price#update_an_itemPrice apis.
* include_deleted have been added in plan#plan_list and addon#addon_list apis.Applicable only for Product Catalog V1.

#### New Enum values:
* subscription_activated_with_backdating, tax_withheld_recorded, tax_withheld_deleted and tax_withheld_refunded has been added to event_type enum.

### v2.4.1 (2021-11-17)
* * *

#### Fixes:
* Added support for custom fields in Item resource.

### v2.4.0 (2021-10-14)
* * *

#### New Attributes:
* entity_id has been added in order_line_item_discounts resource.

#### New Input parameters:
* line_items[tax5_name], line_items[tax5_amount], line_items[tax6_name], line_items[tax6_amount], line_items[tax7_name], line_items[tax7_amount], line_items[tax8_name], line_items[tax8_amount], line_items[tax9_name],  line_items[tax9_amount], line_items[tax10_name], line_items[tax10_amount] have been added in import_invoice api.
* replace_primary_payment_source has been added in create_subscription_for_customer and create_subscription_for_items apis.

#### New Enum values:
* coupon_expired has been added to event_type enum.
* mollie has been added to gateway enum.
* item_level_discount and document_level_discount has been added to discount_type.

### v2.3.0 (2021-08-16)
* * *
* added support for custom http client.

#### New Attributes:
* generated_at has been added in credit_note and invoice resources.
* change_option have been added in quoted_subscription resource.
* changes_scheduled_at has been added in subscription and quoted_subscription resources.
* iin and last4 have been added in transaction resource.

#### New Resource:
* quoted_charge has been added. 

#### New Input parameters:
* invoice_date has been added in estimates##create_subscription_estimate, estimates#estimate_for_creating_a_customer_and_subscription, estimates#estimate_for_creating_a_subscription, estimates#create_subscription_for_a_customer_estimate, estimates#update_subscription_estimate, estimates#estimate_for_updating_a_subscription, estimates#cancel_subscription_estimate, estimates#cancel_subscription_for_items_estimate, estimates#create_invoice_for_items_estimate, estimates#create_invoice_estimate, hosted_pages#checkout_existing_subscription, hosted_pages#create_checkout_to_update_a_subscription, invoices#create_an_invoice, invoices#create_invoice_for_items_and_one-time_charges, subscriptions#create_a_subscription, subscriptions#create_subscription_for_customer, subscriptions#update_a_subscription, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items, subscriptions#reactivate_a_subscription, subscriptions#cancel_subscription_for_items and subscriptions#cancel_a_subscription endpoints.
* coupon_ids has been added to invoices#create_invoice_for_a_one-time_charge, invoice#create_invoice_for_a_non-recurring_addon, quotes#create_quote_for_one-time_charges, quotes#edit_quote_for_one-time_charges, quotes#create_a_quote_for_charge_and_charge_items, quotes#edit_quote_for_charge_items_and_charges endpoints.
* change_option and changes_scheduled_at have been added in quotes#create_quote_for_updating_a_subscription, quotes#edit_quote_for_updating_a_subscription, quotes#create_a_quote_for_update_subscription_items, quotes#edit_update_subscription_quote_for_items, subscriptions#update_a_subscription and subscriptions#create_subscription_for_items endpoints.
* invoice_date, create_pending_invoices and first_invoice_pending have been added in quotes#convert_a_quote endpoint.
* subscription[auto_close_invoices] has been added in quotes#convert_a_quote endpoint.

#### New Enum values:
* subscription_cancel has been added to charge_event enum in quote_line_groups resource.
* subscription_created_with_backdating, subscription_changed_with_backdating, subscription_canceled_with_backdating, subscription_reactivated_with_backdating, invoice_generated_with_backdating and credit_note_created_with_backdating have been added to event_type enum 
* change_option enum has been added.

#### Deprecated parameters:
* coupon has been deprecated in invoices#create_invoice_for_a_one-time_charge and invoices#create_invoice_for_a_non-recurring_addon endpoints.


### v2.2.0 (2021-07-22)
* * *
#### New endpoints:
* hosted_pages#checkout_one_time_for_items and hosted_pages#checkout_gift_for_items have been added in hosted_pages resource. 
* orders#resend_an_order has been added in orders resource.
* quotes#edit_create_subscription_quote_for_items, quotes#edit_update_subscription_quote_for_items and quotes#edit_quote_for_charge_items_and_charges have been added in quotes resource. Applicable only for Product Catalog V2.

#### New attributes:
* accounting_category3 and accounting_category4 have been added to addon, item_price and plan resources. 
* quantity_in_decimal has been added to attached_item resource.
* period and period_unit have been added to coupon resource.
* entity_id has been added to line_item_discounts object of credit_note, credit_note_estimate, invoice, invoice_estimate, quote and quote_line_group resources. 
* vat_number_prefix has been added to credit_note, customer, invoice and quote resources.
* price_in_decimal has been added to differential_price and item_price resources. 
* exchange_rate and new_sales_amount have been added to invoice resource.
* archived_at has been added to item and item_price resources.
* trial_end_action has been added to item_price, plan, subscription and subscription_estimate resources.
* resent_orders, tracking_url, resent_status, is_resent, original_order_id and resend_reasons have been added to order resource. 
* line_item_tiers have been added to quote resource.
* plan_quantity_in_decimal and plan_unit_price_in_decimal have been added to quoted_subscription and subscription resources. Applicable only for Product Catalog V1.
* contract_term_billing_cycle_on_renewal and quoted_contract_term have been added to quoted_subscription resource.
* quantity_in_decimal, metered_quantity, last_calculated_at, unit_price_in_decimal, amount_in_decimal and free_quantity_in_decimal have been added to subscription_items object of quoted_subscription and subscription resources.
* starting_unit_in_decimal, ending_unit_in_decimal and price_in_decimal have been added to item_tiers object in quoted_subscription resources.
* cancel_schedule_created_at has been added to subscription resource.
* exchange_rate and merchant_reference_id have been added to transaction resource.

#### New parameters:
* accounting_category3 and accounting_category4 have been added in addons#create_an_addon, addons#update_an_addon, plans#create_a_plan and plans#update_a_plan endpoints.
* quantity_in_decimal has been added in attached_items#create_an_attached_item, attached_items#update_an_attached_item and gifts#create_a_gift_subscription_for_items endpoints.
* period and period_unit have been added in coupons#create_a_coupon, coupons#update_a_coupon, coupons#create_a_coupon_for_items and coupons#update_a_coupon_for_items endpoints.
* vat_number_prefix has been added in customers#create_a_customer, customers#update_billing_info_for_a_customer, estimates#create_subscription, estimates#create_subscription_estimate, estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#update_subscription_for_items, estimates#update_subscription, hosted_pages#checkout_new_subscription, hosted_pages#checkout_one-time_payments, hosted_pages#create_checkout_for_a_new_subscription hosted_pages#checkout_existing_subscription, hosted_pages#create_checkout_to_update_a_subscription, invoices#import_invoice, invoices#update_invoice_details, subscriptions#create_a_subscription, subscriptions#update_a_subscription, subscriptions#update_subscription_for_items and subscriptions#import_a_subscription endpoints.
* payment_method[additional_information] has been added in customers#create_a_customer, customers#update_payment_method_for_a_customer,  customers#collect_payment_for_customer, invoices#create_an_invoice, invoices#create_invoice_for_items_and_one-time_charges, subscriptions#create_a_subscription, subscriptions#update_a_subscription, subscriptions#update_subscription_for_items and subscriptions#import_a_subscription endpoints.
* card[additional_information] has been added in customers#create_a_customer, customers#collect_payment_for_customer, invoices#create_an_invoice, invoices#create_invoice_for_items_and_one-time_charges, payment_sources#create_a_card_payment_source, subscriptions#create_a_subscription, subscriptions#update_a_subscription, subscriptions#update_subscription_for_items and subscriptions#import_a_subscription endpoints.
* bank_account[billing_address] has been added in customers#create_a_customer, invoices#create_an_invoice, invoices#create_invoice_for_items_and_one-time_charges, payment_sources#create_a_bank_account_payment_source and subscriptions#create_a_subscription endpoints.
* price_in_decimal has been added in differential_prices#create_a_differential_price and differential_prices#update_a_differential_price endpoints.
* tiers[starting_unit_in_decimal][0..N], tiers[ending_unit_in_decimal][0..N] and tiers[price_in_decimal][0..N] have been added in differential_prices#create_a_differential_price, differential_prices#update_a_differential_price, item_prices#create_an_item_price and item_prices#update_an_item_price endpoints.
* remove_general_note and notes_to_remove been added in estimates#create_invoice_estimate, estimates#create_invoice_for_items_estimate and invoices#create_invoice_for_items_and_one-time_charges endpoints.
* trial_end_action has been added in estimates#create_subscription, estimates#create_subscription_estimate, estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#update_subscription_for_items, estimates#update_subscription, plans#create_a_plan, plans#update_a_plan, item_prices#create_an_item_price, item_prices#update_an_item_price, subscriptions#create_a_subscription, subscriptions#create_subscription_for_customer, subscriptions#create_subscription_for_items, subscriptions#update_a_subscription and subscriptions#update_subscription_for_items endpoints.
* subscription_items[unit_price_in_decimal][0..N], item_tiers[starting_unit][0..N], item_tiers[ending_unit][0..N] and item_tiers[price_in_decimal][0..N] have been added in estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#update_subscription_for_items, estimates#update_subscription, estimates#cancel_subscription_for_items_estimate, hosted_pages#create_checkout_for_a_new_subscription, hosted_pages#create_checkout_to_update_a_subscription, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items and subscriptions#import_subscription_for_items endpoints.
* subscription_items[quantity_in_decimal][0..N] has been added in estimates#create_subscription_for_items, estimates#create_subscription_for_items_estimate, estimates#cancel_subscription_for_items_estimate, estimates#update_subscription_for_items, estimates#update_subscription, estimates#gift_subscription_for_items, estimates#cancel_subscription_for_items_estimate, gifts#create_a_gift_subscription_for_items, hosted_pages#create_checkout_for_a_new_subscription, hosted_pages#create_checkout_to_update_a_subscription, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items, subscriptions#import_subscription_for_items and subscriptions#cancel_subscription_for_items endpoints.
* notes_to_remove[entity_type][0..N] and notes_to_remove[entity_id][0..N] have been added in estimates#create_invoice, estimates#create_invoice_for_items and invoices#create_invoice_for_items_and_one-time_charges endpoints.
* item_prices[quantity_in_decimal][0..N] and item_prices[unit_price_in_decimal][0..N] have been added in estimates#create_invoice_for_items, invoices#create_invoice_for_items_and_one-time_charges and unbilled_charges#create_unbilled_charges_for_a_subscription endpoints.
* item_tiers[starting_unit_in_decimal][0..N], item_tiers[ending_unit_in_decimal][0..N] and item_tiers[price_in_decimal][0..N] have been added in estimates#create_invoice_for_items, invoices#create_invoice_for_items_and_one-time_charges and unbilled_charges#create_unbilled_charges_for_a_subscription endpoints.
* payment_intent[additional_information] has been added in customers#create_a_customer, customers#collect_payment_for_customer, gifts#create_a_gift_subscription_for_items, gifts#create_a_gif and payment_sources#create_using_payment_intent endpoints.
* invoice_note has been added in hosted_pages#checkout_one-time_payments endpoint.
* allow_offline_payment_methods has been added in hosted_pages#create_checkout_for_a_new_subscription and hosted_pages#create_checkout_to_update_a_subscription endpoints.
* free_quantity_in_decimal, price_in_decimal and trial_end_action have been added in item_prices#create_an_item_price and item_prices#update_an_item_price endpoints.
* tracking_url has been added in orders#create_an_order, orders#update_an_order and orders#import_an_order endpoints.
* additional_information has been added in payment_sources#create_using_gateway_temporary_token and payment_sources#create_using_permanent_token endpoints.

#### New enum values:
* bancontact has been added to card_type enum.
* card and latam_local_card have been added to powered_by enum.
* item_level_discount and document_level_discount have been added to entity_type enum of credit_note_discount, credit_note_estimate_discount, invoice_discount, line_item_discount and invoice_estimate_discount resources.
* item_level_discount and document_level_discount have been added to discount_type enum.
* add_usages_reminder and order_resent have been added to event_type enum.
* ingenico_direct and exact have been added to gateway enum. 
* max_usage has been added to usage_calculation enum.
* trial_end_action enum has been added.
* order_resent has been added to cancellation_reason enum.
* resent_status enum has been added.
* custom_discount has been added in discount_type enum of order_line_item_discount. 

#### New filter parameters:
* resent_status, is_resent and original_order_id filter parameters have been added in exports#orders and orders#list_orders enpoints.
* updated_at filter parameter has been added in exports#item_families and item_families#list_item_families endpoints.

#### New sort parameters:
* name, id and updated_at sort parameters are added to items#list_items and item_prices#list_item_prices endpoints.
* created_at and updated_at sort parameters are added to payment_sources#list_payment_sources endpoint.
* usage_date sort parameter has been added to usages#list_usages endpoint.

#### Deprecated parameters:
* duration_month parameter has been deprecated in coupon resource.

### v2.1.9 (2021-01-19)
* * *
##### New resources:
* Usages is added. Applicable only for Product Catalog V2

##### New end points:
* hosted_pages#checkout_one-time_payments has been added in hosted_pages resource

##### New attributes:
* auto_close_invoices has been added to customers, subscriptions resources
* metered, usage_calculation have been added to items resources
* create_pending_invoices, auto_close_invoices have been added to subscriptions resources

##### New parameters:
* auto_close_invoices has been added to the endpoint: customers#create_a_customer, customers#list_customers, customers#update_a_customer
* invoice_allocations[invoice_id] has been added to the endpoint: customers#collect_payment_for_customer
* coupon_ids has been added to the endpoint: estimates#create_invoice_for_items_estimate. Applicable only for Product Catalog V2
* subscription[auto_close_invoices] has been added to the endpoint: exports#export_revenue_recognition_reports, exports#export_deferred_revenue_reports, exports#export_subscriptions
* customer[auto_close_invoices] has been added to the endpoint: exports#export_revenue_recognition_reports, exports#export_deferred_revenue_reports, exports#export_customers
* item[metered], item[usage_calculation] have been added to the endpoint: exports#export_items. Applicable only for Product Catalog V2
* subscription_id has been added to the endpoint: invoices#add_one-time_charge_to_a_pending_invoice, invoices#add_non-recurring_addon_to_a_pending_invoice
* subscription_id has been added to the endpoint: invoices#add_a_charge-item_to_a_pending_invoice. Applicable only for Product Catalog V2
* metered, usage_calculation have been added to the endpoint: items#create_an_item, items#list_item. Applicable only for Product Catalog V2
* create_pending_invoices, auto_close_invoices, first_invoice_pending have been added to the endpoint: subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items,subscriptions#import_subscription_for_items. Applicable only for Product Catalog V2
* create_pending_invoices, auto_close_invoices have been added to the endpoint: subscriptions#list_subscriptions

### v2.1.8 (2020-12-15)
* * *
##### New end points:
* estimate_for_creating_a_customer_and_subscription, cancel_subscription_for_items_estimate, gift_subscription_estimate_for_items have been added in estimate resource. Applicable only for Product Catalog V2
* regenerate_invoice_estimate has been added in estimate resource
* create_a_gift_subscription_for_items has been added in gift resource. Applicable only for Product Catalog V2
* regenerate_invoice has been added in subscription resource

##### New attributes:
* show_description_in_invoices, show_description_in_quotes have been added to the resource item_prices
* tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[price_in_decimal] have been added to the resource differential_prices
* show_description_in_invoices, show_description_in_quotes have been added to the resource item_prices. Applicable only for Product Catalog V2

##### New parameters:
* payment_intent[additional_info] has been added to the endpoints customers#create_a_customer, invoices#create_an_invoice, payment_sources#create_using_payment_intent, subscriptions#create_a_subscription, subscriptions#update_a_subscription, subscriptions#create_subscription_for_customer, subscriptions#reactivate_a_subscription, subscriptions#resume_a_subscription
* payment_intent[additional_info] has been added to the endpoints gifts#create_a_gift_subscription_for_items, invoices#create_invoice_for_items_and_one-time_charges, subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items. Applicable only for Product Catalog V2
* contract_term[action_at_term_end], contract_term[cancellation_cutoff_period], subscription[contract_term_billing_cycle_on_renewal] has been added to the endpoint estimates#estimate_for_creating_a_subscription. Applicable only for Product Catalog V2
* cancel_at, contract_term_cancel_option, cancel_reason_code have been added to the endpoint estimates#cancel_subscription_for_items_estimate
* event_based_addons has been added to the endpoint estimates#cancel_subscription_for_items_estimate
* redirect_url has been added to the endpoint hosted_pages#accept_a_quote
* token_id, retain_payment_source, card, bank_account, payment_method added have been added to the endpoint invoices#create_an_invoice
* token_id, retain_payment_source, card, bank_account, payment_method added have been added to the endpoint invoices#create_invoice_for_items_and_one-time_charges. Applicable only for Product Catalog V2
* show_description_in_invoices, show_description_in_quotes have been added to the endpoints item_prices#create_an_item_price, item_prices#update_an_item_price. Applicable only for Product Catalog V2

### v2.1.7 (2020-11-26)
* * *
##### New resources:
item_family, item, item_price, attached_item and differential_price are added. Applicable only for Product Catalog V2

##### New end points:
* coupons#create_a_coupon_for_items and coupons#update_a_coupon_for_items have been added in coupon resource. Applicable only for Product Catalog V2
* estimates#estimate_for_creating_a_subscription, estimates#estimate_for_updating_a_subscription and estimates#create_invoice_for_items_estimate have been added in estimate resource. Applicable only for Product Catalog V2
* estimates#advance_invoice_estimate has been added in estimate resource
* exports#export_item_families, exports#export_items, exports#export_item_prices, exports#export_attached_items and exports#export_differential_price have been added in export api. Applicable only for Product Catalog V2
* checkout_new_for_items and checkout_existing_for_items have been added in hosted_pages api. Applicable only for Product Catalog V2
* invoices#create_invoice_for_items_and_one-time_charges, invoices#create_invoice_for_a_charge-item and invoices#add_a_charge-item_to_a_pending_invoice have been added in invoice resource. Applicable only for Product Catalog V2
* quotes#create_a_quote_for_a_new_subscription_items, quotes#create_a_quote_for_update_subscription_items and quotes#create_a_quote_for_charge_and_charge_items have been added in quote resource. Applicable only for Product Catalog V2
* subscriptions#create_subscription_for_items, subscriptions#update_subscription_for_items, subscriptions#import_subscription_for_items and subscriptions#cancel_subscription_for_items have been added in subscription resource. Applicable only for Product Catalog V2
* subscriptions#edit_advance_invoice_schedule, subscriptions#retrieve_advance_invoice and subscriptions#remove_an_advance_invoice_schedules have been added in subscription resource
* unbilled_charges#create_an_invoice_for_unbilled_charges has been added to unbilled_charge resource

##### New attributes:
* item_constraints and item_constraint_criteria have been added in coupon resource. Applicable only for Product Catalog V2
* success_url and failure_url have been added in payment_intent resource
* subscription_items and item_tiers have been added in quoted_subscription resource. Applicable only for Product Catalog V2
* has_scheduled_advance_invoices has been added in subscription resource
* subscription_items, item_tiers and charged_items have been added in subscription resource. Applicable only for Product Catalog V2

##### New parameters:
* item_id and item_price_id have been added to the end point: subscriptions#list_subscriptions, exports#export_revenue_recognition_reports, exports#export_deferred_revenue_reports, exports#export_subscriptions. Applicable only for Product Catalog V2
* invoice_immediately, schedule_type and fixed_interval_schedule have been added to the end point: subscriptions#charge_future_renewals
*  success_url and failure_url have been added to the end points: payment_intents#create_a_payment_intent, payment_intents#update_a_payment_intent

##### New Enum values:
* PLAN_ITEM_PRICE, ADDON_ITEM_PRICE, CHARGE_ITEM_PRICE are added to Entitytype Enum
* ITEM_FAMILY_CREATED, ITEM_FAMILY_UPDATED, ITEM_FAMILY_DELETED, ITEM_CREATED, ITEM_UPDATED, ITEM_DELETED, ITEM_PRICE_CREATED, ITEM_PRICE_UPDATED, ITEM_PRICE_DELETED, ATTACHED_ITEM_CREATED, ATTACHED_ITEM_UPDATED, ATTACHED_ITEM_DELETED, DIFFERENTIAL_PRICE_CREATED, DIFFERENTIAL_PRICE_UPDATED, DIFFERENTIAL_PRICE_DELETED are added to EventType Enum

### v2.1.6 (2020-11-16)
* * *
* New attributes price_in_decimal, tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[price_in_decimal] have been added to the resource addon
* New input parameters price_in_decimal, tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[price_in_decimal] have been added to addons#create_an_addon, addons#update_an_addon apis.
* New attributes unit_amount_in_decimal, quantity_in_decimal, amount_in_decimal, line_item_tiers[starting_unit_in_decimal], line_item_tiers[ending_unit_in_decimal], line_item_tiers[quantity_used_in_decimal], line_item_tiers[unit_amount_in_decimal] have been added to the resources credit_note, credit_note_estimate, invoice, invoice_estimate
* New input parameters line_items[unit_amount_in_decimal], line_items[quantity_in_decimal] have been added to credit_notes#create_credit_note api.
* New input parameters subscription[plan_unit_price_in_decimal], subscription[plan_quantity_in_decimal], addons[quantity_in_decimal], addons[unit_price_in_decimal], event_based_addons[quantity_in_decimal], event_based_addons[unit_price_in_decimal] have been added to estimates#create_subscription_estimate, estimates#create_subscription_for_a_customer_estimate, estimates#update_subscription_estimate, hosted_pages#checkout_new_subscription, hosted_pages#checkout_existing_subscription, quotes#create_a_quote_for_a_new_subscription,  quotes#edit_create_subscription_quote, quotes#create_a_quote_for_update_subscription, quotes#edit_update_subscription_quote,  apis
* New input parameters subscription[plan_quantity_in_decimal], addons[quantity_in_decimal] have been added to estimates#gift_subscription_estimate, gifts#create_a_gift, hosted_pages#checkout_gift_subscription apis
* New input parameters addons[quantity_in_decimal], addons[unit_price_in_decimal], charges[amount_in_decimal] have been added to estimates#create_invoice_estimate, invoices#create_an_invoice, quotes#create_a_quote_for_one-time_charges, quotes#edit_one-time_quote apis
* New input parameter amount_in_decimal has been added to invoices#create_invoice_for_a_one-time_charge api
* Input parameter amount has been made optional to invoices#create_invoice_for_a_one-time_charge api
* New input parameters addon_quantity_in_decimal, addon_unit_price_in_decimal have been added to invoices#create_invoice_for_a_non-recurring_addon
* New input parameters line_items[unit_amount_in_decimal], line_items[quantity_in_decimal], line_items[amount_in_decimal], line_item_tiers[starting_unit_in_decimal], line_item_tiers[ending_unit_in_decimal], line_item_tiers[quantity_used_in_decimal], line_item_tiers[unit_amount_in_decimal] have been added to invoices#import_invoice api
* Input parameters line_item_tiers[starting_unit], line_item_tiers[ending_unit], line_item_tiers[quantity_used], line_item_tiers[unit_amount] have beed made optional in invoices#import_invoice api
* New input parameters addon_quantity_in_decimal, addon_unit_price_in_decimal have been added to invoices#add_non-recurring_addon_to_a_pending_invoice api
* New input parameter invoice_date has been added to invoices#close_a_pending_invoice api
* New attributes tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[price_in_decimal], attached_addons[quantity_in_decimal], event_based_addons[quantity_in_decimal], free_quantity_in_decimal, price_in_decimal have been added to the resource plan
* New input parameters tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[price_in_decimal], attached_addons[quantity_in_decimal], event_based_addons[quantity_in_decimal], free_quantity_in_decimal, price_in_decimal have been added to plans#create_a_plan, plans#update_a_plan apis
* New attribute amount_in_decimal has been added to the resource promotional_credit
* New input parameter amount_in_decimal has been added to promotional_credits#add_promotional_credits, promotional_credits#deduct_promotional_credits, promotional_credits#set_promotional_credits apis
* Input parameter amount has been made optional in promotional_credits#add_promotional_credits, promotional_credits#deduct_promotional_credits, promotional_credits#set_promotional_credits apis
* New attributes unit_amount_in_decimal, quantity_in_decimal, amount_in_decimal have been added to the resource quote
* New attributes addons[quantity_in_decimal], addons[unit_price_in_decimal], addons[amount_in_decimal], event_based_addons[quantity_in_decimal], event_based_addons[unit_price_in_decimal] have been added to the sub resource quoted_subscription
* New attributes unit_amount_in_decimal, quantity_in_decimal, amount_in_decimal have been added to the sub resource quoted_line_group
* New attributes addons[quantity_in_decimal], addons[unit_price_in_decimal], addons[amount_in_decimal], event_based_addons[quantity_in_decimal], event_based_addons[unit_price_in_decimal], plan_free_quantity_in_decimal, plan_quantity_in_decimal, plan_unit_price_in_decimal, plan_amount_in_decimal have been added to the resource subscription
* New input parameters plan_unit_price_in_decimal, plan_quantity_in_decimal, addons[quantity_in_decimal], addons[unit_price_in_decimal], event_based_addons[quantity_in_decimal], event_based_addons[unit_price_in_decimal] have been added to subscriptions#create_a_subscription, subscriptions#create_subscription_for_customer, subscriptions#update_a_subscription, subscriptions#import_a_subscription, subscriptions#import_subscription_for_customer apis
* New input parameter amount_in_decimal has been added to subscriptions#add_charge_at_term_end apis
* Input parameter amount has been made optional in subscriptions#add_charge_at_term_end apis
* New input parameters addon_quantity_in_decimal, addon_unit_price_in_decimal have been added to subscriptions#charge_addon_at_term_end api
* New attributes tiers[starting_unit_in_decimal], tiers[ending_unit_in_decimal], tiers[quantity_used_in_decimal], tiers[unit_amount_in_decimal], unit_amount_in_decimal, quantity_in_decimal, amount_in_decimal have been added to the resource unbilled_charge
* New endpoint transactions#refund_a_payment has been added to the resource transaction
### v2.1.5 (2020-10-15)
* * *
* New optional attribute quoted_subscriptions has been added to the resource quote
* New optional attributes resource_version and updated_at are added to the resource payment_intent
### v2.1.4 (2020-09-29)
* * *

* New attribute included_in_mrr has been added in addon and coupon resource
* New attribute offline_payment_method has been added in subscription and customer resource
* New input parameter included_in_mrr has been added in create_an_addon, update_an_addon, create_a_coupon and update_a_coupon apis.
* New input parameter offline_payment_method has been added in create_a_customer, list_customers, update_a_customer, create_a_subscription, create_subscription_for_customer and list_subscriptions apis
* New input parameter auto_collection has been added in update_a_subscription
* New input parameter subscription[offline_payment_method] has been added in create_subscription_estimate, create_subscription_for_a_customer_estimate, update_subscription_estimate, export_revenue_recognition_reports, export_deferred_revenue_reports, export_subscriptions, checkout_new_subscription and checkout_existing_subscription apis
* New input parameter subscription[auto_collection] has been added in checkout_existing_subscription and update_subscription_estimate apis
* New input parameter customer[offline_payment_method] has been added in export_revenue_recognition_reports, export_deferred_revenue_reports, export_customers and create_a_subscription

### v2.1.3 (2020-09-09)
* * *

* New input parameter currency_code is added in list_addons, list_coupons, list_plans, export_plans, export_addons, export_coupons apis
* New attributes powered_by has been added in card resource
* New input parameters subscription[free_period], subscription[free_period_unit] have been added in create_subscription_estimate, create_subscription_for_a_customer_estimate, update_subscription_estimate apis
* Input parameter invoice[customer_id] is made optional in create_invoice_estimate api
* Input parameter customer_id is made optional in create_an_invoice
* Attribute created_at in dunning_attempts is made optional in the invoice resource
* New attributes free_period and free_period_unit have been added in the Subscription resource
* New enum type free_period_unit has been added with the values:
	DAY,
	WEEK,
	MONTH,
	YEAR
* New enum type powered_by has been added in card resource with the values: 
	IDEAL,
    SOFORT,
    BANCONTACT,
    GIROPAY,
    NOT_APPLICABLE
* New endpoint import_contract_term has been added to the subscription resource
* Attributes status, amount and date inside the linked_payment are made optional in the transaction resource
* New endpoint delete has been added to the virtual_bank_account resource
* New values GIROPAY and DOTPAY have been added to the payment_method_types enum
* New values PAYMENT_SOURCE_EXPIRING and PAYMENT_SOURCE_EXPIRED have been added to the event_type enums
* New value PAYPAL has been added to the gateway enum

### v2.1.2 (2020-07-15)
* * *

* New attributes show_description_in_invoices, show_description_in_quotes have been added in Addon, Plan resource
* New input parameters show_description_in_invoices, show_description_in_quotes have been added for create_an_addon, update_an_addon, create_a_plan, update_a_plan apis
* New attribute create_reason_code has been added in Credit notes resource
* Attribute reason_code is made optional in the Credit notes resource
* New filter parameter create_reason_code is added in list_credit_notes api
* New input parameter refund_reason_code has been added in Credit notes' Refund, Record a refund apis
* Sub-resources parent_account_access and child_account_access have been added in Customer resource
* New attribute use_default_hierarchy_settings has been added in Customer resource
* New endpoint update_hierarchy_settings has been added in Customer resource
* New input parameters use_default_hierarchy_settings,  parent_account_access, child_account_access have been added in link_a_customer api
* New input parameter invoice_notes, coupon_ids, invoice[subscription_id], charges[taxable], charges[tax_profile_id], charges[avalara_tax_code], charges[taxjar_product_code] has been added to create_invoice_estimate api
* New input param cancel_reason_code has been added to export_revenue_recognition_reports, export_deferred_revenue_reports, export_subscriptions, export_credit_notes apis
* New input param coupon_ids has been added to checkout_new_subscription, checkout_existing_subscription apis
* Input param subscription_coupon has been deprecated in checkout_new_subscription, checkout_existing_subscription apis
* New attribute void_reason_code has been added to the Invoice resource
* Attribute cn_reason_code has been made optional in applied_credits, adjustment_credit_notes, issued_credit_notes, linked_credit_note sub-resources
* New attribute cn_create_reason_code has been added in applied_credits, adjustment_credit_notes, issued_credit_notes, linked_credit_note sub-resources
* New input parameter subscription_id, invoice_note, remove_general_note, coupon_ids, charges[taxable], charges[tax_profile_id], charges[avalara_tax_code], charges[taxjar_product_code] have been added to create_an_invoice api
* Input parameter coupon is deprecated from create_an_invoice api
* New filter parameter void_reason_code has been added to list_invoices api
* New input parameter invoice_note, remove_general_note, notes_to_remove[entity_type], notes_to_remove[entity_id] have been added to close_a_pending_invoice api
* New input parameter void_reason_code has been added to void_an_invoice api
* New input parameter credit_note[create_reason_code] has been added to refund_an_invoice, record_refund_for_an_invoice api
* New endpoints import_an_order and delete_an_imported_order have been added to Order resource
* New filter parameter exclude_deleted_credit_notes has been added to list_orders api
* New attribute payment_method_type has been added to payment_intent resource and active_payment_attempt sub-resource
* New input parameter payment_method_type has been added to create_a_payment_intent, update_a_payment_intent api
* New attributes version, contract_term_start, contract_term_end, contract_term_termination_fee have been added to Quotes resource
* New endpoints edit_create_subscription_quote, edit_update_subscription_quote, edit_one_time_quote, extend_expiry_date have beed added to Quotes resource
* New input parameters contract_term[action_at_term_end], contract_term[cancellation_cutoff_period], subscription[contract_term_billing_cycle_on_renewal] have been added to create_a_quote_for_a_new_subscription, create_a_quote_for_update_subscription apis
* New attribute version has been added to quote_line_group sub-resource
* New attribute cancel_reason_code has been added to Subscriptions resource
* New filter parameter cancel_reason_code has been added to list_subscriptions, cancel_a_subscription apis
* New input parameter contract_term_billing_cycle_on_renewal has been added to import_a_subscription, import_subscription_for_customer apis
* New input parameter event_based_addon has been added to cancel_a_subscription api
* New filter parameter is_voided has been added to list_unbilled_charges api
* New event_types MRR_UPDATED, ORDER_DELETED have been added
* New values SOFORT, BANCONTACT have been added to the payment_method_types enum
* New attributes entity_type_description has been added to the line_items sub-resource

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
