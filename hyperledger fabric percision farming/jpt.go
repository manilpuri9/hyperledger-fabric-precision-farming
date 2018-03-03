package main

type DisplaySpecType struct {
	SpecID string `json:"spec_id"`
	Sdv    string `json:"sdv"`
	Snv    string `json:"snv"`
}

type FilterSpecType struct {
	SpecID string `json:"spec_id"`
	Sdv    string `json:"sdv"`
	Snv    string `json:"snv"`
}
type ProductSpecType struct {
	DisplaySpec []DisplaySpecType `json:"display_spec"`
	FilterSpec  []FilterSpecType  `json:"filter_spec"`
}

type Product struct {
	ProductDisplayname string          `json:"product_displayname"`
	ProductPrice       string          `json:"product_price"`
	Popularity         string          `json:"popularity"`
	Barcode            string          `json:"barcode"`
	ExclusiveFlag      string          `json:"exclusive_flag"`
	ProductID          string          `json:"product_id"`
	ProductName        string          `json:"product_name"`
	BrandName          string          `json:"brand_name"`
	BrandID            string          `json:"brand_id"`
	ProductSpec        ProductSpecType `json:"product_spec"`
}

productData := Product{
	ProductDisplayname: "LG Stylus 2 Plus K535D (16 GB, Brown)",
	ProductPrice: "24000.00", Popularity: "0.00", Barcode: "", 
	ExclusiveFlag: "0",
	ProductID: "17698276", 
	ProductName: "Stylus 2 Plus K535D (Brown)", 
	BrandName: "LG", 
	BrandID: "1", 
	ProductSpec: ProductSpecType{DisplaySpec: []DisplaySpecType{
		DisplaySpecType{
			SpecID: "103", Sdv: "24000", Snv: "24000.0000"}, 
			DisplaySpecType{SpecID: "104", Sdv: "GSM", Snv: "0.0000"}
		}, 
		FilterSpec: []FilterSpecType{
			FilterSpecType{SpecID: "103", Sdv: "24000", Snv: "24000.0000"}, 
			FilterSpecType{SpecID: "105", Sdv: "Touch Screen", Snv: "0.0000"}
		}
	}
	}
