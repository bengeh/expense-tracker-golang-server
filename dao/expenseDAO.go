package dao

import (
	"context"
	"expense_tracker_golang_api_with_docker/models"
	"fmt"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
// CONNECTIONSTRING DB connection string
var CONNECTIONSTRING = goDotEnvVariable("MONGO_URL") + "?retryWrites=false"
// // DBNAME Database name
var DBNAME = goDotEnvVariable("DB_NAME")
// // COLLNAME Collection name
var COLLNAME = goDotEnvVariable("COllECTION_NAME")

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}
// // InsertManyValues inserts many items from byte slice
// func InsertManyValues(people []models.Person) {
// 	var ppl []interface{}
// 	for _, p := range people {
// 		ppl = append(ppl, p)
// 	}
// 	_, err := db.Collection(COLLNAME).InsertMany(context.Background(), ppl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // InsertOneValue inserts one item from Person model
func InsertOneValue(expense models.Expense) {
	fmt.Println("inside insert one value...")
	fmt.Println(expense)

	
	result, insertErr := db.Collection(COLLNAME).InsertOne(context.Background(), expense)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1) // safely exit script on error
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)

		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
	}
}

// GetAllExpense returns all expense from DB
func GetAllExpense() []models.Expense {
	fmt.Println("INSIDE GET ALL Expenses")
	
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Expense

	fmt.Println("this is the data set...")
	fmt.Println(cur)
	fmt.Println(elements)

	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		var elem models.Expense
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("inside the cur")
		fmt.Println(elem)
		elements = append(elements, elem)
		fmt.Println(elements)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return elements
}

// // DeletePerson deletes an existing person
// func DeletePerson(person models.Person) {
// 	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), person, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // UpdatePerson updates an existing person
// func UpdatePerson(person models.Person, personID string) {
// 	doc := db.Collection(COLLNAME).FindOneAndUpdate(
// 		context.Background(),
// 		bson.NewDocument(
// 			bson.EC.String("id", personID),
// 		),
// 		bson.NewDocument(
// 			bson.EC.SubDocumentFromElements("$set",
// 				bson.EC.String("firstname", person.Firstname),
// 				bson.EC.String("lastname", person.Lastname),
// 				bson.EC.String("contactinfo.city", person.City),
// 				bson.EC.String("contactinfo.zipcode", person.Zipcode),
// 				bson.EC.String("contactinfo.phone", person.Phone)),
// 		),
// 		nil)
// 	fmt.Println(doc)
// }
