package modeltests

import (
	"log"
	"testing"

	"github.com/andriikost/burning-gateway/api/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllStops(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatalf("Error refreshing user and stop table %v\n", err)
	}
	_, _, err = seedUsersAndStops()
	if err != nil {
		log.Fatalf("Error seeding user and stop  table %v\n", err)
	}
	stops, err := stopInstance.FindAll(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the stops: %v\n", err)
		return
	}
	assert.Equal(t, len(*stops), 2)
}

func TestSaveStop(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatalf("Error user and stop refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newStop := models.Stop{
		ID:       1,
		Name:     "This is the title",
		Content:  "This is the content",
		AuthorID: user.ID,
	}
	savedStop, err := newStop.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Stop: %v\n", err)
		return
	}
	assert.Equal(t, newStop.ID, savedStop.ID)
	assert.Equal(t, newStop.Name, savedStop.Name)
	assert.Equal(t, newStop.Content, savedStop.Content)
	assert.Equal(t, newStop.AuthorID, savedStop.AuthorID)

}

func TestGetStopByID(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatalf("Error refreshing user and stop table: %v\n", err)
	}
	stop, err := seedOneUserAndOneStop()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundStop, err := stopInstance.Find(server.DB, stop.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundStop.ID, stop.ID)
	assert.Equal(t, foundStop.Name, stop.Name)
	assert.Equal(t, foundStop.Content, stop.Content)
}

func TestUpdateStop(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatalf("Error refreshing user and stop table: %v\n", err)
	}
	stop, err := seedOneUserAndOneStop()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	stopUpdate := models.Stop{
		ID:       1,
		Name:     "modiUpdate",
		Content:  "modiupdate@gmail.com",
		AuthorID: stop.AuthorID,
	}
	updatedStop, err := stopUpdate.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedStop.ID, stopUpdate.ID)
	assert.Equal(t, updatedStop.Name, stopUpdate.Name)
	assert.Equal(t, updatedStop.Content, stopUpdate.Content)
	assert.Equal(t, updatedStop.AuthorID, stopUpdate.AuthorID)
}

func TestDeleteStop(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatalf("Error refreshing user and stop table: %v\n", err)
	}
	stop, err := seedOneUserAndOneStop()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := stopInstance.Delete(server.DB, stop.ID, stop.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
