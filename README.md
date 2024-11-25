# Order server

This project is a gRPC server designed for managing orders with full CRUD (Create, Read, Update, Delete) functionality. It supports RESTful APIs through gRPC-Gateway and integrates with PostgreSQL for persistent storage and Redis for caching. The application is containerized using Docker and designed to run in multiple instances, with NGINX handling load balancing.

---

## Run

1. **Clone the Repository**:

   ```bash
   git clone git@github.com:LLIEPJIOK/orderserver.git
   ```

2. **Navigate to the project folder:**

   ```bash
   cd orderserver
   ```

3. **Run Docker Images**:

   ```bash
   docker-compose up
   ```

This will launch:

- PostgreSQL and Redis services.
- Multiple instances of the gRPC server.
- NGINX for load balancing.

---

## Accessing the Services

- **REST API and gPRC**: Access via NGINX at `http://localhost:80`.

