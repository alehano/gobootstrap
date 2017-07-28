package example

import (
	"time"
	"github.com/asaskevich/govalidator"
)

type ExampleModel struct {
	Id      int
	Title   string `valid:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *ExampleModel) Validate() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	return nil
}

// Add some other methods to transform data ...
