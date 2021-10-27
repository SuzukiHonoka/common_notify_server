package user

type Permission int

const (
	// PermissionUser represents the User Management
	PermissionUser Permission = iota
	// PermissionAPI represents the API Management
	PermissionAPI Permission = iota
	// PermissionWorker represents the Worker Management
	PermissionWorker Permission = iota
	// PermissionSafeGuard represents the SafeGuard Management
	PermissionSafeGuard Permission = iota
	// PermissionGenerator represents the Generator Management
	PermissionGenerator Permission = iota
)
