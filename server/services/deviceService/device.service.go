package deviceService

import (
	"errors"
	"firmware_server/database"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetDevice(id *int) ([]tables.Device, error) {

	var query string
	if id != nil {
		query = fmt.Sprintf("select id, firmware_id, product_id, serial, api_key from %s where id = %d", tablesName.Devices, *id)
	} else {
		query = fmt.Sprintf("select id, firmware_id, product_id, serial, api_key from %s", tablesName.Devices)
	}

	res, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	var devices []tables.Device

	for res.Next() {
		var device tables.Device
		if err := res.Scan(&device.Id, &device.Firmware_Id, &device.Product_Id, &device.Serial, &device.Api_Key); err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return devices, nil

}

func AddDevice(serial string, productId uint, fwId uint) error {
	api_key := uuid.New()
	query := fmt.Sprintf("insert into %s (serial, product_id, firmware_id, api_key) values (%s, %d, %d, \"%s\")", tablesName.Devices, serial, productId, fwId, api_key)

	println(query)

	_, err := database.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func UpdateDevice(
	ID uint,
	Serial *string,
	ProductID *uint,
	FirmwareID *uint,
) error {
	if ID == 0 {
		return errors.New("device id is required")
	}

	updates := []string{}
	args := []any{}

	if Serial != nil {
		updates = append(updates, "serial = ?")
		args = append(args, *Serial)
	}

	if ProductID != nil {
		updates = append(updates, "product_id = ?")
		args = append(args, *ProductID)
	}

	if FirmwareID != nil {
		updates = append(updates, "fw_id = ?")
		args = append(args, *FirmwareID)
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, ID)

	query := fmt.Sprintf(
		"UPDATE devices SET %s WHERE id = ?",
		strings.Join(updates, ", "),
	)

	_, err := database.DB.Exec(query, args...)
	return err
}

func DeleteDevice(id uint) error {
	query := fmt.Sprintf("delete from %s where id = ?", tablesName.Devices)

	_, err := database.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
