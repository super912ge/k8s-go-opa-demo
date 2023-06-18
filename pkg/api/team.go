package api

type TeamService interface {
	CreateTeam(NewTeamRequest) error
	CreateTeamMember(NewTeamMemberRequest) error
	UpdateManager(UpdateManagerRequest) error
	GetTeam(GetTeamRequest) (Team, error)
}

type TeamRepository interface {
	CreateTeam(NewTeamRequest) error
	CreateTeamMember(NewTeamMemberRequest) error
	UpdateManager(UpdateManagerRequest) error
	GetTeam(GetTeamRequest) (Team, error)
}

type teamService struct {
	storage TeamRepository
}

func NewTeamService(storage TeamRepository) TeamService {
	return &teamService{
		storage: storage,
	}
}

func (s *teamService) CreateTeam(team NewTeamRequest) error {
	err := s.storage.CreateTeam(team)
	if err != nil {
		return err
	}
	return nil
}

func (s *teamService) CreateTeamMember(teamMember NewTeamMemberRequest) error {
	err := s.storage.CreateTeamMember(teamMember)
	if err != nil {
		return err
	}
	return nil
}

func (s *teamService) UpdateManager(manager UpdateManagerRequest) error {
	err := s.storage.UpdateManager(manager)
	if err != nil {
		return err
	}
	return nil
}

func (s *teamService) GetTeam(req GetTeamRequest) (Team, error) {
	team, err := s.storage.GetTeam(req)
	if err != nil {
		return Team{}, err
	}
	return team, nil
}
