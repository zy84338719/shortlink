package web

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"shortlink/conf"
	"shortlink/models"
	"shortlink/storage"
)

var s storage.Storage

func init() {
	redis := conf.C.Redis
	s = storage.NewRedis(redis.Addr, redis.Pwd, redis.DB)
}
func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]
	longUrl, err := s.Unshorten(shortUrl)

	if err != nil {
		sendErrorResponse(w, models.Response{Status: http.StatusInternalServerError, Result: models.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	}
}

func getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shortUrl")
	info, err := s.ShortlinkInfo(shortUrl)
	if err != nil {
		sendErrorResponse(w, models.Response{Status: http.StatusInternalServerError, Result: models.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		urlDetail := &storage.URLDetail{}
		_ = json.Unmarshal([]byte(info.(string)), urlDetail)
		sendNormalResponse(w, models.Result{Code: "0", ErrMsg: "OK", Data: urlDetail}, http.StatusOK)
	}
}

func createShortUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := &models.ShortenReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		sendErrorResponse(w, models.ErrorRequestBodyParseFailed)
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		sendErrorResponse(w, models.Response{Status: http.StatusInternalServerError, Result: models.Result{Code: "007", ErrMsg: err.Error()}})
		return
	}
	shorten, err := s.Shorten(req.Url, req.ExpirationInMinutes)
	if err != nil {
		sendErrorResponse(w, models.ErrorStorageError)
		return
	}
	sendNormalResponse(w, models.Result{Code: "0", ErrMsg: "OK", Data: models.ShortenResp{ShortUrl: shorten, LongUrl: req.Url}}, http.StatusOK)
}
