package sst_cloud_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const LANG_RU = "ru"

const (
	DEVICE_TYPE_MC300  = 0
	DEVICE_TYPE_MC350  = 1
	DEVICE_TYPE_NEPTUN = 2

	DEVICE_STATUS_ON  = "on"
	DEVICE_STATUS_OFF = "off"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Language string `json:"language"`
}

type LoginResponse struct {
	Key string
}

type Session struct {
	Id         string
	Key        string
	Cookies    []*http.Cookie
}

type UserProfile struct {
	Id        int64  `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Language  string `json:"language"`
	UserId    int64  `json:"user"`
}

type User struct {
	Pk       int64       `json:"pk"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Profile  UserProfile `json:"profile"`
}

type House struct {
	Id       int64    `json:"id"`
	Owner    string   `json:"owner"`
	Workdays Workdays `json:"workdays"`
}

type Workdays struct {
	Id             int64     `json:"id"`
	CurrentDay     string    `json:"current_day"`
	WorkdaysCount  int       `json:"workdays_count"`
	VacationsCount int       `json:"vacations_count"`
	CurrentWeek    int       `json:"current_week"`
	IsCustom       bool      `json:"is_custom"`
	Vacations      []int     `json:"vacations"`
	StartDate      string    `json:"start_date"`    // undefined type (null)
	NextWorkday    string    `json:"next_workday"`  // undefined type (null)
	NextVacation   string    `json:"next_vacation"` // undefined type (null)
	StartDay       int       `json:"start_day"`
	HouseId        int64     `json:"house"`
	Timezone       string    `json:"timezone"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Uid            string    `json:"uid"`
	Name           string    `json:"name"`
	InHome         bool      `json:"in_home"`
	Behaviour      string    `json:"behaviour"`
	CloseValves    int       `json:"close_valves"`
	ReportDate     int       `json:"report_date"`
	UserIds        []int64   `json:"users"`
}

type TimeSetting struct {
	Id                int64 `json:"id"`
	WorkdayTimeRange  [][]string
	VacationTimeRange [][]string
	DeviceId          int64 `json:"device"`
}

const API_ENDPOINT = "https://api.sst-cloud.com"

/*
	Авторизация в системе
 	POST /auth/user/
    @link https://api.sst-cloud.com/docs/#/auth/login_create
*/
func (s *Session) Login(authUser LoginRequest) (*LoginResponse, error) {
	if authUser.Language == "" {
		authUser.Language = LANG_RU
	}
	respLogin := new(LoginResponse)
	jsonData, err := json.Marshal(authUser)
	if err != nil {
		return nil, err
	}
	bodyReq := bytes.NewReader(jsonData)
	endpoint := fmt.Sprintf("%s/auth/login/", API_ENDPOINT)
	body, err := s.DoRequest("POST", endpoint, bodyReq)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respLogin)
	if err != nil {
		return nil, err
	}
	s.Key = respLogin.Key
	return respLogin, nil
}

/*
	Информация о текущем пользователе
 	GET /auth/user/
    @link https://api.sst-cloud.com/docs/#/auth/user_list
*/
func (s *Session) UserInfo() (*User, error) {
	respUser := new(User)
	endpoint := fmt.Sprintf("%s/auth/user/", API_ENDPOINT)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respUser)
	if err != nil {
		return nil, err
	}
	return respUser, nil
}

/*
	Выход из системы
 	POST /auth/logout/
    @link https://api.sst-cloud.com/docs/#/auth/logout_create
*/
func (s *Session) Logout() (string, error) {
	emptyBody := bytes.NewReader([]byte(``))
	endpoint := fmt.Sprintf("%s/auth/logout/", API_ENDPOINT)
	body, err := s.DoRequest("POST", endpoint, emptyBody)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
	Список домов пользователя
 	POST /houses/
    @link https://api.sst-cloud.com/docs/#/houses/list
*/
func (s *Session) GetHouses() ([]House, error) {
	respHouses := make([]House, 0)
	endpoint := fmt.Sprintf("%s/houses/", API_ENDPOINT)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respHouses)
	if err != nil {
		return nil, err
	}
	return respHouses, nil

}

/*
	Информация о доме по его идентификатору
 	GET /houses/{houseId}/
	@link https://api.sst-cloud.com/docs/#/houses/read
*/
func (s *Session) GetHouse(houseId int64) (*House, error) {
	respHouse := House{}
	endpoint := fmt.Sprintf("%s/houses/%d", API_ENDPOINT, houseId)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respHouse)
	if err != nil {
		return nil, err
	}
	return &respHouse, nil
}

/*
	Список сетей в доме
	GET /houses/{houseId}/networks/
	@link https://api.sst-cloud.com/docs/#/networks/networks_list
*/
// TODO: func (s *Session) GetNetworks(houseId int64) error

/*
	Информация о сети по ее идентификатору
	GET /houses/{houseId}/networks/{networkId}/
 	@link https://api.sst-cloud.com/docs/#/networks/networks_read
*/
// TODO: func (s *Session) GetNetwork(houseId int64, networkId int64) error

/*
	Список устройств в доме
	GET /houses/{houseId}/devices/
	@link https://api.sst-cloud.com/docs/#/devices/devices_list
*/
func (s *Session) GetDevices(houseId int64) ([]Device, error) {
	respDevs := make([]Device, 0)
	endpoint := fmt.Sprintf("%s/houses/%d/devices/", API_ENDPOINT, houseId)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respDevs)
	if err != nil {
		return nil, err
	}
	return respDevs, nil
}

/*
	Информация об устройстве по его идентификатору
	GET /houses/{houseId}/devices/{id}/
	@link https://api.sst-cloud.com/docs/#/devices/devices_read
*/

func (s *Session) GetDevice(houseId int64, deviceId int64) (*Device, error) {
	respDev := new(Device)
	endpoint := fmt.Sprintf("%s/houses/%d/devices/%d", API_ENDPOINT, houseId, deviceId)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respDev)
	if err != nil {
		return nil, err
	}
	return respDev, nil
}

func (s *Session) SetDeviceStatus(d *Device, devStatus string) {
	st := struct {
		Status string `json:"status"`
	}{Status: devStatus}

	status, err := json.Marshal(st)
	if err != nil {
		return
	}
	bodyReq := bytes.NewReader(status)

	respDev := new(Device)
	endpoint := fmt.Sprintf("%s/houses/%d/devices/%d/status/", API_ENDPOINT, d.HouseId, d.Id)
	body, err := s.DoRequest("POST", endpoint, bodyReq)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &respDev)
	if err != nil {
		return
	}
	return
}

/*
	Список беспроводных датчиков, зарегистрированных в устройстве
	GET /houses/{houseId}/wireless_sensors/
	@link https://api.sst-cloud.com/docs/#/devices/devices_wsensors_read
*/

/*
	Информация о счетчиках, зарегистрированных на устройстве
	GET /houses/{houseId}/devices/{deviceId}/counters/
	@link https://api.sst-cloud.com/docs/#/devices/devices_counters_read
*/

/*
	Информация об устройстве по его идентификатору
	GET /houses/{houseId}/devices/{id}/
	@link https://api.sst-cloud.com/docs/#/devices/devices_read
*/

func (s *Session) SetThemperature(houseId int64, deviceId int64) (*Device, error) {
	respDev := new(Device)
	endpoint := fmt.Sprintf("%s/houses/%d/devices/%d", API_ENDPOINT, houseId, deviceId)
	body, err := s.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &respDev)
	if err != nil {
		return nil, err
	}
	return respDev, nil
}

func (s *Session) DoRequest(method string, endpoint string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header["Accept"] = []string{"*/*"}
	req.Header["Content-Type"] = []string{"application/json"}

	s.SetCookies(req)
	s.SetCSRFToken(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	s.GetCookies(resp)
	return ioutil.ReadAll(resp.Body)
}

func (s *Session) SetCookies(r *http.Request) {
	for _, c := range s.Cookies {
		r.AddCookie(c)
	}
}

func (s *Session) GetCookies(r *http.Response) {
	if len(r.Cookies()) > 0 {
		s.Cookies = r.Cookies()
	}
}

func (s *Session) SetCSRFToken(r *http.Request) {
	for _, c := range s.Cookies {
		if c.Name == "csrftoken" {
			r.Header["X-CSRFToken"] = []string{c.Value}
		}
	}
}
