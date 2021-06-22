package testing

import (
	"api/model"
	"api/services"
	"testing"
)

func TestCreate(t *testing.T) {

	var user model.Users
	user.Name = "aa"
	user.Password = "aa"
	user.Email = "aa"
	user.ID = 1

	result := services.CreateService(user)
	if result != "registered successfully" {
		t.Error("invalid data")
	} else {
		t.Log("registered")
	}

}
func TestShowusers(t *testing.T) {

	result := services.ShowusersService()
	if result == nil {
		t.Error("no data in the database")
	} else {
		t.Log(result)
	}

}
func TestShowuser(t *testing.T) {
	id := "1"
	result := services.ShowuserService(id)
	if result.Name == "" {
		t.Error("no data in the database")
	} else {
		t.Log(result)
	}

}
func TestShowuserNoID(t *testing.T) {

	result := services.ShowuserService("")
	if result.Name == "" {
		t.Error("invalid ID")
	} else {
		t.Log("no data found")
	}

}
func TestLogin(t *testing.T) {

	var user model.Users

	user.Password = "aa"
	user.Email = "aa"

	result := services.LoginService(user)
	if result != "logged in" {
		t.Error("invalid data")
	} else {
		t.Log("logged in")
	}

}

func TestUpdate(t *testing.T) {
	id := "1"
	var user model.Users
	user.Name = "sample"
	user.Password = "sample"
	user.Email = "sample"

	result := services.UpdateService(id, user)
	if result != "user successfully updated" {
		t.Error("invalid ID")
	} else {
		t.Log("Updated")
	}

}
func TestDelete(t *testing.T) {
	id := "1"
	result := services.DeleteService(id)
	if result != "Successfully deleted" {
		t.Error("invalid ID")
	} else {
		t.Log("Deleted")
	}

}
