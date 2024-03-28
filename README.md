# VK-TEST

## API Routes

### Authentication

- `POST /signup`: Register a new user.
- `POST /signin`: Login with existing credentials.
- `POST /refresh-token`: Refresh JWT token.

### Ads Management

- `POST /create-ad`: Create a new ad.
- `GET /all-ads`: Get all ads
- `GET /ads-by-page`: Get ads on the page (written param in the query)

#### Examples

```
- localhost:8080/all-ads?date_order=desc

- localhost:8080/all-ads?max_price=1000&price_order=asc

- localhost:8080/all-ads?max_price=1000&price_order=asc&min_price=300

- localhost:8080/ads-by-page?price_order=desc&page=2

- localhost:8080/ads-by-page?date_order=asc&page=2
```

## Running the Application

To run the application, execute the following command in your terminal:

```
go run ./cmd
```

To register use the following json pattern

```json
{
  "email": "chingizkhan@gmail.com",
  "password": "useMeForLoginAlso!"
}
```

In order to add an ad you should be authorized, and if so, use the following json pattern

```json
{
  "title": "Продается дача",
  "text": "Дача в 100 км от города. 6 соток, дом 80 м2.",
  "image_url": "https://www.example.com/dacha.jpg",
  "price": 1500000.0
}
```

Make sure to set the necessary environment variables in a `.env` file before running the application.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework.
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver): Official MongoDB driver for Go.
- [JWT Go](https://github.com/dgrijalva/jwt-go): Library for JSON Web Tokens.

## Environment Variables

The application requires the following environment variables to be set:

- `MONGO_URI`: MongoDB connection URI.
- `PORT`: Port number for the server.
- `JWT_SECRET`: Secret key for JWT token generation.
- `GIN_MODE`: Gin mode (`debug` or `release`).

Ensure these environment variables are correctly set in a `.env` file before running the application.
