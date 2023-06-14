package backoffice

type UserListResponse struct {
	Id          *int    `json:"id"`
	Email       *string `json:"email"`
	UserType    *int    `json:"userType"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	EmailStatus *bool   `json:"emailStatus"`
}
