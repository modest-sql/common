type comms struct {
	Type int 	`json:"Type"`
	Data string	`json:"Data"`
}

type database struct {
	DB_Name string 	`json:"DB_Name"`
	Tables table	`json:"Tables"`
}

type table struct {
	Table_Name string		`json:"Table_Name"`
	ColumnNames []string	`json:"ColumnNames"`
	ColumnTypes []string	`json:"ColumnTypes"`
}