# Load Balancer in Go

This project demonstrates the implementation of a simple load balancer in Go, capable of distributing HTTP requests across multiple backend servers using various algorithms such as round-robin, least-connections, and IP-hash.

## Table of Contents

- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Load Balancing Algorithms](#load-balancing-algorithms)
- [Learning Outcomes](#learning-outcomes)

## Setup and Installation

### Prerequisites

- Go 1.18+ installed on your machine
- Git (for cloning the repository)

### Steps

1. **Clone the repository**:

   ```bash
   git clone https://github.com/nareshNishad/loadbalancer.git
   cd loadbalancer
   ```

2. **Build and run the project:**:

   ```
   go run main.go <algorithm>
   ```

   Replace <algorithm> with one of the available algorithms: round-robin, least-connections, ip-hash.

### Usage

1. **Starting the Load Balancer**
   The load balancer will start on http://localhost:8080 and will distribute incoming HTTP requests to the backend servers running on ports 8081, 8082, and 8083.

2. **Sending Requests**
   You can send requests to the load balancer using a web browser, curl, or Postman:

   '''
   curl http://localhost:8080
   '''
   This request will be routed to one of the backend servers based on the chosen algorithm.

### Load Balancing Algorithms

1. **Round-Robin**
   Distributes requests evenly across all servers in a circular order.
2. **Least-Connections**
   Directs requests to the server with the least number of active connections.
3. **IP-Hash**
   Routes requests based on the client's IP address, ensuring that the same IP always hits the same backend server.
