// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/song.proto

package song

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlaylistAction_ActionType int32

const (
	PlaylistAction_ADD    PlaylistAction_ActionType = 0
	PlaylistAction_REMOVE PlaylistAction_ActionType = 1
	PlaylistAction_MOVE   PlaylistAction_ActionType = 2
)

// Enum value maps for PlaylistAction_ActionType.
var (
	PlaylistAction_ActionType_name = map[int32]string{
		0: "ADD",
		1: "REMOVE",
		2: "MOVE",
	}
	PlaylistAction_ActionType_value = map[string]int32{
		"ADD":    0,
		"REMOVE": 1,
		"MOVE":   2,
	}
)

func (x PlaylistAction_ActionType) Enum() *PlaylistAction_ActionType {
	p := new(PlaylistAction_ActionType)
	*p = x
	return p
}

func (x PlaylistAction_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlaylistAction_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_song_proto_enumTypes[0].Descriptor()
}

func (PlaylistAction_ActionType) Type() protoreflect.EnumType {
	return &file_proto_song_proto_enumTypes[0]
}

func (x PlaylistAction_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlaylistAction_ActionType.Descriptor instead.
func (PlaylistAction_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{5, 0}
}

type Song struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist        string                 `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Album         string                 `protobuf:"bytes,4,opt,name=album,proto3" json:"album,omitempty"`
	Genre         string                 `protobuf:"bytes,5,opt,name=genre,proto3" json:"genre,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Song) Reset() {
	*x = Song{}
	mi := &file_proto_song_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{0}
}

func (x *Song) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Song) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Song) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *Song) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

func (x *Song) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type SongRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SongRequest) Reset() {
	*x = SongRequest{}
	mi := &file_proto_song_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongRequest) ProtoMessage() {}

func (x *SongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongRequest.ProtoReflect.Descriptor instead.
func (*SongRequest) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{1}
}

func (x *SongRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SongList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Songs         []*Song                `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SongList) Reset() {
	*x = SongList{}
	mi := &file_proto_song_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SongList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongList) ProtoMessage() {}

func (x *SongList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongList.ProtoReflect.Descriptor instead.
func (*SongList) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{2}
}

func (x *SongList) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type SongInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Artist        string                 `protobuf:"bytes,2,opt,name=artist,proto3" json:"artist,omitempty"`
	Album         string                 `protobuf:"bytes,3,opt,name=album,proto3" json:"album,omitempty"`
	Genre         string                 `protobuf:"bytes,4,opt,name=genre,proto3" json:"genre,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SongInput) Reset() {
	*x = SongInput{}
	mi := &file_proto_song_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SongInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongInput) ProtoMessage() {}

func (x *SongInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongInput.ProtoReflect.Descriptor instead.
func (*SongInput) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{3}
}

func (x *SongInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SongInput) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *SongInput) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

func (x *SongInput) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_song_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{4}
}

type PlaylistAction struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Type          PlaylistAction_ActionType `protobuf:"varint,1,opt,name=type,proto3,enum=song.PlaylistAction_ActionType" json:"type,omitempty"`
	Song          *SongInput                `protobuf:"bytes,2,opt,name=song,proto3" json:"song,omitempty"`                                   // untuk ADD
	SongId        string                    `protobuf:"bytes,3,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`                 // untuk REMOVE / MOVE
	NewPosition   int32                     `protobuf:"varint,4,opt,name=new_position,json=newPosition,proto3" json:"new_position,omitempty"` // untuk MOVE
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaylistAction) Reset() {
	*x = PlaylistAction{}
	mi := &file_proto_song_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaylistAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistAction) ProtoMessage() {}

func (x *PlaylistAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistAction.ProtoReflect.Descriptor instead.
func (*PlaylistAction) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{5}
}

func (x *PlaylistAction) GetType() PlaylistAction_ActionType {
	if x != nil {
		return x.Type
	}
	return PlaylistAction_ADD
}

func (x *PlaylistAction) GetSong() *SongInput {
	if x != nil {
		return x.Song
	}
	return nil
}

func (x *PlaylistAction) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *PlaylistAction) GetNewPosition() int32 {
	if x != nil {
		return x.NewPosition
	}
	return 0
}

type PlaylistUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Songs         []*Song                `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"` // state playlist terkini
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaylistUpdate) Reset() {
	*x = PlaylistUpdate{}
	mi := &file_proto_song_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaylistUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistUpdate) ProtoMessage() {}

func (x *PlaylistUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistUpdate.ProtoReflect.Descriptor instead.
func (*PlaylistUpdate) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{6}
}

