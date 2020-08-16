package postgres

import (
	"github.com/JamieBShaw/golang-graphql-server/graph/models"
	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
)

type MeetupsRepo struct {
	db  *pg.DB
	Log hclog.Logger
}

func (m *MeetupsRepo) GetAll() ([]*models.Meetup, error) {
	m.Log.Info("Retrieving all meetups")

	var meetups []*models.Meetup

	err := m.db.Model(&meetups).Order("id").Select()
	if err != nil {
		m.Log.Error("Could not retrieve all meetups", err)
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupsRepo) GetMeetUpsForUser(user *models.User) ([]*models.Meetup, error) {
	m.Log.Info("Getting meet up for user", user.ID)

	var meetups []*models.Meetup

	err := m.db.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil

}

func (m *MeetupsRepo) GetByID(id string) (*models.Meetup, error) {
	m.Log.Info("Getting single meetup")

	var meetup models.Meetup

	err := m.db.Model(&meetup).Where("id = ?", id).First()
	if err != nil {
		m.Log.Error("Could not retrieve users", err)
		return nil, err
	}

	return &meetup, nil
}

func (m *MeetupsRepo) Create(newMeetup *models.Meetup) (*models.Meetup, error) {
	m.Log.Info("Creating meetup")
	_, err := m.db.Model(newMeetup).Returning("*").Insert()
	if err != nil {
		m.Log.Error("Error creating new meetup", err)
		return nil, err
	}
	return newMeetup, nil
}

func (m *MeetupsRepo) Update(meetup *models.Meetup) (*models.Meetup, error) {
	m.Log.Info("Updating meetup")
	_, err := m.db.Model(meetup).Where("id = ?", meetup.ID).Update()
	if err != nil {
		m.Log.Error("Error updating meetup", meetup.ID)
		return nil, err
	}
	return meetup, nil
}

func (m *MeetupsRepo) Delete(meetup *models.Meetup) error {
	_, err := m.db.Model(meetup).Where("id = ?", meetup.ID).Delete()
	if err != nil {
		m.Log.Error("Error deleting meetup", meetup.ID)
	}
	return err
}
