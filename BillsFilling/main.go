package main

import (
	"context"
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"time"

	rnd "math/rand"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type JsonObjects map[string]interface{}

func (j JsonObjects) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (a *JsonObjects) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Bill struct {
	idfactura    string
	fechaemision string
	montoTotal   int
	proveedor    struct {
		nombre         string
		identificacion string
	}
	Detalles   JsonObjects `json:"detalles"`
	Miscelaneo JsonObjects `json:"miscelaneo"`
}

func main() {

  dsn := "a"

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close(context.Background())

	var Bills [][]interface{}
	for i := 0; i < 20; i++ {
		min := time.Date(2019, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		max := time.Date(2024, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		delta := max - min

    sec := rnd.Int63n(delta) + min
    randomDate := time.Unix(sec,0).Format(time.RFC3339)

		randomNumber := rnd.Intn(1500) + 550 + rnd.Intn(300)*20/100

		randomNumber2 := randomNumber * 15 / 100

		randomString := make([]byte, 10)
		_, _ = rand.Read(randomString)

		randomString2 := make([]byte, 10)
		_, _ = rand.Read(randomString2)

		Bills = append(Bills, []interface{}{
			uuid.NewString(),
			randomDate,
			randomNumber,
			"ENEE",
			"619bfacd-b198-45cf-bf84-4702e62e4eb4",
			JsonObjects{
				"revenueTax": randomNumber2,
				"clientName": hex.EncodeToString(randomString),
				"clientType": "Residencial",
				"address":    hex.EncodeToString(randomString2),
			},
			JsonObjects{
				"energyConsumption": randomNumber * 30 / 100,
				"tariffType":        "Residencial",
			},
		})
	}

	log.Println(Bills)

	_, err = conn.CopyFrom(
		ctx,
		pgx.Identifier{"bills"},
		[]string{
			"id_factura", "fecha_emision",
			"monto_total", "nombre_proveedor",
			"id_proveedor", "detalles",
			"miscelaneo",
		},
		pgx.CopyFromRows(Bills),
	)
	if err != nil {
		log.Fatal(err)
	}
}
