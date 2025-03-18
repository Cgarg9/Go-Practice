# URL Shortener

A simple URL shortener API built in Go. This project allows you to shorten long URLs and retrieve the original URL using a short code. It uses a basic file-based storage mechanism to save the mapping between short URLs and long URLs.

## Workflow

1. **Shorten URL**: A user provides a long URL via the `/shorten` endpoint. The application generates a unique short URL, stores the mapping, and returns the short URL.
2. **Redirect to Original URL**: When the short URL is accessed, the application redirects the user to the original long URL.
3. **Requests Per Minute Limit** : 5

## How to Run

### Prerequisites

- Go (v1.16 or higher) must be installed on your system.

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Cgarg9/Go-Practice.git
   cd GO_PRACTICE/urlShortener
   go mod init urlShortner
    ```
2. Install dependencies :
    ```bash
    go mod tidy
    ```
3. Run the application:
    ```
    go run main.go
    ```
The server will start on port *8080*.
4. Access the API:
- Shorten a URL: *POST http://localhost:8080/shorten*
- Redirect to the original URL: *GET http://localhost:8080/{short-url}*

## Endpoints

1. */shorten* (POST)
- *Description*: Accepts a long URL and returns a shortened URL.
- *Request Body:*
    ```bash 
    {
  "url": "https://example.com"
    }
    ```
- *Response:*
    ```bash
    {
  "short_url": "http://localhost:8080/abcd1234"
    }
    ```
- Status Codes:
    - *200 OK:* URL shortened successfully.
    - *400 Bad Request:* Invalid or missing URL in the request.
    - *500 Internal Server Error:* Unable to generate a unique short URL.

2. */* (GET)
- *Description:* Redirects to the original long URL using the short URL.
- *Request Example:* GET http://localhost:8080/abcd1234
- *Response:*
    - The server will redirect to the original URL.

## File Storage 

The application stores the mapping between short URLs and long URLs in a file named *store.json*. Each time the application runs, it loads the store data and persists any changes after generating a new short URL.

## Rate Limiting

'''bash
"golang.org/x/time/rate"
'''
-import the required package

'''bash
var rateLimiters = make(map[string]*rate.Limiter)
var rateMu sync.Mutex
const requestsPerMinute = 5
'''
- mapping rate limiter and declaring required rate 

'''bash
// Get client IP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Unable to determine IP", http.StatusInternalServerError)
		return
	}
'''

- get the ip of the user

'''bash
func isRateExceeded(ip string) bool {
	rateMu.Lock()
	defer rateMu.Unlock()

	limiter, exists := rateLimiters[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(requestsPerMinute)), requestsPerMinute)
		rateLimiters[ip] = limiter
	}

	return !limiter.Allow()
}
'''

-isRateExceeded function takes ip as input check whether rate is exceeded or not and returns a boolean


## Ways to Contribute

**📌 Task Description**
We welcome contributions to enhance the functionality and performance of this project. If you would like to contribute, here are some ways you can help:

**🛠 Steps to Implement**
1. Fork the repository.
2. Create a new branch: git checkout -b feature-name.
3. Implement the feature or fix the bug.
4. Run tests and ensure everything works.
5. Commit changes and push: git push origin feature-name.
6. Submit a Pull Request.

**📌 Additional Notes**
- Ensure code quality and follow best practices.
- Update this README file if needed to reflect any changes.
- Add tests where applicable.
- Ensure that the application handles edge cases, such as invalid URLs, gracefully.



