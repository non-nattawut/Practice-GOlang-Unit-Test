package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url~url invalid"`
}

func TestVideoValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("the data is correct", func(t *testing.T) {
		v := Video{
			Name: "123",
			Url:  "123.com",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).To(gomega.BeTrue()) // ข้อมูลถูก ok จะเป็น true

		g.Expect(err).To(gomega.BeNil()) // ข้อมูลถูก error จะเป็น nil

		//g.Expect(err.Error()).To(gomega.Equal("")) // comment ทิ้งเนื่องจากไม่มี error ก็ย่อมไม่มี error message
	})

	t.Run("check Name cannot be blank", func(t *testing.T) {
		v := Video{
			Name: "",
			Url:  "123.com",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(gomega.BeTrue()) // ข้อมูลผิด ok จะเป็น flase

		g.Expect(err).ToNot(gomega.BeNil()) // ข้อมูลผิด error จะมีค่า

		g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank")) // เช็ค error message
	})

	t.Run("check url", func(t *testing.T) {
		v := Video{
			Name: "123",
			Url:  "1",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(gomega.BeTrue()) // ข้อมูลผิด ok จะเป็น flase

		g.Expect(err).ToNot(gomega.BeNil()) // ข้อมูลผิด error จะมีค่า

		g.Expect(err.Error()).To(gomega.Equal("url invalid")) // เช็ค error message
	})
}
