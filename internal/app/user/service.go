package user

type Service interface {
	Get() ([]User, error)
	FindByID(id string) (User, error)
	Create(input CreateUserRequest) (User, error)
	Delete(id string) error
	Update(id string, input UpdateUserRequest) (User, error)
	Login(input LoginRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Get() ([]User, error) {
	users, err := s.repository.Get()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) FindByID(id string) (User, error) {

	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (s *service) Create(input CreateUserRequest) (User, error) {
	user := User{}

	user.Email = input.Email
	user.Username = input.Username
	user.Password = input.Password

	user, err := s.repository.Create(user)

	if err != nil {
		return user, nil
	}

	return user, nil
}

func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)

	return err
}

func (s *service) Update(id string, input UpdateUserRequest) (User, error) {
	user := User{}
	user.Email = input.Email
	user.Username = input.Username
	user.Password = input.Password

	user, err := s.repository.UpdateByID(id, user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Login(input LoginRequest) (User, error) {
	user := User{}
	user.Username = input.Username
	user.Password = input.Password
	user, err := s.repository.Login(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
