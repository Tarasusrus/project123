package database

import (
	"BaseApi/internal/models"
	"context"
	"errors"
)

func (d *Database) AddSong(ctx context.Context, songData models.NewSong) error {
	if err := d.client.WithContext(ctx).Create(&songData).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateSong(ctx context.Context, songID int, songData models.SongUpdate) error {
	if err := d.client.WithContext(ctx).
		Model(&models.Song{}).
		Where("id = ?", songID).
		Updates(songData).
		Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteSong(ctx context.Context, songID int) error {
	if err := d.client.WithContext(ctx).
		Where("id = ?", songID).
		Delete(&models.Song{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetLibrary(ctx context.Context, filter models.LibraryFilter, page, pageSize int) ([]models.Song, error) {
	var songs []models.Song
	if err := d.client.WithContext(ctx).
		Where("group = ? AND song = ?", filter.Group, filter.Song).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&songs).
		Error; err != nil {
		return nil, err
	}
	return songs, nil
}
func (d *Database) GetSongText(ctx context.Context, songID int, verse int) (string, error) {
	var song models.Song
	if err := d.client.WithContext(ctx).
		Where("id = ?", songID).
		First(&song).
		Error; err != nil {
		return "", err
	}
	switch verse {
	case 1:
		return song.Text, nil
	case 2:
		return song.Text, nil
	default:
		return "", errors.New("verse not found")
	}
}
