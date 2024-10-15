package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/building_base/model"
)

func (r *Repository) AddBuilding(building model.Building) (int, error) {
	var id int

	sql, args, err := sq.
		Insert("buildings").
		Columns("name", "city", "year_completed", "floors").
		Values(building.Name, building.City, building.YearCompleted, building.Floors).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create name group creation request: %w", err)
	}

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to proccess news create query: %w", err)
	}

	return id, nil

}
