package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetNeed(ctx context.Context, id int) (*domain.Need, error) {
	var resp domain.Need
	err := d.Conn.QueryRowContext(ctx, `select id, client_id, realtor_id, property_type,
                  property_address_city, property_address_street, property_address_building, 
                  property_address_floor, min_floor, max_floor, min_room_count, max_room_count,
                  min_square, max_square, min_floor_count, max_floor_count, min_price, max_price from needs where id = ?`, id).Scan(
		&resp.Id,
		&resp.ClientId,
		&resp.RealtorId,
		&resp.PropertyType,
		&resp.Address.City,
		&resp.Address.Street,
		&resp.Address.Building,
		&resp.Address.Floor,
		&resp.MinFloor,
		&resp.MaxFloor,
		&resp.MinRooms,
		&resp.MaxRooms,
		&resp.MinSquare,
		&resp.MaxSquare,
		&resp.MinFloorCount,
		&resp.MaxFloorCount,
		&resp.MinPrice,
		&resp.MaxPrice,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
