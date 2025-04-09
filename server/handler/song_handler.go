package handler

import (
	"context"
	song "grpc-song-manager/proto"
	"grpc-song-manager/server/model"
	"grpc-song-manager/server/repository"
)

type SongService struct {
	song.UnimplementedSongServiceServer
	Repo *repository.SongRepository
}

func (s *SongService) CreateSong(ctx context.Context, in *song.SongInput) (*song.Song, error) {
	song := &model.Song_Model{
		Title:  in.Title,
		Artist: in.Artist,
		Album:  in.Album,
		Genre:  in.Genre,
	}
	res, err := s.Repo.Create(ctx, song)
	if err != nil {
		return nil, err
	}
	return toProto(res), nil
}

func (s *SongService) GetSong(ctx context.Context, in *song.SongRequest) (*song.Song, error) {
	song, err := s.Repo.GetByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return toProto(song), nil
}

func (s *SongService) ListSongs(ctx context.Context, _ *song.Empty) (*song.SongList, error) {
	songs, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}
	var res []*song.Song
	for _, song := range songs {
		res = append(res, toProto(song))
	}
	return &song.SongList{Songs: res}, nil
}

func (s *SongService) UpdateSong(ctx context.Context, in *song.Song) (*song.Song, error) {
	song := &model.Song_Model{
		ID:     in.Id,
		Title:  in.Title,
		Artist: in.Artist,
		Album:  in.Album,
		Genre:  in.Genre,
	}
	updated, err := s.Repo.Update(ctx, song)
	if err != nil {
		return nil, err
	}
	return toProto(updated), nil
}

func (s *SongService) DeleteSong(ctx context.Context, in *song.SongRequest) (*song.Empty, error) {
	err := s.Repo.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &song.Empty{}, nil
}

func toProto(songModel *model.Song_Model) *song.Song {
	return &song.Song{
		Id:     songModel.ID,
		Title:  songModel.Title,
		Artist: songModel.Artist,
		Album:  songModel.Album,
		Genre:  songModel.Genre,
	}
}
