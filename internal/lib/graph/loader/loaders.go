package loader

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders / 各DataLoaderを取りまとめるstruct
type Loaders struct{}

// BatchFuncMap / 外部から渡されるBatchFunc型を束ねたもの
type BatchFuncMap map[string]*dataloader.BatchFunc

func NewLoaders() *Loaders {
	loaders := &Loaders{}
	return loaders
}

// Middleware LoadersをcontextにインジェクトするHTTPミドルウェア
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// GetLoaders / ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
