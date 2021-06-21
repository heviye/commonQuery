package commonQuery

import (
	"database/sql"
	"log"
)

func Query(rows *sql.Rows) ([]Rower, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Printf("[query] rows.Columns err:%s\n", err.Error())
		return nil, err
	}

	colLen := len(columns)
	values := make([]interface{}, colLen)
	valuePtrs := make([]interface{}, colLen)

	for i := 0; i < colLen; i++ {
		valuePtrs[i] = &values[i]
	}

	result := make([]Rower, 0)
	for rows.Next() {
		err = rows.Scan(valuePtrs...)
		if err != nil {
			log.Printf("[query] rows.Scan err:%s\n", err.Error())
			return nil, err
		}

		retRow := make(Row)

		for i, col := range columns {
			val := values[i]
			switch realVal := val.(type) {
			case []uint8:
				retRow[col] = string(realVal)
			default:
				retRow[col] = val
			}
		}

		result = append(result, retRow)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[query] rows.Err err:%s\n", err.Error())
		return nil, err
	}

	return result, nil
}
