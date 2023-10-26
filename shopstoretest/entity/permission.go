package entity

type Permission struct {
	ID    uint
	Title PermissionTitle
}

type PermissionTitle string

const (
	AddCategoryPermission = PermissionTitle("add_category")
	AddProductPermission  = PermissionTitle("add_product")
)
