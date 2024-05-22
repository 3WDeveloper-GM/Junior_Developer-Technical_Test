package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
)

type billModel struct {
	DB *sql.DB
}

func (b *billModel) Create(payload *domain.Bill) error {
	stmt := `
    INSERT INTO bills (id_factura, fecha_emision, monto_total, nombre_proveedor,id_proveedor,detalles,miscelaneo)
    VALUES($1,$2,$3,$4,$5,$6,$7)
    RETURNING id, created_at
  `

	args := []interface{}{
		payload.BillID,
		payload.Date,
		payload.TotalAmmount,
		payload.Provider.Name,
		payload.Provider.ProviderID,
		payload.Details,
		payload.Misc,
	}

	return b.DB.QueryRow(stmt, args...).Scan(&payload.SysID, &payload.Created_at)
}

func (b *billModel) Delete(id int) error {
	stmt := `
    DELETE FROM bills 
    WHERE id = $1
  `

	result, err := b.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return ErrNoRows
	}

	return nil
}

func (b *billModel) Fetch(bill *domain.Bill, id int64) error {
	stmt := `
    SELECT id, id_factura, fecha_emision, monto_total,detalles, miscelaneo
    FROM bills
    WHERE id = $1 AND id_proveedor = $2
  `
  
  args := []interface{}{id, bill.Provider.ProviderID}

	return b.DB.QueryRow(stmt, args...).Scan(
		&bill.SysID,
		&bill.BillID,
		&bill.Date,
		&bill.TotalAmmount,
		&bill.Details,
		&bill.Misc,
	)
}

func (b *billModel) DateFetch(startingDate string, endingDate string, user *domain.Users) ([]*domain.Bill, error) {
	stmt := `
    SELECT id,id_factura,fecha_emision,monto_total,detalles, miscelaneo
    FROM bills
    WHERE created_at >= $1 AND created_at <= $2 AND id_proveedor = $3 
  `

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	args := []interface{}{startingDate, endingDate,user.ProviderID}

	rows, err := b.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	bills := []*domain.Bill{}

	for rows.Next() {
    
    var bill domain.Bill  

    bill.Provider.ProviderID = user.ProviderID
    bill.Provider.Name = user.Name

		err := rows.Scan(
			&bill.SysID,
			&bill.BillID,
			&bill.Date,
			&bill.TotalAmmount,
			&bill.Details,
			&bill.Misc,
		)
		if err != nil {
			return nil, err
		}

		bills = append(bills, &bill)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bills, nil
}

func (b *billModel) Update(bill *domain.Bill) error {
	stmt := `
    UPDATE bills
    SET id_factura =$1, fecha_emision = $2, monto_total = $3,
    id_proveedor =$4, nombre_proveedor = $5, detalles =$6,
    miscelaneo =$7, version = version+1
    WHERE id = $8
    RETURNING version
  `

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	args := []interface{}{
		bill.BillID,
		bill.Date,
		bill.TotalAmmount,
		bill.Provider.ProviderID,
		bill.Provider.Name,
		bill.Details,
		bill.Misc,
		bill.SysID,
	}

	return b.DB.QueryRowContext(ctx,stmt,args...).Scan(&bill.Version)
}
