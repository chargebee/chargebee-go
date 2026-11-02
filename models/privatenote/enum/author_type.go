package enum

type AuthorType string

const (
	AuthorTypeSiteUser AuthorType = "site_user"
	AuthorTypeAgent    AuthorType = "agent"
	AuthorTypeOther    AuthorType = "other"
)
