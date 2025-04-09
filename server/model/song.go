package model

type Song_Model struct {
	ID     string `bson:"_id,omitempty"`
	Title  string `bson:"title"`
	Artist string `bson:"artist"`
	Album  string `bson:"album"`
	Genre  string `bson:"genre"`
}
