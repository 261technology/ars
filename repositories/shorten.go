package repositories

import (
	"context"

	"github.com/harisaginting/gwyn/common/log"
	"github.com/harisaginting/gwyn/common/utils/helper"
	model "github.com/harisaginting/gwyn/models"
	dao "github.com/harisaginting/gwyn/models/dao"
	httpModel "github.com/harisaginting/gwyn/models/http"
)

func (repo *Database) Get(ctx context.Context, p *model.Shorten) (err error) {
	qx := Connection()
	defer Close(qx)

	var table dao.Shorten

	if p.ID != 0 {
		table.ID = p.ID
		r := qx.First(&table)
		err = r.Error
		log.Error(ctx, err)
	} else {
		r := qx.Debug().Where("shortcode = ?", p.Shortcode).First(&table)
		if !ErrDb(r.Error) {
			err = r.Error
			log.Error(ctx, err)
		}
	}
	if err != nil {
		return
	}
	helper.AdjustStructToStruct(table, &p)
	if table.StartDate != nil {
		p.StartDate = table.StartDate.Format(helper.FormatYmdHis)
	}

	if table.LastSeenDate != nil {
		p.LastSeenDate = table.LastSeenDate.Format(helper.FormatYmdHis)
	}
	return
}

func (repo *Database) FindAll(ctx context.Context) (data []model.Shorten, err error) {
	qx := Connection()
	defer Close(qx)

	var table []dao.Shorten
	qx.Find(&table)
	if qx.Error != nil {
		err = qx.Error
		log.Error(ctx, err)
	}

	if len(table) == 0 {
		data = make([]model.Shorten, 0)
	} else {
		for i, v := range table {
			if v.StartDate != nil {
				table[i].StartDateFormatted = v.StartDate.Format(helper.FormatYmdHis)
			}

			if v.LastSeenDate != nil {
				table[i].LastSeenDateFormatted = v.LastSeenDate.Format(helper.FormatYmdHis)
			}

		}
		helper.AdjustStructToStruct(table, &data)
	}
	return
}

func (repo *Database) Create(ctx context.Context, req httpModel.RequestCreate) (shorten dao.Shorten, err error) {
	qx := Connection()
	defer Close(qx)

	tx := qx.Begin()
	shorten.Shortcode = req.Shortcode
	shorten.URL = req.URL
	now := helper.Now()
	shorten.StartDate = &now
	tx.Create(&shorten)
	if tx.Error != nil {
		err = tx.Error
		log.Error(ctx, err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func (repo *Database) Execute(ctx context.Context, p model.Shorten) (err error) {
	qx := Connection()
	defer Close(qx)

	var shorten dao.Shorten
	helper.AdjustStructToStruct(p, &shorten)
	shorten.StartDate, err = helper.FormatToDateTime(p.StartDate)
	if err != nil {
		return
	}
	tx := qx.Begin()
	now := helper.Now()
	shorten.LastSeenDate = &now
	shorten.RedirectCount++
	tx.Save(&shorten)
	if tx.Error != nil {
		err = tx.Error
		log.Error(ctx, err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
