package domain

const (
	SuperAdminRole = "superadmin" // работает в рамках всех компаний
	AdminRole      = "admin"      // работает в рамках своей компании(crud all)
	UserRole       = "user"       // работает в рамках своей компании(read only)
)

var AllowedRoles = []string{AdminRole, UserRole}
