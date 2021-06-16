package testing

import (
	"api/model"
	"api/routes"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShowusers(t *testing.T) {

	assert.Equal(t, 120, 120.1, "equal values")
	if false {
		t.Error("not equal valsues")
	}
}

type MyMockedObject struct{
	mock.Mock
  }

  func (m *MyMockedObject) hashing(routes) (bool, error) {

	args := m.Called(hashing)
	return args.Bool(0), args.Error(1)
  
  }
  func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)
  
	// setup expectations
	testObj.On("hashing", mock.Anything).Return(true, nil)
  
	routes.Create(testObj)

  
	

  
  
  }
  