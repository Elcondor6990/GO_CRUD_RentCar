export MONGO_URI=mongodb://root:pass@mongo:27017/?authSource=admin
export MONGO_DB_NAME=rent_car_crud

run-debug:
	go run main.go