# Transactions Routine
A service to store all the customer transactions. Each cardholder (customer) has an account with their data. For each operation done by the customer a transaction is created and associated with their respective account. Each transaction has a specific type (normal purchase, withdrawal, credit voucher or purchase with installments)

Transactions of type purchase and withdrawal are registered with negative amounts, while transactions of credit voucher are registered with positive value.

### API documentation  
For detail API documentation refer file `swagger/swagger.yaml`.  
Open this file in swagger editor.

### How to run the application locally
Test the application: `make test`  
Run the containerized application locally: `make run`  
Note: For dev environment it will create app docker image every time we run the app.

### TODO:
1. Write integration tests  
   For now this application has only unit test cases which are testing core business logic. Write an integration tests which will test end to end flow of every API. In integration tests use real DB and not mocked one.
2. Have a proper logger  
   For now application is logging the information on console directly. Have a logger which will store logs in log file along with metadata info like log level, request UUID etc. Log file should persist even if container dies.
3. Generate the modules using swagger code generator.  
   API requests and responses modules are manually written. Use swagger code generated to generate the modules.
