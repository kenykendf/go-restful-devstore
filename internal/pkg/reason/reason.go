package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"
)

var (
	CategoryNotFound        = "category not found"
	CategoryCannotCreate    = "cannot create category"
	CategoryCannotBrowse    = "cannot browse category"
	CategoryCannotUpdate    = "cannot update category"
	CategoryCannotDelete    = "cannot delete category"
	CategoryCannotGetDetail = "cannot get detail"
)

var (
	ProductNotFound        = "product not found"
	ProductCannotCreate    = "cannot create product"
	ProductCannotBrowse    = "cannot browse product"
	ProductCannotUpdate    = "cannot update product"
	ProductCannotDelete    = "cannot delete product"
	ProductCannotGetDetail = "cannot get detail"
)

var (
	UserAlreadyExist   = "user already exist"
	RegisterFailed     = "cannot register user."
	FailedLogin        = "login failed, make sure credentials are correct"
	Unauthorized       = "user unauthorized"
	InvalidAccess      = "unable to access this resource."
	FailedLogout       = "failed logout."
	FailedRefreshToken = "failed to refresh token, please check your token." //nolint
)
