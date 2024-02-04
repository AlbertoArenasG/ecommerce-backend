# ecommerce-backend
Backend of an e-commerce application built on Golang with PostgreSQL and Fiber.

## Requirements

- Docker
- Docker Compose

## Environment Setup

1. Clone this repository to your local machine:

```bash
git clone https://github.com/AlbertoArenasG/ecommerce-backend.git
cd ecommerce-backend
```

2. Create a `.env` file based on the provided example:

```bash
cp .env.example .env
```

3. Edit the `.env` file and configure the environment variables as needed, such as the database connection.

## Running the Project

To run the project, simply execute the following command:

```bash
docker-compose up --build
```

The project will run at http://localhost:3000.

## Available Endpoints

- **GET /products**: Retrieve a list of all products.
- **GET /products/{id}**: Retrieve details of a specific product.
- **POST /products**: Add a new product.
- **PUT /products/{id}**: Edit an existing product.
- **DELETE /products/{id}**: Delete an existing product.
- **GET /shopping-carts/{id}**: Retrieve a shopping cart's contents.
- **POST /shopping-carts**: Add a new shopping cart.
- **PUT /shopping-carts/items**: Add an item to shopping cart.
- **DELETE /shopping-carts/items/{cartId}/{productId}**: Remove an item from shopping cart.
