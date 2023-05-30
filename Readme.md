 สร้าง db ก่อน

docker compose up //to create db

go mod tidy
go run .


go test -v


gorm 
- Create
- Read
- Update
- Delete
- Preload
- Modeling

Layer Separation
- Business Logic
- Database Access Layer (DAO)

routing
- Handler Request
- Add Course
- Add Class
- Enroll
- Register
- JWT
- Password Hashing


Today
- Login
- Testing
- Left over feature
- CORS
- Deploy
- Middleware
- Context




## Deploy

create Dockerfile

docker build -t peago1

docker run -rm -p 8080:8080 peago1

ต้องสร้าง db ใน container ก่อนด้วย docker-compose.yml

docker compose up -d