package service

func (s *Service) NewTournament(IDs string) (int64, error) {
	return s.db.NewTournament(IDs)
}
