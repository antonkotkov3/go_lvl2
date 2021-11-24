package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Contacts struct {
	Email            string
	Phone            string
	AdditionalPhones map[string]string
}

type Person struct {
	Name     string
	Age      int64
	Contacts Contacts
}

type Person2 struct {
	Name     string
	Age      int64
	Contacts *Contacts
}

type Person3 struct {
	Name     string
	age      int64
	Contacts *Contacts
}

func buildData() map[string]interface{} {
	data := make(map[string]interface{})
	data["Name"] = "Nick"
	data["Age"] = int64(45)
	contacts := make(map[string]interface{})
	contacts["Phone"] = "71234567890"
	contacts["Email"] = "nick@gmail.com"
	contacts["AdditionalPhones"] = map[string]string{
		"Home": "71234566",
		"Work": "12345667",
	}
	data["Contacts"] = contacts
	return data
}

func TestFillStruct(t *testing.T) {
	data := buildData()
	valid := &Person{
		Name: "Nick",
		Age:  45,
		Contacts: Contacts{
			Phone: "71234567890",
			Email: "nick@gmail.com",
			AdditionalPhones: map[string]string{
				"Home": "71234566",
				"Work": "12345667",
			},
		},
	}
	result := &Person{}
	err := FillStruct(result, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid)
	}
	if !reflect.DeepEqual(result, valid) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result, valid)
	}
}

func TestFieldNotExists(t *testing.T) {
	data := buildData()
	result := &Person{}

	data["Invalid"] = "Invalid"
	validError := fmt.Errorf("No such field: %s in obj", "Invalid")
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
func TestInvalidType(t *testing.T) {
	data := buildData()
	result := &Person{}

	validError := fmt.Errorf("Provided value type didn't match obj field type")
	data["Age"] = float64(3.1415926)
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}

func TestPassInnerStructViaPtr(t *testing.T) {
	data := buildData()
	data["Age"] = int64(23)
	valid := &Person2{
		Name: "Nick",
		Age:  23,
		Contacts: &Contacts{
			Phone: "71234567890",
			Email: "nick@gmail.com",
			AdditionalPhones: map[string]string{
				"Home": "71234566",
				"Work": "12345667",
			},
		},
	}
	result := &Person2{}
	err := FillStruct(result, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid)
	}
	if !reflect.DeepEqual(result, valid) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result, valid)
	}
}

func TestSetRemovedField(t *testing.T) {
	data := buildData()

	delete(data, "Age")
	data["age"] = 23
	result := &Person3{}
	validError := fmt.Errorf("Cannot set %s field value", "age")
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
