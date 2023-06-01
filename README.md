# Theater and Cinema GraphQL System

The Theater and Cinema System is a software application that provides a comprehensive solution for managing theater and cinema operations. It allows users to manage movie screenings, ticket sales, seat reservations, customer data, and other related functionalities.

## Features

- Movie Management: Admins can add, update, and delete movie information, including title, description, duration, genre, and showtimes.
- Screen Management: Admins can manage screens and their seating arrangements, including the number of seats and their availability.
- Showtime Management: Admins can create and schedule movie showtimes for different screens.
- Ticket Sales: Customers can browse available movies and showtimes, select seats, and purchase tickets securely.
- Seat Reservation: Customers can reserve seats for upcoming movie screenings.
- Customer Management: Admins can manage customer information, including personal details, purchase history, and loyalty program status.
- Reporting: Generate reports on ticket sales, movie popularity, and other relevant metrics.

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
git clone https://github.com/BaseMax/CinemaGraphQLAPI.git
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

- Rename the `config.example.json` file to `config.json`. (or `.env` file)
- Modify the values in config file to match your database configuration.

**Run the application:**

```bash
go run main.go
```

**Access the application:**

Open a web browser and visit http://localhost:8080.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please create an issue or submit a pull request.

## License

The Theater and Cinema System is open-source software licensed under the GPL-3.0 license.

## Authors

- ...
- Max Base

Copyright 2023, Max Base
