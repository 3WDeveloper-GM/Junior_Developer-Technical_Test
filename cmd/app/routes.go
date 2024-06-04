package app

import (
	"github.com/go-chi/chi/v5"
)

const (
	USER_WRITE_PERMISSION  = "users:write"
	USER_READ_PERMISSION   = "users:read"
	BILLS_READ_PERMISSION  = "bills:read"
	BILLS_WRITE_PERMISSION = "bills:write"
)

func (a *Application) setRoutes() {
	a.Server.Post("/public/login", a.Dependencies.Handlers.LoginJWTtokenPOST)

	a.Server.Group(func(r chi.Router) {
		r.Use(a.JWTAuthentication)

		r.Get("/public/whoAmI", a.Dependencies.Handlers.WhoAmIGET)
		r.Get("/public/healthCheck", a.RequireUserAuth(a.Dependencies.Handlers.HealthCheckGET))
		r.Get("/public/bills/fetch", a.RequirePermissions(BILLS_READ_PERMISSION, a.Dependencies.Handlers.BillsFetchByDateGET))

		r.Put("/public/bills/update", a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.BillUpdateClientPUT))

		r.Delete("/public/bills/delete/{id}", a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.SingleBillDELETE))

		r.Post("/public/bills/create",
			a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.BillingPOST))
	})

	a.Server.Group(func(r chi.Router) {
		r.Use(a.Authenticate)
		r.Get("/v1/healthCheck", a.RequireUserAuth(a.Dependencies.Handlers.HealthCheckGET))

		r.Get("/v1/bills/fetch/{id}",
			a.RequirePermissions(BILLS_READ_PERMISSION, a.Dependencies.Handlers.FetchBillGET))
		r.Get("/v1/bills/fetchAll",
			a.RequirePermissions(BILLS_READ_PERMISSION, a.Dependencies.Handlers.BillsFetchByDateGET))
		r.Get("/v1/users/fetch",
			a.RequirePermissions(USER_READ_PERMISSION, a.Dependencies.Handlers.FetchUserByMailGET))

		r.Post("/v1/sendJsonTest",
			a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.SendJson))
		r.Post("/v1/bills/create",
			a.RequirePermissions(BILLS_WRITE_PERMISSION, a.Dependencies.Handlers.BillingPOST))
		r.Post("/v1/users/create",
			a.RequirePermissions(USER_WRITE_PERMISSION, a.Dependencies.Handlers.UserCreatePOST))

		r.Post("/v1/tokens/authenticate", a.Dependencies.Handlers.CreateAuthTokenPOST)

		r.Delete("/v1/bills/delete/{id}", a.Dependencies.Handlers.SingleBillDELETE)

		r.Put("/v1/updateBill", a.Dependencies.Handlers.UpdateBillPUT)
	})
}
