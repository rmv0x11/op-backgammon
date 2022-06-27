package service

func (s *Service) CreateRoundsTable() error {
	return s.db.CreateRoundsTable()
}
