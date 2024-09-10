Create a crypto notification service as an HTTP Rest API server


Create a Notification (Input parameters: Current Price of Bitcoin, Daily Percentage Change, Trading Volume, etc)

Send a notification to email/emails

List notifications (Sent, Pending, Failed)

Update/Delete notification

To run the code:
`go run cmd/main.go`

DB used: postgres

created the table using
```
CREATE TABLE notifications
(
    id            SERIAL PRIMARY KEY,
    current_price FLOAT NOT NULL,
    daily_change_percentage FLOAT NOT NULL ,
    trading_volume INT NOT NULL ,
    emails TEXT[] NOT NULL ,
    status VARCHAR(50) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```