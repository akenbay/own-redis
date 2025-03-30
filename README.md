
# **Own-Redis Project**

## **Description**
This project implements a **simple key-value store** using **UDP protocol**. It mimics a basic version of Redis, allowing the client to store and retrieve data with commands like `SET`, `GET`, and `PING`. This server operates entirely in memory and can optionally expire keys after a specified duration using the `PX` flag for the `SET` command.

## **Learning Objectives**
- Understand and implement **UDP communication**.
- Work with **concurrent programming** using goroutines.
- Implement a **key-value store** and simulate key expiration using the `PX` flag.
- Learn basic **data synchronization** and thread safety using `sync.Map`.

## **Key Features**
- **`SET` Command**: Store a key-value pair with an optional expiration time (`PX` flag).
- **`GET` Command**: Retrieve a value by its key.
- **`PING` Command**: Check if the server is alive.
- **Concurrent Handling**: The server can handle multiple clients concurrently using **goroutines**.

## **Commands Supported**
1. **`PING`**:
   - Sends back a `"PONG"` response.
   - **Example**:
     ```sh
     PING
     PONG
     ```

2. **`SET`**:
   - Sets a key-value pair.
   - **With expiration (`PX`)**: The key will expire after a specified time in milliseconds.
   - **Example**:
     ```sh
     SET foo bar
     OK
     ```
     ```sh
     SET foo bar PX 1000
     OK
     ```
     After 1000 milliseconds, the key `foo` will be deleted.

3. **`GET`**:
   - Retrieves the value for a given key.
   - **Example**:
     ```sh
     GET foo
     bar
     ```

## **Project Setup**

### **1. Prerequisites**
- **Go 1.9+** is required to run the server, as we are using `sync.Map`.
- A **Linux**, **macOS**, or **Windows** environment.

### **2. Clone the Repository**
Clone this repository to your local machine:
```sh
git clone https://github.com/yourusername/own-redis.git
cd own-redis
```

### **3. Install Dependencies**
This project doesn't require external dependencies, but make sure you have **Go 1.9+** installed.

To check your Go version:
```sh
go version
```

### **4. Build the Project**
To build the project, run the following command from the root directory of the project:
```sh
go build -o own-redis .
```

### **5. Run the Server**
Start the server using:
```sh
./own-redis
```
The server will start on **port 8080** by default. You can specify a different port using the `--port` flag:
```sh
./own-redis --port 9090
```

### **6. Test the Server with Netcat**
You can use **netcat (`nc`)** to test the server by sending UDP messages:
```sh
echo -n "PING" | nc -u 127.0.0.1 8080
```

For **SET** and **GET** commands:
```sh
echo -n "SET foo bar" | nc -u 127.0.0.1 8080
echo -n "GET foo" | nc -u 127.0.0.1 8080
```

### **7. Server Logging**
- The server logs **errors** and **information** using the `slog` package.
- Logs are generated for successful `SET` operations, errors like wrong number of arguments, and `PING` responses.

---

## **Concurrency and Data Storage**
The server uses **`sync.Map`** to store key-value pairs in memory, ensuring safe concurrent access. When a key has an expiration (`PX`), a **goroutine** is spawned to delete the key after the specified time.

### **`sync.Map` in Action**
- The `SET` command with an expiration uses a goroutine that sleeps for the specified duration and then deletes the key.
- This allows for **non-blocking operations** and efficient handling of multiple clients.

---

## **Code Structure**

- **`main.go`**: The entry point where the UDP server is started and the request handler is set up.
- **`handler/`**: Contains the `Set`, `Get`, and `Ping` functions that handle the commands.
- **`data/`**: Defines the `sync.Map` for storing the key-value pairs.

### **Main Functions**

- **`Set`**: Stores a key-value pair and handles the `PX` expiration logic.
- **`Get`**: Retrieves a value by key.
- **`Ping`**: Responds with `PONG` to check if the server is alive.

---

## **Future Enhancements**
- **Persistence**: Implement data persistence (e.g., save the key-value pairs to a file) to retain data across server restarts.
- **Authentication**: Add basic authentication for accessing the key-value store.
- **More Commands**: Implement additional Redis-like commands, such as `DEL` (delete key) and `EXPIRE`.

---

## **License**
This project is open-source under the **MIT License**.

---

### **Conclusion**
This is a simple, lightweight **key-value store** built with **Go** and **UDP**. It helps you understand **networking**, **concurrent programming**, and **data storage** using in-memory maps.

Let me know if you need any further details or changes to the README! 😊
