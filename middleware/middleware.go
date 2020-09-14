package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"
)

// middleware rate limiter prevent for DDoS attack
// concurrent limit, limit all request for 5 request/minutes this is just example simple rate limiter by juju
func TokenBucketLimiter(bucket *ratelimit.Bucket, logger log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if bucket.TakeAvailable(1) == 0 {
				ErrorEncoder(ctx, common.ErrLimitExceed, w)
				ErrorHandlerFunc(logger)(ctx, common.ErrLimitExceed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// this middleware has function only to log request body
// when request doesn't has request body, this middleware not nessesary needed
func LogRequestBody(next httptransport.DecodeRequestFunc, logger log.Logger) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (decoded interface{}, err error) {
		defer func() {
			level.Info(logger).Log(
				"requestBody", fmt.Sprintf("%+v", decoded),
			)
		}()

		return next(ctx, r)
	}
}
