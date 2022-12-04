/*
	Student Entities
*/
package entities

//Data Student
type Address struct {
	State string `json:"state"`
	City  string `json:"city"`
}

type Student struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int64   `json:"age"`
	Address Address `json:"address"`
}

//Fungsi ini untuk mendefinisikan hasil return data
func ToStudentEntity(u *Student) *Student {
	return &Student{
		ID:      u.ID,
		Name:    u.Name,
		Age:     u.Age,
		Address: u.Address,
	}
}
