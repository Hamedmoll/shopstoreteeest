package mysql

import (
	"shopstoretest/entity"
)

func (mysql MySQLDB) GetUserPermissionTitles(userID uint) ([]entity.PermissionTitle, error) {
	rows, qErr := mysql.DB.Query("select permission_id from access_controls where actor_id = ?", userID)
	if qErr != nil {

		return nil, qErr
	}

	permissionTitleIDs := make([]uint, 0)
	var tmpPermissionTitleID uint

	for rows.Next() {
		sErr := rows.Scan(&tmpPermissionTitleID)
		if sErr != nil {

			return nil, sErr
		}

		permissionTitleIDs = append(permissionTitleIDs, tmpPermissionTitleID)
	}

	permissionTitles := make([]entity.PermissionTitle, 0)
	var tmpPermissionTitle entity.PermissionTitle

	for _, id := range permissionTitleIDs {
		row := mysql.DB.QueryRow("select title from permissions where id = ?", id)
		sErr := row.Scan(&tmpPermissionTitle)
		if sErr != nil {

			return nil, sErr
		}

		permissionTitles = append(permissionTitles, tmpPermissionTitle)
	}

	return permissionTitles, nil
}
