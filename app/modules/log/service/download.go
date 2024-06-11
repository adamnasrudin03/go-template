package service

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (srv *logSrv) Download(ctx *gin.Context, params *payload.ListLogReq) (err error) {
	opName := "LogService-Download"
	defer helpers.PanicRecover(opName)

	err = params.Validate()
	if err != nil {
		return err
	}

	records, err := srv.Repo.GetList(ctx, *params)
	if err != nil {
		srv.Logger.Errorf("%v error get records: %v ", opName, err)
		return helpers.ErrDB()
	}

	if len(records) == 0 {
		return helpers.ErrDataNotFound("Daftar log histori aktifitas pengguna", "List log activity history user")
	}
	nameUser := records[0].User.Name
	emailUser := records[0].User.Email

	today := time.Now().Unix()
	sheetName := "Sheet1"
	fileName := fmt.Sprintf("%s%d%s", "activity-history-", today, ".xlsx")
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			srv.Logger.Errorf("%v failed close file: %v ", opName, err)
		}
	}()

	// set info header
	index := 0
	for idx, row := range [][]interface{}{
		{"Nama Pengguna", nameUser},
		{"Email Pengguna", emailUser},
		{nil},
		{"Waktu", "Tipe", "Nama Aktifitas"},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		index = idx
		if err != nil {
			srv.Logger.Errorf("%v failed set info header: %v ", opName, err)
			return err
		}
		f.SetSheetRow(sheetName, cell, &row)
	}
	index++

	// set data
	for _, val := range records {
		cell, err := excelize.CoordinatesToCellName(1, index+1)
		if err != nil {
			srv.Logger.Errorf("%v failed set records to row: %v ", opName, err)
			return err
		}

		activityTime := val.LogDateTime
		if activityTime.IsZero() {
			activityTime = val.CreatedAt
		}

		exportedData := []interface{}{
			helpers.CheckTimeIsZeroToString(activityTime, helpers.FormatDateTime),
			val.Action,
			val.Name,
		}
		f.SetSheetRow(sheetName, cell, &exportedData)
		index++
	}

	// set style
	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err := f.SetCellStyle(sheetName, "A1", "A4", style); err != nil {
		srv.Logger.Errorf("%v failed set style cell A1 s/d A4: %v ", opName, err)
		return err
	}
	if err := f.SetCellStyle(sheetName, "B4", "C4", style); err != nil {
		srv.Logger.Errorf("%v failed set style cell B4 s/d C4: %v ", opName, err)
		return err
	}

	// download
	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		srv.Logger.Errorf("%v failed write: %v ", opName, err)
		return err
	}

	downloadName := time.Now().UTC().Format(fileName)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename="+downloadName)
	ctx.Data(http.StatusOK, "application/octet-stream", b.Bytes())
	return nil
}
