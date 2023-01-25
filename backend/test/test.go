package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func TestUserValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check", func(t *testing.T) {
		v := Video{
			Name: "",
			Url:  "",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
	})
}
