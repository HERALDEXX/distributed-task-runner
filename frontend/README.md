# 🚀 Distributed Task Runner - Frontend

A modern, responsive React TypeScript frontend for managing and monitoring distributed task execution. Submit commands, scripts, and jobs to a distributed processing system with real-time status tracking and comprehensive job history management.

![React](https://img.shields.io/badge/React-19.1.0-61DAFB?style=flat-square&logo=react)
![TypeScript](https://img.shields.io/badge/TypeScript-5.8.3-3178C6?style=flat-square&logo=typescript)
![Vite](https://img.shields.io/badge/Vite-7.0.4-646CFF?style=flat-square&logo=vite)

## ✨ Features

### 🎯 **Job Submission Interface**

- **Intuitive Command Input**: Clean textarea interface for submitting shell commands, scripts, or any executable payload
- **Real-time Feedback**: Immediate job ID assignment and status confirmation upon submission
- **Input Validation**: Smart form validation prevents empty submissions and provides user feedback

### 📊 **Comprehensive Job Management**

- **Live Job History**: View all submitted jobs with real-time status updates
- **Status-Coded Visualization**: Color-coded job entries (🟢 completed, 🔴 failed, 🟡 running, ⚪ pending)
- **Chronological Sorting**: Most recent jobs displayed first for better workflow management
- **One-Click History Management**: Bulk clear job history with confirmation protection

### 🎨 **Modern User Experience**

- **Responsive Design**: Optimized for desktop and mobile viewing
- **Clean Material-Inspired UI**: Professional styling with subtle shadows and smooth transitions
- **Intuitive Navigation**: Seamless routing between job submission and history views
- **Loading States**: Proper loading indicators and error handling throughout

## 🛠️ Technology Stack

### **Core Framework**

- **React 19.1.0** - Latest React with concurrent features
- **TypeScript 5.8.3** - Full type safety and enhanced developer experience
- **React Router DOM 7.7.1** - Client-side routing for SPA navigation

### **Build & Development**

- **Vite 7.0.4** - Lightning-fast build tool with HMR
- **ESLint 9.30.1** - Code quality and consistency enforcement
- **Modern ESNext/ES2022** - Latest JavaScript features and syntax

### **API Integration**

- **RESTful Backend Communication** - Clean HTTP client for job management
- **Proxy Configuration** - Seamless development experience with backend integration
- **Error Handling** - Comprehensive error states and user feedback

## Quick Start

> Install dependencies

  ```bash
  npm install
  ```

> Start dev server (requires backend on [`localhost:8080`](http://localhost:8080))

  ```bash
  npm run dev
  ```

> Build for production

  ```bash
  npm run build
  ```

## API Integration

Expects backend running on [`localhost:8080`](http://localhost:8080) with endpoints:

- `POST /jobs` - Submit job
- `GET /jobs` - Get job history
- `DELETE /jobs` - Clear history

## Project Structure

```
src/
├── App.tsx           # Main app + routing
├── JobSubmitter.tsx  # Job submission form
├── JobList.tsx       # Job history view
└── *.css            # Component styles
```

---

Part of the distributed task runner system.
