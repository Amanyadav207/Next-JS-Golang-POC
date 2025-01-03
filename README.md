# HiddenAlpha: Dockerized Full-Stack Application

This project demonstrates a simple full-stack application using **Next.js** (frontend) and **Golang** (backend), Dockerized for easy deployment with Docker Compose. The frontend fetches product data from the backend API.

---

## Table of Contents
1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Prerequisites](#prerequisites)


---

## Features

- **Backend**: A RESTful API built with Golang.
- **Frontend**: A Next.js-based application for fetching and displaying product data.
- **Dockerized**: Both services are containerized for easy setup and deployment.
- **Inter-Container Networking**: The frontend communicates with the backend using Docker's internal networking.

---

## Tech Stack

- **Frontend**: [Next.js](https://nextjs.org/) (React framework)
- **Backend**: [Golang](https://golang.org/)
- **Containerization**: [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)

---

## Prerequisites

Ensure you have the following installed:

1. **Docker**: [Download Docker](https://docs.docker.com/get-docker/)
2. **Docker Compose**: Comes bundled with Docker Desktop.
3. **Node.js** (for local development, if needed): [Download Node.js](https://nodejs.org/)
4. **Golang** (for local backend development, if needed): [Download Golang](https://golang.org/)

---

- Open the **frontend** at [http://localhost:3000](http://localhost:3000). The frontend fetches product data from the backend.

- Test the backend APIs directly:
  - [http://localhost:8080/api/ssg-products](http://localhost:8080/api/ssg-products)
  - [http://localhost:8080/api/ssr-products](http://localhost:8080/api/ssr-products)
  - [http://localhost:8080/api/isr-products](http://localhost:8080/api/isr-products)