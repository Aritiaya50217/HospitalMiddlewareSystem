# Hospital Middleware System

Hospital Middleware System is a **production-ready backend service** designed to manage **patients, staff, hospitals, and users** in a hospital ecosystem.  
It is built with **Golang**, **Gin**, **PostgreSQL**, and follows **Hexagonal (Ports & Adapters) and Clean Architecture** to ensure scalability, maintainability, and testability.

This project demonstrates how to design a **modern backend system** with strong separation of concerns, authentication, and containerized deployment.

---

## Key Features

- JWT-based Authentication & Authorization  
- User, Staff, Patient, Hospital, and Gender Management  
- PostgreSQL persistence with migrations  
- Hexagonal (Ports & Adapters) Architecture  
- Clean Architecture principles  
- Fully Dockerized for local & production use  
- Unit tests with mocks for business logic and handlers  

---

## Technology Stack

| Category | Technology |
|--------|------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| Database | PostgreSQL |
| ORM | GORM |
| Authentication | JWT |
| Containerization | Docker, Docker Compose |
| Testing | Go Test, Testify |
| Architecture | Hexagonal (Ports & Adapters), Clean Architecture |

---

## System Architecture

This project is designed using **Hexagonal Architecture**, which separates the system into **core business logic** and **external adapters**.

