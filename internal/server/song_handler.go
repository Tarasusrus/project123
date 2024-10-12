package server

import (
	"BaseApi/internal/models"
	"BaseApi/tools"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) AddSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.NewSong
	h.logger.Debug("Начало декодирования")
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		h.logger.Error("ошибка при добавлении песни", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h.logger.Info("Успешное декодирование", newSong)

	err = h.MusicService.AdminService().AddSong(r.Context(), newSong)
	if err != nil {
		h.logger.Error("error adding song", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tools.JsonRespond(w, http.StatusOK, []byte("Песня добавлена успешно!"))
}

func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	var (
		updateData models.SongUpdate
		songId     int
	)

	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		h.logger.Error("ошибка при обновлении песни", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.MusicService.AdminService().UpdateSong(r.Context(), songId, updateData)
	if err != nil {
		h.logger.Error("error updating song", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tools.JsonRespond(w, http.StatusOK, []byte("Песня успешно обновлена!"))

}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	var songId int
	_, err := fmt.Sscanf(r.URL.Path, "/api/v1/songs/%d", &songId)
	if err != nil {
		h.logger.Error("invalid song ID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.MusicService.AdminService().DeleteSong(r.Context(), songId)
	if err != nil {
		h.logger.Error("error deleting song", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tools.JsonRespond(w, http.StatusOK, []byte("Песня успешно удалена!"))
}

func (h *Handler) GetSongText(w http.ResponseWriter, r *http.Request) {
	var songId int
	_, err := fmt.Sscanf(r.URL.Path, "/api/v1/songs/%d/text", &songId)
	if err != nil {
		h.logger.Error("invalid song ID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text, err := h.MusicService.UserService().GetSongText(r.Context(), songId, 1)
	if err != nil {
		h.logger.Error("error getting song text", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tools.JsonRespond(w, http.StatusOK, []byte(text))
}

func (h *Handler) GetLibrary(w http.ResponseWriter, r *http.Request) {
	filter := models.LibraryFilter{}
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	if group != "" {
		filter.Group = group
	}
	if song != "" {
		filter.Song = song
	}

	page := 1
	pageSize := 10
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	library, err := h.MusicService.UserService().GetLibrary(r.Context(), filter, page, pageSize)
	if err != nil {
		h.logger.Error("error getting library", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(library)
	tools.JsonRespond(w, http.StatusOK, response)

}
