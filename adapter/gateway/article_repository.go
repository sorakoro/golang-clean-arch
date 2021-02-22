package gateway

import (
	"github.com/jmoiron/sqlx"
	"github.com/sorakoro/golang-clean-arch/domain/usecase"
)

// ArticleRepositoryImpl ArticleRepositoryImpl
type ArticleRepositoryImpl struct {
	db *sqlx.DB
}

// NewArticleRepository ArticleRepositoryを作成する
func NewArticleRepository(db *sqlx.DB) usecase.ArticleRepository {
	return &ArticleRepositoryImpl{db: db}
}

// Store 記事を作成する
func (r *ArticleRepositoryImpl) Store() {}

// Fetch 記事を取得する
func (r *ArticleRepositoryImpl) Fetch() {}
