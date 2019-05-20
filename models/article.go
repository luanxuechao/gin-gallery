package models

type Article struct {
	Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	Note       string `json:"note"`
	ImageUri   string `json:"image_uri"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	ActiveFlag int    `json:"active_flag"`
}

// rest
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Debug().Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticlesTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func AddArticle(title string, activeFlag int, createBy string, imageUri string) bool {
	db.Create(&Article{
		Title:      title,
		ActiveFlag: activeFlag,
		CreatedBy:  createBy,
		ImageUri:   imageUri,
	})
	return true
}
