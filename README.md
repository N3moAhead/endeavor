# Endeavor

A personal journaling web application built with Go and PostgreSQL, designed to make daily journaling and activity tracking effortless while providing powerful insights into your personal growth and goal progress.

## Features

- **Daily Journaling**: Quick and intuitive interface for daily entries
- **Activity Tracking**: Log and categorize daily activities
- **Mood Tracking**: Monitor emotional patterns over time
- **Historical Analysis**: View and analyze past entries with detailed insights
- **Category System**: Organize activities with customizable categories and colors
- **Web Interface**: Clean, responsive design built with Go templates and Tailwind CSS

## Tech Stack

- **Backend**: Go 1.24+
- **Database**: PostgreSQL 16
- **Frontend**: HTML templates with Tailwind CSS
- **Migration**: golang-migrate/migrate
- **Live Reload**: Air for development

## Quick Start

### Prerequisites

- Go 1.24+
- PostgreSQL 16
- Node.js & npm (for CSS compilation)
- Make

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/N3moAhead/endeavor.git
   cd endeavor
   ```

2. **Setup environment**
   ```bash
   cp .env_example .env
   # Edit .env with your database configuration
   ```

3. **Install dependencies and tools**
   ```bash
   make setup
   ```

4. **Start PostgreSQL**
   ```bash
   docker-compose up -d
   ```

5. **Seed the database**
   ```bash
   make seed
   ```

6. **Start development server**
   ```bash
   make dev
   ```

The application will be available at `http://localhost:9090`

## Development

### Available Commands

```bash
make help           # Show all available commands
make dev            # Start development environment with live reload
make build-prod     # Build optimized production binary
make run            # Build and run the application
make seed           # Build and run database seeder
make clean          # Remove generated files and dependencies
make fmt            # Format Go code
make vet            # Run Go vet analysis
```

### Project Structure

```
endeavor/
├── cmd/
│   ├── endeavor/          # Main application entry point
│   └── seed/              # Database seeding utility
├── internal/
│   ├── controller/        # HTTP handlers
│   ├── db/               # Database connection and migrations
│   ├── model/            # Data models and business logic
│   └── router/           # HTTP routing and templates
├── web/
│   ├── templates/        # HTML templates
│   ├── static/          # Static assets (CSS, JS)
│   └── input.css        # Tailwind CSS source
├── docker-compose.yaml   # PostgreSQL container setup
└── Makefile             # Build and development commands
```

### Database Schema

The application uses the following main entities:

- **Days**: Daily journal entries
- **Activities**: Trackable activities with categories
- **Categories**: Activity categorization with colors
- **Moods**: Emotional state tracking
- **Day_Activities**: Many-to-many relationship between days and activities

## Configuration

### Environment Variables

Create a `.env` file based on `.env_example`:

```env
DATABASE_URL=postgres://user:password@localhost:2509/endeavor?sslmode=disable
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=endeavor
```

### Database Setup

The application includes a Docker Compose configuration for PostgreSQL:

```bash
docker-compose up -d  # Start PostgreSQL
docker-compose down   # Stop PostgreSQL
```

## Production Deployment

1. **Build production binary**
   ```bash
   make build-prod
   ```

2. **Setup production database**
   - Ensure PostgreSQL is running
   - Run migrations: `migrate -path internal/db/migrations -database "$DATABASE_URL" up`
   - Seed initial data: `make seed`

3. **Run the application**
   ```bash
   ./endeavor
   ```

The server runs on port 9090 by default.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting: `make fmt && make vet`
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## About

Endeavor is a personal project focused on making journaling and self-reflection a daily habit. It's designed to be simple, fast, and insightful - helping you track your progress towards personal goals while maintaining a record of your journey.

Built with ❤️ using Go and PostgreSQL.
