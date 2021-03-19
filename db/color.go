package db

import (
	"encoding/hex"

	"github.com/aikon001/colorapiserver/models"
)

func (db Database) GetAllColors() (*models.ColorList, error) {
	list := &models.ColorList{}
	rows, err := db.Conn.Query("SELECT * FROM colors ORDER BY id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var color models.Color
		err := rows.Scan(&color.ID, &color.Name, &color.Hexadecimal, &color.R, &color.G, &color.B, &color.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Colors = append(list.Colors, color)

	}
	return list, nil
}

func (db Database) AddColor(color *models.Color) error {
	var id int

	query := `INSERT INTO colors (name, hexadecimal,r,g,b) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	if len(color.Hexadecimal) != 0 {
		byt, _ := hex.DecodeString(color.Hexadecimal)
		err := db.Conn.QueryRow(query, color.Name, color.Hexadecimal, byt[0], byt[1], byt[2]).Scan(&id)
		if err != nil {
			return err
		}

	} else {
		rgb := []byte{byte(color.R), byte(color.G), byte(color.B)}
		hexadecimal := hex.EncodeToString(rgb)
		err := db.Conn.QueryRow(query, color.Name, hexadecimal, color.R, color.G, color.B).Scan(&id)
		if err != nil {
			return err
		}

	}
	color.ID = id
	return nil

}

func (db Database) GetColorById(colorId int) (models.Color, error) {
	color := models.Color{}
	query := `SELECT * FROM colors WHERE id = $1;`
	row := db.Conn.QueryRow(query, colorId)
	err := row.Scan(&color.ID, &color.Name, &color.Hexadecimal, &color.R, &color.G, &color.B, &color.CreatedAt)
	return color, err
}

func (db Database) DeleteColor(colorId int) error {
	query := `DELETE FROM colors WHERE id = $1;`
	_, err := db.Conn.Exec(query, colorId)
	return err
}

func (db Database) UpdateColor(colorId int, colorData models.Color) (models.Color, error) {
	color := models.Color{}
	query := `UPDATE colors SET name=$1, hexadecimal=$2, r=$3, g=$4, b=$5 WHERE id=$6 RETURNING id, name, hexadecimal,R,G,B`
	if len(color.Hexadecimal) != 0 {
		byt, _ := hex.DecodeString(color.Hexadecimal)
		err := db.Conn.QueryRow(query, color.Name, color.Hexadecimal, byt[0], byt[1], byt[2], color.ID).Scan(&color.Name, &color.Hexadecimal, &color.R, &color.G, &color.B, &color.ID)
		if err != nil {
			return color, err
		}
	} else {
		rgb := []byte{byte(color.R), byte(color.G), byte(color.B)}
		hexadecimal := hex.EncodeToString(rgb)
		err := db.Conn.QueryRow(query, color.Name, hexadecimal, color.R, color.G, color.B, color.ID).Scan(&color.ID, &color.Name, &color.Hexadecimal, &color.R, &color.G, &color.B)
		if err != nil {
			return color, err
		}
	}
	return color, nil
}
