package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateProperty(ctx context.Context, req *domain.UpdatePropertyRequest) error {
	var (
		args   = make([]any, 0)
		values []string
	)

	if req.Type != nil {
		args = append(args, *req.Type)
		values = append(values, "type = ?")
	}

	if req.Address != nil {
		if req.Address.City != nil {
			args = append(args, *req.Address.City)
			values = append(values, "address_city = ?")
		}
		if req.Address.Street != nil {
			args = append(args, *req.Address.Street)
			values = append(values, "address_street = ?")
		}
		if req.Address.Building != nil {
			args = append(args, *req.Address.Building)
			values = append(values, "address_building = ?")
		}
		if req.Address.Floor != nil {
			args = append(args, *req.Address.Building)
			values = append(values, "address_floor = ?")
		}
	}

	if req.Coordinates != nil {
		if req.Coordinates.Longitude != nil {
			args = append(args, *req.Coordinates.Longitude)
			values = append(values, "cords_longitude = ?")
		}
		if req.Coordinates.Latitude != nil {
			args = append(args, *req.Coordinates.Latitude)
			values = append(values, "cords_latitude = ?")
		}
	}

	if req.Floor != nil {
		args = append(args, *req.Floor)
		values = append(values, "floor = ?")
	}

	if req.RoomCount != nil {
		args = append(args, *req.RoomCount)
		values = append(values, "room_count = ?")
	}

	if req.Square != nil {
		args = append(args, *req.Square)
		values = append(values, "square = ?")
	}

	if req.FloorCount != nil {
		args = append(args, *req.FloorCount)
		values = append(values, "floor_count = ?")
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update properties set %s where id = ?", q)
	args = append(args, req.PropertyId)

	_, err := d.Conn.ExecContext(ctx, query, args...)
	return err
}
