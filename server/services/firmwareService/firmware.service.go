package firmwareService

import (
	"firmware_server/database"
	firmwareDtos "firmware_server/dtos/firmware"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"fmt"
)

func GetFirmwares(product_id uint, id *uint) ([]tables.Firmwares, error) {
	var query string
	if id != nil {
		query = fmt.Sprintf("select id, product_id, version from %s where product_id = %d and id = %d", tablesName.Firmwares, product_id, *id)
	} else {
		query = fmt.Sprintf("select id, product_id, version from %s where product_id = %d", tablesName.Firmwares, product_id)
	}

	res, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	var firmwares []tables.Firmwares

	for res.Next() {
		var firmware tables.Firmwares
		if err := res.Scan(&firmware.Id, &firmware.Product_Id, &firmware.Version); err != nil {
			return nil, err
		}

		firmwares = append(firmwares, firmware)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return firmwares, nil

}

func AddFirmware(args firmwareDtos.AddFirmwareBody) error {
	query := fmt.Sprintf("insert into %s (version, product_id) values (\"%s\" , %d)", tablesName.Firmwares, *args.Version, *args.Product_Id)

	_, err := database.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func UpdateFirmware(args firmwareDtos.UpdateFirmwareBody) error {

	query := fmt.Sprintf("update %s set ", tablesName.Firmwares)

	if args.Version != nil {
		query += fmt.Sprintf("version = \"%s\" ,", *args.Version)

	}

	if args.Product_Id != nil {
		query += fmt.Sprintf("product_id = %d ", *args.Product_Id)

	}

	query = fmt.Sprintf("where id = %d", *args.Id)

	_, err := database.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func DeleteFirmware(args firmwareDtos.DeleteFirmwareBody) error {
	query := fmt.Sprintf("delete from %s where id =%d", tablesName.Firmwares, *args.Id)

	_, err := database.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
