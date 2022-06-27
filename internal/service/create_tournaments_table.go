package service

func (s *Service) CreateTournamentsTable() error {
	return s.db.CreateTournamentTables()
}
