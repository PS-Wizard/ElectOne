```mermaid
erDiagram
CITIZENS {
    VARCHAR citizenID PK "Citizen ID"
        VARCHAR fullName "Full Name"
        DATE dateOfBirth "Date of Birth"
        TEXT placeOfResidence "Place of Residence"
}
USERS {
    INTEGER userID PK "User ID"
        VARCHAR citizenID FK "Citizen ID"
        VARCHAR password "Password"
}
ADMINS {
    INTEGER adminID PK "Admin ID"
        VARCHAR email "Email"
        VARCHAR password "Password"
}
ELECTIONS {
    INTEGER electionID PK "Election ID"
        VARCHAR title "Title"
        DATE startDate "Start Date"
        DATE endDate "End Date"
        VARCHAR votingRestrictions "Voting Restrictions"
}
CANDIDATES {
    INTEGER candidateID PK "Candidate ID"
        VARCHAR citizenID FK "Citizen ID"
        VARCHAR post "Post"
        INTEGER electionID FK "Election ID"
}

CITIZENS ||--o| USERS : has
CITIZENS ||--o| CANDIDATES : is_citizen_of
ELECTIONS ||--o| CANDIDATES : has_election
USERS }|--|| ADMINS : is_admin
```

### Admins:

| CID | NAME     | TYPE         | NOTNULL | DFLT VALUE | PK |
| --- | -------- | ------------ | ------- | ---------- | --- |
| 0   | adminID  | INTEGER      | 0       | NULL       | 1   |
| 1   | email    | VARCHAR(255) | 1       | NULL       | 0   |
| 2   | password | VARCHAR(255) | 1       | NULL       | 0   |

---

### Citizens:

| CID | NAME              | TYPE         | NOTNULL | DFLT VALUE | PK |
| --- | ----------------- | ------------ | ------- | ---------- | --- |
| 0   | citizenID         | VARCHAR(20)  | 0       | NULL       | 1   |
| 1   | fullName          | VARCHAR(255) | 1       | NULL       | 0   |
| 2   | dateOfBirth       | DATE         | 1       | NULL       | 0   |
| 3   | placeOfResidence  | TEXT         | 1       | NULL       | 0   |

---

### Users:

| CID | NAME      | TYPE         | NOTNULL | DFLT VALUE | PK |
| --- | --------- | ------------ | ------- | ---------- | --- |
| 0   | userID    | INTEGER      | 0       | NULL       | 1   |
| 1   | citizenID | VARCHAR(20)  | 1       | NULL       | 0   |
| 2   | password  | VARCHAR(255) | 1       | NULL       | 0   |

---

### Election:

| CID | NAME               | TYPE         | NOTNULL | DFLT VALUE | PK |
| --- | ------------------ | ------------ | ------- | ---------- | --- |
| 0   | electionID         | INTEGER      | 0       | NULL       | 1   |
| 1   | title              | VARCHAR(255) | 1       | NULL       | 0   |
| 2   | startDate          | DATE         | 1       | NULL       | 0   |
| 3   | endDate            | DATE         | 1       | NULL       | 0   |
| 4   | votingRestrictions | VARCHAR(255) | 1       | NULL       | 0   |

---

### Candidates:

| CID | NAME        | TYPE         | NOTNULL | DFLT VALUE | PK |
| --- | ----------- | ------------ | ------- | ---------- | --- |
| 0   | candidateID | INTEGER      | 0       | NULL       | 1   |
| 1   | citizenID   | VARCHAR(20)  | 1       | NULL       | 0   |
| 2   | post        | VARCHAR(255) | 1       | NULL       | 0   |
| 3   | electionID  | INTEGER      | 1       | NULL       | 0   |
| 4   | GroupName   | TEXT         | 0       | NULL       | 0   |

---

# API Routes:
[Postman Collection](./rester-export-postman.json)
![routes](assets/routes.png)
