package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Error("DivByZero", String)
	HTTP(func() {
		Response("DivByZero", StatusBadRequest)
	})
})

var _ = Service("calc", func() {
	Method("div", func() {
		Error("DivByZero", String)
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/div/{a}/{b}")
		})
	})
})