func (x *PlaylistUpdate) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type RecommendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mood          string                 `protobuf:"bytes,1,opt,name=mood,proto3" json:"mood,omitempty"` // e.g. "happy", "chill"
	SkipGenres    []string               `protobuf:"bytes,2,rep,name=skip_genres,json=skipGenres,proto3" json:"skip_genres,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendRequest) Reset() {
	*x = RecommendRequest{}
	mi := &file_proto_song_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendRequest) ProtoMessage() {}

func (x *RecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendRequest.ProtoReflect.Descriptor instead.
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{7}
}

func (x *RecommendRequest) GetMood() string {
	if x != nil {
		return x.Mood
	}
	return ""
}

func (x *RecommendRequest) GetSkipGenres() []string {
	if x != nil {
		return x.SkipGenres
	}
	return nil
}

type RecommendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Song          *Song                  `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"` // misal: "because you said chill"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendResponse) Reset() {
	*x = RecommendResponse{}
	mi := &file_proto_song_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendResponse) ProtoMessage() {}

func (x *RecommendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendResponse.ProtoReflect.Descriptor instead.
func (*RecommendResponse) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{8}
}

func (x *RecommendResponse) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

func (x *RecommendResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type GenreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Genre         string                 `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenreRequest) Reset() {
	*x = GenreRequest{}
	mi := &file_proto_song_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreRequest) ProtoMessage() {}

func (x *GenreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreRequest.ProtoReflect.Descriptor instead.
func (*GenreRequest) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{9}
}

func (x *GenreRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type ArtistRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Artist        string                 `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArtistRequest) Reset() {
	*x = ArtistRequest{}
	mi := &file_proto_song_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistRequest) ProtoMessage() {}

func (x *ArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_song_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtistRequest.ProtoReflect.Descriptor instead.
func (*ArtistRequest) Descriptor() ([]byte, []int) {
	return file_proto_song_proto_rawDescGZIP(), []int{10}
}

func (x *ArtistRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

var File_proto_song_proto protoreflect.FileDescriptor

const file_proto_song_proto_rawDesc = "" +
	"\n" +
	"\x10proto/song.proto\x12\x04song\"p\n" +
	"\x04Song\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x16\n" +
	"\x06artist\x18\x03 \x01(\tR\x06artist\x12\x14\n" +
	"\x05album\x18\x04 \x01(\tR\x05album\x12\x14\n" +
	"\x05genre\x18\x05 \x01(\tR\x05genre\"\x1d\n" +
	"\vSongRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\",\n" +
	"\bSongList\x12 \n" +
	"\x05songs\x18\x01 \x03(\v2\n" +
	".song.SongR\x05songs\"e\n" +
	"\tSongInput\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x16\n" +
	"\x06artist\x18\x02 \x01(\tR\x06artist\x12\x14\n" +
	"\x05album\x18\x03 \x01(\tR\x05album\x12\x14\n" +
	"\x05genre\x18\x04 \x01(\tR\x05genre\"\a\n" +
	"\x05Empty\"\xd3\x01\n" +
	"\x0ePlaylistAction\x123\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1f.song.PlaylistAction.ActionTypeR\x04type\x12#\n" +
	"\x04song\x18\x02 \x01(\v2\x0f.song.SongInputR\x04song\x12\x17\n" +
	"\asong_id\x18\x03 \x01(\tR\x06songId\x12!\n" +
	"\fnew_position\x18\x04 \x01(\x05R\vnewPosition\"+\n" +
	"\n" +
	"ActionType\x12\a\n" +
	"\x03ADD\x10\x00\x12\n" +
	"\n" +
	"\x06REMOVE\x10\x01\x12\b\n" +
	"\x04MOVE\x10\x02\"2\n" +
	"\x0ePlaylistUpdate\x12 \n" +
	"\x05songs\x18\x01 \x03(\v2\n" +
	".song.SongR\x05songs\"G\n" +
	"\x10RecommendRequest\x12\x12\n" +
	"\x04mood\x18\x01 \x01(\tR\x04mood\x12\x1f\n" +
	"\vskip_genres\x18\x02 \x03(\tR\n" +
	"skipGenres\"K\n" +
	"\x11RecommendResponse\x12\x1e\n" +
	"\x04song\x18\x01 \x01(\v2\n" +
	".song.SongR\x04song\x12\x16\n" +
	"\x06reason\x18\x02 \x01(\tR\x06reason\"$\n" +
	"\fGenreRequest\x12\x14\n" +
	"\x05genre\x18\x01 \x01(\tR\x05genre\"'\n" +
	"\rArtistRequest\x12\x16\n" +
	"\x06artist\x18\x01 \x01(\tR\x06artist2\xb5\x04\n" +
	"\vSongService\x12)\n" +
	"\n" +
	"CreateSong\x12\x0f.song.SongInput\x1a\n" +
	".song.Song\x12(\n" +
	"\aGetSong\x12\x11.song.SongRequest\x1a\n" +
	".song.Song\x12(\n" +
	"\tListSongs\x12\v.song.Empty\x1a\x0e.song.SongList\x12$\n" +
	"\n" +
	"UpdateSong\x12\n" +
	".song.Song\x1a\n" +
	".song.Song\x12,\n" +
	"\n" +
	"DeleteSong\x12\x11.song.SongRequest\x1a\v.song.Empty\x12+\n" +
	"\bSongChat\x12\x0f.song.SongInput\x1a\n" +
	".song.Song(\x010\x01\x12E\n" +
	"\x13CollaboratePlaylist\x12\x14.song.PlaylistAction\x1a\x14.song.PlaylistUpdate(\x010\x01\x12@\n" +
	"\tRecommend\x12\x16.song.RecommendRequest\x1a\x17.song.RecommendResponse(\x010\x01\x12+\n" +
	"\x0eStreamAllSongs\x12\v.song.Empty\x1a\n" +
	".song.Song0\x01\x126\n" +
	"\x12StreamSongsByGenre\x12\x12.song.GenreRequest\x1a\n" +
	".song.Song0\x01\x128\n" +
	"\x13StreamSongsByArtist\x12\x13.song.ArtistRequest\x1a\n" +
	".song.Song0\x01B\x1eZ\x1cgrpc-song-manager/proto;songb\x06proto3"

var (
	file_proto_song_proto_rawDescOnce sync.Once
	file_proto_song_proto_rawDescData []byte
)

func file_proto_song_proto_rawDescGZIP() []byte {
	file_proto_song_proto_rawDescOnce.Do(func() {
		file_proto_song_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_song_proto_rawDesc), len(file_proto_song_proto_rawDesc)))
	})
	return file_proto_song_proto_rawDescData
}

