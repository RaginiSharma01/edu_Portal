package service

import (
	"context"
	"errors"
	"fmt"
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

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword

	userID, err := s.userRepo.OnboardingUser(ctx, user)
	if err != nil {
		return "", err
	}

	otp := utils.GenerateOTP()

	err = utils.StoreOTP(ctx, s.redis, user.Email, otp)
	if err != nil {
		return "", err
	}

	err = utils.SendOTPEmail(user.Email, otp)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// THIS WAS THE BUG — function name was missing
func (s *UserService) VerifyOTP(ctx context.Context, email string, otp string) error {

	// must match exactly how StoreOTP saves it
	fmt.Println("Redis key:", "otp:"+email)
	storedOTP, err := s.redis.Get(ctx, "otp:"+email).Result()
	fmt.Println("storedOTP", storedOTP)
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("err redis: ", err)
		return errors.New("otp expired or not found")
	}

	if storedOTP != otp {
		return errors.New("invalid otp")
	}

	s.redis.Del(ctx, "otp:"+email)

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

	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.IsBlocked {
		return "", errors.New("your account has been blocked. contact admin")
	}

	if !user.IsVerified && user.Role != "admin" {

		otp := utils.GenerateOTP()

		err = utils.StoreOTP(ctx, s.redis, email, otp)
		if err != nil {
			fmt.Println("err", err)
			return "", err
		}

		storedOTP, err := s.redis.Get(ctx, "otp:"+email).Result()
		if err != nil {
			fmt.Println("Redis err", err)
			return "", err
		}

		fmt.Println("Redis storedOTP", storedOTP)

		err = utils.SendOTPEmail(email, otp)
		if err != nil {
			return "", err
		}

		return "", errors.New("email not verified. OTP sent to your email")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) OnboardStudent(ctx context.Context, student models.StudentOnboarding) (string, error) {

	if student.Email == "" {
		return "", errors.New("email required")
	}

	if student.Password == "" || len(student.Password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	hashedPassword, err := utils.HashPassword(student.Password)
	if err != nil {
		return "", err
	}

	student.Password = hashedPassword

	userID, err := s.userRepo.OnboardStudent(ctx, student)
	if err != nil {
		return "", err
	}

	otp := utils.GenerateOTP()

	err = utils.StoreOTP(ctx, s.redis, student.Email, otp)
	if err != nil {
		return "", err
	}

	err = utils.SendOTPEmail(student.Email, otp)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *UserService) GetAllTeachers(ctx context.Context) (interface{}, error) {

	rows, err := s.userRepo.GetAllTeachers(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []map[string]interface{}

	for rows.Next() {
		var (
			id, firstName, lastName, email, phone, dob, address string
			qualification, subjects                             string
			age                                                 int
		)

		err := rows.Scan(
			&id,
			&firstName,
			&lastName,
			&email,
			&phone,
			&age,
			&dob,
			&address,
			&qualification,
			&subjects,
		)

		if err != nil {
			return nil, err
		}

		teacher := map[string]interface{}{
			"id":               id,
			"firstName":        firstName,
			"lastName":         lastName,
			"email":            email,
			"phone":            phone,
			"age":              age,
			"dob":              dob,
			"address":          address,
			"qualification":    qualification,
			"subjectsTeaching": subjects,
		}

		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (s *UserService) GetAllStudents(ctx context.Context) (interface{}, error) {

	rows, err := s.userRepo.GetAllStudents(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []map[string]interface{}

	for rows.Next() {
		var (
			id, firstName, lastName, email, phone, dob, address string
			father, mother, guardian, occupation                string
			age, height, weight                                 int
		)

		err := rows.Scan(
			&id,
			&firstName,
			&lastName,
			&email,
			&phone,
			&age,
			&dob,
			&address,
			&father,
			&mother,
			&guardian,
			&occupation,
			&height,
			&weight,
		)

		if err != nil {
			return nil, err
		}

		student := map[string]interface{}{
			"id":           id,
			"firstName":    firstName,
			"lastName":     lastName,
			"email":        email,
			"phone":        phone,
			"age":          age,
			"dob":          dob,
			"address":      address,
			"fatherName":   father,
			"motherName":   mother,
			"guardianName": guardian,
			"occupation":   occupation,
			"height":       height,
			"weight":       weight,
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *UserService) DeleteStudent(ctx context.Context, studentID string) error {

	if studentID == "" {
		return errors.New("student ID required")
	}

	return s.userRepo.DeleteStudent(ctx, studentID)
}

func (s *UserService) UpdateStudent(ctx context.Context, studentID string, data models.UpdateStudent) error {

	if studentID == "" {
		return errors.New("student ID required")
	}

	if data.FirstName == "" || data.LastName == "" {
		return errors.New("first name and last name are required")
	}

	return s.userRepo.UpdateStudent(ctx, studentID, data)
}

func (s *UserService) BlockStudent(ctx context.Context, studentID string, block bool) error {

	if studentID == "" {
		return errors.New("student ID required")
	}

	return s.userRepo.BlockStudent(ctx, studentID, block)
}

func (s *UserService) DeleteTeacher(ctx context.Context, teacherID string) error {

	if teacherID == "" {
		return errors.New("teacher ID required")
	}

	return s.userRepo.DeleteTeacher(ctx, teacherID)
}

func (s *UserService) UpdateTeacher(ctx context.Context, teacherID string, data models.UpdateTeacher) error {

	if teacherID == "" {
		return errors.New("teacher ID required")
	}

	if data.FirstName == "" || data.LastName == "" {
		return errors.New("first name and last name are required")
	}

	return s.userRepo.UpdateTeacher(ctx, teacherID, data)
}

// forget and reset password

func (s *UserService) ForgotPassword(ctx context.Context, email string) error {

	if email == "" {
		return errors.New("email is required")
	}

	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil // silent — don't reveal if email exists
	}

	otp := utils.GenerateOTP()

	err = utils.StoreOTP(ctx, s.redis, "reset:"+email, otp)
	if err != nil {
		return err
	}

	return utils.SendOTPEmail(email, otp)
}

func (s *UserService) ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error {

	if req.Email == "" || req.OTP == "" || req.NewPassword == "" {
		return errors.New("email, otp and new password are required")
	}

	if len(req.NewPassword) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	storedOTP, err := s.redis.Get(ctx, "otp:reset:"+req.Email).Result()
	if err != nil {
		return errors.New("otp expired or not found")
	}

	if storedOTP != req.OTP {
		return errors.New("invalid otp")
	}

	s.redis.Del(ctx, "otp:reset:"+req.Email)

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(ctx, req.Email, hashedPassword)
}