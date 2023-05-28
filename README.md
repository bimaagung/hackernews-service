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
