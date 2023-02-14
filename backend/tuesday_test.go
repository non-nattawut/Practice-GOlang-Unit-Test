package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Tuesday struct {
	gorm.Model
	Student_ID string `valid:"matches([BCM][0-9]*)~wrong student id format"`
	Name       string `valid:"required~name cannot be blank"`
	Email      string `valid:"email~wrong email format"`
}

func TestTuesday(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	tu := Tuesday{
		Student_ID: "B6308636",
		Name:       "Nattawut Rodthong",
		Email:      "nattavutnon@gmail.com",
	}

	t.Run("All OK", func(t *testing.T) {
		ok, err := govalidator.ValidateStruct(tu)

		g.Expect(ok).To(gomega.BeTrue())

		g.Expect(err).To(gomega.BeNil())
	})

	t.Run("Wrong Srudent ID format", func(t *testing.T) {
		newTU := tu
		newTU.Student_ID = "Z6308636"
		ok, err := govalidator.ValidateStruct(newTU)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("wrong student id format"))
	})

	t.Run("Name cannot blank", func(t *testing.T) {
		newTU := tu
		newTU.Name = ""
		ok, err := govalidator.ValidateStruct(newTU)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("name cannot be blank"))
	})

	t.Run("Wrong Email format", func(t *testing.T) {
		newTU := tu
		newTU.Email = "1234234"
		ok, err := govalidator.ValidateStruct(newTU)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("wrong email format"))
	})
}
