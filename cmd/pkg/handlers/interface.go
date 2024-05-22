package handlers

import (
	"net/http"
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
	context     ContextRequest
	help        Helper.Helper
}

func NewHandlerInstance(portNumber int, billMod BillModel, usrMod UsrModel, tokMod TokenModel, permits PermissionsModel, context ContextRequest) *Handler {
	return &Handler{
		portNumber:  portNumber,
		bills:       billMod,
		users:       usrMod,
		tokens:      tokMod,
		permissions: permits,
		context:     context,
	}
}

type BillModel interface {
	Create(payload *domain.Bill) error
	Delete(id int) error
	Fetch(bill *domain.Bill, id int64) error
	DateFetch(string, string, *domain.Users) ([]*domain.Bill, error)
	Update(*domain.Bill) error
}

type UsrModel interface {
	Fetch(email string) (*domain.Users, error)
	Create(*domain.Users) error
}

type TokenModel interface {
	Insert(token *auth.Token) error
	New(user int64, ttl time.Duration, scope string) (*auth.Token, error)
	DeleteAllTokensFromUser(scope string, userID int) error
}

type PermissionsModel interface {
	GrantPermissionToUser(userID int64, codes ...string) error
	GenerateUserPermissions() []string
}

type ContextRequest interface {
	ContextGetUser(r *http.Request) *domain.Users
}
