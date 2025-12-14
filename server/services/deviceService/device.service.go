package deviceService

import (
	"database/sql"
	"errors"
	"firmware_server/database"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"fmt"
	"strings"
)

func GetDevice(id *int) ([]tables.Device, error) {
	var res *sql.Rows
	var err error
	var query string
	if id != nil {
		query = fmt.Sprintf("select * from %s where id = ?", tablesName.Devices)
		res, err = database.DB.Query(query, id)
	} else {
		query = fmt.Sprintf("select * from %s where", tablesName.Devices)
		res, err = database.DB.Query(query)
	}

	defer res.Close()

	if err != nil {
		return nil, err
	}

	var devices []tables.Device

	for res.Next() {
		var device tables.Device
		if err := res.Scan(&device.Id, &device.Firmware_Id, &device.Product_Id, &device.Serial); err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return devices, nil

}

func AddDevice(id uint, serial uint, productId uint, fwId uint) error {
	_, err := database.DB.Exec("insert into ? (id, serial, product_id, firmware_id) values (? ? ? ?)", tablesName.Devices, id, serial, productId, fwId)

	if err != nil {
		return err
	}

	return nil
}

func UpdateDevice(p UpdateDeviceParams) error {
	if p.ID == 0 {
		return errors.New("device id is required")
	}

	updates := []string{}
	args := []any{}

	if p.Serial != nil {
		updates = append(updates, "serial = ?")
		args = append(args, *p.Serial)
	}

	if p.ProductID != nil {
		updates = append(updates, "product_id = ?")
		args = append(args, *p.ProductID)
	}

	if p.FirmwareID != nil {
		updates = append(updates, "fw_id = ?")
		args = append(args, *p.FirmwareID)
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, p.ID)

	query := fmt.Sprintf(
		"UPDATE devices SET %s WHERE id = ?",
		strings.Join(updates, ", "),
	)

	_, err := database.DB.Exec(query, args...)
	return err
}
