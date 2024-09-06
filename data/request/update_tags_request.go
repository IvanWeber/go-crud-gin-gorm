package request

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Progress string `validate:"required,max=200,min=1" json:"progress"`
}
