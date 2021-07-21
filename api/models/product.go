package models

import (
	"errors"
)

type ProductStatus uint8

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

type Product struct {
	Model
	Name       string        `gorm:"size:512;not null;unique" json:"name"`
	Price      float64       `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	Quantity   uint16        `gorm:"default:0;unsigned" json:"quantity"`
	Status     ProductStatus `gorm:"char(1);default:0" json:"status"`
	CategoryID uint64        `gorm:"not null" json:"category_id"`
 	Category   Category		 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CategoryID;references:ID;"`
}

/*
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(512) NOT NULL,
  `price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `quantity` smallint unsigned DEFAULT '0',
  `status` tinyint unsigned DEFAULT '0',
  `category_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `fk_categories_products` (`category_id`),
  CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
)
*/

var (
	ErrProductEmptyName = errors.New("product.name can't be empty")
)

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}

func (p *Product) CheckStatus() {
	p.Status = ProductStatus_Unavailable
	if p.Quantity > 0 {
		p.Status = ProductStatus_Available
	}
}