# Working Branch For The Backend

## Authentication
- [x] Implemented user signup (with hashed password)
- [x] Implemented user login (JWT token generation)
- [x] **Check**: Add validation to check if user already exists before inserting
- [ ] **Check**: Add JWT validation middleware for protected routes

## Database (Authentication)
- [ ] Add a verifying function that handles any requests to any of the `/api/***` routes
    - Needs to verify the JWT token, to check if it's of type `admin`? <- Need To Figure Out

## Database (Elections)
```sql

~ Mock National Citizens Table

CREATE TABLE citizens (
    citizen_id INTEGER PRIMARY KEY,  -- Unique Citizen ID
    full_name VARCHAR(255) NOT NULL,  -- Full name of the citizen
    dob DATE NOT NULL,  -- Date of birth
    pob VARCHAR(255) NOT NULL,  -- Place of birth
    permanent_address VARCHAR(255) NOT NULL  -- Permanent address
);
```

>

```sql

~ Users Table

CREATE TABLE users (
    citizen_id INTEGER PRIMARY KEY,  -- Links to citizen's ID
    password TEXT NOT NULL,  -- Hashed password (for security)
    voted_in TEXT,  -- Comma-separated list of election_ids the user has voted in
    FOREIGN KEY (citizen_id) REFERENCES citizens(citizen_id) ON DELETE CASCADE  -- Enforces relationship to citizens
);
```

>

```
~ Elections Table

CREATE TABLE elections (
    election_id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Auto-incremented unique ID for each election
    name VARCHAR(255) NOT NULL,  -- Name of the election (e.g., "Presidential Election")
    start_date DATETIME NOT NULL,  -- Start date of the election
    end_date DATETIME NOT NULL,  -- End date of the election
    vote_restriction VARCHAR(255),  -- Restriction on who can vote (e.g., "@everyone", "@provincial", "@federal")
    FOREIGN KEY (election_id) REFERENCES elections(election_id) ON DELETE CASCADE
);
```

>

```sql

~ Runners:

CREATE TABLE runners (
    citizen_id INTEGER,  -- The ID of the citizen running (links to the citizens table)
    election_id INTEGER,  -- The election they are running in
    post VARCHAR(255) NOT NULL,  -- The position they are running for (e.g., "President", "Senator")
    total_votes INTEGER DEFAULT 0,  -- The total number of votes for this candidate (starts at 0)
    PRIMARY KEY (citizen_id, election_id),  -- Composite primary key (citizen + election)
    FOREIGN KEY (citizen_id) REFERENCES citizens(citizen_id) ON DELETE CASCADE,  -- Links to citizens table
    FOREIGN KEY (election_id) REFERENCES elections(election_id) ON DELETE CASCADE  -- Links to elections table
);
```

---

Foreign Keys and Hierarchy:
- Users: A user must have a valid citizen_id that links to a record in the citizens table. You cannot create a user without first having a citizen in the citizens table. If a citizen is deleted, the associated user is also deleted (thanks to ON DELETE CASCADE).
- Runners: A runner must have both a valid citizen_id (from the citizens table) and an election_id (from the elections table). You cannot create a runner without these records. If a citizen or election is deleted, the associated runner will also be deleted.

Creating New Records:
- Create a Citizen: Insert the citizen’s information into the citizens table. This generates a unique citizen_id that can be used in other tables.
- Create a User: After creating the citizen, you can insert a user record that links to the citizen_id. This establishes the relationship between a user and their citizen profile.

- Create an Election: Insert a new election into the elections table. This election will have a unique election_id that can be used when associating runners.

- Create a Runner: After creating both the citizen and the election, you can create a runner record. This links the citizen_id and election_id and specifies the post (e.g., President, Senator) and initial vote count.

Data Integrity: The foreign key constraints ensure that you cannot create a user or runner without the necessary linked records. This maintains consistency and prevents data from becoming disjointed.

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
- [x] Add API documentation (Swagger or Postman collection)
- [ ] Write README with project setup, usage, and endpoints
- [x] Document database schema for elections and users

## Tests
- [ ] Write unit tests for each handler and function
- [ ] Implement integration tests for the voting system
- [ ] Set up test environment and use mocking for Redis and DB

---

