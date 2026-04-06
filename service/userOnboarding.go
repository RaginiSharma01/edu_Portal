package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
	"smp/utils"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	userRepo *repository.UserRepo
	redis    *redis.Client
}

func NewUserService(repo *repository.UserRepo, redis *redis.Client) *UserService {
	return &UserService{
		userRepo: repo,
		redis:    redis,
	}
}

func (s *UserService) OnboardUsers(ctx context.Context, user models.TeacherOnboarding) (string, error) {

	if user.Email == "" {
		return "", errors.New("email required")
	}

	if user.Password == "" || len(user.Password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword

	// save user in DB
	userID, err := s.userRepo.OnboardingUser(ctx, user)
	if err != nil {
		return "", err
	}

	// generate OTP
	otp := utils.GenerateOTP()

	// store OTP in redis
	err = utils.StoreOTP(ctx, s.redis, user.Email, otp)
	if err != nil {
		return "", err
	}

	// send email
	err = utils.SendOTPEmail(user.Email, otp)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *UserService) VerifyOTP(ctx context.Context, email string, otp string) error {

	storedOTP, err := s.redis.Get(ctx, "otp:"+email).Result()
	if err != nil {
		return errors.New("otp expired or not found")
	}

	if storedOTP != otp {
		return errors.New("invalid otp")
	}

	// delete OTP
	s.redis.Del(ctx, "otp:"+email)

	// verify user in database
	err = s.userRepo.VerifyUser(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
func (s *UserService) Login(ctx context.Context, email string, password string) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// check password
	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// if email not verified
	if !user.IsVerified {

		otp := utils.GenerateOTP()

		// store OTP in redis
		err = utils.StoreOTP(ctx, s.redis, email, otp)
		if err != nil {
			return "", err
		}

		// send email
		err = utils.SendOTPEmail(email, otp)
		if err != nil {
			return "", err
		}

		return "", errors.New("email not verified. OTP sent to your email")
	}

	// generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
