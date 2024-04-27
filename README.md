# How the service works

1. Receives a request with a password
2. Checks local DB for a record if such password was already been checked
3. Returns `true` if a password was found - means it appeared in leakages
4. Sends request to https://haveibeenpwned.com/API/v3 to check for appearance in leakages
5. In case not found - returns `false`

# Endpoints

OpenAPI link here

GET `/checkPassword`
GET `/ping`