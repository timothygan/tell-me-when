# Tell Me When

First attempt at vibe coding. The other day I saw the news that Figma planned to go public, and I was interested in buying stock. The exact date when the stock would drop was unknown, and I knew it was very unlikely I would check in every day just to check if their stock was available. Idea is to make an application that will take queries whose answers are not yet known, and push a notification to the user once that information is available. It'll query some LLM for the information, and use another LLM to verify that information is correct/satisfactory before notifying the user, and run on a schedule each day.

---

## Table of Contents

1. [Features](#features)
2. [System Architecture](#system-architecture)
3. [Data Flow & User Journey](#data-flow--user-journey)
4. [AI Integration](#ai-integration)
5. [Data Model](#data-model)
6. [Tech Stack](#tech-stack)
7. [Security & Privacy](#security--privacy)
8. [Extensibility](#extensibility)
9. [Example Workflow](#example-workflow)
10. [Future Enhancements](#future-enhancements)

---

## Features

- **User Registration & Authentication**
- **Submit & Manage Queries**
- **Automated Daily (or configurable) Checks for Updates**
- **AI-powered Information Retrieval**
- **AI-powered Answer Verification**
- **Multi-channel Notifications (Email/SMS)**
- **Dashboard for Query Management**
- **Notification History**

---

## System Architecture

### High-Level Diagram

```
+----------------+        +-------------------+        +-------------------+
|                |        |                   |        |                   |
|   Frontend     +------->+    Backend API    +------->+    Database       |
| (React/Next.js)|        | (Node/Python)     |        |   (PostgreSQL)    |
|                |        |                   |        |                   |
+----------------+        +-------------------+        +-------------------+
                                |   ^  |   ^
                                v   |  v   |
                  +-------------------+   |
                  |   Scheduler/Worker|   |
                  +-------------------+   |
                                |         |
         +----------------------+         +--------------------+
         |                                                |
         v                                                v
+-------------------+                          +----------------------+
| AI Query Engine   |                          | Notification Service |
| (LLM/Search API)  |                          | (Email/SMS)         |
+-------------------+                          +----------------------+
         |
         v
+----------------------+
| AI Verification      |
| Engine (LLM)         |
+----------------------+
```

---

### Component Deep Dive

#### 1. **Frontend Webapp**
- Built with React (Next.js recommended for SSR/SSG)
- User authentication (OAuth/JWT)
- Query submission form
- Dashboard to view/manage queries and notifications

#### 2. **Backend API**
- RESTful or GraphQL API
- Handles user management, query CRUD, notification logs
- Exposes endpoints for frontend and worker
- Enforces authentication and authorization

#### 3. **Database**
- Stores users, queries, notification logs, etc.
- PostgreSQL recommended for relational structure and reliability

#### 4. **Scheduler/Worker**
- Runs as a background process (e.g., Celery, BullMQ, or serverless cron)
- Periodically fetches active queries and triggers the AI pipeline

#### 5. **AI Query Engine**
- Uses LLM or search APIs (OpenAI, Bing, Google Search, etc.)
- Retrieves up-to-date information relevant to each query

#### 6. **AI Verification Engine**
- Uses a separate LLM or prompt to validate if the retrieved information satisfactorily answers the query
- Extracts structured data (e.g., dates, facts) from unstructured text

#### 7. **Notification Service**
- Sends notifications via email (SendGrid, Mailgun) or SMS (Twilio)
- Logs all notifications for auditing and user review

---

## Data Flow & User Journey

1. **User submits a query** via the frontend.
2. **Backend stores the query** in the database, linked to the user.
3. **Scheduler/Worker triggers daily** (or at a set interval):
   - Pulls all active queries.
   - For each query:
     - **AI Query Engine** searches for new information.
     - **AI Verification Engine** validates the result.
     - If satisfactory, **Notification Service** alerts the user.
4. **User receives notification** (email/SMS) with the answer and a link to the source.
5. **User can manage or archive queries** via the dashboard.

---

## AI Integration

### AI Query Engine

- **Purpose:** Retrieve the latest information about a user’s query.
- **How:**  
  - Use a search API (Google, Bing, or custom web scraping) to find relevant news or announcements.
  - Optionally, use an LLM to summarize or extract key points from search results.

- **Sample Prompt:**  
  ```
  Find the latest news or official announcement about [QUERY]. Return the most relevant snippet and its source.
  ```

### AI Verification Engine

- **Purpose:** Determine if the retrieved information satisfactorily answers the query.
- **How:**  
  - Use an LLM with a verification prompt.
  - Extract structured data (e.g., dates, event confirmations).

- **Sample Prompt:**  
  ```
  Given the following text, does it clearly state when [QUERY]? If so, extract the date and summarize the finding. If not, respond ‘No’.
  ```

---

## Data Model

### User

| Field                 | Type        | Description                      |
|-----------------------|-------------|----------------------------------|
| id                    | UUID        | Primary key                      |
| email                 | String      | User email                       |
| phone_number          | String      | Optional, for SMS notifications  |
| notification_prefs    | JSON        | Email/SMS/both                   |
| password_hash/auth_id | String      | For authentication               |
| created_at            | Timestamp   |                                  |

### Query

| Field               | Type      | Description                                    |
|---------------------|-----------|------------------------------------------------|
| id                  | UUID      | Primary key                                    |
| user_id             | UUID      | Foreign key to User                            |
| query_text          | String    | User’s question (e.g., “When is Figma IPO?”)   |
| status              | Enum      | active, answered, expired                      |
| last_checked        | Timestamp | Last AI check                                  |
| answer              | Text      | Answer if found                                |
| answer_verified     | Bool      | Was answer verified by AI?                     |
| notification_sent   | Bool      | Has user been notified?                        |
| created_at          | Timestamp |                                                |

### NotificationLog

| Field           | Type      | Description                          |
|-----------------|-----------|--------------------------------------|
| id              | UUID      | Primary key                          |
| user_id         | UUID      | Foreign key to User                  |
| query_id        | UUID      | Foreign key to Query                 |
| sent_at         | Timestamp | When notification was sent           |
| channel         | Enum      | email, sms, etc.                     |
| message         | Text      | Notification content                 |

---

## Tech Stack

- **Frontend:** React (Next.js)
- **Backend:** Node.js (Express/NestJS) or Python (FastAPI)
- **Database:** PostgreSQL
- **Scheduler/Worker:** Celery (Python), BullMQ (Node.js), or serverless cron
- **AI Query Engine:** OpenAI, Google Search API, Bing Search API, or custom web scraping
- **AI Verification Engine:** OpenAI GPT-4/Claude with custom prompts
- **Notifications:** SendGrid (email), Twilio (SMS)
- **Authentication:** Auth0, Firebase Auth, or custom JWT
- **Hosting:** Vercel/Netlify (frontend), AWS/GCP/Azure (backend/workers)

---

## Security & Privacy

- **Authentication:** Secure user authentication (OAuth/JWT)
- **Data Protection:** Encrypt sensitive data at rest and in transit
- **Rate Limiting:** Prevent abuse of query/notification endpoints
- **User Privacy:** Allow users to delete their data and queries

---

## Extensibility

- Add more notification channels (push, Slack, Discord)
- Support for recurring or more complex queries
- Admin dashboard for monitoring and moderation
- Integration with more advanced AI agents for deeper verification
- Allow users to refine/update queries after submission

---

## Example Workflow

1. **User submits:**  
   “Notify me when Figma’s stock is available for public purchase.”

2. **System stores query.**

3. **Each day:**
   - AI Query Engine searches for latest news on “Figma IPO.”
   - AI Verification Engine is prompted:  
     “Does this text specify a public stock release date for Figma? If so, extract the date.”
   - If a date is found and verified, notification is sent to the user.

---

## Future Enhancements

- **User-defined check frequency** (hourly, daily, weekly)
- **Smart notifications** (only notify if information is new or changed)
- **Community-sourced queries/answers**
- **Integration with calendar apps (Google Calendar, Outlook)**
- **Mobile app version**

---

## Getting Started (Example)

1. **Clone the repository**
2. **Install dependencies**
3. **Set up environment variables** (API keys for AI, email/SMS providers)
4. **Run database migrations**
5. **Start backend and frontend servers**
6. **Start the scheduler/worker process**

---

## Contributing

PRs and suggestions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## License

MIT
