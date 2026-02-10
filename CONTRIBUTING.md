# Contributing to E-Commerce API

First off, thank you for considering contributing to this project! It's people like you that make this such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps which reproduce the problem**
- **Provide specific examples to demonstrate the steps**
- **Describe the behavior you observed after following the steps**
- **Explain which behavior you expected to see instead and why**
- **Include logs or screenshots if possible**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

- **Use a clear and descriptive title**
- **Provide a step-by-step description of the suggested enhancement**
- **Provide specific examples to demonstrate the steps**
- **Describe the current behavior and expected behavior**
- **Explain why this enhancement would be useful**

### Pull Requests

- Follow the Go code conventions and style guide
- Include appropriate test cases
- Update documentation as needed
- End all files with a newline

## Development Setup

1. **Fork and clone the repository**:
   ```bash
   git clone https://github.com/your-username/ecom-api.git
   cd ecom-api
   ```

2. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Set up development environment**:
   ```bash
   cp .env.example .env
   docker-compose up -d
   ```

4. **Make your changes and test**:
   ```bash
   make test
   make lint
   ```

5. **Commit with clear messages**:
   ```bash
   git commit -m "feat: add descriptive message"
   ```

6. **Push and create a pull request**:
   ```bash
   git push origin feature/your-feature-name
   ```

## Code Style

- Run `gofmt` on all files: `make fmt`
- Run linters: `make lint`
- Keep functions small and focused
- Add comments for exported functions
- Write tests for new functionality

## Testing

- Write unit tests for all new features
- Ensure all tests pass: `make test`
- Aim for >80% code coverage for new code

## Git Commit Messages

- Use the present tense ("add feature" not "added feature")
- Use the imperative mood ("move cursor to..." not "moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
Scopes: `orders`, `products`, `api`, `db`, etc.

Example:
```
feat(orders): add order pagination support

- Implement limit/offset pagination
- Add query parameters to GET /orders
- Update OpenAPI documentation

Closes #123
```

## Documentation

- Update README.md when adding features
- Add/update API documentation
- Include examples in docstrings
- Update CHANGELOG.md

## Community

- Be respectful and inclusive
- Help others when possible
- Provide constructive feedback

Thank you for contributing! 🎉
