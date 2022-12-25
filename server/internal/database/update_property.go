package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateProperty(ctx context.Context, req *domain.UpdatePropertyRequest) error {
	var (
		argId  = 1
		args   = make([]any, 0)
		values []string
	)

	if req.Type != nil {
		args = append(args, *req.Type)
		values = append(values, fmt.Sprintf("type = $%d", argId))
		argId++
	}

	if req.Address != nil {
		if req.Address.City != nil {
			args = append(args, *req.Address.City)
			values = append(values, fmt.Sprintf("address_city = $%d", argId))
			argId++
		}
		if req.Address.Street != nil {
			args = append(args, *req.Address.Street)
			values = append(values, fmt.Sprintf("address_street = $%d", argId))
			argId++
		}
		if req.Address.Building != nil {
			args = append(args, *req.Address.Building)
			values = append(values, fmt.Sprintf("address_building = $%d", argId))
			argId++
		}
		if req.Address.Floor != nil {
			args = append(args, *req.Address.Building)
			values = append(values, fmt.Sprintf("address_floor = $%d", argId))
			argId++
		}
	}

	if req.Coordinates != nil {
		if req.Coordinates.Longitude != nil {
			args = append(args, *req.Coordinates.Longitude)
			values = append(values, fmt.Sprintf("cords_longitude = $%d", argId))
			argId++
		}
		if req.Coordinates.Latitude != nil {
			args = append(args, *req.Coordinates.Latitude)
			values = append(values, fmt.Sprintf("cords_latitude = $%d", argId))
			argId++
		}
	}

	if req.Floor != nil {
		args = append(args, *req.Floor)
		values = append(values, fmt.Sprintf("floor = $%d", argId))
		argId++
	}

	if req.RoomCount != nil {
		args = append(args, *req.RoomCount)
		values = append(values, fmt.Sprintf("room_count = $%d", argId))
		argId++
	}

	if req.Square != nil {
		args = append(args, *req.Square)
		values = append(values, fmt.Sprintf("square = $%d", argId))
		argId++
	}

	if req.FloorCount != nil {
		args = append(args, *req.FloorCount)
		values = append(values, fmt.Sprintf("floor_count = $%d", argId))
		argId++
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update properties set %s where id = $%d", q, argId)
	args = append(args, req.PropertyId)

	_, err := d.Conn.Exec(ctx, query, args...)
	return err
}
