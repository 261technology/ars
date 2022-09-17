package shorten

import (
	"context"

	"github.com/harisaginting/guin/common/log"
	"github.com/harisaginting/guin/common/utils/helper"
	database "github.com/harisaginting/guin/db"
	"github.com/harisaginting/guin/db/table"
)

type Repository struct{}

func (repo *Repository) Get(ctx context.Context, p *Shorten) (err error) {
	qx := database.Connection()
	defer database.Close(qx)

	var table table.Shorten

	if p.ID != 0 {
		table.ID = p.ID
		r := qx.First(&table)
		err = r.Error
		log.Error(ctx, err)
	} else {
		r := qx.Debug().Where("shortcode = ?", p.Shortcode).First(&table)
		if !database.ErrDb(r.Error) {
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

func (repo *Repository) FindAll(ctx context.Context) (data []Shorten, err error) {
	qx := database.Connection()
	defer database.Close(qx)

	var table []table.Shorten
	qx.Find(&table)
	if qx.Error != nil {
		err = qx.Error
		log.Error(ctx, err)
	}

	if len(table) == 0 {
		data = make([]Shorten, 0)
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

func (repo *Repository) Create(ctx context.Context, req RequestCreate) (shorten table.Shorten, err error) {
	qx := database.Connection()
	defer database.Close(qx)

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

func (repo *Repository) Execute(ctx context.Context, p Shorten) (err error) {
	qx := database.Connection()
	defer database.Close(qx)

	var shorten table.Shorten
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
