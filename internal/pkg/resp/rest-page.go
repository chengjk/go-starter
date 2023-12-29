package resp

type RestPage struct {
	TotalPage   int64       `json:"total_page,omitempty"`
	TotalRecord int64       `json:"total_record,omitempty"`
	PageNum     int         `json:"page_num,omitempty"`
	PageSize    int         `json:"page_size,omitempty"`
	Records     interface{} `json:"records,omitempty"`
}

func Page(pageNum, pageSize int, totalRecord int64, records interface{}) RestPage {
	page := RestPage{
		PageNum:     pageNum,
		PageSize:    pageSize,
		TotalPage:   totalRecord / int64(pageSize),
		Records:     records,
		TotalRecord: totalRecord,
	}
	if totalRecord%int64(pageSize) > 0 {
		page.TotalPage = page.TotalPage + 1
	}
	return page
}
