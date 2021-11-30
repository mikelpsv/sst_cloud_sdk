# sst-cloud-sdk

Внимание! Это неофициальная реализация sst-cloud API  
Attention! This is an unofficial implementation of the sst-cloud API


SDK для работы с термостатами теплого пола [Equation Wi-Fi](assets/equation_wifi.jpg) через API SST Cloud (https://sstcloud.ru/).

Модуль разрабатывался для подключения устройств с другими системами умного дома. 

Реализован только необходимый минимум методов для интеграции.
  

**Реализованные методы**

| Метод | Опиcание |
|-------|----------|
| Login | Аутентификация и авторизация в системе. |
| Logout | Выход из системы |
| UserInfo | Получение данных профиля текущего пользователя |
| GetHouses | Получение домовладений текущего пользователя |
| GetHouse(houseId) | Получение домовладения по идентификатору |
| GetDevices | Получение списка устройств домовладения |
| GetDevice(houseId, deviceId) | Получение устройства по его идентификатору |
| SetDeviceStatus(device, const.DEVICE_STATUS_ON/DEVICE_STATUS_OFF) | Включение/выключение устройства. * после экспериментов со сменой статусов устройства потерялись в приложении, но работали и в API остались и управлялись как прежде. Помогла переустановка приложения |
| device.ReadConfigThermostat()| Чтение конфигурации устройства в структуру. Данные текущих температур и состояния хранится в самом объекте устройства, полученные методом `GetDevice`. Метод предназначен для структурирования информации и не требует обращения к API |
