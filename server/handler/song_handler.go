package handler

import (
	"context"
	"io"
	"fmt"
	"sync"

	pb "grpc-song-manager/proto"
	"grpc-song-manager/server/model"
	"grpc-song-manager/server/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"strings"
)

type SongService struct {
	pb.UnimplementedSongServiceServer
	Repo *repository.SongRepository
	mu        sync.Mutex
  	streams   map[int]pb.SongService_CollaboratePlaylistServer
  	nextID    int
  	playlist  []*pb.Song
}

var moodToGenres = map[string][]string{
	"chill":     {"Ambient", "Lo-fi", "Acoustic", "Jazz"},
	"energetic": {"Rock", "Pop", "EDM", "Dance"},
	"sad":       {"Blues", "Soul"},
	"happy":     {"Pop", "Funk", "Disco"},
}

func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{
	  Repo:     repo,
	  streams:  make(map[int]pb.SongService_CollaboratePlaylistServer),
	  playlist: []*pb.Song{},
	}
  }

func (s *SongService) CreateSong(ctx context.Context, in *pb.SongInput) (*pb.Song, error) {
	if in.Title == "" || in.Artist == "" || in.Album == "" || in.Genre == "" {
		return nil, status.Error(codes.InvalidArgument, "all song fields must be filled")
	}

	song := &model.Song_Model{
		Title:  in.Title,
		Artist: in.Artist,
		Album:  in.Album,
		Genre:  in.Genre,
	}
	res, err := s.Repo.Create(ctx, song)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create song: %v", err)
	}
	return toProto(res), nil
}


func (s *SongService) GetSong(ctx context.Context, in *pb.SongRequest) (*pb.Song, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "song ID must be provided")
	}
	song, err := s.Repo.GetByID(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "song not found: %v", err)
	}
	return toProto(song), nil
}


func (s *SongService) ListSongs(ctx context.Context, _ *pb.Empty) (*pb.SongList, error) {
	songs, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}
	var res []*pb.Song
	for _, song := range songs {
		res = append(res, toProto(song))
	}
	return &pb.SongList{Songs: res}, nil
}

func (s *SongService) UpdateSong(ctx context.Context, in *pb.Song) (*pb.Song, error) {
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

func (s *SongService) DeleteSong(ctx context.Context, in *pb.SongRequest) (*pb.Empty, error) {
	err := s.Repo.Delete(ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func toProto(songModel *model.Song_Model) *pb.Song {
	return &pb.Song{
		Id:     songModel.ID,
		Title:  songModel.Title,
		Artist: songModel.Artist,
		Album:  songModel.Album,
		Genre:  songModel.Genre,
	}
}

// SongChat: untuk setiap SongInput yang diterima, server
// akan menyimpan (Create) lalu langsung mengirim balik Song yang baru dibuat.
func (s *SongService) SongChat(stream pb.SongService_SongChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "error receiving from stream: %v", err)
		}

		// Validasi input
		if in.Title == "" || in.Artist == "" {
			return status.Error(codes.InvalidArgument, "title and artist are required")
		}

		song := &model.Song_Model{
			Title:  in.Title,
			Artist: in.Artist,
			Album:  in.Album,
			Genre:  in.Genre,
		}
		created, err := s.Repo.Create(stream.Context(), song)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to save song: %v", err)
		}

		if err := stream.Send(toProto(created)); err != nil {
			return status.Errorf(codes.Unknown, "failed to send song to client: %v", err)
		}
	}
}

// AddSongToPlaylist: untuk menambahkan lagu ke playlist
// server. Playlist ini adalah list lagu yang akan dikirim ke semua client yang terhubung.
func (s *SongService) AddSongToPlaylist(ctx context.Context, songID string) (*pb.Song, error) {
	// Mengonversi ID lagu dari string ke ObjectID MongoDB
	oid, err := primitive.ObjectIDFromHex(songID)
	if err != nil {
		return nil, fmt.Errorf("invalid song ID: %v", err)
	}

	// Ambil lagu berdasarkan ID dari repository
	song, err := s.Repo.GetByID(ctx, oid.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not find song: %v", err)
	}

	// Mengubah data song ke format yang akan digunakan dalam playlist
	newSong := &pb.Song{
		Id:     song.ID,
		Title:  song.Title,
		Artist: song.Artist,
		Album:  song.Album,
		Genre:  song.Genre,
	}

	// Menambahkan lagu ke playlist
	s.playlist = append(s.playlist, newSong)

	// Kembalikan response
	return newSong, nil
}

