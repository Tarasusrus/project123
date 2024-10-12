package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JsonRespond(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func JsonRespondCreated(w http.ResponseWriter, id string) {
	JsonRespond(w, http.StatusCreated, []byte(fmt.Sprintf(`{"id":"%s"}`, id)))
}

func JsonRespondCreatedInt(w http.ResponseWriter, id int) {
	JsonRespond(w, http.StatusCreated, []byte(fmt.Sprintf(`{"id":"%d"}`, id)))
}

func RespondBadRequestError(w http.ResponseWriter, err error) {
	errMsg := err.Error()
	response, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Ошибка чтения ID",
	})
	JsonRespond(w, http.StatusBadRequest, response)
}

func RespondNotFoundError(w http.ResponseWriter, err error) {
	errMsg := err.Error()
	response, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Запрос не найден",
	})
	JsonRespond(w, http.StatusNotFound, response)
}

func RespondConflictError(w http.ResponseWriter, err error) {
	errMsg := err.Error()
	response, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Конфликтующие данные",
	})
	JsonRespond(w, http.StatusConflict, response)
}

func RespondInvalidFormatError(w http.ResponseWriter, err error) {
	errMsg := err.Error()
	response, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Неверный формат данных",
	})
	JsonRespond(w, http.StatusBadRequest, response)
}

func RespondValidationError(w http.ResponseWriter, err error) {
	response, _ := json.Marshal(map[string]interface{}{
		"status": http.StatusUnprocessableEntity,
		"data":   "Ошибка валидации",
		"errors": err,
	})
	JsonRespond(w, http.StatusUnprocessableEntity, response)
}

func RespondForbiddenError(w http.ResponseWriter, err error) {
	errMsg := "Forbidden"
	if err != nil {
		errMsg = err.Error()
	}

	responce, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Доступ запрещен",
	})
	JsonRespond(w, http.StatusForbidden, responce)
}

func HtmlRespond(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func RespondUnprocessableContent(w http.ResponseWriter, err error) {
	errMsg := err.Error()
	response, _ := json.Marshal(map[string]string{
		"error":   errMsg,
		"message": "Неверные данные",
	})
	JsonRespond(w, http.StatusUnprocessableEntity, response)
}
