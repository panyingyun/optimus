![Optimus logo](./.github/optimus-logo--960x540.png)

# Optimus

> Image compression, optimization and conversion desktop app.

## Overview

Optimus is a desktop image optimization application. It supports conversion and compression between WebP, JPEG, and PNG image formats.

## Features

- Convert to and from JPEG, PNG, and WebP formats.
- Compress JPEG, PNG (lossy), and WebP (lossy and lossless) formats.
- Resize images to various sizes in a single batch operation.
- View simple stats on session and all-time use.

![Screenshot of Optimus primary image editor view](./.github/optimus_screenshot_editor--1200x742.png)

![Screenshot of Optimus options view](./.github/optimus_screenshot_options--1200x936.png)

## Installation

### Downloads

Download the latest version from the [releases page](https://github.com/panyingyun/optimus/releases).

Optimus is available for Windows, macOS, and Linux.

## Development

Optimus is built using [Wails](https://wails.app/) v2.11.0 and uses JavaScript on the frontend and Go on the backend.

### Prerequisites

- Go 1.22.0 or later
- Node.js and npm
- Wails v2.11.0

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/panyingyun/optimus.git
   cd optimus
   ```

2. Install Wails:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

3. Install system dependencies (Linux):
   
   For Ubuntu 22.04:
   ```bash
   sudo apt install libwebkit2gtk-4.0-dev libgtk-3-dev -y
   ```
   
   For Ubuntu 24.04:
   ```bash
   sudo apt install libwebkit2gtk-4.1-dev libgtk-3-dev -y
   ```

4. Install Go dependencies:
   ```bash
   go mod tidy
   ```

5. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   cd ..
   ```

### Development

Run the development server:

For Ubuntu 22.04:
```bash
make dev_ubuntu2204
# or
wails dev -tags webkit2_40
```

For Ubuntu 24.04:
```bash
make dev_ubuntu2404
# or
wails dev -tags webkit2_41
```

### Building

Build the application:

For Ubuntu 22.04:
```bash
make build_ubuntu2204
# or
wails build -tags webkit2_40
```

For Ubuntu 24.04:
```bash
make build_ubuntu2404
# or
wails build -tags webkit2_41
```

For Windows:
```bash
make build-windows
# or
wails build
```

## License

MIT &copy; 2020 Christopher Murphy
