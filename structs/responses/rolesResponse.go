package responses

// type RolesResponseDTO struct {
// 	StatusCode int
// 	Roles      *[]models.Roles
// 	StatusDesc string
// }

type Actions struct {
	ActionId    int64
	Action      string
	Description string
}

type Permissions struct {
	PermissionId          int64
	Permission            string
	PermissionCode        string
	PermissionDescription string
}

type Role_permissions struct {
	RolePermissionId int64
	Role             *Roles
	Permission       *Permissions
	Action           *Actions
}

type Roles struct {
	RoleId          int64  `orm:"auto"`
	Role            string `orm:"size(100)"`
	Description     string `orm:"size(500)"`
	Active          int
	RolePermissions []*Role_permissions
}

type RoleResponseDTO struct {
	StatusCode int
	Role       *Roles
	StatusDesc string
}

type RolesAllResponseDTO struct {
	StatusCode int
	Roles      *[]Roles
	StatusDesc string
}

type RolesAllGatewayResponseDTO struct {
	Success    bool
	Result     *[]Roles
	StatusDesc string
}

type RoleGatewayResponseDTO struct {
	Success    bool
	Result     *Roles
	StatusDesc string
}

type UserPermission struct {
	PermissionCode string
	ActionCode     string
}
