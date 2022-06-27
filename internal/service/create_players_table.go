package service

func (s *Service) CreatePlayersTable() error {
	return s.db.CreatePlayersTable()
}
