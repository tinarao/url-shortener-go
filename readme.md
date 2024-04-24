# URL Shortener

<hr />

<p align="center">
    <img alt="docker" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-plain.svg" />
    <img alt="go" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" />
    <img alt="postgresql" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-plain.svg" />
</p>

<hr />

| Path             | Method | Возвращает | Body |
|------------------|--|--|--|
| /shorten         | POST | Сгенерированную короткую ссылку | { link: string, alias: string } |
| /get-all         | GET |  Массив со всеми ссылками из базы | |
| /get-one/{alias} | GET  | Информацию о ссылке | |
| /l/{alias}       | GET | Редирект на оригинальную ссылку | |

## TODO
- Swagger