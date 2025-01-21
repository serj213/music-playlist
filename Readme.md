## 1. Разработка музыкального плейлиста

### Часть 1. Разработка основного модуля работы с плейлистом

Требуется разработать модуль для обеспечения работы с плейлистом. Модуль должен обладать следующими возможностями:
 - Play - начинает воспроизведение
 - Pause - приостанавливает воспроизведение
 - AddSong - добавляет в конец плейлиста песню
 - Next воспроизвести след песню
 - Prev воспроизвести предыдущую песню


### Часть 2: Построение API для музыкального плейлиста

Реализовать сервис, который позволит управлять музыкальным плейлистом. Доступ к сервису должен осуществляться с помощью API, который имеет возможность выполнять CRUD операции с песнями в плейлисте, а также воспроизводить, приостанавливать, переходить к следующему и предыдущему трекам. Конфигурация может храниться в любом источнике данных, будь то файл на диске, либо база данных. Для удобства интеграции с сервисом может быть реализована клиентская библиотека.

### Технические требования

* реализация задания может быть выполнена на любом языке программирования
* сервис должен обеспечивать персистентность данных
* сервис должен поддерживать все CRUD операции 
* удалять трек допускается только если он не воспроизводится в данный момент
* API должен иметь необходимые методы для взаимодействия с плейлистом.
* API должен возвращать значимые коды ошибок и сообщения в случае ошибок.


### Будет здорово, если:
* в качестве протокола взаимодействия сервиса с клиентами будете использовать gRPC
* напишите Dockerfile и docker-compose.yml
* покроете проект unit-тестами
* сделаете тестовый пример использования написанного сервиса