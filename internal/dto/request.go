package dto

type SearchRequestParams struct {
	Name     string `form:"name" validate:"omitempty,max=100"`
	Page     int    `form:"page,default=1" validate:"min=1"`
	PageSize int    `form:"pageSize,default=10" validate:"min=1,max=100"`
}

// IDRequestParams for ID-based handlers (path params)
type IDRequestParams struct {
	ID uint `uri:"id" validate:"required,min=1"`
}
