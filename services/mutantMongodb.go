package services

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MutantMongodb struct {
	stringConn string
	nameColl   string
	nameDb     string
}

func NewMutantMongodb() *MutantMongodb {
	m := new(MutantMongodb)
	m.stringConn = os.Getenv("DBCONN")
	m.nameDb = "mutants"
	m.nameColl = "dnas"
	return m
}

func (m *MutantMongodb) Find(dna EDna) (EDna, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		return EDna{}, &ConexError{}
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	filter := bson.D{{"dna", dna.Dna}}
	var result EDna
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (m *MutantMongodb) Add(dna EDna) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		return &ConexError{}
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	insertResult, err := collection.InsertOne(context.TODO(), dna)
	_ = insertResult
	return err
}

func (m *MutantMongodb) Update(dna EDna, newDna EDna) error {
	return nil
}

func (m *MutantMongodb) Delete(dna EDna) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		return &ConexError{}
	}
	defer client.Disconnect(ctx)

	collection := client.Database(m.nameDb).Collection(m.nameColl)

	delFilter := bson.D{{"dna", dna.Dna}}
	deleteResult, err := collection.DeleteOne(context.TODO(), delFilter)
	_ = deleteResult
	return err
}

func (m *MutantMongodb) GetAll() ([]EDna, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.stringConn))
	if err != nil {
		return []EDna{}, &ConexError{}
	}
	defer client.Disconnect(ctx)

	var results []EDna
	collection := client.Database(m.nameDb).Collection(m.nameColl)
	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return results, err
	}

	for cur.Next(context.TODO()) {

		var s EDna
		err := cur.Decode(&s)
		if err != nil {
			return results, err
		}

		results = append(results, s)
	}

	return results, err
}

type ConexError struct {
}

func (e *ConexError) Error() string {
	return "Error de conexion a Mongodb"
}
