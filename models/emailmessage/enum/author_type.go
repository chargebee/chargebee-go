package enum

type AuthorType string

const (
	AuthorTypeContact  AuthorType = "contact"
	AuthorTypeSiteUser AuthorType = "site_user"
	AuthorTypeAgent    AuthorType = "agent"
	AuthorTypeOther    AuthorType = "other"
)
