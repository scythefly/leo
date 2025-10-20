package data

import (
	"leo/api/snippet/v1"
	"leo/internal/data/model"

	"gorm.io/gorm"
)

func (d *Data) SnippetQuery(key, value string) *snippet.Response {
	var (
		resp = &snippet.Response{
			Code: 0,
		}
		snippets []model.Snippet
		pkey     = "%" + key + "%"
		pvalue   = "%" + value + "%"
	)

	if value != "" {
		db := d.db.Model(&model.Snippet{}).Where("key LIKE ? AND value LIKE ?", pkey, pvalue).Limit(10).Find(&snippets)
		if db.Error != nil {
			resp.Code = 1
			resp.Message = db.Error.Error()
			return resp
		}
	} else {
		ids := make(map[uint]struct{})
		var kss, vss []model.Snippet
		db := d.db.Model(&model.Snippet{}).Where("key LIKE ? ", pkey).Limit(10).Find(&kss)
		if db.Error != nil {
			resp.Code = 1
			resp.Message = db.Error.Error()
			return resp
		}
		snippets = append(snippets, kss...)
		for _, s := range kss {
			ids[s.ID] = struct{}{}
		}
		db = d.db.Model(&model.Snippet{}).Where("value LIKE ? ", pkey).Limit(10).Find(&vss)
		if db.Error != nil {
			resp.Code = 1
			resp.Message = db.Error.Error()
			return resp
		}
		for _, s := range vss {
			if _, ok := ids[s.ID]; !ok {
				snippets = append(snippets, s)
			}
		}
	}

	var pairs []*snippet.Response_Pair
	for _, s := range snippets {
		pairs = append(pairs, &snippet.Response_Pair{
			Key:   s.Key,
			Value: s.Value,
		})
	}

	resp.Pairs = pairs
	return resp
}

func (d *Data) SnippetPut(key, value string) *snippet.Response {
	var (
		resp = &snippet.Response{
			Code: 0,
		}
	)

	result := d.db.Create(&model.Snippet{
		Key:   key,
		Value: value,
	})

	if result.Error != nil {
		resp.Code = 1
		resp.Message = result.Error.Error()
	}

	return resp
}

func (d *Data) SnippetDelete(key, value string) *snippet.Response {
	var (
		resp = &snippet.Response{
			Code: 0,
		}
		db *gorm.DB
	)

	db = d.db.Where("key = ? AND value = ?", key, value).Delete(&model.Snippet{})

	if db.Error != nil {
		resp.Code = 1
		resp.Message = db.Error.Error()
	}

	return resp
}
