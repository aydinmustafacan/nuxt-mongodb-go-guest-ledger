package middleware

import (
	"awesomeProject/models"
	"awesomeProject/morse"
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
	_ "time"
)


// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	print(payload)
	json.NewEncoder(w).Encode(payload)
}





func MorseCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var morsecode models.Morse
	_ = json.NewDecoder(r.Body).Decode(&morsecode)
	fmt.Println(morse.EncodeITU(morsecode.MorseCode))
	morsecode.MorseCode = morse.EncodeITU(morsecode.MorseCode)
	insertOneMorse(morsecode)
	fmt.Println(" morse code is --> ")
	fmt.Println(morsecode)
	json.NewEncoder(w).Encode(morsecode)


}


func FizzBuzz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var fizzbuzz models.FizzBuzz
	_ = json.NewDecoder(r.Body).Decode(&fizzbuzz)
	params := mux.Vars(r)
	fmt.Println(params["count"])
	x := params["count"]
	intVar, _ := strconv.Atoi(x)
	fmt.Println(intVar)
	var myarr[16]string
	mylist := list.New()
	for i := 0; i < intVar; i++ {
		if (i+1) % 15 == 0 {
			mylist.PushBack("FizzBuzz")
			myarr[i] = "FizzBuzz"
		} else if (i+1) % 3 == 0 {
			mylist.PushBack("Fizz")
			myarr[i] = "Fizz"
		} else if (i+1) % 5 == 0 {
			mylist.PushBack("Buzz")
			myarr[i] = "Buzz"
		} else {
			var mystr string = strconv.Itoa(i + 1)
			mylist.PushBack(mystr)

			myarr[i] = mystr
		}
	}
	for element := mylist.Front(); element != nil; element = element.Next() {
		//fizzbuzz.FizzBuzz += element.Value
		fmt.Println(reflect.TypeOf(element.Value))
		fizzbuzz.FizzBuzz += element.Value.(string)
		fizzbuzz.FizzBuzz += " "
		fmt.Println(element.Value)
	}
	fmt.Println(myarr)
	fmt.Println(fizzbuzz)
	json.NewEncoder(w).Encode(fizzbuzz)
}

type people struct {
	Author string `json:"author"`
	Text string `json:"text"`
}
type resp struct {
	Text string `json:"text"`
	Author string `json:"author"`

}


func Quotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var quotes models.Quotes
	_ = json.NewDecoder(r.Body).Decode(&quotes)
	fmt.Println(quotes.Author)
	fmt.Println(quotes.Text)
//	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"},
//{"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"},
//{"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"},
//{"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
//	textBytes := []byte(text)
//
//	people1 := people{}
//	err := json.Unmarshal(textBytes, &people1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(people1.Number)
	url := "https://type.fit/api/quotes"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	dataJson := `[
  {
    "text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
    "author": "Thomas Edison"
  },
  {
    "text": "You can observe a lot just by watching.",
    "author": "Yogi Berra"
  },
  {
    "text": "A house divided against itself cannot stand.",
    "author": "Abraham Lincoln"
  }]`
	var arr []resp
	_ = json.Unmarshal([]byte(dataJson), &arr)
	log.Printf("Unmarshaled: %v", arr)

	var myarr[]resp
	people1 := myarr
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(people1)



}


func GuestLedger(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var guestledger models.GuestLedger
	_ = json.NewDecoder(r.Body).Decode(&guestledger)
	fmt.Println("counter: ")
	fmt.Println(guestledger.Email)
	fmt.Println(guestledger.Message)
	insertOneGuestLedger(guestledger)
	payload := getAllTask()
	print(payload)
	json.NewEncoder(w).Encode(payload)
}

func Card(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var card models.Card
	_ = json.NewDecoder(r.Body).Decode(&card)
	fmt.Println("counter: ")
	fmt.Println(card.Title)
	fmt.Println(card.SubTitle)
	fmt.Println(card.Text)
	fmt.Println(card.Author)
	insertOneCard(card)
	payload := getAllTask()
	print(payload)
	json.NewEncoder(w).Encode(payload)
}

func Counter(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var counter models.Counter
	_ = json.NewDecoder(r.Body).Decode(&counter)
	fmt.Println("counter: ")
	fmt.Println(counter.Counter)
	insertOneCounter(counter)

}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	fmt.Println(task, r.Body)
	fmt.Println("task: "+task.Task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// TaskComplete update task route
func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}




// Insert one task in the DB
func insertOneCounter(counter models.Counter) {
	insertResult, err := collection.InsertOne(context.Background(), counter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single counter ", insertResult.InsertedID)
}


func insertOneGuestLedger(guest models.GuestLedger) {
	insertResult, err := collection.InsertOne(context.Background(), guest)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single email and message ", insertResult.InsertedID)
}
func insertOneCard(card models.Card) {
	insertResult, err := collection.InsertOne(context.Background(), card)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single email and message ", insertResult.InsertedID)
}

func insertOneMorse(morse models.Morse) {
	insertResult, err := collection.InsertOne(context.Background(), morse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
