package serializer

import "school_manager/model"

type Article struct {
	ID        uint   `json:"ID"`
	Tile      string `json:"tile"`
	Content   string `json:"content"`
	Publisher uint   `json:"publisher"`
}

func BuildArticle(article model.Article) Article {
	return Article{
		ID:        article.ID,
		Tile:      article.Tile,
		Content:   article.Content,
		Publisher: article.Publisher,
	}
}

func BuildArticles(ArticleIn []model.Article) (ArticleOut []Article) {
	for _, forArticle := range ArticleIn {
		article := BuildArticle(forArticle)
		ArticleOut = append(ArticleOut, article)
	}
	return ArticleOut
}
