package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"  validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"required,gte=1,lte=120"`
}
