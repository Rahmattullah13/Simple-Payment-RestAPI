package repository

import (
	"simple-payment/model"
	"simple-payment/util"

	"github.com/jmoiron/sqlx"
)

type MerchantRepository interface {
	Insert(merchant *model.Merchant) error
	Merchants() (*[]model.Merchant, error)
	MerchantById(id int) (*model.Merchant, error)
	Delete(id int) error
}

type merchantRepository struct {
	db *sqlx.DB
}

func (mr *merchantRepository) Insert(merchant *model.Merchant) error {
	createdMerchant := new(model.Merchant)
	row := mr.db.QueryRowx(util.CREATE_MERCHANT, merchant.UserId, merchant.Name, merchant.CreatedAt)
	if err := row.StructScan(createdMerchant); err != nil {
		return err
	}

	return nil
}

func (mr *merchantRepository) Merchants() (*[]model.Merchant, error) {
	merchants := new([]model.Merchant)

	if err := mr.db.Select(merchants, util.ALL_MERCHANT); err != nil {
		return nil, err
	}

	return merchants, nil
}

func (mr *merchantRepository) MerchantById(id int) (*model.Merchant, error) {
	merchant := new(model.Merchant)

	if err := mr.db.Get(merchant, util.READ_MERCHANT, id); err != nil {
		return nil, err
	}

	return merchant, nil
}

func (mr *merchantRepository) Delete(id int) error {
	if _, err := mr.db.Exec(util.DELETE_MERCHANT, id); err != nil {
		return err
	}

	return nil
}

func NewMerchantRerpository(db *sqlx.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
