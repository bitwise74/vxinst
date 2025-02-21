/*
VxInstagram - Blazing fast embedder for instagram posts
Copyright (C) 2025 Bash06

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	Id           string `gorm:"primaryKey;unique;not null;index"`
	Title        string
	PostURL      string `gorm:"not null"`
	ThumbnailURL string
	VideoURL     string `gorm:"not null"`
	ExpiresAt    int64  `gorm:"not null"`
}

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data.db"))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Post{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
