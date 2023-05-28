# HackerNews Service

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
