package main

import (
	"log"
	models "./models"
)

func main() {
	defer models.CloseConnections()
	trainerUday := models.Trainer{"Uday Singh", 27, "Mumbai"}
	trainerJohn := models.Trainer{"John Cena", 45, "New York"}
	trainerKapil := models.Trainer{"Kapil Sharma", 33, "Nalla Sopara"}
	
	// Insert a record in DB
	insertedID := models.InsertOneDoc(trainerUday, "goldGym", "trainers")
	log.Println(insertedID)

	// Now try to find what we have inserted
	query := `{"name": "Uday Singh"}`;
	trainer := models.FindOneDoc("goldGym", "trainers", query)
	log.Println("trainer: ", trainer)

	// Now insert many records at a time
	trainers := []interface{}{trainerJohn, trainerKapil}
	insertedIDs := models.InsertManyDoc(trainers, "goldGym", "trainers")
	log.Println(insertedIDs)

	// Retrive records on the basis of a query
	query = `{}`
	records := models.FindAllDocs("goldGym", "trainers", query)
	for _, trainer := range records {
		log.Println("trainer: ", trainer)
	}

	// Update a record
	updatedData := `{"age": 10}`
	query = `{"name": "Uday Singh"}`
	recordsUpdated := models.UpdateOneDoc("goldGym", "trainers",updatedData, query)
	log.Println("records updated count:", recordsUpdated)

	// Update many records
	updatedData = `{"age": 10}`
	query = `{"name": "Uday Singh"}`
	recordsUpdated = models.UpdateAllDoc("goldGym", "trainers",updatedData, query)
	log.Println("records updated count:", recordsUpdated)

	// Delete a record
	query = `{"name": "Uday Singh"}`
	recordsDelted := models.DeleteOneDoc("goldGym", "trainers", query)
	log.Println("removed count:", recordsDelted)


	// Delete Many records
	query = `{"name": "Kapil Sharma"}`
	recordsDelted = models.DeleteAllDoc("goldGym", "trainers", query)
	log.Println("removed count:", recordsDelted)

}

func init() {
	// Connect with DB
	connectionString := "mongodb://localhost:27017"
	models.GetClient(connectionString)

}
