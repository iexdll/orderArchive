package postgreSQL

import (
	"log"
	"orderArchive/models"
	"time"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Save(o *models.Order) error {

	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id=$1", o.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM \"orderRows\" WHERE \"order\"=$1", o.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO orders ("+
			"id,"+
			"number,"+
			"date,"+
			"\"shippingDate\","+
			"\"deliveryDate\","+
			"customer,"+
			"sum,"+
			"\"tradePoint\","+
			"\"paymentType\","+
			"\"deliveryType\","+
			"\"transportType\","+
			"status,"+
			"comment,"+
			"\"warehouseShipping\""+
			") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		o.ID,
		o.Number,
		o.Date,
		o.ShippingDate,
		o.DeliveryDate,
		o.Customer,
		o.Sum,
		o.TradePoint,
		o.PaymentType,
		o.DeliveryType,
		o.TransportType,
		o.Status,
		o.Comment,
		o.WarehouseShipping)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO "orderRows" (` +
		`"order",` +
		`"rowNumber",` +
		`"rowID",` +
		"goods," +
		`"group",` +
		"price," +
		"count," +
		`"countCancel",` +
		"comment" +
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, item := range o.GoodsList {
		_, err = stmt.Exec(o.ID, item.RowNumber, item.RowID, item.Goods, item.Group, item.Price, item.Count, item.CountCancel, item.Comment)
		if err != nil {
			log.Fatal(err)
		}
	}

	return tx.Commit()

}

func (r *OrderRepository) Get(id string) (*models.Order, error) {
	o := &models.Order{}
	return o, nil
}

func (r *OrderRepository) FindByDate(dateSart time.Time, dateEnd time.Time) ([]string, error) {
	var orders []string
	return orders, nil
}
