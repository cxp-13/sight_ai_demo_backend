# Pre requirements
- Start local 3000 port computing service  
# Quick start

1. Rename the `config/config.ini.example` file by removing the `.example` suffix, making it `config/config.ini`.

2. In the `config/config.ini` file, fill in the following parameters:
   ```
   test_addr_1=<your_wallet_address>
   test_private_1=<your_private_key>
   ```
   Replace `<your_wallet_address>` with your wallet address and `<your_private_key>` with your private key.

3. To start the project, run the following command:
   ```bash
   go run cmd/main.go
   ```
4. Interact with API on api.md
`http://localhost:8080/compute`


