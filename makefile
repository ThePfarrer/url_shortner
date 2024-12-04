# Makefile for URL Shortener App

# Variables
FRONTEND_DIR := frontend
PYTHON_BACKEND_DIR := python_backend
GO_BACKEND_DIR := golang_backend
PYTHON_PORT := 5000
GO_PORT := 8080
FRONTEND_PORT := 5173

# Commands
.PHONY: go py build-go build-py run-go run-py stop clean frontend-go frontend-py python-backend go-backend

# Default target: Build and run the application with Golang
go: build-go run-go

# Default target: Build and run the application with Python
py: build-py run-py

# Build all components wit Golang
build-go: frontend-go go-backend
	@echo "All components built successfully for Go backend."

# Build all components wit Golang
build-py: frontend-py python-backend
	@echo "All components built successfully for Python backend."

# Build Vue.js frontend with Golang
frontend-go:
	@echo "Building frontend with..."
	cd $(FRONTEND_DIR) && npm install && npm run build-only:golang
	@echo "Frontend built successfully."

# Build Vue.js frontend with Python
frontend-py:
	@echo "Building frontend with..."
	cd $(FRONTEND_DIR) && npm install && npm run build-only:python
	@echo "Frontend built successfully."

# Prepare Python backend
python-backend:
	@echo "Setting up Python backend..."
	cd $(PYTHON_BACKEND_DIR) && python3 -m venv venv && ./venv/bin/pip install -r requirements.txt
	@echo "Python backend is ready."

# Build Go backend
go-backend:
	@echo "Building Go backend..."
	cd $(GO_BACKEND_DIR) && go build -o url-shortener
	@echo "Go backend built successfully."

# Run the application using Golang
run-go: frontend-run go-backend-run
	@echo "Application is running."

# Run the application using Python
run-py: frontend-run python-backend-run
	@echo "Application is running."

# Run Vue.js frontend
frontend-run:
	@echo "Starting frontend..."
	cd $(FRONTEND_DIR) && npm run preview -- --port $(FRONTEND_PORT) &

# Run Python backend
python-backend-run:
	@echo "Starting Python backend..."
	cd $(PYTHON_BACKEND_DIR) && flask --app api run &

# Run Go backend
go-backend-run:
	@echo "Starting Go backend..."
	cd $(GO_BACKEND_DIR) && ./url-shortener &

# Stop the application
stop:
	@echo "Stopping application..."
	killall npm || true
	killall python || true
	killall url-shortener || true
	@echo "Application stopped."

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(FRONTEND_DIR)/dist
	rm -rf $(PYTHON_BACKEND_DIR)/venv
	rm -f $(GO_BACKEND_DIR)/url-shortener
	@echo "Cleanup complete."
