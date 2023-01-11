package database

import (
	"context"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (d *Database) GetProperty(ctx context.Context, id int) (*domain.Property, error) {
	var prop domain.Property

	row := d.Conn.QueryRowContext(ctx, `select id, type, address_city, address_street, address_building, address_floor, cords_latitude, cords_longitude, floor, room_count, square, floor_count from properties where id = ?`, id)
	err := row.Scan(
		&prop.Id,
		&prop.Type,
		&prop.Address.City,
		&prop.Address.Street,
		&prop.Address.Building,
		&prop.Address.Floor,
		&prop.Coordinates.Latitude,
		&prop.Coordinates.Longitude,
		&prop.Floor,
		&prop.RoomCount,
		&prop.Square,
		&prop.FloorCount,
	)
	if err != nil {
		return nil, err
	}

	return &prop, nil
}
