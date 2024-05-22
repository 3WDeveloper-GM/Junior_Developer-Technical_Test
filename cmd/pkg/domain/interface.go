package domain

import (
	"time"

	jsonbobjects "github.com/3WDeveloper-GM/billings/cmd/pkg/domain/jsonbObjects"
)

// Bill is a structure to model the bills that are coming from the clients
type Bill struct {
	SysID        int64       `json:"numeroRegistro"`
	Created_at   time.Time `json:"-"`
	BillID       string    `json:"idFactura"`
	Date         string    `json:"fechaEmision"`
	TotalAmmount *int      `json:"montoTotal"`
	Provider     struct {
		Name       string `json:"nombre"`
		ProviderID string `json:"identificacion"`
	} `json:"proveedor"`
	Details jsonbobjects.JsonObjects `json:"detalles,omitempty"`
	Misc    jsonbobjects.JsonObjects `json:"miscelaneo,omitempty"`
	Version int                      `json:"version,omitempty"`
}

type validate interface {
	Valid() bool
	Check(bool, string, string)
	AddErrorKey(string, string)
}

// User is a structure that model the end user
type Users struct {
	SysID      int64       `json:"-"`
	Created_at time.Time `json:"fechaCreacion"`
	ProviderID string    `json:"idProveedor"`
	Name       string    `json:"nombre"`
	Email      string    `json:"email"`
	Password   passWord  `json:"-"`
	Activated  bool      `json:"activado"`
	Version    int       `json:"version,omitempty"`
}

type passWord struct {
	Plaintext *string
	Hash      []byte
}
