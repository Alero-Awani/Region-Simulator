package dto

type CreateCategoryRequest struct {
	Name         string `json:"name"`
	ParentId     uint   `json:"parent_id"`
	ImageUrl     string `json:"image_url"`
	DisplayOrder uint   `json:"display_order"`
}
