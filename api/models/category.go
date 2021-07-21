package models

import "errors"

type Category struct {
	Model
	Description string     `gorm: "size:256;not null; unique" json:"description"`
	//Products    []*Product `gorm:"foreignKey:CategoryID" json:"products"`
}

/*
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `description` longtext,
  PRIMARY KEY (`id`)
)
*/

var (
	ErrCategoryEmptyDescription = errors.New("category.description can't be empty")
)

func (c *Category) Validate() error {
	if c.Description == "" {
		return ErrCategoryEmptyDescription
	}

	return nil
}
