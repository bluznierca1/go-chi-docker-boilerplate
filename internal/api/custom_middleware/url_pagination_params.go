package custommiddleware

import (
	"context"
	"fmt"
	"myapp/internal/api/apiresponse"
	"myapp/internal/api/models"
	"myapp/internal/apperrors"
	"net/http"
	"strconv"
)

type CtxKeyUrlPaginationParams string

const PaginationUrlParams CtxKeyUrlPaginationParams = "pagination"

var (
	pageParam         = "page"
	itemsPerPageParam = "itemsPerPage"
)

func ExtractPaginationUrlParams() func(next http.Handler) http.Handler {
	var validationErrors []apperrors.Error
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {
			var err error

			// default to 1
			page := 1
			if queryPage := r.URL.Query().Get(pageParam); queryPage != "" {
				page, err = strconv.Atoi(queryPage)
				if err != nil || page < 1 {
					validationErrors = append(validationErrors, apperrors.GetErrorDetails(apperrors.IntegerGreaterThanZero, &pageParam, nil, nil))
					apiresponse.ErrorResponse(validationErrors, http.StatusBadRequest, w)
					return
				}
			}

			// default to 30
			itemsPerPage := 30
			if perPage := r.URL.Query().Get(itemsPerPageParam); perPage != "" {
				itemsPerPage, err = strconv.Atoi(perPage)
				if err != nil || itemsPerPage < 1 {
					validationErrors = append(validationErrors, apperrors.GetErrorDetails(apperrors.IntegerGreaterThanZero, &itemsPerPageParam, nil, nil))
					apiresponse.ErrorResponse(validationErrors, http.StatusBadRequest, w)
					return
				}
			}

			ctx := context.WithValue(r.Context(), PaginationUrlParams, models.PaginationUrlParams{
				Page:         page,
				ItemsPerPage: itemsPerPage,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func GetPaginationUrlParams(ctx context.Context) (models.PaginationUrlParams, error) {
	var params models.PaginationUrlParams
	var ok bool
	params, ok = ctx.Value(PaginationUrlParams).(models.PaginationUrlParams)
	if !ok {
		return params, fmt.Errorf("could not extract pagination params from given context")
	}

	return params, nil
}
