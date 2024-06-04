package domain

func (b *Bill) ValidateBill(v validate) bool {
	v.Check(b.BillID != "", "idFactura", "must be non-empty")
	v.Check(len(b.BillID) <= 40, "idFactura", "must not be greater than 40 characters")

	v.Check(b.Date != "", "fechaEmision", "must be non-empty")
	v.Check(len(b.Date) <= 34, "fechaEmision", "must not be greater than 35 characters")

	v.Check(b.TotalAmmount != nil, "montoTotal", "must be provided")
	v.Check(*b.TotalAmmount > 0, "montoTotal", "must be a positive ammount")

	v.Check(b.Provider.Name != "", "nombreProveedor", "must be provided")
	v.Check(len(b.Provider.Name) <= 50, "nombreProveedor", "must be less than 50 characters")

	v.Check(b.Provider.ProviderID != "", "idProveedor", "must be provided")
	v.Check(len(b.Provider.ProviderID) <= 50, "idProveedor", "must be less than 50 characters")

	return v.Valid()
}


