# Retry Mechanism with Exponential Backoff in Go

This folder provides a reusable implementation of a **retry mechanism** with **exponential backoff** in Go. It demonstrates how to handle transient errors (e.g., network issues or server errors) gracefully by retrying operations with increasing delays between attempts.

## Features

- **Retry Logic**: Retries failed operations up to a specified number of attempts.
- **Exponential Backoff**: Implements exponential delays between retries to avoid overwhelming the system.
- **Customizable Retry Conditions**: Allows defining which errors are retryable.
- **Decoupled Design**: The retry logic is decoupled from the core operation, making it reusable across different use cases.

## Usage

### Example: Retrying an HTTP Request