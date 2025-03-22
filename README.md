# Working Branch For The Backend

## Authentication
- [x] Implemented user signup (with hashed password)
- [x] Implemented user login (JWT token generation)
- [x] **Check**: Add validation to check if user already exists before inserting
- [ ] **Check**: Add JWT validation middleware for protected routes

## Database (Authentication)
- [x] User table for storing CitizenID, Password, Phone
- [ ] **Check**: Add unique constraints or indexes to prevent duplicate users
- [ ] **Check**: Store JWT tokens for user sessions (optional, depends on use case)

## Database (Elections)
- [ ] Create tables for storing election data
  - ElectionID
  - Start Time / End Time
  - Election Name
  - Candidates (possibly in a separate table with relationships)
  - Vote Count
  - Election Status (active, closed, etc.)

## Redis (Real-time vote counting)
- [ ] Set up Redis for real-time vote count tracking
- [ ] Implement pub/sub mechanism to listen to vote changes
- [ ] Integrate Redis with Go for tracking and updating vote counts

## Voting System
- [ ] Implement voting endpoint to register votes
- [ ] Integrate with Redis to keep vote counts up-to-date in real-time
- [ ] Ensure that each user can vote only once per election
- [ ] Add restrictions (e.g., voter eligibility, one vote per citizen)

## WebSockets (Real-time updates)
- [ ] Set up WebSocket connections for real-time vote updates
- [ ] Broadcast vote changes to all connected clients
- [ ] Implement client-side WebSocket handling to receive updates
- [ ] Add a WebSocket server to handle connections and broadcast messages

## Load Balancing & Scaling
- [ ] Implement load balancing for the backend (using tools like NGINX or HAProxy)
- [ ] Set up horizontal scaling (possibly via Docker containers or Kubernetes)
- [ ] Monitor server load and adjust resources as necessary

## Logging & Error Handling
- [ ] Implement centralized logging (e.g., with ELK stack or something simpler like Logrus)
- [ ] Add error handling middleware for consistent error responses
- [ ] Set up monitoring for errors and uptime (e.g., Prometheus, Grafana)

## Security
- [ ] Implement input validation and sanitization to prevent SQL injection and XSS
- [ ] Implement rate limiting for API calls to prevent abuse
- [ ] Use HTTPS (SSL/TLS) for secure communication between frontend and backend
- [ ] Store sensitive data securely (e.g., JWT secret, database credentials)

## Documentation
- [ ] Add API documentation (Swagger or Postman collection)
- [ ] Write README with project setup, usage, and endpoints
- [ ] Document database schema for elections and users

## Tests
- [ ] Write unit tests for each handler and function
- [ ] Implement integration tests for the voting system
- [ ] Set up test environment and use mocking for Redis and DB
