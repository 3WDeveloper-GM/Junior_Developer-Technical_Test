package app

const (
	USER_WRITE_PERMISSION  = "users:write"
	USER_READ_PERMISSION   = "users:read"
	BILLS_READ_PERMISSION  = "bills:read"
	BILLS_WRITE_PERMISSION = "bills:write"
)

func (a *Application) setRoutes() {
	a.Server.Get("/v1/healthCheck", a.RequireUserAuth(a.Dependencies.Handlers.HealthCheckGET))

	a.Server.Get("/v1/bills/fetch/{id}",
		a.RequirePermissions(BILLS_READ_PERMISSION, a.Dependencies.Handlers.FetchBillGET))
	a.Server.Get("/v1/bills/fetchAll",
		a.RequirePermissions(BILLS_READ_PERMISSION, a.Dependencies.Handlers.BillsFetchByDateGET))
	a.Server.Get("/v1/users/fetch",
		a.RequirePermissions(USER_READ_PERMISSION, a.Dependencies.Handlers.FetchUserByMailGET))

	a.Server.Post("/v1/sendJsonTest",
		a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.SendJson))
	a.Server.Post("/v1/bills/create",
		a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.BillingPOST))
	a.Server.Post("/v1/users/create",
		a.RequirePermissions(USER_WRITE_PERMISSION, a.Dependencies.Handlers.UserCreatePOST))

	a.Server.Post("/v1/tokens/authenticate", a.Dependencies.Handlers.CreateAuthTokenPOST)

	a.Server.Delete("/v1/bills/delete/{id}", a.Dependencies.Handlers.SingleBillDELETE)

	a.Server.Put("/v1/updateBill", a.Dependencies.Handlers.UpdateBillPUT)
}
