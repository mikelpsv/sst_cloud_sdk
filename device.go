package sst_cloud_sdk

import (
	"encoding/json"
	"time"
)

type Device struct {
	Id                         int64       `json:"id"`
	Configuration              string      `json:"configuration"` // base64
	ParsedConfiguration        string      `json:"parsed_configuration"`
	Timeout                    int         `json:"timeout"`
	TimeSetting                TimeSetting `json:"time_setting"`
	Group                      string      `json:"group"` // undefined type null
	ActiveNetworkId            int64       `json:"active_network"`
	SpecificSettings           interface{} `json:"specific_settings"`
	CreatedAt                  time.Time   `json:"created_at"`
	UpdatedAt                  time.Time   `json:"updated_at"`
	Name                       string      `json:"name"`
	Type                       int         `json:"type"`
	PreviousMode               string      `json:"previous_mode"`
	IsActive                   bool        `json:"is_active"`
	IsConnected                bool        `json:"is_connected"`
	MacAddress                 string      `json:"mac_address"`
	Power                      int         `json:"power"`            // Мощность модуля
	PowerRelayTime             string      `json:"power_relay_time"` // Время рабоы устройства (readonly)
	ChartTemperatureComfort    int         `json:"chart_temperature_comfort"`
	ChartTemperatureEconomical int         `json:"chart_temperature_economical"`
	WirelessSensorsNames       []string    `json:"wireless_sensors_names"` // undefined type - пустой массив имен?
	LineNames                  []string    `json:"line_names"`
	LinesEnable                []string    `json:"lines_enable"`
	HouseId                    int64       `json:"house"`
}

type DevConfThermostatSetttings struct {
	Mode         string `json:"mode"`
	Status       string `json:"status"`
	SelfTraining struct {
		Air        string `json:"air"`
		Floor      string `json:"floor"`
		Status     string `json:"status"`
		OpenWindow string `json:"open_window"`
	} `json:"self_training"`
	TemperatureAir           int `json:"temperature_air"`
	TemperatureManual        int `json:"temperature_manual"`
	TemperatureCorrectionAir int `json:"temperature_correction_air"`
}

type DevConfThermostat struct {
	Detector           int                        `json:"detector"`
	Settings           DevConfThermostatSetttings `json:"settings"`
	ModelId            string                     `json:"device_id"`
	MacAddress         string                     `json:"mac_address"`
	RelayStatus        string                     `json:"relay_status"`
	SignalLevel        int                        `json:"signal_level"`
	AccessStatus       string                     `json:"access_status"`
	CurrentTemperature struct {
		Event            int `json:"event"`
		DayOfWeek        int `json:"day_of_week"`
		TemperatureAir   int `json:"temperature_air"`
		TemperatureFloor int `json:"temperature_floor"`
	} `json:"current_temperature"`
	OpenWindowMinutes int `json:"open_window_minutes"`
}

type TemperatureSet struct {
	TemperatureManual        int `json:"temperature_manual"`
	TemperatureVacation      int `json:"temperature_vacation"`
	TemperatureAir           int `json:"temperature_air"`
	TemperatureLove          int `json:"temperature_love"`
	TemperatureCorrectionAir int `json:"temperature_correction_air"`
	CurrTemperature          int `json:"curr_temperature"`
}

/*
Читает в структуру строку настроек
*/
func (d *Device) ReadConfigThermostat() (*DevConfThermostat, error) {
	res := new(DevConfThermostat)
	err := json.Unmarshal([]byte(d.ParsedConfiguration), &res)
	return res, err
}
