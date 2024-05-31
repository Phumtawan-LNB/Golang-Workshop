package entities

type WeatherService interface {
	WeatherCreate(command WeatherCreateEvent) (Weather, error)
	WeatherSearch(command WeatherSearchsEvent) (weather Weather, err error)
	WeatherUpdate(command WeatherUpdateEvent) (weather []Weather, err error)
	WeatherDelete(command WeatherDeleteEvent) error
}

type WeatherRepository interface {
	Create(user WeatherUserAuth) error
	Save(weather Weather) error
	Update(id string, updatedWeather Weather) error
	UserUpdate(id string, user WeatherUserAuth) error
	Delete(id string) error
	Search(weatherSearch WeatherSearch) (weather Weather, err error)
	CheckId(weatherSearch WeatherSearch) (user WeatherUserAuth, err error)
	FindAll() (weather []Weather, err error)
	WeatherDelete(id string) error
}

type WeatherSearchsEvent struct {
	ID   string
	Name string
}


type WeatherCreateEvent struct {
	Lat  float64
	Long float64
}

type WeatherUpdateEvent struct {
	Key string
}

type WeatherDeleteEvent struct {
	ID string
}

type WeatherUserAuth struct {
	ID   string
	Name string
}

type WeatherSearch struct {
	ID   string
	Name string
}

type WeatherCreate struct {
	ID   string
	Lat  float64
	Long float64
}

type Weather struct {
	ID       string
	Name     string
	Lat      float64
	Long     float64
	Country  string
	Temp     float64
	Wind_dir string
	Wind_kph float64
	UV       float64
}

type Location struct {
	Region    string
	Country   string
	LocalTime string
}

type Current struct {
	Temp_c   float64
	Wind_dir string
	Wind_kph float64
	UV       float64
}

type WeatherResponse struct {
	Weather  Weather
	Location Location
	Current  Current
}
