package app

func (a *Application) setRoutes() {
	a.Server.Get("/v1/healthCheck", a.RequireUserAuth(a.Dependencies.Handlers.HealthCheckGET))
	a.Server.Get("/v1/bills/fetch/{id}", a.Dependencies.Handlers.FetchBillGET)
	a.Server.Get("/v1/bills/fetchAll", a.Dependencies.Handlers.BillsFetchByDateGET)
  a.Server.Get("/v1/users/fetch", a.Dependencies.Handlers.FetchUserByMailGET)

	a.Server.Post("/v1/sendJsonTest", a.Dependencies.Handlers.SendJson)
	a.Server.Post("/v1/bills/create", a.Dependencies.Handlers.BillingPOST)
  a.Server.Post("/v1/users/create", a.Dependencies.Handlers.UserCreatePOST)
  a.Server.Post("/v1/tokens/authenticate", a.Dependencies.Handlers.CreateAuthTokenPOST)

	a.Server.Delete("/v1/bills/delete/{id}", a.Dependencies.Handlers.SingleBillDELETE)

	a.Server.Put("/v1/updateBill", a.Dependencies.Handlers.UpdateBillPUT)
}
