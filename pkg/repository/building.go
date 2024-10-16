package repository

import (
	"fmt"
	"github.com/kahuri1/building_base/model"
	"github.com/spf13/viper"
)

func (r *Repository) Building(filter model.Filter) (model.BuildingResponse, error) {
	var buildings model.BuildingResponse

	query := viper.GetString("sql.query")
	args := []interface{}{}
	paramIndex := viper.GetInt("param")

	if filter.City != "" {
		query += fmt.Sprintf(viper.GetString("sql.city"), paramIndex)
		args = append(args, filter.City)
		paramIndex++
	}
	if filter.Year != nil {
		query += fmt.Sprintf(viper.GetString("sql.year"), paramIndex)
		args = append(args, filter.Year)
		paramIndex++
	}
	if filter.Floors != nil {
		query += fmt.Sprintf(viper.GetString("sql.floors"), paramIndex)
		args = append(args, filter.Floors)
		paramIndex++
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return buildings, err
	}
	defer rows.Close()
	for rows.Next() {
		var building model.Building
		if err := rows.StructScan(&building); err != nil {
			return buildings, err
		}
		buildings.Buildings = append(buildings.Buildings, building)
	}

	return buildings, nil
}
