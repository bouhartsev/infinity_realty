package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) UpdateNeed(ctx context.Context, req *domain.UpdateNeedRequest) error {
	var (
		args   = make([]any, 0)
		values []string
	)

	if req.ClientId != nil {
		args = append(args, *req.ClientId)
		values = append(values, "client_id = ?")
	}

	if req.RealtorId != nil {
		args = append(args, *req.RealtorId)
		values = append(values, "realtor_id = ?")
	}

	if req.PropertyType != nil {
		args = append(args, *req.PropertyType)
		values = append(values, "property_type = ?")

	}

	if req.Address != nil {
		if req.Address.City != "" {
			args = append(args, req.Address.City)
			values = append(values, "address_city = ?")

		}
		if req.Address.Street != "" {
			args = append(args, req.Address.Street)
			values = append(values, "address_street = ?")

		}
		if req.Address.Building != "" {
			args = append(args, req.Address.Building)
			values = append(values, "address_building = ?")

		}
		if req.Address.Floor != "" {
			args = append(args, req.Address.Floor)
			values = append(values, "address_floor = ?")
		}
	}

	if req.MinSquare != nil {
		args = append(args, *req.MinSquare)
		values = append(values, "min_square = ?")
	}

	if req.MaxSquare != nil {
		args = append(args, *req.MaxSquare)
		values = append(values, "max_square = ?")
	}

	if req.MinRooms != nil {
		args = append(args, *req.MinRooms)
		values = append(values, "min_rooms = ?")
	}

	if req.MaxRooms != nil {
		args = append(args, *req.MaxRooms)
		values = append(values, "max_rooms = ?")
	}

	if req.MinFloor != nil {
		args = append(args, *req.MinFloor)
		values = append(values, "min_floor = ?")
	}

	if req.MaxFloor != nil {
		args = append(args, *req.MaxFloor)
		values = append(values, "max_floor = ?")
	}

	if req.MaxFloorCount != nil {
		args = append(args, *req.MaxFloorCount)
		values = append(values, "max_floor_count = ?")
	}

	if req.MinFloorCount != nil {
		args = append(args, *req.MinFloorCount)
		values = append(values, "min_floor_count = ?")
	}

	if req.MaxPrice != nil {
		args = append(args, *req.MaxPrice)
		values = append(values, "max_price = ?")
	}

	if req.MinPrice != nil {
		args = append(args, *req.MinPrice)
		values = append(values, "min_price = ?")
	}

	q := strings.Join(values, ",")
	query := fmt.Sprintf("update needs set %s where id = ?", q)
	args = append(args, req.NeedId)

	_, err := d.Conn.ExecContext(ctx, query, args...)
	return err
}
