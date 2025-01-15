package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/Kongsavanh12/testFinal.git/entity"
	
	
)

func TestBooking(t *testing.T){
	
	g := NewGomegaWithT(t)

	t.Run(`CustomerName is required`, func(t *testing.T){
		booking := entity.BookingHall{
			CustomerName: "Kongsavanh",
			CustomerPhone:	"0123456789",
			CustomerEmail:	"jame@gmail.com",
		}
		ok,err := govalidator.ValidateStruct(booking)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestBooking2(t *testing.T){
	g:=NewGomegaWithT(t)

	t.Run(`CustomerPhone is required`, func(t *testing.T){
		booking := entity.BookingHall{
			CustomerName:	"",
			CustomerPhone:	"0123456789",
			CustomerEmail:	"j@gmail.com",
		}
		ok,err := govalidator.ValidateStruct(booking)
	})
}