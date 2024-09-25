# projectNews
Бэк сайта для новостей, реализовано: отдача новостей, добавление новостей с картинкой (кладется в папку image), аутенфикация.
GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (1 handlers)

GET    /auth/sign-up             --> Регистрация переход на сайт /n
GET    /auth/sign-in             --> Аутенфикация
POST   /auth/sign-up             --> Отправка json с данными для регистрации
POST   /auth/sign-in             --> Отправка json с данными для входа, в ответ приходит токен
GET    /api/:id                  --> Получение новости по id
POST   /api/addnews              --> Добавление новости
GET    /api/addnews              --> Переход на сайт для добавления новости
DELETE /api/del_news             --> Удаление
PUT    /api/update_news          --> Обновление
