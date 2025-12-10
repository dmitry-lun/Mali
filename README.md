# Mali - PE File Analysis CLI

CLI tool for analyzing PE files with risk level assessment.

## Architecture

The project is built on **Clean Architecture** principles with layer separation:

```
mali/
cmd/                    # Application entry point
     main.go            # Dependency initialization and CLI startup
internal/
     domain/            # Domain layer (business logic)
         entity/        # Entities (PEFile, Section, Import, Export)
         repository/    # Repository interfaces
         usecase/       # Use cases (AnalyzePEUseCase)
     infrastructure/    # Infrastructure layer (implementations)
         parser/        # PE file parsers (headers, sections, imports, exports)
         repository/    # Repository implementations
    presentation/       # Presentation layer
         cli/           # CLI commands (Cobra)
pkg/                    # Reusable components
    binary/             # Binary data utilities (RVA)
    detector/           # Detectors (packers, risk level)
    entropy/            # Entropy calculation
    file/               # File reading interface and implementation
    hash/               # Hashing (MD5, SHA256)
    reports/            # Report generation (JSON)
```

### Architecture Layers

1. **Domain Layer** (`internal/domain/`)
   - Contains business logic and entities
   - No dependencies on external libraries
   - Defines repository interfaces

2. **Infrastructure Layer** (`internal/infrastructure/`)
   - Implements PE file parsing
   - Implements repositories
   - Depends on domain layer

3. **Presentation Layer** (`internal/presentation/`)
   - CLI interface (Cobra)
   - Uses use cases from domain layer

4. **Package Layer** (`pkg/`)
   - Reusable utilities
   - Independent components

## Usage

```bash
# Analyze PE file
./mali analyze -f <file.exe> -o report.json

# Create test PE file
./scripts/create_test_pe.sh
```

## Development

### Commands

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Build
make build

# Format code
make fmt

# Run vet
make vet

# Run all checks (fmt, vet, test)
make check

# Clean build artifacts
make clean
```

## Features

- PE header parsing (DOS, File, Optional)
- Sections, imports, and exports extraction
- Section entropy calculation
- Packer detection (UPX, Themida, VMProtect)
- Hash computation (MD5, SHA256)
- Risk level assessment (SAFE, LOW, MEDIUM, HIGH)
- JSON report generation

## CI/CD

The project includes GitHub Actions workflows:

- **CI** (`.github/workflows/ci.yml`): Runs tests and builds on multiple platforms
- **Release** (`.github/workflows/release.yml`): Creates cross-platform releases on tag push

### Docker

Build and run with Docker:

```bash
docker build -t mali .
docker run --rm -v $(pwd):/data mali analyze -f /data/file.exe -o /data/report.json
```

