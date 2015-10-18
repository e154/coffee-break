Coffee Break v1
------------

Эта небольшая программа позволяет установить таймер, который предупредит Вас, чтобы сделать перерыв 10 минут.
Доступные таймеры: 5м,10м,15м,20м,25м,30м,45м,2ч,3ч,4ч. Присутсвует возможность включить/отключить звуковой сигнал.
Визуальное оповещение о последующей блокировке экрана, и клавиатуры. Блокировка экрана, сопровождается незательевой заставкой, в стиле матрицы.


Прграмма в работе, таймер рабочего времени включен:

<img src="images/tray1.jpg" alt="Coffee Break icon active" width="630" height="88">


Программа в состоянии блокировка, либо ожидается блок через одну минуту, о чем оповещает красный значек:

<img src="images/tray2.jpg" alt="Coffee Break icon inactive" width="630" height="88">


Программа находится в состоянии паузы, все процессы, и таймеры заморожены:

<img src="images/tray3.jpg" alt="Coffee Break icon paused" width="630" height="88">


Быстрый доступ к настройкам через значёк в системном лотке.

<img src="images/menu.jpg" alt="Coffee Break icon paused" width="630" height="358">

Экран блокировки, включающий следующую информацию:
    1) суммарно работы рабочей станции
    2) суммарно простоя рабочей станции
    3) суммарно блокировки
    4) блокировки
    5) суммарно работы
    6) суммарно простоя
    7) обратный отсчёт до разблокировки

matrix

<img src="images/screenshot.png" alt="Coffee Break screenshot" width="630" height="394">

windows bsod screen

<img src="images/screenshot2.png" alt="Coffee Break screenshot2" width="630" height="394">

ide work simulate

<img src="images/screenshot3.png" alt="Coffee Break screenshot3" width="630" height="394">


#### Сборка из исходников

```bash
git clone git@github.com:e154/coffee-break.git
cd coffee-break
make clean && make
```
на выходе получаем deb пакет, который устанавливаем системным пакетным менеджером.

#### Установка deb пакета, debian|ubuntu

```bash
sudo dpkg -i coffee-break*.deb
sudo apt-get install -f
```

#### Удаление

```bash
sudo apt-get remove coffee-break
```

#### Зависимости

```bash
go get golang.org/x/mobile/exp/audio
go get github.com/astaxie/beego/config
go get github.com/gorilla/websocket
go get github.com/c9s/goprocinfo/linux
go get github.com/looplab/fsm
go get github.com/mattn/go-gtk/glib

apt-get install libopenal-dev libnotify-dev libxss-dev md5deep libglib2.0-dev libqt5webkit5-dev
```

#### LICENSE

Coffee Break is licensed under the [MIT License (MIT)](https://github.com/e154/coffee-break/blob/master/LICENSE).