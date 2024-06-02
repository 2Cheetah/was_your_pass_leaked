```mermaid
sequenceDiagram
    autonumber
    actor client
    client->>deriv: password
    deriv->>service: isLeaked?
    service->>DB: available in DB?
    service->>service: password available in DB?
    DB->>HIBP: get list of hashes
    HIBP->>service: list of hashes
    service->>service: match with the list
    service->>deriv: response
    deriv->>client: message
```