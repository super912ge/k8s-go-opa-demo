package api

type TeamService interface{}

type TeamRepository interface{}

type teamService struct {
	storage TeamRepository
}

func NewTeamService(storage TeamRepository) TeamService {
	return &teamService{
		storage: storage,
	}
}
