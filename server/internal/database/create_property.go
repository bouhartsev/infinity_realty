package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) CreateProperty(ctx context.Context, req *domain.CreatePropertyRequest) (int, error) {
	var id int

	row := d.Conn.QueryRow(ctx, "insert into properties(type, address_city, address_street, address_building, address_floor, cords_latitude, cords_longitude, floor, room_count, square, floor_count) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id",
		req.Type,
		req.Address.City,
		req.Address.Street,
		req.Address.Building,
		req.Address.Floor,
		req.Coordinates.Latitude,
		req.Coordinates.Longitude,
		req.Floor,
		req.RoomCount,
		req.Square,
		req.FloorCount,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
