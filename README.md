# Restaurant Management System – Go + Gin + MongoDB

A production-ready Restaurant Management backend built with Go, Gin and MongoDB.  
This backend manages Users, Menus, Foods, Orders, Order Items, Tables and Invoices. It uses JWT (access + refresh tokens) for authentication and supports Role-Based Access Control (Admin / User). The project follows an MVC-like, modular structure.

---

## 🚀 Tech stack

- Go 1.24.3
- Gin Web Framework
- MongoDB (official Go driver)
- JWT Authentication (HS256)
- validator.v10
- Role-Based Access Control (Admin / User)
- Architecture: MVC-style folder structure

---

## ✨ Features

Authentication
- JWT-based authentication (HS256)
- Access token: valid for 24 hours
- Refresh token: valid for 7 days
- Role-based access (Admin / User)
- Token validation middleware

Restaurant management
- User signup / login
- Menu CRUD
- Food CRUD
- Table CRUD
- Order CRUD
- Order Item CRUD
- Invoice generation with auto-calculated totals
- MongoDB collections for all models

Additional
- Modular routing
- Environment-based configuration
- Clean, scalable architecture ready for deployment

---

## 📂 Project structure

Restaurant_Managment_System/  
├── go.mod  
├── go.sum  
├── main.go  
├── controllers/  
│   ├── foodController.go  
│   ├── invoiceController.go  
│   ├── menuController.go  
│   ├── orderController.go  
│   ├── orderItemController.go  
│   ├── tableController.go  
│   └── userController.go  
├── database/  
│   └── dbConnect.go  
├── helpers/  
│   └── tokenHelper.go  
├── middlewares/  
│   └── authMiddleware.go  
├── models/  
│   ├── FoodModel.go  
│   ├── invoiceModel.go  
│   ├── menuModel.go  
│   ├── noteModel.go  
│   ├── orderItemModel.go  
│   ├── orderModel.go  
│   ├── tableModel.go  
│   └── userModel.go  
└── routes/  
    ├── foodRouter.go  
    ├── invoiceRouter.go  
    ├── menuRouter.go  
    ├── orderItemRouter.go  
    ├── orderRouter.go  
    ├── tableRouter.go  
    └── userRouter.go

---

## 🧩 Architecture (high level)

Client / Frontend -> Gin Router -> Authentication Middleware -> Controllers -> Models -> MongoDB

Helper functions (e.g. token creation/verification) are located in `helpers/tokenHelper.go`.

---

## ⚙️ Installation & setup

1. Clone the repository
```bash
git clone https://github.com/Aditya7880900936/Restaurant_Managment_System.git
cd Restaurant_Managment_System
```

2. Install dependencies
```bash
go mod tidy
```

3. Environment variables

Create a `.env` file in the project root or set the following system environment variables. Example `.env`:

```
PORT=8000
SECRET_KEY=your_jwt_secret_here
MONGO_URL=mongodb://localhost:27017
DB_NAME=restaurant_db
```

Notes:
- SECRET_KEY is used to sign both access and refresh tokens (HS256). Use a secure, long random string in production.
- MONGO_URL may include credentials if needed, e.g. `mongodb://user:pass@host:27017`.

5. Run the server

Development:
```bash
go run main.go
```

Build for production:
```bash
go build -o server
./server
```

---

## 🔐 JWT authentication flow

- Access token: valid for 24 hours — used for protected routes.
- Refresh token: valid for 7 days — used to obtain a new access token.
- Token creation and verification are implemented in `helpers/tokenHelper.go`.
- Protected routes are validated by `middlewares/authMiddleware.go`.

Example token payload (used in repo):
```go
type SignedDetails struct {
    Email      string
    First_name string
    Last_name  string
    Uid        string
    User_Type  string // "Admin" or "User"
}
```

---

## 👤 Role-based access

Roles supported:
- Admin — full management: users, menus, foods, tables, orders, invoices
- User — view menus, create orders, view tables

The role is encoded in the JWT and enforced by middleware.

---

## 📌 API Routes (overview)

User routes
```
GET    /users
GET    /users/:user_id
POST   /users/signup
POST   /users/login
```

Menu routes
```
GET    /menus
GET    /menus/:menu_id
POST   /menus
PATCH  /menus/:menu_id
```

Food routes
```
GET    /foods
GET    /foods/:food_id
POST   /foods
PATCH  /foods/:food_id
```

Invoice routes
```
GET    /invoices
GET    /invoices/:invoice_id
POST   /invoices
PATCH  /invoices/:invoice_id
```

Orders
```
GET    /orders
GET    /orders/:order_id
POST   /orders
PATCH  /orders/:order_id
```

Order items
```
GET    /orderItems
GET    /orderItems/:orderItem_id
GET    /orderItems-order/:order_id
POST   /orderItems
PATCH  /orderItems/:orderItem_id
```

Tables
```
GET    /tables
GET    /tables/:table_id
POST   /tables
PATCH  /tables/:table_id
```

Note: Check controller and route files for exact request and response payloads, validation rules and required fields.

---

## 🧪 Testing

- Run unit tests (if present):
```bash
go test ./...
```
- For integration tests that require MongoDB, point to a test instance (separate DB_NAME) or use a test container.

---
## 🤝 Contributing

Contributions are welcome. Suggested workflow:
1. Fork the repo
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Implement changes and add tests
4. Run tests locally
5. Open a pull request with a clear description of changes

Please keep code modular and follow the existing project structure.

---

## 📜 License

Recommended: MIT License. If you want a different license, I can add it.

---

If you'd like, I can now:
- generate a Dockerfile and docker-compose.yml, or
- produce an OpenAPI (Swagger) spec from the route/controller signatures, or
- create a .env.example file and a simple config loader — tell me which one you prefer.
