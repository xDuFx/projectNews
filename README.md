# projectNews
Бэк сайта для новостей, реализовано: отдача новостей, добавление новостей с картинкой (кладется в папку image), аутенфикация.
<br>
<br>
GET    /auth/sign-up             --> Регистрация переход на сайт <br>
GET    /auth/sign-in             --> Аутенфикация <br>
POST   /auth/sign-up             --> Отправка json с данными для регистрации <br>
POST   /auth/sign-in             --> Отправка json с данными для входа, в ответ приходит токен <br>
GET    /api/:id                  --> Получение новости по id <br>
POST   /api/addnews              --> Добавление новости <br>
GET    /api/addnews              --> Переход на сайт для добавления новости <br>
DELETE /api/del_news             --> Удаление <br>
PUT    /api/update_news          --> Обновление <br>
