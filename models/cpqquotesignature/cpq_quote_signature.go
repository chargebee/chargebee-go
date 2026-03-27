package cpqquotesignature

import (
	cpqQuoteSignatureEnum "github.com/chargebee/chargebee-go/v3/models/cpqquotesignature/enum"
)

type CpqQuoteSignature struct {
	Id                       string                                         `json:"id"`
	Status                   cpqQuoteSignatureEnum.Status                   `json:"status"`
	Name                     string                                         `json:"name"`
	DocumentName             string                                         `json:"document_name"`
	CustomerAcceptanceMethod cpqQuoteSignatureEnum.CustomerAcceptanceMethod `json:"customer_acceptance_method"`
	QuoteType                cpqQuoteSignatureEnum.QuoteType                `json:"quote_type"`
	ExpiresAt                int64                                          `json:"expires_at"`
	Timezone                 string                                         `json:"timezone"`
	ProviderRequestId        string                                         `json:"provider_request_id"`
	ProviderDocumentId       string                                         `json:"provider_document_id"`
	CreatedAt                int64                                          `json:"created_at"`
	ModifiedAt               int64                                          `json:"modified_at"`
	Object                   string                                         `json:"object"`
}
