package reqres

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/julienschmidt/httprouter"
)

func DecodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entity.TestTable
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// func DecodeGetByParamRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	if data, ok := vars["param"]; !ok && data == "" {
// 		return nil, common.ErrDataNotFound
// 	}

// 	var request entity.TestTable
// 	{
// 		id, err := strconv.Atoi(vars["param"])
// 		if err != nil {
// 			return nil, err
// 		}

// 		request.ID = id
// 	}

// 	return request, nil
// }

func DecodeGetByPathRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	data := params.ByName("param")

	if data == "" {
		return nil, common.ErrDataNotFound
	}

	var request entity.TestTable
	{
		id, err := strconv.Atoi(data)
		if err != nil {
			return nil, err
		}

		request.ID = id
	}

	return request, nil
}
