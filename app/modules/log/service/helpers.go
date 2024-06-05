package service

import (
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
)

func (srv *logSrv) convertModelsToListResponse(data []models.Log) []payload.LogRes {
	var (
		records = []payload.LogRes{}
	)
	for i := 0; i < len(data); i++ {
		temp := payload.LogRes{
			ID:        data[i].ID,
			DateTime:  data[i].LogDateTime,
			Name:      data[i].Name,
			Action:    data[i].Action,
			UserID:    data[i].UserID,
			UserName:  data[i].User.Name,
			CreatedAt: data[i].CreatedAt,
			UpdatedAt: data[i].UpdatedAt,
		}
		records = append(records, temp)
	}

	return records
}