var file_proto_song_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_song_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_song_proto_goTypes = []any{
	(PlaylistAction_ActionType)(0), // 0: song.PlaylistAction.ActionType
	(*Song)(nil),                   // 1: song.Song
	(*SongRequest)(nil),            // 2: song.SongRequest
	(*SongList)(nil),               // 3: song.SongList
	(*SongInput)(nil),              // 4: song.SongInput
	(*Empty)(nil),                  // 5: song.Empty
	(*PlaylistAction)(nil),         // 6: song.PlaylistAction
	(*PlaylistUpdate)(nil),         // 7: song.PlaylistUpdate
	(*RecommendRequest)(nil),       // 8: song.RecommendRequest
	(*RecommendResponse)(nil),      // 9: song.RecommendResponse
	(*GenreRequest)(nil),           // 10: song.GenreRequest
	(*ArtistRequest)(nil),          // 11: song.ArtistRequest
}
var file_proto_song_proto_depIdxs = []int32{
	1,  // 0: song.SongList.songs:type_name -> song.Song
	0,  // 1: song.PlaylistAction.type:type_name -> song.PlaylistAction.ActionType
	4,  // 2: song.PlaylistAction.song:type_name -> song.SongInput
	1,  // 3: song.PlaylistUpdate.songs:type_name -> song.Song
	1,  // 4: song.RecommendResponse.song:type_name -> song.Song
	4,  // 5: song.SongService.CreateSong:input_type -> song.SongInput
	2,  // 6: song.SongService.GetSong:input_type -> song.SongRequest
	5,  // 7: song.SongService.ListSongs:input_type -> song.Empty
	1,  // 8: song.SongService.UpdateSong:input_type -> song.Song
	2,  // 9: song.SongService.DeleteSong:input_type -> song.SongRequest
	4,  // 10: song.SongService.SongChat:input_type -> song.SongInput
	6,  // 11: song.SongService.CollaboratePlaylist:input_type -> song.PlaylistAction
	8,  // 12: song.SongService.Recommend:input_type -> song.RecommendRequest
	5,  // 13: song.SongService.StreamAllSongs:input_type -> song.Empty
	10, // 14: song.SongService.StreamSongsByGenre:input_type -> song.GenreRequest
	11, // 15: song.SongService.StreamSongsByArtist:input_type -> song.ArtistRequest
	1,  // 16: song.SongService.CreateSong:output_type -> song.Song
	1,  // 17: song.SongService.GetSong:output_type -> song.Song
	3,  // 18: song.SongService.ListSongs:output_type -> song.SongList
	1,  // 19: song.SongService.UpdateSong:output_type -> song.Song
	5,  // 20: song.SongService.DeleteSong:output_type -> song.Empty
	1,  // 21: song.SongService.SongChat:output_type -> song.Song
	7,  // 22: song.SongService.CollaboratePlaylist:output_type -> song.PlaylistUpdate
	9,  // 23: song.SongService.Recommend:output_type -> song.RecommendResponse
	1,  // 24: song.SongService.StreamAllSongs:output_type -> song.Song
	1,  // 25: song.SongService.StreamSongsByGenre:output_type -> song.Song
	1,  // 26: song.SongService.StreamSongsByArtist:output_type -> song.Song
	16, // [16:27] is the sub-list for method output_type
	5,  // [5:16] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_song_proto_init() }
func file_proto_song_proto_init() {
	if File_proto_song_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_song_proto_rawDesc), len(file_proto_song_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_song_proto_goTypes,
		DependencyIndexes: file_proto_song_proto_depIdxs,
		EnumInfos:         file_proto_song_proto_enumTypes,
		MessageInfos:      file_proto_song_proto_msgTypes,
	}.Build()
	File_proto_song_proto = out.File
	file_proto_song_proto_goTypes = nil
	file_proto_song_proto_depIdxs = nil
}
