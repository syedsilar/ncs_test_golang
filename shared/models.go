// shared/models.go

package shared

type ParkingSpot struct {
	ID           int
	Type         string
	Availability bool
	Floor        int
	Location     string
}

type Transaction struct {
	ID          int
	Timestamp   string
	Amount      float64
	VehicleInfo string
	SpotInfo    ParkingSpot
}

type User struct {
	ID             int
	Username       string
	PasswordHash   string
	AccountBalance float64
	VehicleInfo    string
	Transactions   []Transaction
}
