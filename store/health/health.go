package health

import (
	"context"
	"net/http"

	"github.com/PatrickChagastavares/church_backend/model"
	"github.com/PatrickChagastavares/church_backend/utils/do"
	"github.com/PatrickChagastavares/church_backend/utils/logger"

	"github.com/jmoiron/sqlx"
)

// Store interface para implementação do health
type Store interface {
	Ping(ctx context.Context) do.ChanResult
	Check(ctx context.Context) do.ChanResult
}

// NewStore cria uma nova instancia do repositorio de health
func NewStore(reader *sqlx.DB) Store {
	return &storeImpl{reader}
}

type storeImpl struct {
	reader *sqlx.DB
}

// Ping checa se o banco está online
func (r *storeImpl) Ping(ctx context.Context) do.ChanResult {
	return do.Do(func(res *do.Result) {
		err := r.reader.PingContext(ctx)
		if err != nil {
			logger.ErrorContext(ctx, "store.health.ping", err.Error())
			res.Error = model.NewError(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		res.Data = &model.Health{DatabaseStatus: "OK"}
	})
}

// Check checa se o banco está com status OK
func (r *storeImpl) Check(ctx context.Context) do.ChanResult {
	return do.Do(func(res *do.Result) {
		data := new(model.Health)

		err := r.reader.GetContext(ctx, data, `SELECT 'DB OK' AS database_status`)
		if err != nil {
			logger.ErrorContext(ctx, "store.health.check", err.Error())
			res.Error = model.NewError(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		res.Data = data
	})
}