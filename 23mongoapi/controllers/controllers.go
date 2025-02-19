package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "main/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongo
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("COllection reference is coming")
}

//mongodb helpers - file

// insert 1 record
// why model? bcz it depends on wht is intialised within package
func insertOneMovie(movie_name model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie_name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted successfully", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} //added filter
	update := bson.M{"$set": bson.M{"watched": true}}
	result, _ := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("modified count", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Println("Movie got deleted with delete count", deleteCount)
}

// delete all records from mongodb
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual Controllers -> fetch the functions all above
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //can also do in routes
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkasWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("deleted successfully")
}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

// func main() {
// 	// Example usage in main
// 	if collection != nil {
// 		fmt.Println("Successfully connected to the collection:", collection.Name())
// 	} else {
// 		fmt.Println("Failed to connect to the collection.")
// 	}
// }

//bson.M and bson.D

//updation and deletion in mongodb happens on the basis of filter
//bson.D{{}} -> means that no filter is needed
