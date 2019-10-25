package article_service

import (
	"encoding/json"
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/logging"
	"gin-blog/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			fmt.Println(data)
			fmt.Println(fmt.Sprintf("%T", data))
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(article)
	fmt.Println(fmt.Sprintf("%T", article))
	gredis.Set(key, article, 3600)
	return article, nil
}
