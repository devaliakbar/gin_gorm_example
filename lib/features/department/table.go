package department

type Department struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
