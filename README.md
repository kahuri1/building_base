Из заданий тестового выполнены все эндпоинты:
  1. POST запрос на добавления строения;
  2. GET запрос на получения строения/строений;
  3. Реализован swagger http://localhost:8000/swagger/index.html#/.

Файл с конфигам отправлен сообщением в тг Яне

Конфиг читается с помощью viper

в файл конфига вынесена дополнительная информация, которую лучше в коде не хранить (порт, запросы в sql для фильтра, параметр для фильтрации)


Для запуска:
1. файл config поместить в папку config;
2. поднять docker-compose;
3. запустить main.go)).


по поводу БД:
1. таблица создается автоматически гусем(1 раз с первым запуском);
2. Тестовые данные так же подтягиваются с помощью гуся;
3. если следующие города: Москва, Санкт-Петербург, Казань, Екатеринбург, Нижний Новгород, Ростов-на-Дону, Волгоград, Уфа, Челябинск;
4. дата постройки с 2016 по 2024;
5. этажи до 25;
6. Для значений фильра сделан индекс.

использовал gin для api

Если есть вопросы https://t.me/kahuri1


P.s Так из того что я вижу, нужно исправить поиск или хранение данных в дб. Вкратце Если напишем москва, мы ничего не найдем



