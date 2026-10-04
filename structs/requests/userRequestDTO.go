package requests

type UpdateUserRequestDTO struct {
	FirstName   string `orm:"size(255)"`
	LastName    string `orm:"size(255)"`
	Username    string `orm:"size(255); omitempty; null"`
	PhoneNumber string `orm:"size(255); omitempty; null"`
	Gender      string `orm:"size(10); omitempty; null"`
	Dob         string `orm:"size(50); omitempty; null"`
	Address     string `orm:"size(255); omitempty; null"`
	BranchId    int64
	RoleId      int64
}

type UpdateUserRoleRequestDTO struct {
	RoleId int64
}

type UpdateUserBranchRequestDTO struct {
	BranchId int64
}

type AddRoleRequestDTO struct {
	Name        string
	Description string
}

type AddRoleRequest struct {
	Role        string
	Description string
}

type UpdateRolePermissionRequestDTO struct {
	Role           string
	Action         string
	PermissionCode string
}

type UpdateRolePermissionRequest struct {
	Role           string
	Action         string
	PermissionCode string
	ActionCode     string
}
