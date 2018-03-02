// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "survey": Application User Types
//
// Command:
// $ goagen
// --design=github.com/goadesign/examples/form_encoding/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/form_encoding
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// yesNoPayload user type.
type yesNoPayload struct {
	// Voter name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Yes or no
	Vote *string `form:"vote,omitempty" json:"vote,omitempty" xml:"vote,omitempty"`
}

// Validate validates the yesNoPayload type instance.
func (ut *yesNoPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Vote == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "vote"))
	}
	if ut.Vote != nil {
		if !(*ut.Vote == "yes" || *ut.Vote == "no") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.vote`, *ut.Vote, []interface{}{"yes", "no"}))
		}
	}
	return
}

// Publicize creates YesNoPayload from yesNoPayload
func (ut *yesNoPayload) Publicize() *YesNoPayload {
	var pub YesNoPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Vote != nil {
		pub.Vote = *ut.Vote
	}
	return &pub
}

// YesNoPayload user type.
type YesNoPayload struct {
	// Voter name
	Name string `form:"name" json:"name" xml:"name"`
	// Yes or no
	Vote string `form:"vote" json:"vote" xml:"vote"`
}

// Validate validates the YesNoPayload type instance.
func (ut *YesNoPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Vote == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "vote"))
	}
	if !(ut.Vote == "yes" || ut.Vote == "no") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.vote`, ut.Vote, []interface{}{"yes", "no"}))
	}
	return
}
