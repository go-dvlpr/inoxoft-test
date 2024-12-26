# **Project README**

## **Project Overview**

This project is a job management service that processes job logs and exposes RESTful endpoints to interact with these logs. You can run the application locally or using Docker for convenience.

---

## **Getting Started**

### **1. Run with Docker**
#### **Prerequisites**
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

#### **Command**
To build and run the application in a Docker container:
```bash
docker-compose up --build -d
```

#### **Key Notes**
- The application runs on **port 8080**.
- The log files are stored in the `logs/logs.log` file, which is shared with the host via a Docker volume.

---

### **2. Run Locally**
#### **Prerequisites**
- [Go (latest)](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/)

#### **Command**
To run the application locally:
```bash
make up
```

#### **Key Notes**
- The application requires `Make` for automation.
- Logs are written to the `logs/logs.log` file in the project directory.

---

## **Project Structure**

```plaintext
inoxoft-test/
├── cmd/                     # Main entry point for the application
│   └── main.go              # Bootstraps the server
├── config/                  # Configuration handling
│   └── config.go
├── jobs/                    # Job-related logic
│   ├── job.go               # Job struct and logic
│   └── processor.go         # Handles job processing
├── logs/                    # Log storage
│   └── logs.log             # Log file (shared with Docker volume)
├── server/                  # Server and API routing
│   ├── handlers/            # API handlers
│   │   └── router.go        # Route definitions
│   └── server.go            # Server setup
├── .env                     # Environment variables
├── .env.example             # Example environment variables
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile               # Multi-stage Dockerfile
├── Makefile                 # Local automation commands
└── README.md                # Documentation (this file)
```

---

## **Environment Variables**

The following variables can be configured in the `.env` file:

| **Variable**      | **Default Value**         | **Description**                     |
|-------------------|---------------------------|-------------------------------------|
| `LOG_FILE_PATH`   | `./logs/logs.log`         | Path to the log file               |
| `SERVER_PORT`     | `8080`                    | The port where the server runs     |

---

## **Endpoints**

### **1. Get Logs for a Specific Job**
- **Method:** `GET`
- **URL:** `http://localhost:8080/jobs/{id}/logs`
- **Description:** Fetch logs for a specific job by its ID.

#### Example:
```bash
curl -X GET http://localhost:8080/jobs/5/logs
```

---

### **2. Get All Logs**
- **Method:** `GET`
- **URL:** `http://localhost:8080/jobs/logs`
- **Description:** Retrieve all job logs.

#### Example:
```bash
curl -X GET http://localhost:8080/jobs/logs
```

---

### **3. Add a New Job**
- **Method:** `POST`
- **URL:** `http://localhost:8080/jobs/logs`
- **Description:** Add a new job log entry with a name and duration.

#### Request Body:
```json
{
  "name": "test",
  "millisecond_duration": 35000
}
```

#### Example:
```bash
curl -X POST http://localhost:8080/jobs/logs \
-H "Content-Type: application/json" \
-d '{
  "name": "test",
  "millisecond_duration": 35000
}'
```

---

## **Building the Application**

### **Using Docker**
1. Build and run the Docker containers:
   ```bash
   docker-compose up --build -d
   ```

2. Stop and remove the containers:
   ```bash
   docker-compose down
   ```

---

### **Using Make**
1. Run the application locally:
   ```bash
   make up
   ```

2. Stop the application:
   ```bash
   make stop
   ```

---

## **Development Notes**
- Logs are stored in `logs/logs.log`.
- When using Docker, logs are shared with the host via a volume mounted to the `logs` directory.
- Multi-stage builds in the Dockerfile ensure a smaller runtime image.

---

## **Contributing**
1. Fork this repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Make your changes and commit:
   ```bash
   git commit -m "Add new feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Create a pull request.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for more details.