package services

import (
	"database/sql"
	"fmt"

	"github.com/hadipranoto/stock_opname/repository"
)



type serviceHandler struct {	
	repo repository.OpnameRepositoryInterface
	connMySql *sql.DB
}


func NewOpnameService(connMySql *sql.DB) OpnameServiceInterface {
	opnameRepository := repository.NewOpnameRepository(connMySql)
	return &serviceHandler{opnameRepository, connMySql}
}


type OpnameServiceInterface interface {

	SummaryOpname() (interface{}, error)
	
}

//I only want to state that service could access many repositories here
func (s *serviceHandler) SummaryOpname() (interface{}, error) {
	var (
		err					error
		dataResult 			[]map[string]interface{}		
		dataProduct 		[]repository.ResponseProduct
	)

	dataProduct, err = s.repo.GetProducts()

	if err != nil{
		return dataResult, err
	}

	for _, product := range dataProduct {
		
		report, err := s.repo.GetProductReport(product.ProductCode)

		if err != nil {
			fmt.Println(err)
		}

		dataResult = append(dataResult, map[string]interface{}{			
			"product_code":  	product.ProductCode,
			"product_name":		product.ProductName,
			"uom":				product.UOM,
			"price":			product.Price,
			"sale_price":		product.SalePrice,		
			"stock_inventory": 	report,
		})
		
	}	
	
	return dataResult,nil
}