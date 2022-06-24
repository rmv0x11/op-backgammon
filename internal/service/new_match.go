package service

import "log"

func (s *Service) NewMatch(playerOneID, playerTwoID int) error {

	playerOne, err := s.db.GetPlayerInfo(playerOneID)
	if err != nil {
		log.Fatalln("unable create new match, error", err.Error())
		return err
	}

	playerTwo, err := s.db.GetPlayerInfo(playerTwoID)
	if err != nil {
		log.Fatalln("unable create new match, error", err.Error())
		return err
	}

	//TODO write main logic of calculation prematch stats from players info

	return nil
}
