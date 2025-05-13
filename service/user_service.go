package service

import (
	"gin-api-blog/api/dto"
	"gin-api-blog/config"
	"gin-api-blog/data/db"
	"gin-api-blog/data/models"
	"gin-api-blog/pkg/logging"
	"gin-api-blog/pkg/service_errors"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	tokenService *TokenService
	db           *sqlx.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:          cfg,
		logger:       logger,
		tokenService: NewTokenService(cfg),
		db:           db.GetDB(),
	}
}

func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	query := `SELECT id, username, email, password FROM users WHERE username = $1`
	err := s.db.QueryRowx(query, req.Username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	tdto := tokenDto{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Name: req.Name, Username: req.Username, Email: req.Email}

	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)

	tx := s.db.MustBegin()
	query := `INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(query, u.Name, u.Username, u.Email, u.Password)
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	query := `SELECT count(*) > 0 FROM users WHERE email = $1`
	if err := s.db.QueryRowx(query, email).Scan(&exists); err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	query := `SELECT count(*) > 0 FROM users WHERE username = $1`
	if err := s.db.QueryRowx(query, username).Scan(&exists); err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}
