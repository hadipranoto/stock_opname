package middleware



type ReqProductType struct {
	ProductCode  	string 	`json:"product_code"`
	ProductName 	string 	`json:"product_name"`
	UOM   			string 	`json:"uom"`
	Price			int		`json:"price"`
	SalePrice		int		`json:"sale_price"`	
}

//kind : Persediaan Awal / Penjualan / Barang Masuk /  Persediaan Akhir / Persediaan Gudang 
type ReportPayloadType struct {
	ProductCode  	string  `json:"product_code"`		
	Qty				int		`json:"qty"`
	Kind 			string 	// kind_initial_inventory, kind_sale, kind_incoming_product, kind_final_inventory, kind_warehouse_inventory
}

var (
	ReqProduct 					ReqProductType
	ReqReport				 	ReportPayloadType	
)