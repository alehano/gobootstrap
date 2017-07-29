package example

import (
	"time"
	"github.com/asaskevich/govalidator"
)

type ExampleModel struct {
	ID        int `db:"id"`
	Title     string `db:"title" valid:"required"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t *ExampleModel) Validate() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	return nil
}

// Add some other methods to transform data ...
