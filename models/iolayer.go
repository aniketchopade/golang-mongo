package models

import (
	"context"
	"log"
	"encoding/json" // Use JSON encoding for bson. M string
	"go.mongodb.org/mongo-driver/bson"
)

// Trainer struct represent a trainer details
type Trainer struct {
	Name string
	Age  int
	City string
}

// InsertOneDoc insert a new document
func InsertOneDoc(trainer Trainer, dbName, collectionName string) interface{} {
	collection := Connclient.Database(dbName).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		log.Fatalln("Error on inserting new trainer", err)
	} else {
		log.Println("Trainer inserted sucessfully: ", insertResult)
	}
	return insertResult.InsertedID
}

// InsertManyDoc insert a new document
func InsertManyDoc(trainer []interface{}, dbName, collectionName string) interface{} {
	collection := Connclient.Database(dbName).Collection(collectionName)
	insertResult, err := collection.InsertMany(context.TODO(), trainer)
	if err != nil {
		log.Fatalln("Error on inserting new trainer", err)
	} else {
		log.Println("Trainer inserted sucessfully: ", insertResult)
	}
	return insertResult.InsertedIDs
}

// FindOneDoc will retune find results
func FindOneDoc(dbName, collectionName, query string) Trainer {
	var trainer Trainer
	var err error
	var bsonMap bson.M
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &bsonMap)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	documentReturned := collection.FindOne(context.TODO(), bsonMap)
	documentReturned.Decode(&trainer)
	return trainer
}

// FindAllDocs return all documents from the collection Heroes
func FindAllDocs(dbName, collectionName, query string) []*Trainer {
	var trainers []*Trainer
	var err error
	var bsonMap bson.M
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &bsonMap)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), bsonMap)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var trainer Trainer
		err = cur.Decode(&trainer)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		trainers = append(trainers, &trainer)
	}
	return trainers
}


// UpdateOneDoc update the info of a informed Hero
func UpdateOneDoc(dbName, collectionName, updatedData, query string) int64 {
	var filter  bson.M
	var updateData bson.M
	var err error
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &filter)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(updatedData), &updateData)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	atualizacao := bson.D{{Key: "$set", Value: updateData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating", err)
	}
	return updatedResult.ModifiedCount
}

// UpdateAllDoc update the info of a informed Hero
func UpdateAllDoc(dbName, collectionName, updatedData, query string) int64 {
	var filter  bson.M
	var updateData bson.M
	var err error
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &filter)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(updatedData), &updateData)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	atualizacao := bson.D{{Key: "$set", Value: updateData}}
	updatedResult, err := collection.UpdateMany(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating", err)
	}
	return updatedResult.ModifiedCount
}


// DeleteOneDoc will retune find results
func DeleteOneDoc(dbName, collectionName, query string) int64 {
	var err error
	var bsonMap bson.M
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &bsonMap)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	deleteResult, errDelete := collection.DeleteOne(context.TODO(), bsonMap)
	if err != nil {
		log.Fatal("Error on deleting one Doc", errDelete)
	}
	return deleteResult.DeletedCount
}

// DeleteAllDoc will retune find results
func DeleteAllDoc(dbName, collectionName, query string) int64 {
	var err error
	var bsonMap bson.M
	// Use the JSON package's Unmarshal() method
	err = json. Unmarshal([]byte(query), &bsonMap)
	if err != nil {
	log. Fatal("json. Unmarshal() ERROR:", err)
	}
	collection := Connclient.Database(dbName).Collection(collectionName)
	deleteResult, errDelete := collection.DeleteMany(context.TODO(), bsonMap)
	if err != nil {
		log.Fatal("Error on deleting many Doc", errDelete)
	}
	return deleteResult.DeletedCount
}


