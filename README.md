# Go Programming Language Exam

Fork this repository and make all changes to your own fork.

## Instructions:

Complete the following tasks in the Go programming language.

-   Write unit tests for all functions you implement.
-   There is no time limit.
-   Submit your code along with the unit tests.

## Task 1: Implement a Simple API

Create a small HTTP API with the following endpoints:

1. GET /hello
    - Returns a JSON response with a message: `{"message": "Hello, World!"}`.
2. POST /reverse:
    - Accepts a JSON payload with a single field text (string).
    - Returns a JSON response with the reversed string: `{"reversed": "<reversed_text>"}`.
3. GET /factorial/{n}:
    - Returns the factorial of a given integer n passed as a URL parameter.
    - If n is negative, return an error message: `{"error": "Input must be a non-negative integer"}`.

### Requirements:

-   Use the built-in net/http package to implement the API.
-   Handle errors gracefully and return appropriate HTTP status codes.
-   Ensure your code is clean, idiomatic, and easy to read/maintain.

## Task 2: Implement Unit Tests

Write unit tests for each endpoint in the API.

Ensure your tests cover:

-   Valid inputs.
-   Invalid inputs (e.g., negative numbers for the factorial).
-   Edge cases (e.g., empty strings, very large numbers).

### Requirements:

-   Use Go's testing package (testing) for writing the unit tests.
-   Ensure that your tests are easy to read and maintain.

## Task 3: Bonus

-   Implement a simple in-memory cache for the factorial results.
-   Ensure the API can handle multiple requests concurrently and the cache is thread-safe.

## Submission:

Submit a pull request to this repository with your changes.
