package service

import (
	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
)

func (srv *logSrv) convertModelsToListResponse(data []models.Log) []dto.LogRes {
	var (
		records = []dto.LogRes{}
	)
	for i := 0; i < len(data); i++ {
		temp := dto.LogRes{
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
