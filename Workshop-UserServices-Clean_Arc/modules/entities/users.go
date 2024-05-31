package entities

type UserService interface {
	UserCreated(command *UserCreatedCommand) (id string, err error)
	UserReaded(command *UserReadedCommand) (history []UserHistory, user *User, err error)
	UserUpdate(command *UserUpdateCommand) (user *User, err error)
	UserDelete(command *UserDeleteCommand) (id string, err error)
}

type UserRepository interface {
	Create(user *User) error
	Save(history *UserHistory) error
	Readed(id string) (history []UserHistory, err error)
	Update(id string, user *User) (userResponse *User, err error)
	UpdateHistory(id string, weather_id string, history *UserHistory) error
	FindHistory(id string, weather_id string) (history *UserHistory, err error)
	FindById(id string) (user *User, err error)
	Delete(id string) error
}

type WeatherData struct {
	Weather_id string                 `json:"weather_id"`
	Details    map[string]interface{} `json:"details"`
}

type UserCreatedCommand struct {
	ID         string
	Name       string
	Lat        float64
	Long       float64
	Age        int
	First_name string
	Last_name  string
}

type UserReadedCommand struct {
	User_id      string
	Weather_id   string
	User_name    string
	Weather_name string
}

type UserUpdateCommand struct {
	ID         string
	Name       string
	Lat        float64
	Long       float64
	Age        int
	First_name string
	Last_name  string
}

type UserDeleteCommand struct {
	ID string
}

type User struct {
	ID         string
	Name       string
	Lat        float64
	Long       float64
	Age        int
	First_name string
	Last_name  string
}

type UserHistory struct {
	User_id      string
	Weather_id   string
	User_name    string
	Weather_name string
	Quantity     int
}
