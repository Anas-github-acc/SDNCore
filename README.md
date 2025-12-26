# SDN Core

> A high-performance OpenFlow 1.3 SDN controller built with Go

*A Software-Defined Networking (SDN) controller build using Clean Architecture and Domain-Driven Design. It features comprehensive OpenFlow 1.3 support and advanced Layer 2 switching, making it a prime choice for network virtualization and programmable network applications.*

## Key Features

**OpenFlow 1.3 Protocol**

- Full OpenFlow 1.3 specification compliance
- Secure handshake and feature negotiation
- Flow table management and packet processing

**Layer 2 Switching**

- Dynamic MAC address learning
- Intelligent forwarding decisions
- Broadcast flood handling
- Loop prevention mechanisms

**Clean Architecture**

- Domain-driven design principles
- Separation of concerns
- Testable and maintainable codebase
- Package-by-component organization

**Security & Configuration**

- MAC address-based switch authorization
- Configurable datapath whitelisting
- Runtime monitoring and logging

## Quick Start

### Prerequisites

Ensure you have the following installed:

- **Go 1.19+** - [Download here](https://golang.org/dl/)
- **Mininet** - For network topology testing

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Anas-github-acc/SDNCore.git
   cd SDNCore
   ```
2. **Initialize Go modules**

   ```bash
   go mod init sdn/sdncore
   go mod tidy
   ```
3. **Install dependencies**

   ```bash
   go get github.com/google/gopacket
   go get github.com/netrack/openflow
   go get github.com/golang/glog
   ```

### Running

Start the controller with default settings:

```bash
go run main.go
```

The controller will:

- Initialize session and load configuration
- Set up OpenFlow 1.3 control plane
- Configure Layer 2 switching logic
- Listen on port 6633 for switch connections
- Display real-time status indicator

### Testing with Qemu

Create a test network topology:

```bash
# Start Mininet with OpenFlow 1.3 support
sudo mn --controller=remote,ip=127.0.0.1,port=6633 --switch ovsk,protocols=OpenFlow13

# Test connectivity
mininet> pingall
```

### Switch Authorization

Configure which switches can connect by modifying the `Config` entity:

```go
conf := entity.Config{
     DPIDs: []common.EthAddr{
          {0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, // Switch 1
          {0x00, 0x00, 0x00, 0x00, 0x00, 0x02}, // Switch 2
     },
}
```

### Port Configuration

By default, Krios listens on port 6633. To change this:

```go
cp.Start(6653) // Custom port
```

## Development

### Building

```bash
# Build for current platform
go build -o sdn main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o sdn-linux main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o sdn.exe main.go
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code (requires golangci-lint)
golangci-lint run

# Vet code
go vet ./...
```

## Contributing

We welcome contributions! Please see our contributing guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following Go conventions
4. Add tests for new functionality
5. Run tests and ensure they pass
6. Submit a pull request

## 🔗 Links

- [OpenFlow Specification](https://opennetworking.org/sdn-resources/openflow/)
- [Mininet Documentation](http://mininet.org/)
- [Go Documentation](https://golang.org/doc/)
