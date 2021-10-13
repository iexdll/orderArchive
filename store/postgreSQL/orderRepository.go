package postgreSQL

import (
	"database/sql"
	"log"
	"orderArchive/models"
	"strconv"
	"strings"
	"time"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Save(o *models.Order) error {

	tx, err := r.store.db.Begin()
	if err != nil {
		log.Println(1)
		return err
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id=$1", o.ID)
	if err != nil {
		log.Println(2)
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM \"orderRows\" WHERE \"order\"=$1", o.ID)
	if err != nil {
		log.Println(3)
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
		log.Println(4)
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
		log.Println(5)
		_ = tx.Rollback()
		return err
	}

	for _, item := range o.GoodsList {
		_, err = stmt.Exec(o.ID, item.RowNumber, item.RowID, item.Goods, item.Group, item.Price, item.Count, item.CountCancel, item.Comment)
		if err != nil {
			log.Println(6)
			log.Fatal(err)
		}
	}

	return tx.Commit()

}

func (r *OrderRepository) Get(id string) (*models.Order, error) {

	order := &models.Order{}

	row := r.store.db.QueryRow("SELECT "+
		`id,`+
		`number,`+
		`date,`+
		`"shippingDate",`+
		`"deliveryDate",`+
		`customer,`+
		`sum,`+
		`"tradePoint",`+
		`"paymentType",`+
		`"deliveryType",`+
		`"transportType",`+
		`status,`+
		`comment,`+
		`"warehouseShipping"`+
		` FROM orders WHERE id = $1`, id)

	err := row.Scan(
		&order.ID,
		&order.Number,
		&order.Date,
		&order.ShippingDate,
		&order.DeliveryDate,
		&order.Customer,
		&order.Sum,
		&order.TradePoint,
		&order.PaymentType,
		&order.DeliveryType,
		&order.TransportType,
		&order.Status,
		&order.Comment,
		&order.WarehouseShipping)

	if err == sql.ErrNoRows {
		return order, nil
	}

	if err != nil {
		return nil, err
	}

	rows, err := r.store.db.Query("SELECT "+
		`"rowNumber",`+
		`"rowID",`+
		`goods,`+
		`"group",`+
		`price,`+
		`count,`+
		`"countCancel",`+
		`comment`+
		` FROM "orderRows" WHERE "order" = $1`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		orderRow := models.OrderRow{}

		err = rows.Scan(
			&orderRow.RowNumber,
			&orderRow.RowID,
			&orderRow.Goods,
			&orderRow.Group,
			&orderRow.Price,
			&orderRow.Count,
			&orderRow.CountCancel,
			&orderRow.Comment)

		if err != nil {
			return nil, err
		}

		order.GoodsList = append(order.GoodsList, orderRow)
	}

	order.AfterGet()
	return order, nil
}

func (r *OrderRepository) FindIDByDate(dateSart time.Time, dateEnd time.Time) ([]string, error) {
	var orders []string
	return orders, nil
}

func (r *OrderRepository) FindIDByCustomer(customer string, tradePoint string, limit int32, skip int32) ([]string, error) {

	var param []interface{}

	query := `SELECT "id" FROM orders WHERE customer = $1`
	param = append(param, customer)

	if models.EmptyRef != tradePoint {
		query = query + ` AND "tradePoint" = $` + strconv.Itoa(len(param)+1)
		param = append(param, tradePoint)
	}

	query = query + ` ORDER BY "date" DESC`

	if limit > -1 {
		query = query + ` LIMIT $` + strconv.Itoa(len(param)+1)
		param = append(param, limit)
	}

	if skip > -1 {
		query = query + ` OFFSET $` + strconv.Itoa(len(param)+1)
		param = append(param, skip)
	}

	rows, err := r.store.db.Query(query, param...)

	if err != nil {
		return nil, err
	}

	var orders []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, id)
	}

	return orders, nil
}

func (r *OrderRepository) FindIDByGoods(customer string, tradePoint string, goodsID []string) ([]string, error) {

	var param []interface{}

	query := `SELECT DISTINCT 
                o.id 
			  FROM orders AS o 
			  INNER JOIN "orderRows" AS r
			      ON r.order = o.id 
			  WHERE o.customer = $1`

	param = append(param, customer)

	if models.EmptyRef != tradePoint {
		query = query + ` AND o."tradePoint" = $` + strconv.Itoa(len(param)+1)
		param = append(param, tradePoint)
	}

	var p []string
	for _, goods := range goodsID {
		p = append(p, "$"+strconv.Itoa(len(param)+1))
		param = append(param, goods)
	}
	query = query + ` AND r.goods IN (` + strings.Join(p, ",") + `)`

	rows, err := r.store.db.Query(query, param...)

	if err != nil {
		return nil, err
	}

	var orders []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, id)
	}

	return orders, nil
}
