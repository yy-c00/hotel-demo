package model

//History interface implemented by all structures that manage completed sale logs
type History interface{
	//GetLogsByRange used to get logs for a range
	GetLogsByRange(uint, uint) ([]Sale, error)
	//GetSaleLogById used to get a sale by id
	GetSaleLogById(uint) (Sale, error)
	//SearchByRange used to search history by range
	SearchByRange(string, uint, uint) ([]Sale, error)
}