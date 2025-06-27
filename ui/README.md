# AdFlowHQ UI

A modern React TypeScript application for managing AdFlowHQ workflows and dashboards.

## 🚀 Features

- **React 18** with TypeScript for type safety
- **React Router** for client-side routing
- **TanStack Query** for efficient data fetching and caching
- **Tailwind CSS** for modern, responsive styling
- **Vite** for fast development and building
- **Docker** support for containerized deployment
- **Nginx** for production serving

## 📁 Project Structure

```
ui/
├── src/
│   ├── components/
│   │   └── Layout.tsx          # Main layout component
│   ├── navigation.ts          # Navigation configuration
│   ├── pages/
│   │   ├── Dashboard.tsx       # Dashboard page
│   │   ├── Workflows.tsx       # Workflows management page
│   │   └── Settings.tsx        # Application settings page
│   ├── App.tsx                 # Main app component with routing
│   ├── main.tsx               # Application entry point
│   └── index.css              # Global styles and Tailwind imports
├── Dockerfile                 # Multi-stage Docker build
├── nginx.conf                 # Nginx configuration for production
├── package.json               # Dependencies and scripts
├── tailwind.config.js         # Tailwind CSS configuration
├── tsconfig.json              # TypeScript configuration
├── vite.config.ts             # Vite build configuration
└── index.html                 # HTML template
```

## 🛠️ Development

### Prerequisites

- Node.js 18 or higher
- npm or yarn package manager

### Installation

1. Navigate to the UI directory:
   ```bash
   cd ui
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

### Development Server

Start the development server:
```bash
npm run dev
```

The application will be available at `http://localhost:3000`

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run lint` - Run ESLint
- `npm run preview` - Preview production build

## 🎨 Styling

The application uses **Tailwind CSS** for styling with custom utility classes:

- `.btn-primary` - Primary button styling
- `.btn-secondary` - Secondary button styling
- `.card` - Card component styling

### Custom Colors

The theme includes custom primary colors:
- Primary 50: `#eff6ff`
- Primary 500: `#3b82f6`
- Primary 600: `#2563eb`
- Primary 700: `#1d4ed8`

## 🐳 Docker Deployment

### Building the Image

```bash
docker build -t adflowhq-ui .
```

### Running the Container

```bash
docker run -p 3000:3000 adflowhq-ui
```

The application will be available at `http://localhost:3000`

### Docker Compose

The UI is designed to work with the main `docker-compose.yml` file in the root directory, which includes:
- Temporal server for workflow management
- API proxy configuration through Nginx

## 🔧 Configuration

### Vite Configuration

The development server includes a proxy configuration for API calls:
- API requests to `/api/*` are proxied to `http://localhost:8080`
- This allows seamless integration with the backend services

### Nginx Configuration

The production Nginx configuration includes:
- React Router support with fallback to `index.html`
- API proxy to Temporal server
- Static asset caching
- Gzip compression

## 📱 Pages

### Dashboard (`/`)
The main dashboard page showing an overview of the system.

### Workflows (`/workflows`)
The workflows management page for creating and managing AdFlowHQ workflows.

### Settings (`/settings`)
Manage application configuration.

## 🔌 Dependencies

### Core Dependencies
- `react` - React library
- `react-dom` - React DOM rendering
- `react-router-dom` - Client-side routing
- `@tanstack/react-query` - Data fetching and caching
- `axios` - HTTP client
- `lucide-react` - Icon library
- `clsx` & `tailwind-merge` - Conditional styling utilities

### Development Dependencies
- `typescript` - TypeScript compiler
- `vite` - Build tool and dev server
- `tailwindcss` - CSS framework
- `eslint` - Code linting
- `@vitejs/plugin-react` - React plugin for Vite

## 🚀 Production Build

1. Build the application:
   ```bash
   npm run build
   ```

2. The built files will be in the `dist/` directory

3. For production deployment, use the Docker image which includes Nginx for serving the static files.

## 🔗 Integration

This UI is designed to integrate with:
- **Temporal** - For workflow orchestration
- **AdFlowHQ Backend** - For business logic and data management
- **Docker Compose** - For local development and testing

## 📄 License

This project is part of AdFlowHQ and follows the same license as the main project. 