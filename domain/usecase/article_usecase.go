package usecase

import "time"

// ArticleRepository ArticleRepository
type ArticleRepository interface {
	Store()
	Fetch()
}

// ArticleUseCase ArticleUseCase
type ArticleUseCase interface {
	Store()
	Fetch()
}

// ArticleUseCaseImpl ArticleUseCaseImpl
type ArticleUseCaseImpl struct {
	repository  ArticleRepository
	contextTime time.Duration
}

// NewArticleUseCase ArticleUseCaseを作成する
func NewArticleUseCase(repository ArticleRepository, timeout time.Duration) ArticleUseCase {
	return &ArticleUseCaseImpl{repository: repository, contextTime: timeout}
}

// Store 記事を作成する
func (u *ArticleUseCaseImpl) Store() {}

// Fetch 記事を取得する
func (u *ArticleUseCaseImpl) Fetch() {}
