# HackerNews Service

**For the HackerNews App built with Laravel, please visit the [HackerNews App repository](https://github.com/bimaagung/hackernews-app.git).**

## Installation

To get started with the HackerNews Service project, follow the steps below:

1. **Clone the Repository**

   Clone the project repository using the following command:

   ```
   git clone https://github.com/bimaagung/hackernews-service.git
   ```

2. **Navigate to the Project Directory**

   Navigate to the project directory using the following command:

   ```
   cd hackernews-service
   ```

3. **Install Dependencies**

   The project uses Go modules for dependency management. Use the following command to install the dependencies:

   ```
   go mod download
   ```

4. **Start the Project**

   Start the project using the provided Makefile target:

   ```
   make start
   ```

   This command will build and run the HackerNews Service.

   The HackerNews Service will be accessible at `http://localhost:8080`.

---


# API Documentation

This documentation provides details on how to use the API for retrieving news information.

**For open via Postman HackerNews API, please visit the [HackerNews API Postman](https://documenter.getpostman.com/view/20841766/2s93m8xeup).**

## Endpoint

### Retrieve News

**Request**

- URL: `GET /api/news`
- Headers:
  - `Content-Type: application/json`

**Response**

The API response will be in the following format:

```json
{
    "status": "ok",
    "message": "success",
    "data": [
        {
            "by": "truth_seeker",
            "descendants": 24,
            "id": 36106196,
            "total_comment": 9,
            "score": 60,
            "time": 1685296272,
            "title": "io_uring support for libuv – 8x increase in throughput",
            "url": "https://github.com/libuv/libuv/pull/3952"
        }
    ]
}
```

### Retrieve Story

**Request**

- URL: `GET /api/news/story/{storyId}`
- Headers:
  - `Content-Type: application/json`

**Response**

The API response will be in the following format:

```json
{
    "status": "ok",
    "message": "success",
    "data": {
        "by": "gilad",
        "descendants": 22,
        "id": 31636642,
        "comments": [
            {
                "by": "tony-allan",
                "id": 31637813,
                "comments": null,
                "parent": 31636642,
                "text": "Is Wi-Fi for ESP32 devices working yet? I didn't see anything specific on the supported boards page.",
                "time": 1654497627,
                "type": "comment"
            },
        ],
        "score": 182,
        "time": 1654480623,
        "title": "Tinygo: LLVM-based Go compiler for microcontrollers, WASM, and CLI tools",
        "type": "story",
        "url": "https://github.com/tinygo-org/tinygo"
    }
}
```

### Retrieve Comment

**Request**

- URL: `GET /api/news/story/comment/{commentId}`
- Headers:
  - `Content-Type: application/json`

**Response**

The API response will be in the following format:

```json
{
    "status": "ok",
    "message": "success",
    "data": {
        "by": "tony-allan",
        "id": 31637813,
        "comments": [
            {
                "by": "sp33k3rph433k",
                "id": 31643285,
                "comments": null,
                "parent": 31637813,
                "text": "I was recently working on a project and bumped into this. It's not currently supported, but here are the relevant Github links if you want to learn more (it doesn't look like there's much movement recently):\n- [Link 1](https://github.com/tinygo-org/tinygo/issues/1427)\n- [Link 2](https://github.com/tinygo-org/drivers/pull/320)",
                "time": 1654536763,
                "type": "comment"
            }
        ],
        "parent": 31636642,
        "text": "Is Wi-Fi for ESP32 devices working yet? I didn't see anything specific on the supported boards page.",
        "time": 1654497627,
        "type": "comment"
    }
}
```

## Documentation: Integration of Laravel Framework and Golang Service

This documentation provides an overview of how the HackerNews application integrates the Laravel framework with a Golang service for data processing. The integration is done using the HTTPS protocol, and the Golang service implements caching to enhance the retrieval of data from the HackerNews API.

### Architecture Overview

The HackerNews application follows a microservices architecture, with the frontend built using Laravel and the data processing handled by a Golang service. Here's an overview of the integration:

1. The Laravel application serves as the user interface, providing a user-friendly interface to browse, read, and interact with HackerNews articles and discussions.

2. When a user interacts with the application, Laravel sends requests to the Golang service for data processing. These requests include fetching top stories, story, comments from the HackerNews API.

3. The Golang service acts as a middleware between the Laravel application and the HackerNews API. It handles the communication with the API, performs data processing tasks, and caches the retrieved data.

4. To enhance performance, the Golang service implements caching using a caching mechanism such as Memcached. This caching layer helps speed up data retrieval from the HackerNews API by storing frequently accessed data locally.

5. The Golang service retrieves data from the HackerNews API and stores it in the cache. Subsequent requests for the same data can be served directly from the cache, reducing the response time and the load on the HackerNews API.

6. The Golang service then sends the processed data back to the Laravel application, which renders it for display to the user.

### Integration Steps

To integrate the Laravel framework with the Golang service, follow these steps:

1. Set up the Laravel application by following the installation instructions provided in the previous section. Make sure the Laravel application is configured to communicate with the Golang service via HTTPS.

2. Set up the Golang service by cloning the repository from [https://github.com/bimaagung/hackernews-service.git](https://github.com/bimaagung/hackernews-service.git).

3. Configure the Golang service to connect to the HackerNews API. Update the necessary configuration files with the API credentials and endpoints.

4. Implement the caching mechanism in the Golang service using a caching library like Memcached. Configure the caching options according to your preferences and requirements.

5. Integrate the Golang service into the Laravel application by making HTTP requests to the Golang service for data processing. Use the appropriate endpoints and payloads as defined in the Golang service's API documentation.

6. Handle the responses from the Golang service in the Laravel application and render the data to the user interface.

## Design Decision

For the Hackernews service built using Golang, I decided to implement the clean architecture. This decision was based on the fact that clean architecture provides a clear and organized structure, making it easy to scale and maintain the code. Clean architecture also enables the separation of business logic, infrastructure, and the use of the dependency inversion principle.

In the case of the Laravel application, I chose to use a simpler architecture, focusing on the existing MVC (Model-View-Controller) pattern provided by Laravel. This approach allows for faster and simpler development for this project.

## Challenges Faced

1. Dealing with a large and continuous data API required an efficient approach. To address this, I implemented asynchronous handling in Golang, allowing the application to fetch data faster.

2. One challenge was improving response performance using caching. I opted to use Memcached as the caching solution for this project. However, there were challenges in properly configuring and managing the cache to ensure consistent and accurate cached data.

3. Another challenge faced was the tight deadline for this project. In tackling this challenge, I focused on prioritizing and ensuring that core features functioned well within the given timeframe.

## Potential Improvements

1. One potential improvement is to switch the communication protocol from HTTP to gRPC. gRPC provides lower latency compared to HTTP and enables more efficient communication between services.

2. To enhance caching management, considering the use of Redis as the caching solution. Redis has built-in clustering features, allowing for horizontal data distribution across multiple Redis nodes. This can improve cache scalability and performance.

Through careful design evaluation, overcoming challenges faced, and exploring potential solution improvements, this project can continue to be developed and enhanced to deliver a better and more efficient application.

