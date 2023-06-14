package backoffice

type GetUserListRequest struct {
	Page  int
	Limit int
	Email string
}
