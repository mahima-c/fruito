package database

import (
	"errors"
	"fmt"

	"github.com/Mrhb787/hospital-ward-manager/model"
)

func (s *service) GetUserById(userId uint32) (resp model.User, err error) {

	tx, txErr := s.client.DB.Begin()
	if txErr != nil {
		fmt.Println("Failed to begin transaction", txErr)
		return resp, txErr
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	dbResp, dbErr := model.GetUserById(tx, int(userId))
	if dbErr != nil {
		return resp, dbErr
	}

	if dbResp == nil {
		return resp, errors.New("user not found")
	}

	resp = *dbResp
	return resp, nil
}

func (s *service) GetUserByPhone(phone string) (resp model.User, err error) {

	tx, txErr := s.client.DB.Begin()
	if txErr != nil {
		fmt.Println("Failed to begin transaction", txErr)
		return resp, txErr
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	dbResp, dbErr := model.GetUserByPhone(tx, phone)
	if dbErr != nil {
		return resp, dbErr
	}

	if dbResp == nil {
		return resp, errors.New("user not found")
	}

	resp = *dbResp
	return resp, nil
}

func (s *service) CreateUserSession(session model.UserSession) (err error) {

	tx, txErr := s.client.DB.Begin()
	if txErr != nil {
		fmt.Println("Failed to begin transaction", txErr)
		return txErr
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	dbErr := model.CreateUserSession(tx, &session)
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func (s *service) GetUserSession(token string, userId int) (session model.UserSession, err error) {

	tx, txErr := s.client.DB.Begin()
	if txErr != nil {
		fmt.Println("Failed to begin transaction", txErr)
		return session, txErr
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	dbResp, dbErr := model.GetUserSessionByToken(tx, token)
	if dbErr != nil {
		return session, dbErr
	}

	if dbResp == nil {
		return session, errors.New("user session not found")
	}

	session = *dbResp
	return session, nil
}
