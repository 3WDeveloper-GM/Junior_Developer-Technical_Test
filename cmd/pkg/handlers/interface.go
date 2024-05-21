package handlers

import (
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/auth"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
	Helper "github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/helper"
)

type Handler struct {
	portNumber  int
	bills       BillModel
	users       UsrModel
	tokens      TokenModel
	permissions PermissionsModel
	help        Helper.Helper
}

func NewHandlerInstance(portNumber int, billMod BillModel, usrMod UsrModel, tokMod TokenModel, permits PermissionsModel) *Handler {
	return &Handler{
		portNumber: portNumber,
		bills:      billMod,
		users:      usrMod,
		tokens:     tokMod,
    permissions: permits,
	}
}

type BillModel interface {
	Create(payload *domain.Bill) error
	Delete(id int) error
	Fetch(bill *domain.Bill, id int) error
	DateFetch(string, string) ([]*domain.Bill, error)
	Update(*domain.Bill) error
}

type UsrModel interface {
	Fetch(email string) (*domain.Users, error)
	Create(*domain.Users) error
}

type TokenModel interface {
	Insert(token *auth.Token) error
	New(user int, ttl time.Duration, scope string) (*auth.Token, error)
	DeleteAllTokensFromUser(scope string, userID int) error
}

type PermissionsModel interface {
	GrantPermissionToUser(userID int, codes ...string) error
	GenerateUserPermissions() []string
}
