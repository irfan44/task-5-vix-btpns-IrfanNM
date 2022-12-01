package models

import (
	"html"
	"strings"

	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/types"
)

type Photo struct {
	ID       int          `gorm:"primary_key;auto_increment" json:"id"`
	Title    string       `gorm:"size:100;not null" json:"title"`
	Caption  string       `gorm:"size:255;not null" json:"caption"`
	PhotoUrl string       `gorm:"size:255;not null;" json:"photo_url"`
	UserId   string       `gorm:"not null" json:"user_id"`
	Author   types.Author `gorm:"author"`
}

func (p *Photo) Initialize() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Caption = html.EscapeString(strings.TrimSpace(p.Caption))
	p.PhotoUrl = html.EscapeString(strings.TrimSpace(p.PhotoUrl))
}
