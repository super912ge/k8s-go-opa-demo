package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"k8s-go-opa/pkg/api"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate(connectionString string) error
	CreateUser(req api.NewUserRequest) error
	CreateTeam(req api.NewTeamRequest) error
	CreateTeamMember(req api.NewTeamMemberRequest) error
	UpdateManager(req api.UpdateManagerRequest) error
	GetTeam(req api.GetTeamRequest) (api.Team, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) Migrate(connectionString string) error {

	p, err := filepath.Abs("../..")
	p = filepath.ToSlash(p)
	p = filepath.Join(p, "pkg/repository/migrations/")

	migrationsPath := fmt.Sprintf("file:%s", p)

	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}

func (s *storage) CreateUser(req api.NewUserRequest) error {
	newUserStatement := `INSERT INTO "user" (name, email) VALUES ($1, $2);`
	err := s.db.QueryRow(newUserStatement, req.Name, req.Email).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *storage) CreateTeam(req api.NewTeamRequest) error {
	newTeamStatement := `INSERT INTO team (name, description) VALUES ($1, $2);`
	err := s.db.QueryRow(newTeamStatement, req.Name, req.Description).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *storage) CreateTeamMember(req api.NewTeamMemberRequest) error {
	newTeamMemberStatement := `INSERT INTO user_team (team_id, user_id) VALUES ($1, $2);`
	err := s.db.QueryRow(newTeamMemberStatement, req.TeamID, req.UserID).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *storage) UpdateManager(req api.UpdateManagerRequest) error {
	updateManagerStatement := `UPDATE team SET manager_id = $1 WHERE id = $2;`
	err := s.db.QueryRow(updateManagerStatement, req.UserID, req.TeamID).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *storage) GetTeam(req api.GetTeamRequest) (api.Team, error) {
	queryTeamStatement := `SELECT t.id, t.name, t.description, u.id as manager_id, u.name as manager_name, u2.id as member_id, u2.name as member_name FROM team t INNER JOIN "user" u 
	ON t.manager_id = u.id INNER JOIN user_team ut ON t.id = ut.team_id INNER JOIN "user" u2 ON ut.user_id = u2.id WHERE t.id = $1;`
	rows, err := s.db.Query(queryTeamStatement, req.TeamID)
	if err != nil {
		return api.Team{}, err
	}
	var team api.Team
	var manager api.User
	for rows.Next() {
		//var ignore any
		var member api.User
		err = rows.Scan(&team.ID, &team.Name, &team.Description, &manager.ID, &manager.Name, &member.ID, &member.Name)
		if err != nil {
			return api.Team{}, err
		}
		team.Manager = manager
		team.Members = append(team.Members, member)
	}
	return team, nil
}
