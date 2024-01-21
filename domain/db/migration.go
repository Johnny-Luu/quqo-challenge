package db

import "quqo_challenge/domain/entity"

func (p *Persistence) AutoMigrate() {
	_ = p.AppDb.AutoMigrate(&entity.Product{})
}
