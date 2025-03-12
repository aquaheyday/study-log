package repositories

import (
	"log"
	"my-production-app/database"
	"my-production-app/models"
)

// FetchUserMenu 사용자 이메일과 그룹을 기반으로 메뉴 목록 조회
func FetchUserMenu(email, group string) ([]models.MenuItem, error) {
	query := `
		SELECT
		    parent_menu.id,
			parent_menu.name,
			parent_menu.url,
			parent_menu.level,
			parent_menu.sort,
			COALESCE(parent_menu.parent_id, '') as parent_id,
			'' AS role
		FROM admin_user_permissions aup
		JOIN permissions per ON aup.permission_id = per.id AND per.type = 'url'
		JOIN menu ON per.detail = menu.url
		JOIN menu parent_menu ON menu.parent_id = parent_menu.id
		WHERE menu.group = ?
		AND aup.email = ?
		AND menu.is_active = 1
		GROUP BY parent_menu.id
		
		UNION
		
		SELECT
		    menu.id,
			menu.name,
			COALESCE(menu.url, '') as url,  
			menu.level, 
			menu.sort,
			menu.parent_id,
			MAX(CASE 
				WHEN per.role = 'write' THEN 'write'
				WHEN per.role = 'read' THEN 'read'
				ELSE NULL
			END) AS role
		FROM admin_user_permissions aup
		JOIN permissions per ON aup.permission_id = per.id AND per.type = 'url'
		JOIN menu ON per.detail = menu.url
		WHERE menu.group = ?
		AND aup.email = ?
		AND menu.is_active = 1
		GROUP BY menu.url, menu.level;`

	rows, err := database.DB.Query(query, group, email, group, email)
	if err != nil {
		log.Println("Database query failed:", err)
		return nil, err
	}
	defer rows.Close()

	var menus []models.MenuItem
	for rows.Next() {
		var menu models.MenuItem
		if err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.URL,
			&menu.Level,
			&menu.Sort,
			&menu.ParentID,
			&menu.Role,
		); err != nil {
			log.Println("Failed to scan menu:", err)
			return nil, err
		}
		menus = append(menus, menu)
	}

	return menus, nil
}
