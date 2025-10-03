You can test APIs using **cURL** by executing command-line requests that simulate the various methods (GET, POST, PUT, DELETE) typically used in API interactions. Here's how to do it:

## Basic cURL Commands

### 1. **GET Request**
To fetch data from an API:
```bash
curl -X GET "https://api.example.com/resource"
```

### 2. **POST Request**
To send data to an API:
```bash
curl -X POST "https://api.example.com/resource" -H "Content-Type: application/json" -d '{"key1":"value1", "key2":"value2"}'
```

### 3. **PUT Request**
To update existing data:
```bash
curl -X PUT "https://api.example.com/resource/1" -H "Content-Type: application/json" -d '{"key1":"newValue1"}'
```

### 4. **DELETE Request**
To remove data:
```bash
curl -X DELETE "https://api.example.com/resource/1"
```

## Additional Options

### 1. **Custom Headers**
To include custom headers, use the `-H` option:
```bash
curl -X GET "https://api.example.com/resource" -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### 2. **Verbose Output**
To see detailed information about the request/response:
```bash
curl -X GET "https://api.example.com/resource" -v
```

### 3. **Saving Response**
To save the response to a file:
```bash
curl -X GET "https://api.example.com/resource" -o response.json
```

### 4. **Using Query Parameters**
To pass query parameters in a GET request:
```bash
curl -X GET "https://api.example.com/resource?param1=value1&param2=value2"
```

### 5. **Handling HTTPS**
If the API uses HTTPS and you encounter SSL issues, you can bypass SSL verification (not recommended for production):
```bash
curl -X GET "https://api.example.com/resource" -k
```

---

Using cURL provides a robust and lightweight way to test APIs directly from the command line without needing a web interface like Postman. Each command can be modified to fit the specific API you’re working with. Would you like additional examples or specific scenarios?