// CollaboratePlaylist: untuk setiap action yang diterima dari client, server akan
// mengupdate playlist dan mengirimkan update ke semua client yang terhubung.
// Setiap client yang terhubung akan mendapatkan update playlist yang sama.
func (s *SongService) CollaboratePlaylist(stream pb.SongService_CollaboratePlaylistServer) error {
	s.mu.Lock()
	id := s.nextID
	s.nextID++
	s.streams[id] = stream
	s.mu.Unlock()
  
	defer func() {
	  s.mu.Lock()
	  delete(s.streams, id)
	  s.mu.Unlock()
	}()
  
	for {
	  action, err := stream.Recv()
	  if err == io.EOF {
		return nil
	  }
	  if err != nil {
		return err
	  }
  
	  s.mu.Lock()
	  switch action.Type {
	case pb.PlaylistAction_ADD:
		// Cek apakah action.Song.Id kosong atau tidak
		if action.SongId != "" {
			// Ambil lagu berdasarkan ID dari database
			oid, err := primitive.ObjectIDFromHex(action.SongId)
			if err != nil {
				s.mu.Unlock()
				return fmt.Errorf("invalid song ID: %v", err)
			}
			song, err := s.Repo.GetByID(stream.Context(), oid.Hex())
			if err != nil {
				s.mu.Unlock()
				return fmt.Errorf("could not find song with id %s: %v", action.SongId, err)
			}
	
			// Mengubah data song ke format yang akan digunakan dalam playlist
			newSong := &pb.Song{
				Id:     song.ID,
				Title:  song.Title,
				Artist: song.Artist,
				Album:  song.Album,
				Genre:  song.Genre,
			}
			// Menambahkan lagu ke playlist
			s.playlist = append(s.playlist, newSong)
		} else {
			// Jika tidak ada ID, buat lagu baru
			song := &model.Song_Model{
				Title:  action.Song.Title,
				Artist: action.Song.Artist,
				Album:  action.Song.Album,
				Genre:  action.Song.Genre,
			}
			saved, err := s.Repo.Create(stream.Context(), song)
			if err != nil {
				s.mu.Unlock()
				return err
			}
	
			// Konversi data lagu baru ke dalam format yang diperlukan untuk playlist
			newSong := &pb.Song{
				Id:     saved.ID,
				Title:  saved.Title,
				Artist: saved.Artist,
				Album:  saved.Album,
				Genre:  saved.Genre,
			}
			// Menambahkan lagu baru ke playlist
			s.playlist = append(s.playlist, newSong)
		}
	
		// Broadcast update playlist ke semua client
		update := &pb.PlaylistUpdate{Songs: s.playlist}
		for _, peer := range s.streams {
			_ = peer.Send(update)
		}	
  
	  case pb.PlaylistAction_REMOVE:
		for i, song := range s.playlist {
		  if song.Id == action.SongId {
			s.playlist = append(s.playlist[:i], s.playlist[i+1:]...)
			break
		  }
		}
  
	  case pb.PlaylistAction_MOVE:
		var moved *pb.Song
		for i, song := range s.playlist {
		  if song.Id == action.SongId {
			moved = song
			s.playlist = append(s.playlist[:i], s.playlist[i+1:]...)
			break
		  }
		}
		pos := int(action.NewPosition)
		if pos < 0 { pos = 0 }
		if pos > len(s.playlist) { pos = len(s.playlist) }
		s.playlist = append(s.playlist[:pos], append([]*pb.Song{moved}, s.playlist[pos:]...)...)
	  }
	  update := &pb.PlaylistUpdate{Songs: s.playlist}
  
	  // broadcast
	  for _, peer := range s.streams {
		_ = peer.Send(update)
	  }
	  s.mu.Unlock()
	}
  }
  
