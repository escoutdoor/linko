package dto

type UpdateUserParams struct {
	ID          string
	FirstName   *string
	LastName    *string
	Email       *string
	PhoneNumber *string
	Password    *string
}

type CreateUserParams struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
}
