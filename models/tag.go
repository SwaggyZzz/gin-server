package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	MysqlDb.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func getTagToatal(maps interface{}) (count int) {
	MysqlDb.Model(&Tag{}).Where(maps).Count(&count)
	return
}
