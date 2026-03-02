package tests

import (
	"Go_Rent_Car/src/dto"
	"Go_Rent_Car/src/repository"
	"Go_Rent_Car/src/service"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var endpoint string

func TestMain(m *testing.M) {

	ctx := context.Background()
	mongodbContainer, err := mongodb.Run(ctx, "mongo:latest")
	if err != nil {
		log.Printf("failed to start container: %s", err)
		os.Exit(1)
	}
	defer func() {
		if err := testcontainers.TerminateContainer(mongodbContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	endpointVar, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		log.Printf("failed to get endpoint: %s", err)
		os.Exit(1)
	}
	endpoint = endpointVar

	os.Exit(m.Run())
}

func TestMongoConnection(t *testing.T) {
	dbName := "test_db"
	mongoUri := endpoint

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoUri).SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	require.NoError(t, err)

	db := client.Database(dbName)
	require.NotNil(t, db)
}

//func TestServiceReservationCar_ValidFindByID(t *testing.T) {
//	ctx := context.Background()
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
//	require.NoError(t, err)
//	defer client.Disconnect(ctx)
//
//	db := client.Database("test-db")
//	coll := db.Collection("reservations")
//	//coll.DeleteMany(ctx, bson.D{{}})
//
//	repo := repository.NewReservationCarRepository(db)
//	svc := service.NewReservationCarService(repo)
//
//	test_id:= primitive.NewObjectID()
//
//	model := models.ReservationCar{
//		ID: test_id,
//		StartReservation: primitive.NewDateTimeFromTime(time.Date(2026, 3, 20, 12, 0, 0, 0, time.UTC)),
//		EndReservation: primitive.NewDateTimeFromTime(time.Date(2026, 3, 15, 10, 0, 0, 0, time.UTC)),
//		NumberPlate: "test_plate",
//	}
//
//	coll.InsertOne(ctx, model.ID)
//	string_id := model.ID.Hex()
//
//
//	result, err := svc.GetByID(ctx, string_id)
//	require.NoError(t, err)
//	require.NotNil(t, result)
//
//}

func TestServiceReservationCar_ValidCreate(t *testing.T) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	db := client.Database("test-db")
	//coll := db.Collection("reservations")
	//coll.DeleteMany(ctx, bson.D{{}})

	repo := repository.NewReservationCarRepository(db)
	svc := service.NewReservationCarService(repo)

	input := dto.ReservationCarDto{
		StartReservation: "2026-02-20T10:00:00Z",
		EndReservation:   "2026-03-20T12:00:00Z",
		NumberPlate:      "test_plate",
	}

	result, err := svc.Create(ctx, input)
	require.NoError(t, err)
	require.NotNil(t, result)

	require.Equal(t, input.NumberPlate, result.NumberPlate)
}