// Recommend: untuk setiap mood yang diterima dari client, server akan mengirimkan
// lagu-lagu yang sesuai dengan mood tersebut. Server akan mengabaikan genre yang dikirimkan
// oleh client. Jika mood tidak dikenali, server akan mengirimkan semua lagu yang ada di database.
func (s *SongService) Recommend(stream pb.SongService_RecommendServer) error {
	for {
	  req, err := stream.Recv()
	  if err == io.EOF {
		return nil
	  }
	  if err != nil {
		return err
	  }
  
	  all, err := s.Repo.List(stream.Context())
	  if err != nil {
		return err
	  }
  
	  // ambil genre yang cocok untuk mood ini
	  allowedGenres, ok := moodToGenres[req.Mood]
	  if !ok {
		// kalau mood tidak dikenali, fallback ke semua genre
		allowedGenres = []string{}
	  }
	  // bikin set untuk cek cepat
	  genreSet := make(map[string]bool)
	  for _, g := range allowedGenres {
		genreSet[g] = true
	  }
  
	  for _, m := range all {
		// skip genre yang user mau skip
		skip := false
		for _, sg := range req.SkipGenres {
		  if m.Genre == sg {
			skip = true
			break
		  }
		}
		if skip {
		  continue
		}
  
		// jika mood dikenali, hanya kirim yang genrenya ada di mapping
		if len(genreSet) > 0 && !genreSet[m.Genre] {
		  continue
		}
  
		reason := fmt.Sprintf("Recommended for mood: %s", req.Mood)
		resp := &pb.RecommendResponse{
		  Song: &pb.Song{
			Id:     m.ID,
			Title:  m.Title,
			Artist: m.Artist,
			Album:  m.Album,
			Genre:  m.Genre,
		  },
		  Reason: reason,
		}
		if err := stream.Send(resp); err != nil {
		  return err
		}
	  }
	}
  }

func (s *SongService) StreamAllSongs(req *pb.Empty, stream pb.SongService_StreamAllSongsServer) error {
	ctx := stream.Context()
	songs, err := s.Repo.List(ctx)
	if err != nil {
	  return err
	}
	for _, song := range songs {
	  res := &pb.Song{
		Id:     song.ID,
		Title:  song.Title,
		Artist: song.Artist,
		Album:  song.Album,
		Genre:  song.Genre,
	  }
	  if err := stream.Send(res); err != nil {
		return err
	  }
	}
	return nil
  }
  
func (s *SongService) StreamSongsByArtist(req *pb.ArtistRequest, stream pb.SongService_StreamSongsByArtistServer) error {
	ctx := stream.Context()
	artistQuery := strings.ToLower(req.Artist)
	songs, err := s.Repo.FindByArtist(ctx, artistQuery)
	if len(songs) == 0 {
		return status.Errorf(codes.NotFound, "no songs found for artist: %s", req.Artist)
	}
	if err != nil {
		return err
	}
	for _, song := range songs {
		res := &pb.Song{
			Id:     song.ID,
			Title:  song.Title,
			Artist: song.Artist,
			Album:  song.Album,
			Genre:  song.Genre,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *SongService) StreamSongsByGenre(req *pb.GenreRequest, stream pb.SongService_StreamSongsByGenreServer) error {
	ctx := stream.Context()
	genreQuery := strings.ToLower(req.Genre)
	songs, err := s.Repo.FindByGenre(ctx, genreQuery)
	if len(songs) == 0 {
		return status.Errorf(codes.NotFound, "no songs found for genre: %s", req.Genre)
	}
	if err != nil {
		return err
	}
	for _, song := range songs {
		res := &pb.Song{
			Id:     song.ID,
			Title:  song.Title,
			Artist: song.Artist,
			Album:  song.Album,
			Genre:  song.Genre,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}