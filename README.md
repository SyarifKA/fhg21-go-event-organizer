<h2>Project Description</h2>

Let me introduce my awesome online-based ticket booking, KuyTick. This repository utilized PGX for package and Gin-gonic as framework API for development.

Tickhub is an easy way to order an event from a distance place. Without order on the spot, we can choose the event as we want with the simple process and various payment. Also, we can save the listed event as the wishlist.

<h2>Tech Stack</h2>

- Backend's Programming Language: Go
- Framework : Gin-Gonic
- Package Manager : Golang PGX
- Data Migration : Golang Migrate
- RDBMS : PostgreSQL
- API Testing : ThunderClient
- Containerization : Docker

<h2>Config / installation process</h2>

<h3>1. Clone this repository</h3>

```sh
  git clone https://github.com/SyarifKA/fgh21-go-event-organizer.git
  cd <project-name>
```

<h3>2. Open in VSCode</h3>

```sh
  code .
```

<h3>3. Install all the dependencies</h3>

```sh
  go mod tidy
```

<h3>4. Run the program</h3>

```sh
  go run main.go
```

<h2>API References</h2>

## Login

```http
  POST auth/login
```

## Register

```http
  POST auth/register
```

| Parameter               | Type     | Description                                            |
| :---------------------- | :------- | :----------------------------------------------------- |
| `users`                 | `GET`    | `Get a list of users data`                             |
| `users/:id`             | `GET`    | `Select the user data according to registered id`      |
| `users`                 | `POST`   | `Create new user data`                                 |
| `users/:id`             | `PATCH`  | `Edit the selected user data`                          |
| `users/:id`             | `DELETE` | `Remove the selected user data`                        |
| `users/change-password` | `PATCH`  | `Change the user's password`                           |
| `events`                | `GET`    | `Get a list of events data`                            |
| `events/:id`            | `GET`    | `Select the event data according to registered id`     |
| `events/`               | `POST`   | `Create new event`                                     |
| `events/:id`            | `PATCH`  | `Edit the selected event data`                         |
| `events/:id`            | `DELETE` | `Remove the selected event data`                       |
| `events/payment_method` | `GET`    | `Get a list of payment methods data`                   |
| `events/section/:id`    | `GET`    | `Get a list of event sections data`                    |
| `events/section`        | `POST`   | `Create new event sections`                            |
| `transactions`          | `POST`   | `Create new transactions`                              |
| `transactions`          | `GET`    | `Get a list of transactions by registered user`        |
| `profile`               | `GET`    | `Select the profile data according to registered user` |
| `profile`               | `PATCH`  | `Change the profile data from registered user`         |
| `profile/upload-img`    | `PATCH`  | `Change the profile's image  from registered user`     |
| `categories`            | `GET`    | `Get a list of categories data`                        |
| `categories/:id`        | `GET`    | `Select the category data according to registered id`  |
| `categories`            | `POST`   | `Create new category data`                             |
| `categories/:id`        | `PATCH`  | `Edit the selected category data`                      |
| `categories/:id`        | `DELETE` | `Remove the selected category data`                    |
| `locations`             | `GET`    | `Get a list of locations data`                         |
| `nationalities`         | `GET`    | `Get a list of nationalities data`                     |
| `partners`              | `GET`    | `Get a list of partners data`                          |
| `wishlist`              | `GET`    | `Get a list of wishlist data`                          |
| `wishlist`              | `POST`   | `Create new wishlist data`                             |
| `wishlist/:id`          | `DELETE` | `Remove the selected wishlist data`                    |

## Contributing

Feel free to contribute the repo for better code!

## Authors

- Me

## Feedback

If you have any feedback, please reach out to us at syarif.khalid@gmail.com