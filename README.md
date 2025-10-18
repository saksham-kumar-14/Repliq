# Repliq
My submission for Inter IIT Dev

## Tech Stack
### Backend
- Language: Go
- Framework: Echo
- ORM: GORM
- Authentication: JWT

### Database
- Postgresql

### Frontend
- Framework: Svelte
- Runtime: Bun
- Build Tool: Vite

### Reverse Proxy
- Nginx: Route traffic to frontend and backend services

### Containerization
- Docker

## Features
### Backend
- Features
    - API routes for communication
    - Rate limiter: server protection by limiting request frequency per IP
    - Dockerfile for Go server setup
- Api Routes
    - `GET /v1/health` : basic check for server health
    - `GET /v1/user/:id` : get specific user information
    - `POST /v1/user`: user registration
    - `POST /v1/user/login`: user login
    - `GET /v1/api/token`: JWT verification
    </br></br>
    ** The following routes use JWT middleware for user verification.
    * API request must contain `headers:{Authorization: "Bearer <jwt-token>"}`
    - `GET /v1/post`: get all posts
    - `POST /v1/post`: create post
    - `GET /v1/post/:id`: get specific post
    - `PATCH /v1/post/:id`: update specific post
    - `DELETE /v1/post/:id`: delete specific post

### Frontend
- Features
    - `/`:
        - Displays Login/Registration if not loggedin
        - Else displays posts.
    - Sorting comments by Oldest, Newest, Upvotes
    - Collapsable Replies
    - Notifications
    - Avatar for users

### Project Structure
```
├── backend/
    ├── cmd/
        ├── api/
            ├── api.go          # for api and server config
            ├── main.go         # main file to run
            ├── errors.go       # api resp errors are defined
            ├── json.go         # for writing json responses and reading json requests
            ├── middleware.go   # for JWT auth
            ├── user.go         # user api functions
            ├── comment.go      # posts api functions
            └── health.go       # api function for basic health check
    ├── internal/
        ├── auth/                # JWT auth
        ├── db/                  # initialize db
        ├── rateLimiter/         # rate limiting
        └── store/               # db store for user and comments
    └── Dockerfile

├── frontend/
│   ├── src/
│   │   ├── lib/                # Svelte components
│   │   ├── store/              # API service functions
│   │   ├── utils/              # util functions
│   │   └── App.svelte

├── nginx
|   ├── Dockerfile              # This first builds svelte frontend, then host it with nginx on PORT 80
|   ├── nginx.conf              # configuration for nginx
├── docker-compose.yml
└── README.md
```


## Run

### Common commands
```bash
git clone https://github.com/saksham-kumar-14/Repliq
cd ./Repliq
```
- Create `./backend/.env` according to `./backend/.env.example` provided

### Using Docker
```bash
docker-compose up --build
```
- The website will run on `http://localhost`
- Backend service will run on `http://localhost:8000`

### Without Using Docker
```bash
git clone https://github.com/saksham-kumar-14/repliq
```
```bash
cd ./backend
go mod tidy
go run ./cmd/api/*.go
```
In another terminal,
```bash
cd ./frontend
bun install
bun run dev
```
- The website will run on `http://localhost:5173`
- Backend service will run on `http://localhost:8000`
