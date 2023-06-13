# Theater and Cinema API System

The Theater and Cinema System is a software application that provides a comprehensive solution for managing theater and cinema operations. It allows users to manage movie screenings, ticket sales, seat reservations, customer data, and other related functionalities.

## Features

**Category Management:**

- Manage categories by adding, editing, and deleting them.
- Retrieve a list of all categories.

**Movie Management:**

- Add new movies to the system.
- Edit movie details such as title, description, and genre.
- Delete movies from the system.
- Activate or deactivate movies for screening in the cinema.
- Retrieve a list of all movies.

**Hall Management:**

- Create new halls with defined seating capacity.
- Edit hall details, such as the number of seats.
- Delete halls from the system.
- Retrieve a list of all halls.

**Movie Schedule Management:**

- Set the date and time for movie screenings.
- Assign movies to specific halls and schedule them accordingly.
- Retrieve the schedule for a particular movie.

These features provide a high-level overview of the functionality offered by the Theater and Cinema System. You can further enhance and expand upon these features based on your specific requirements.

## Technologies Used

- Backend: Go programming language
- Frontend: HTML, CSS, JavaScript
- Database: MySQL
- Framework: Gin (a web framework for Go)
- API Design: RESTful architecture

## Getting Started

**Prerequisites**

- Go (version 1.16 or above) installed
- MySQL database server installed

**Clone the repository:**

```bash
git clone https://github.com/BaseMax/TheaterCinemaAPISystem.git
cd TheaterCinemaAPISystem
```

**Change to the project directory:**

```bash
cd theater-cinema-system
```

**Install the dependencies:**

```go
go mod download
```

Set up the database:

- Create a new MySQL database.
- Import the SQL schema provided in database/schema.sql into the database.

**Configure the application:**

- Rename the `.env_sample` file to `.env`.
- Modify the values in config file to match your database configuration.

**Run the application:**

```bash
go run main.go
```
**Run the test application:**

```bash
cd test &&  go test
```
**Run the migrate application:**

```bash
go run .\migration\migrate.go
```

**Access the application:**

Open a web browser and visit http://localhost:8080.

## API Routes

API routes for your Theater and Cinema System:

### Category Routes

- `GET /categories`: Retrieve all categories.
- `POST /categories`: Add a new category.
- `PUT /categories/{categoryId}`: Edit a category.
- `DELETE /categories/{categoryId}`: Delete a category.

### Movie Routes

- `GET /movies`: Retrieve all movies.
- `POST /movies`: Add a new movie.
- `PUT /movies/{movieId}`: Edit a movie.
- `DELETE /movies/{movieId}`: Delete a movie.
- `PUT /movies/{movieId}/active`: Activate or deactivate a movie in the cinema.

### Hall Routes

- `GET /halls`: Retrieve all halls.
- `POST /halls`: Create a new hall.
- `PUT /halls/{hallId}`: Edit a hall.
- `DELETE /halls/{hallId}`: Delete a hall.

### Movie Schedule Routes

- `POST /movies/{movieId}/schedule`: Set the time and date for a movie.
- `GET /movies/{movieId}/schedule`: Retrieve the schedule for a movie.
### user Routes 

- `POST /auth/login `: Login to Website use jwt.
- `POST /auth/register `: register to Website User.
- `POST /auth/fargetPassword `: fargetpassword user and send email  .


### admin Routes 

- `POST /api/admin ` :group url
- `POST /profile `: show profile user.
- `POST /dashbord `: count all usecse projeact.


In addition to the API routes, you'll need to enforce the following rules:

- Ensure that two movies cannot be scheduled in the same hall at the same time.
- Allow one movie to be scheduled in multiple halls at the same time.

To enforce these rules, you can add validation logic in the respective route handlers for scheduling movies.

Make sure to implement proper authentication and authorization mechanisms to secure your API routes and restrict access as needed.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please create an issue or submit a pull request.

## License

The Theater and Cinema System is open-source software licensed under the GPL-3.0 license.

## Authors

- KhaleghiDev
- Max Base

Copyright 2023, Max Base
