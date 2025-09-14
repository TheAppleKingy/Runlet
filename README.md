# Runlet


**Runlet** is a system for remote code execution.  
It provides an API for solving programming problems by running code on remote services and retrieving execution results.  

There are 4 main domains: **student, teacher, course, problems**.  
- Students belong to classes (identified by a number).  
- Each class can be subscribed to multiple courses (many-to-many).  
- Courses contain problems, which students can solve if their class is subscribed.  
- Teachers can review solutions, edit problems (including test cases), and leave comments.  

---

- [Install](#️-install)
- [Quickstart](#-quickstart)
- [Usage](#-usage)
- [Stack](#-stack)

---

## ⚙️ Install


Before running the app, set the following environment variables:

### 🔑 JWT Auth

| Variable                 | Description                        |
|--------------------------|------------------------------------|
| `SECRET_KEY`             | Secret key to sign JWT tokens      |
| `JWT_TOKEN_EXPIRE_TIME`  | Token expiration time (in seconds) |

### 🗄 PostgreSQL

| Variable            | Description             |
|---------------------|-------------------------|
| `POSTGRES_DB`       | Database name           |
| `POSTGRES_USER`     | PostgreSQL username     |
| `POSTGRES_PASSWORD` | PostgreSQL password     |

### 🏃 Runners

| Variable            | Description                             |
|---------------------|-----------------------------------------|
| `RUNNERS_CONF_PATH` | Path to runners config (runners.yaml)   |


---

To start:

```bash
make runlet.rebuild.start
```

---

## 🚀 Quickstart

Once containers are up, the API will be ready to accept requests.  
Swagger documentation is available at:

🔗 [http://localhost:8081/docs/index.html](http://localhost:8081/docs/index.html)

---

## 📦 Usage

The API provides basic authentication features:

- Registration teacher
- Registration student
- Login / Logout

>Refer to Swagger UI for request details.

---

### 🧩 Sending solutions


Runlet itself is just a **gateway**.  
Code execution happens in remote runners ([Runners repository](https://github.com/TheAppleKingy/Runlet_runners)) via **gRPC**.  

To add a new runner for a programming language, just configure it in `runners.yaml`.

**Available endpoints:**

| Method | Endpoint                                | Description                          |
|--------|-----------------------------------------|--------------------------------------|
| GET    | `/api/student/courses`                  | Get all courses for the student      |
| GET    | `/api/student/courses/{course_id}/problems` | Get all problems in a course         |
| POST   | `/api/student/problems/{problem_id}/sent_solution` | Submit a solution to a problem |




---

## 🧰 Stack

- **Gin** – web framework  
- **PostgreSQL** – database  
- **goqu** – quary builder  
- **[Migrate](github.com/golang-migrate/migrate)** – DB migrations  
- **Docker** – containerization
- **Swaggo** - swagger for Gin framework
- **[gRPC](google.golang.org/grpc) - gRPC client(server on runners side)**
---

Feel free to contribute or open issues.
