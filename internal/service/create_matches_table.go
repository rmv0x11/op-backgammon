package service

func (s *Service) CreateMatchesTable() error {
	return s.db.CreateMatchesTable()
}
