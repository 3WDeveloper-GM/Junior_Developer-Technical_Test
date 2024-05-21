package models

import (
	"context"
	"database/sql"
	"time"
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermitModel struct {
	DB *sql.DB
}

func (pm *PermitModel) GetPermissionsFromUser(userID int) (Permissions, error) {
	stmt := `
    SELECT permissions.code
    FROM permissions
    INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
    INNER JOIN users ON users_permissions.user_id = users.id WHERE users.id = $1
    WHERE users.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := pm.DB.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}

	var permissions []string

	for rows.Next() {
		var permission string

		err = rows.Scan(&permission)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}