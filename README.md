# SDN Core🌐

> A high-performance OpenFlow 1.3 SDN controller built with Go

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![OpenFlow](https://img.shields.io/badge/OpenFlow-1.3-green.svg)](https://opennetworking.org/sdn-resources/openflow/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Krios is a modern Software-Defined Networking (SDN) controller designed with Clean Architecture principles and Domain-Driven Design patterns. It provides robust OpenFlow 1.3 protocol support with intelligent Layer 2 switching capabilities, making it ideal for network virtualization and programmable networking scenarios.

## ✨ Key Features

🔌 **OpenFlow 1.3 Protocol**

- Full OpenFlow 1.3 specification compliance
- Secure handshake and feature negotiation
- Flow table management and packet processing

🔄 **Layer 2 Switching**

- Dynamic MAC address learning
- Intelligent forwarding decisions
- Broadcast flood handling
- Loop prevention mechanisms

🏗️ **Clean Architecture**

- Domain-driven design principles
- Separation of concerns
- Testable and maintainable codebase
- Package-by-component organization

🔒 **Security & Configuration**

- MAC address-based switch authorization
- Configurable datapath whitelisting
- Runtime monitoring and logging

## 🚀 Quick Start

### Prerequisites

Ensure you have the following installed:

- **Go 1.19+** - [Download here](https://golang.org/dl/)
- **Mininet** - For network topology testing
- **Git** - For version control

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/krios.git
   cd krios
   ```
2. **Initialize Go modules**

   ```bash
   go mod init arieoldman/arieoldman/krios
   go mod tidy
   ```
3. **Install dependencies**

   ```bash
   go get github.com/google/gopacket
   go get github.com/netrack/openflow
   go get github.com/golang/glog
   ```

### Running Krios

Start the controller with default settings:

```bash
go run main.go
```

The controller will:

- ✅ Initialize session and load configuration
- ✅ Set up OpenFlow 1.3 control plane
- ✅ Configure Layer 2 switching logic
- ✅ Listen on port 6633 for switch connections
- ✅ Display real-time status indicator

### Testing with Mininet

Create a test network topology:

```bash
# Start Mininet with OpenFlow 1.3 support
sudo mn --controller=remote,ip=127.0.0.1,port=6633 --switch ovsk,protocols=OpenFlow13

# Test connectivity
mininet> pingall
```

## 🏛️ Architecture Overview

Krios follows Clean Architecture principles with distinct layers:

```
┌─────────────────────────────────────────────────────────┐
│                    Infrastructure Layer                 │
│  ┌─────────────────┐  ┌─────────────────────────────┐   │
│  │   OpenFlow 1.3  │  │    Network Adapters         │   │
│  │   Protocol      │  │    (TCP, WebSocket)         │   │
│  └─────────────────┘  └─────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                               │
┌─────────────────────────────────────────────────────────┐
│                    Interface Layer                      │
│  ┌─────────────────┐  ┌─────────────────────────────┐   │
│  │   Controllers   │  │      Session Management     │   │
│  │   (REST/CLI)    │  │      (WebSocket, gRPC)      │   │
│  └─────────────────┘  └─────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                               │
┌─────────────────────────────────────────────────────────┐
│                    Use Case Layer                       │
│  ┌─────────────────┐  ┌─────────────────────────────┐   │
│  │   Handshake     │  │   Configuration Loading     │   │
│  │   Logic         │  │   Flow Management           │   │
│  └─────────────────┘  └─────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                               │
┌─────────────────────────────────────────────────────────┐
│                    Entity Layer                         │
│  ┌─────────────────┐  ┌─────────────────────────────┐   │
│  │   Config        │  │    Control Plane            │   │
│  │   Domain        │  │    Flow Tables              │   │
│  └─────────────────┘  └─────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
krios/
├── 📄 main.go                      # Application entry point
├── 📄 README.md                    # This file
├── 📄 go.mod                       # Go module definition
├── 📄 go.sum                       # Dependency checksums
├── 📂 common/                      # Shared utilities
│   └── 📄 types.go                 # Common type definitions
├── 📂 controller/                  # Interface adapters
│   └── 📄 session.go               # Session management
├── 📂 entity/                      # Domain entities
│   ├── 📄 config.go                # Configuration domain
│   └── 📄 controlplane.go          # Control plane interface
├── 📂 infrastructure/              # External interfaces
│   ├── 📄 base.go                  # Infrastructure base
│   └── 📄 ofp13cp.go               # OpenFlow 1.3 implementation
├── 📂 usecase/                     # Application business logic
│   ├── 📄 base.go                  # Use case base
│   ├── 📄 handshake.go             # Handshake orchestration
│   └── 📄 loadconfig.go            # Configuration loading
└── 📂 docs/                        # Documentation
    ├── 📄 L2SWITCH.md              # Layer 2 switching guide
    ├── 📄 DESIGNING_CONTROLLER.md  # Architecture documentation
    └── 📄 NOTES.md                 # Implementation notes
```

## 🔧 Configuration

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

## 🧪 Development

### Building

```bash
# Build for current platform
go build -o krios main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o krios-linux main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o krios.exe main.go
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestHandshake ./usecase/
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

## 📚 Documentation

- 📖 [**Layer 2 Switch Implementation**](L2SWITCH.md) - Detailed switching logic
- 🏗️ [**Controller Design**](DESIGNING_CONTROLLER.md) - Architecture decisions
- 📝 [**OpenFlow 1.3 Notes**](NOTES.md) - Protocol implementation details

## 🤝 Contributing

We welcome contributions! Please see our contributing guidelines:

1. 🍴 Fork the repository
2. 🌿 Create a feature branch (`git checkout -b feature/amazing-feature`)
3. 📝 Make your changes following Go conventions
4. ✅ Add tests for new functionality
5. 🔍 Run tests and ensure they pass
6. 📤 Submit a pull request

### Code Standards

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Write comprehensive tests
- Document public APIs
- Use meaningful commit messages

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **OpenFlow 1.3 Specification** - [Open Networking Foundation](https://opennetworking.org/)
- **Clean Architecture** - Robert C. Martin
- **Domain-Driven Design** - Eric Evans
- **Go Programming Language** - The Go Team

## 🔗 Links

- [OpenFlow Specification](https://opennetworking.org/sdn-resources/openflow/)
- [Mininet Documentation](http://mininet.org/)
- [Go Documentation](https://golang.org/doc/)

---

<div align="center">
  <sub>Built with ❤️ by the Krios team</sub>
</div>
