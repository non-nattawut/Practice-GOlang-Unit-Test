package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `valid:"email~email must true format,required~email cannot be blank"`
	Url   string `valid:"url~url must true format"`
}

func TestUserValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	u := User{
		Email: "natt@gmail.com",
		Url:   "123.com",
	}

	t.Run("all OK", func(t *testing.T) {

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).To(gomega.BeTrue())

		g.Expect(err).To(gomega.BeNil())
	})

	t.Run("url bad format", func(t *testing.T) {
		u2 := u
		u2.Url = "123"

		ok, err := govalidator.ValidateStruct(u2)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("url must true format"))
	})

	t.Run("email can not blank", func(t *testing.T) {
		u2 := u
		u2.Email = ""

		ok, err := govalidator.ValidateStruct(u2)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("email cannot be blank"))
	})

	t.Run("email must true format", func(t *testing.T) {
		u2 := u
		u2.Email = "natt@gmail"

		ok, err := govalidator.ValidateStruct(u2)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("email must true format"))
	})
}
