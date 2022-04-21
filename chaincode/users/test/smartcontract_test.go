package test

import (
	"chaincode/smartcontract"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
)

var Stub *shimtest.MockStub
var Scc *contractapi.ContractChaincode
var user1 smartcontract.User = smartcontract.User{
	ID:    "1",
	Name:  "John Lee",
	Email: "john.lee@g.com",
}
var user2 smartcontract.User = smartcontract.User{
	ID:    "2",
	Name:  "Amy Lin",
	Email: "amy.lin@g.com",
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	log.SetOutput(ioutil.Discard)
}

func NewStub() {
	Scc, err := contractapi.NewChaincode(new(smartcontract.SmartContract))
	if err != nil {
		log.Println("NewChaincode failed", err)
		os.Exit(0)
	}
	Stub = shimtest.NewMockStub("main", Scc)
}

func Test_CreateUser(t *testing.T) {
	fmt.Println("Test_CreateUser-----------------")
	NewStub()

	err := MockCreateUser(user1.ID, user1.Name, user1.Email)
	if err != nil {
		t.FailNow()
	}
}

func Test_UserExists(t *testing.T) {
	fmt.Println("Test_UserExists-----------------")
	NewStub()

	err := MockCreateUser(user1.ID, user1.Name, user1.Email)
	if err != nil {
		t.FailNow()
	}

	result, err := MockUserExists(user1.ID)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, result, true)
}

func Test_GetUser(t *testing.T) {
	fmt.Println("Test_GetUser-----------------")
	NewStub()

	err := MockCreateUser(user1.ID, user1.Name, user1.Email)
	if err != nil {
		t.FailNow()
	}

	userJson, err := MockGetUser(user1.ID)
	if err != nil {
		fmt.Println("get User error", err)
	}

	assert.Equal(t, userJson.ID, user1.ID)
	assert.Equal(t, userJson.Name, user1.Name)
	assert.Equal(t, userJson.Email, user1.Email)
}

func Test_UpdateUser(t *testing.T) {
	fmt.Println("Test_UpdateUser-----------------")
	NewStub()

	err := MockCreateUser(user1.ID, user1.Name, user1.Email)
	if err != nil {
		t.FailNow()
	}

	MockUpdateUser(user1.ID, "change name", "change email")

	userJson, err := MockGetUser(user1.ID)
	if err != nil {
		fmt.Println("get User", err)
	}

	assert.Equal(t, userJson.ID, user1.ID)
	assert.Equal(t, userJson.Name, "change name")
	assert.Equal(t, userJson.Email, "change email")

}

func Test_DeleteUser(t *testing.T) {
	fmt.Println("Test_DeleteUser-----------------")
	NewStub()
	err := MockCreateUser(user1.ID, user1.Name, user1.Email)
	if err != nil {
		t.FailNow()
	}

	MockDeleteUser(user1.ID)

	userJson, err := MockGetUser(user1.ID)
	if err != nil {
		fmt.Println("get User", err)
	}
	fmt.Println(userJson)
	assert.Equal(t, err, errors.New("GetUser error"))
}

func Test_GetAllUsers(t *testing.T) {
	fmt.Println("MockGetAllUsers-----------------")
	NewStub()

	MockCreateUser(user1.ID, user1.Name, user1.Email)
	MockCreateUser(user2.ID, user2.Name, user2.Email)

	users, err := MockGetAllUsers()
	if err != nil {
		fmt.Println("GetAllUsers error", err)
	}

	assert.Equal(t, len(users), 2)
}

func MockUserExists(id string) (bool, error) {
	res := Stub.MockInvoke("uuid", [][]byte{[]byte("UserExists"), []byte(id)})
	if res.Status != shim.OK {
		return false, errors.New("UserExists error")
	}
	var result bool = false
	json.Unmarshal(res.Payload, &result)
	return result, nil
}

func MockCreateUser(id string, name string, email string) error {
	res := Stub.MockInvoke("uuid",
		[][]byte{
			[]byte("CreateUser"),
			[]byte(id),
			[]byte(name),
			[]byte(email),
		})

	if res.Status != shim.OK {
		fmt.Println("CreateUser failed", string(res.Message))
		return errors.New("CreateUser error")
	}
	return nil
}

func MockGetUser(id string) (*smartcontract.User, error) {
	var result smartcontract.User
	res := Stub.MockInvoke("uuid",
		[][]byte{
			[]byte("GetUser"),
			[]byte(id),
		})
	if res.Status != shim.OK {
		fmt.Println("GetUser failed", string(res.Message))
		return nil, errors.New("GetUser error")
	}
	json.Unmarshal(res.Payload, &result)
	return &result, nil
}

func MockUpdateUser(id string, name string, email string) error {
	res := Stub.MockInvoke("uuid",
		[][]byte{
			[]byte("UpdateUser"),
			[]byte(id),
			[]byte(name),
			[]byte(email),
		})
	if res.Status != shim.OK {
		fmt.Println("UpdateUser failed", string(res.Message))
		return errors.New("UpdateUser error")
	}
	return nil
}

func MockDeleteUser(id string) error {
	res := Stub.MockInvoke("uuid",
		[][]byte{
			[]byte("DeleteUser"),
			[]byte(id),
		})
	if res.Status != shim.OK {
		fmt.Println("DeleteUser failed", string(res.Message))
		return errors.New("DeleteUser error")
	}
	return nil
}

func MockGetAllUsers() ([]*smartcontract.User, error) {
	res := Stub.MockInvoke("uuid", [][]byte{[]byte("GetAllUsers")})
	if res.Status != shim.OK {
		fmt.Println("GetAllUsers failed", string(res.Message))
		return nil, errors.New("GetAllUsers error")
	}
	var users []*smartcontract.User
	json.Unmarshal(res.Payload, &users)
	return users, nil
}
