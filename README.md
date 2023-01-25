<a name="readme-top"></a>

[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<br />
<div align="center">

<h3 align="center">📅 Date Difference Calculator 📅</h3>

  <p align="center">
    Калькулятор разницы в днях между текущей датой и 1 января заданного года
  </p>
</div>



<details>
  <summary>Содержание</summary>
  <ol>
    <li>
        <a href="#about">О проекте</a>
    </li>
    <li>
        <a href="#built-with">Используемые технологии</a>
    </li>
    <li>
        <a href="#start">Запуск</a>
      <ul>
        <li><a href="#prerequisites">Требования</a></li>
      </ul>
    </li>
    <li><a href="#contacts">Контакты</a></li>
  </ol>
</details>



<a name="about"></a>
## О проекте

### Часть 1 - приложение на фреймворке gin

Разработайте веб-приложение на фреймворке gin, которое содержит один маршрут `/when/:year`. В обработчике этого маршрута вам нужно вывести строку, которая покажет, сколько дней осталось или прошло до 1-го января года, указанного в параметре маршрута,

Например: `/when/2000`
Должно вывести сколько дней прошло с 01.01.2000 до сегодня, в формате: `Days gone: 1234`

Например: `/when/2030`
Должно вывести сколько дней осталось до 01.01.2030, в формате: `Days left: 5432`

Статус HTTP-ответа должен быть 200.

#### Обработка ошибок

Ошибки работы приложения должны попадать в консольный лог.

#### Запуск приложения

Для запуска приложения создайте Makefile с двумя командами:
- `run` - запускает приложение
- `build` - компилирует приложение

### Часть 2 - middleware на фреймворке gin

Создайте свой собственный middleware в этом приложении, который проверяет наличие HTTP-заголовка. Если он содержит заголовок `X-PING` со значением `ping`, то в ответе вашего сервиса добавьте заголовок ответа `X-PONG` со значением `pong`.

### Необязательная Часть 3

Постарайтесь организовать структуру и код проекта, следуя лучшим практикам организации кода приложения и структуры приложения на языке Go.

<p align="right">(<a href="#readme-top">в начало</a>)</p>



<a name="built-with"></a>
## Используемые технологии

* [![Go][Go]][Go-url]
* [![Docker][Docker]][Docker-url]

<p align="right">(<a href="#readme-top">в начало</a>)</p>



<a name="start"></a>
## Запуск

```bash
git clone https://github.com/tanya-lyubimaya/date_difference_calculator && make run
```

### Или через docker:
```bash
git clone https://github.com/tanya-lyubimaya/date_difference_calculator && make image
```

<a name="prerequisites"></a>
### Требования

* go-1.20rc
* make

<p align="right">(<a href="#readme-top">в начало</a>)</p>



<a name="contacts"></a>
## Контакты
<div align="center">
    <a href="https://t.me/tatyana_lyubimaya">
        <img src="https://img.shields.io/badge/-telegram-2CA5E0?style=for-the-badge&logo=telegram&labelColor=white">
    </a>
    <a href="mailto:tatyana.lyubimaya.work@gmail.com">
        <img src="https://img.shields.io/badge/-gmail-EA4335?style=for-the-badge&logo=gmail&labelColor=white">
    </a>
</div>

<p align="right">(<a href="#readme-top">в начало</a>)</p>



[stars-shield]: https://img.shields.io/github/stars/tanya-lyubimaya/date_difference_calculator.svg?style=for-the-badge
[stars-url]: https://github.com/tanya-lyubimaya/date_difference_calculator/stargazers
[issues-shield]: https://img.shields.io/github/issues/tanya-lyubimaya/date_difference_calculator.svg?style=for-the-badge
[issues-url]: https://github.com/tanya-lyubimaya/date_difference_calculator/issues
[product-screenshot]: images/screenshot.png
[Go]: https://img.shields.io/badge/-go-00ADD8?style=for-the-badge&logo=go&labelColor=white
[Go-url]: https://go.dev/
[Gin]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[Gin-url]: https://gin-gonic.com/
[Docker]: https://img.shields.io/badge/-docker-2496ED?style=for-the-badge&logo=docker&labelColor=white
[Docker-url]: https://www.docker.com