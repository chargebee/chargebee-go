package chargebee

type CommentEntityType string

const (
	CommentEntityTypeCustomer       CommentEntityType = "customer"
	CommentEntityTypeSubscription   CommentEntityType = "subscription"
	CommentEntityTypeInvoice        CommentEntityType = "invoice"
	CommentEntityTypeQuote          CommentEntityType = "quote"
	CommentEntityTypeCreditNote     CommentEntityType = "credit_note"
	CommentEntityTypeTransaction    CommentEntityType = "transaction"
	CommentEntityTypePlan           CommentEntityType = "plan"
	CommentEntityTypeAddon          CommentEntityType = "addon"
	CommentEntityTypeCoupon         CommentEntityType = "coupon"
	CommentEntityTypeOrder          CommentEntityType = "order"
	CommentEntityTypeBusinessEntity CommentEntityType = "business_entity"
	CommentEntityTypeItemFamily     CommentEntityType = "item_family"
	CommentEntityTypeItem           CommentEntityType = "item"
	CommentEntityTypeItemPrice      CommentEntityType = "item_price"
	CommentEntityTypePriceVariant   CommentEntityType = "price_variant"
)

type CommentType string

const (
	CommentTypeUser   CommentType = "user"
	CommentTypeSystem CommentType = "system"
)

// just struct
type Comment struct {
	Id               string            `json:"id"`
	EntityType       CommentEntityType `json:"entity_type"`
	AddedBy          string            `json:"added_by"`
	Notes            string            `json:"notes"`
	CreatedAt        int64             `json:"created_at"`
	Type             CommentType       `json:"type"`
	EntityId         string            `json:"entity_id"`
	BusinessEntityId string            `json:"business_entity_id"`
	Object           string            `json:"object"`
}

// sub resources
// operations
// input params
type CommentCreateRequest struct {
	EntityType CommentEntityType `json:"entity_type"`
	EntityId   string            `json:"entity_id"`
	Notes      string            `json:"notes"`
	AddedBy    string            `json:"added_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CommentCreateRequest) payload() any { return r }

type CommentListRequest struct {
	Limit      *int32            `json:"limit,omitempty"`
	Offset     string            `json:"offset,omitempty"`
	EntityType CommentEntityType `json:"entity_type,omitempty"`
	EntityId   string            `json:"entity_id,omitempty"`
	CreatedAt  *TimestampFilter  `json:"created_at,omitempty"`
	SortBy     *SortFilter       `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CommentListRequest) payload() any { return r }

// operation response
type CommentCreateResponse struct {
	Comment *Comment `json:"comment,omitempty"`
	apiResponse
}

// operation response
type CommentRetrieveResponse struct {
	Comment *Comment `json:"comment,omitempty"`
	apiResponse
}

// operation sub response
type CommentListCommentResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}

// operation response
type CommentListResponse struct {
	List       []*CommentListCommentResponse `json:"list,omitempty"`
	NextOffset string                        `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CommentDeleteResponse struct {
	Comment *Comment `json:"comment,omitempty"`
	apiResponse
}
