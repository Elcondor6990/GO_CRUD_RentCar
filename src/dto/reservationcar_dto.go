package dto


type ReservationCarDto struct{
	StartReservation 	string	`json:"start_reserve"`
	EndReservation 	 	string	`json:"end_reserve"`
	NumberPlate			string	`json:"number_plate"`
}