# Webhook Project in Go

This project implements a simple and efficient webhook server in Go. The webhook server listens for incoming HTTP POST requests, processes the payload, and provides a customizable response. This setup is ideal for integrating with third-party services or triggering workflows in your own systems.

---

## Features

- **Simple Setup**: Lightweight and easy to deploy.
- **Customizable Payload Processing**: Add your own logic for handling incoming data.
- **Efficient HTTP Server**: Built using Go's `net/http` package for high performance.
- **Logging**: Logs incoming requests and responses for easy debugging.
- **Secure**: Option to validate payload signatures (e.g., HMAC).

---

## Requirements

- **Go**: Version 1.19 or higher

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/webhook-go.git
   cd webhook-go
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the project:

   ```bash
   go build -o webhook-server
   ```

---

## Usage

1. Run the server:

   ```bash
   ./webhook-server
   ```

2. By default, the server listens on port `8080`. You can configure the port by setting the `PORT` environment variable:

   ```bash
   PORT=9090 ./webhook-server
   ```

3. The webhook endpoint is available at `/webhook`. Example:

   ```bash
   curl -X POST http://localhost:8080/webhook \
        -H "Content-Type: application/json" \
        -d '{"message": "Hello, Webhook!"}'
   ```

---

## Configuration

You can customize the server behavior by editing the `config.json` file:

```json
{
  "port": 8080,
  "logRequests": true,
  "validateSignature": true,
  "secret": "your-secret-key"
}
```

- **port**: Port on which the server listens.
- **logRequests**: Enable or disable request logging.
- **validateSignature**: Enable or disable payload signature validation.
- **secret**: Shared secret for validating signatures (HMAC).

---

## File Structure

```
.
├── main.go              # Entry point of the application
├── handler.go           # Logic for handling incoming requests
├── config.json          # Server configuration file
├── go.mod               # Go module file
├── go.sum               # Dependencies checksum file
└── README.md            # Project documentation
```

---

## Development

To run the server in development mode with hot-reloading:

1. Install `air` (a live reloading tool for Go):

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Start the server with:

   ```bash
   air
   ```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message"
   ```
4. Push the branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

Special thanks to the Go community for providing excellent resources and inspiration for this project.

