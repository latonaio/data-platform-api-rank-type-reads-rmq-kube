package dpfm_api_output_formatter

import (
	"data-platform-api-rank-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToRankType(rows *sql.Rows) (*[]RankType, error) {
	defer rows.Close()
	rankType := make([]RankType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.RankType{}

		err := rows.Scan(
			&pm.RankType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &rankType, nil
		}

		data := pm
		rankType = append(rankType, RankType{
			RankType:				data.RankType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &rankType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.RankType,
			&pm.Language,
			&pm.RankTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			RankType:     			data.RankType,
			Language:          		data.Language,
			RankTypeName:			data.RankTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
