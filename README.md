# LTP — Local Time Protocol (Port 2)
Network protocol to request and synchronize time across machines.

### Explanation
The master server has the **correct** time and distributes it on request. A slave server can ask the master server for the correct time and will synchronize accordingly. Max request size is **12 bytes**.

### Types of requests
- DATETIME
- DATE
- TIME
- MONTH
- DAY
- YEAR
- HOUR
- MINUTE
- SECOND
- NANOSECOND

#### Responses
- **DATETIME** responds with the date and time. Response:
    ```
    07-26-2021 02:09:33:572
    ```
- **DATE** responds with the date. Response:
    ```
    07-26-2021
    ```
- **TIME** responds with the time. Response:
    ```
    02:09:33:572
    ```
- **MONTH** responds with the time. Response:
    ```
    7
    ```
- **DAY** responds with the time. Response:
    ```
    26
    ```
- **YEAR** responds with the time. Response:
    ```
    2021
    ```
- **HOUR** responds with the time. Response:
    ```
    2
    ```
- **MINUTE** responds with the time. Response:
    ```
    9
    ```
- **SECOND** responds with the time. Response:
    ```
    33
    ```

- **NANOSECOND** responds with the time. Response:
    ```
    572
    ```