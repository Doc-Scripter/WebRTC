- [WebRTC Implementation in Go](#webrtc-implementation-in-go)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Project Structure](#project-structure)
  - [Dependencies](#dependencies)
  - [Setup](#setup)
  - [Usage](#usage)
  - [Technical Details](#technical-details)
    - [Signaling Server](#signaling-server)
    - [WebRTC Implementation](#webrtc-implementation)
  - [Contributing](#contributing)
  - [License](#license)
  - [Acknowledgments](#acknowledgments)


# WebRTC Implementation in Go

A real-time communication application built with Go and WebRTC technology, enabling peer-to-peer video/audio streaming and data sharing between browsers.

## Features

- Peer-to-peer video/audio streaming
- WebSocket-based signaling server
- ICE candidate handling
- Modern and responsive UI
- Real-time connection status updates

## Prerequisites

- Go 1.22.2 or higher
- Modern web browser with WebRTC support
- Basic understanding of WebRTC concepts

## Project Structure

```
.
├── handlers/         # Server-side request handlers
│   └── server.go    # WebRTC signaling and ICE candidate handling
├── static/          # Frontend assets
│   ├── index.html   # Main application page
│   ├── scripts.js   # Client-side WebRTC implementation
│   └── styles.css   # Application styling
├── main.go          # Application entry point
└── go.mod           # Go module definition
```

## Dependencies

- github.com/gorilla/websocket v1.5.3 - WebSocket implementation for signaling

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd webRTC
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `localhost:8080`

## Usage

1. Open your web browser and navigate to `http://localhost:8080`
2. Allow camera and microphone access when prompted
3. Share the URL with another peer to establish a connection
4. The application will automatically handle the WebRTC connection setup

## Technical Details

### Signaling Server
The application uses WebSocket for signaling, handling:
- Session Description Protocol (SDP) offers and answers
- ICE candidate exchange
- Connection state management

### WebRTC Implementation
- Implements the WebRTC API for peer-to-peer communication
- Handles media streams (video/audio)
- Manages ICE candidates for NAT traversal

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms specified in the LICENSE file.

## Acknowledgments

- WebRTC team for the amazing technology
- Gorilla WebSocket for the excellent WebSocket implementation
