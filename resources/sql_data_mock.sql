
CREATE DATABASE stock_opname;

create table stock_opname.products (	
    product_code  	VARCHAR(40) NOT NULL,
	product_name 	VARCHAR(255) NOT NULL,
	uom   			VARCHAR(255) NOT NULL,
	price			INT,
	sale_price		INT,
	created_at		DATE,	
    PRIMARY KEY (product_code)
);

create table stock_opname.reports (	
    _id             INT NOT NULL AUTO_INCREMENT,
    product_code  	VARCHAR(40),
	qty				INT,
	amount			INT,
    kind            VARCHAR(100),
	created_at		DATE,	
    PRIMARY KEY (_id),
    FOREIGN KEY (product_code) REFERENCES products(product_code),
    CONSTRAINT UC_Reports UNIQUE (product_code,kind,created_at)
);


-- mock data from https://i0.wp.com/zahiraccounting.com/id/blog/wp-content/uploads/2017/07/contoh-laporan-stock-opname-excel-10.png

insert into stock_opname.products values ("LP101", "ASUS X401", "UNIT", 3200000, 3500000, "2021-06-04");
insert into stock_opname.products values ("LP102", "TOSIBA P25W", "UNIT", 2250000, 2500000, "2021-06-04");
insert into stock_opname.products values ("LP103", "DELL VOSTRO 3468", "UNIT", 4000000, 4300000, "2021-06-04");
insert into stock_opname.products values ("LP104", "LENOVO IDEAPAD 310", "UNIT", 3000000, 3250000, "2021-06-04");
insert into stock_opname.products values ("LP105", "HP14-AN004AU", "UNIT", 2300000, 2500000, "2021-06-04");

-- kind_initial_inventory, 
insert into stock_opname.reports value (1,"LP101",22,70400000,"kind_initial_inventory", "2021-06-04");
insert into stock_opname.reports value (2,"LP102",53,119250000,"kind_initial_inventory", "2021-06-04");
insert into stock_opname.reports value (3,"LP103",25,100000000,"kind_initial_inventory", "2021-06-04");
insert into stock_opname.reports value (4,"LP104",47,141000000,"kind_initial_inventory", "2021-06-04");
insert into stock_opname.reports value (5,"LP105",24,55200000,"kind_initial_inventory", "2021-06-04");
-- kind_sale, 
insert into stock_opname.reports value (6,"LP101",22,70000000,"kind_sale", "2021-06-04");
insert into stock_opname.reports value (7,"LP102",34,85000000,"kind_sale", "2021-06-04");
insert into stock_opname.reports value (8,"LP103",23,98900000,"kind_sale", "2021-06-04");
insert into stock_opname.reports value (9,"LP104",12,39000000,"kind_sale", "2021-06-04");
insert into stock_opname.reports value (10,"LP105",21,52500000,"kind_sale", "2021-06-04");

-- kind_incoming_product, 
insert into stock_opname.reports value (11,"LP101",50,160000000,"kind_incoming_product", "2021-06-04");
insert into stock_opname.reports value (12,"LP102",50,112500000,"kind_incoming_product", "2021-06-04");
insert into stock_opname.reports value (13,"LP103",50,200000000,"kind_incoming_product", "2021-06-04");
insert into stock_opname.reports value (14,"LP104",50,150000000,"kind_incoming_product", "2021-06-04");
insert into stock_opname.reports value (15,"LP105",50,115000000,"kind_incoming_product", "2021-06-04");

--kind_final_inventory, 
insert into stock_opname.reports value (16,"LP101",52,166400000,"kind_final_inventory", "2021-06-04");
insert into stock_opname.reports value (17,"LP102",69,155250000,"kind_final_inventory", "2021-06-04");
insert into stock_opname.reports value (18,"LP103",52,208000000,"kind_final_inventory", "2021-06-04");
insert into stock_opname.reports value (19,"LP104",85,255000000,"kind_final_inventory", "2021-06-04");
insert into stock_opname.reports value (20,"LP105",53,121900000,"kind_final_inventory", "2021-06-04");

--kind_warehouse_inventory
insert into stock_opname.reports value (21,"LP101",51,163200000,"kind_warehouse_inventory", "2021-06-04");
insert into stock_opname.reports value (22,"LP102",67,150750000,"kind_warehouse_inventory", "2021-06-04");
insert into stock_opname.reports value (23,"LP103",55,220000000,"kind_warehouse_inventory", "2021-06-04");
insert into stock_opname.reports value (24,"LP104",85,255000000,"kind_warehouse_inventory", "2021-06-04");
insert into stock_opname.reports value (25,"LP105",53,121900000,"kind_warehouse_inventory", "2021-06-04");