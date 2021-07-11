package repository

import (
	"database/sql"
)


type ResponseProduct struct {
	ProductCode  	string
	ProductName 	string
	UOM   			string
	Price			int
	SalePrice		int	
	CreatedAt		string
}


type ResponseReport struct {
	Id 				int
	ProductCode  	string		
	Qty				int
	Amount			int
	Kind 			string	
	CreatedAt 		string
}



type opnameRepository struct {	
	connMySql *sql.DB
}

func NewOpnameRepository(connMySql *sql.DB) OpnameRepositoryInterface {
	return &opnameRepository{connMySql}
}

type OpnameRepositoryInterface interface {
	
	GetProducts() ([]ResponseProduct, error)
	GetProductReport(productCode string) ([]ResponseReport, error)

}



