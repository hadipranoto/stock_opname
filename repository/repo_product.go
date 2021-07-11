package repository





func (repo *opnameRepository) GetProducts() ([]ResponseProduct, error) {
	var (
		result []ResponseProduct
	)
	
	db := repo.connMySql	
	
	rows, err := db.Query("select product_code, product_name, uom, price, sale_price, created_at from products")

	if err != nil {		
		return result, err
	}

	defer rows.Close()
	
	for rows.Next() {
		each 	:= ResponseProduct{}
		err 	:= rows.Scan(&each.ProductCode, &each.ProductName, &each.UOM, &each.Price, &each.SalePrice, &each.CreatedAt)

		if err != nil {			
			return result, err
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {		
		return result, err
	}	

	return result, nil
}


func (repo *opnameRepository) GetProductReport(productCode string) ([]ResponseReport, error) {
	var (
		result []ResponseReport
	)
	
	db := repo.connMySql	
	
	rows, err := db.Query("select _id, product_code, qty, amount, kind, created_at from reports where product_code = ?", productCode)

	if err != nil {		
		return result, err
	}

	defer rows.Close()
	
	for rows.Next() {
		each 	:= ResponseReport{}
		err 	:= rows.Scan(&each.Id, &each.ProductCode, &each.Qty, &each.Amount, &each.Kind, &each.CreatedAt)

		if err != nil {			
			return result, err
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {		
		return result, err
	}	

	return result, nil
		
}


