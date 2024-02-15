package main

import (
	"fmt"
	sst "github.com/mikelpsv/sst_cloud_sdk"
)

const (
	USERNAME  = ""
	PASSWORD  = ""
	YOUR_MAIL = ""
)

func main() {
	s := new(sst.Session)

	// Авторизация
	_, err := s.Login(sst.LoginRequest{Username: USERNAME, Password: PASSWORD, Email: YOUR_MAIL})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Logout()

	// Чтение профиля пользователя
	ui, err := s.UserInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ui)

	// Чтение домовладений текущего аккаунта
	houses, err := s.GetHouses()
	if err != nil {
		fmt.Println(err)
	}

	for _, h := range houses {
		// домовладение по id
		house, err := s.GetHouse(h.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(house)

		// устройства
		devs, err := s.GetDevices(h.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, d := range devs {
			dev, err := s.GetDevice(h.Id, d.Id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			conf, _ := dev.ReadConfigThermostat()
			fmt.Println(dev, conf)

			//s.SetDeviceStatus(&d, sst.DEVICE_STATUS_ON)
		}
	}

}
