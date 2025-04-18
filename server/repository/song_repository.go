package repository

import (
    "context"
    "fmt"
    "grpc-song-manager/server/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type SongRepository struct {
    collection *mongo.Collection
}

func NewSongRepository(db *mongo.Database) *SongRepository {
    return &SongRepository{
        collection: db.Collection("songs"),
    }
}

func (r *SongRepository) Create(ctx context.Context, song *model.Song_Model) (*model.Song_Model, error) {
    res, err := r.collection.InsertOne(ctx, song)
    if err != nil {
        return nil, err
    }
    // konversi ObjectID ke hex string
    oid, ok := res.InsertedID.(primitive.ObjectID)
    if !ok {
        return nil, fmt.Errorf("failed to convert InsertedID to ObjectID: %v", res.InsertedID)
    }
    song.ID = oid.Hex()
    return song, nil
}

func (r *SongRepository) GetByID(ctx context.Context, id string) (*model.Song_Model, error) {
    // ubah hex string ke ObjectID
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var song model.Song_Model
    if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&song); err != nil {
        return nil, err
    }
    // pastikan field ID terisi dengan string
    song.ID = id
    return &song, nil
}

func (r *SongRepository) List(ctx context.Context) ([]*model.Song_Model, error) {
    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var songs []*model.Song_Model
    for cursor.Next(ctx) {
        var s model.Song_Model
        if err := cursor.Decode(&s); err != nil {
            return nil, err
        }
        // jika model.Song_Model.ID bertipe string, pastikan sudah berisi hex string
        // (driver akan decode ObjectID ke string kosong, jadi ini opsional)
        songs = append(songs, &s)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }
    return songs, nil
}

func (r *SongRepository) Update(ctx context.Context, song *model.Song_Model) (*model.Song_Model, error) {
    oid, err := primitive.ObjectIDFromHex(song.ID)
    if err != nil {
        return nil, err
    }

    set := bson.M{}

    if song.Title != "" {
        set["title"] = song.Title
    }
    if song.Artist != "" {
        set["artist"] = song.Artist
    }
    if song.Album != "" {
        set["album"] = song.Album
    }
    if song.Genre != "" {
        set["genre"] = song.Genre
    }

    if len(set) == 0 {
        return nil, fmt.Errorf("no valid fields to update")
    }

    update := bson.M{"$set": set}

    _, err = r.collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
    if err != nil {
        return nil, err
    }

    return song, nil
}

func (r *SongRepository) Delete(ctx context.Context, id string) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
    return err
}

func (r *SongRepository) FindByArtist(ctx context.Context, artist string) ([]*model.Song_Model, error) {
	filter := bson.M{
		"artist": bson.M{
			"$regex":   artist,
			"$options": "i", // <== ignore case
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var songs []*model.Song_Model
	for cursor.Next(ctx) {
		var song model.Song_Model
		if err := cursor.Decode(&song); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}

func (r *SongRepository) FindByGenre(ctx context.Context, genre string) ([]*model.Song_Model, error) {
	filter := bson.M{
		"genre": bson.M{
			"$regex":   genre,
			"$options": "i", // <== ignore case
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var songs []*model.Song_Model
	for cursor.Next(ctx) {
		var song model.Song_Model
		if err := cursor.Decode(&song); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}

