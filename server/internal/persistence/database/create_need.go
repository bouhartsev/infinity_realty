package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateNeed(ctx context.Context, req *domain.CreateNeedRequest) (int, error) {
	res, err := d.Conn.ExecContext(ctx, `insert into needs(client_id, realtor_id, property_type,
                  property_address_city, property_address_street, property_address_building, 
                  property_address_floor, min_floor, max_floor, min_room_count, max_room_count,
                  min_square, max_square, min_floor_count, max_floor_count, min_price, max_price)
							 values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.ClientId,
		req.RealtorId,
		req.PropertyType,
		req.Address.City,
		req.Address.Street,
		req.Address.Building,
		req.Address.Floor,
		req.MinFloor,
		req.MaxFloor,
		req.MinRooms,
		req.MaxRooms,
		req.MinSquare,
		req.MaxSquare,
		req.MinFloorCount,
		req.MaxFloorCount,
		req.MinPrice,
		req.MaxPrice,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
