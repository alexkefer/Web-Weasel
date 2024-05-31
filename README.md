# P2P Web Cache

## Overview

Welcome to the P2P Web Cache project! This project is developed by Alex Kefer, Keagan Edwards, Ryan Martin, and Khang Tran as part of our Senior Project at WWU. Originally built by Keagan Edwards and later modified by Alex Kefer, the current version of the project is 1.0.0.

The P2P Web Cache project aims to create a decentralized web caching system where peers can share cached web pages, improving the efficiency and speed of web access.

Sure, here's how you can specify the installation instructions for Go and npm in your README:

## Dependencies

Before running the program, make sure you have the following dependencies installed:

### Go

To install Go, follow these steps:

1. Download the Go distribution from the [official Go website](https://golang.org/dl/).
2. Follow the installation instructions for your operating system.
3. Verify the installation by running the following command in your terminal:
   ```sh
   go version
   ```
   If installed correctly, this command should display the installed Go version.
   
   **Note:** A working version for the project is Go 1.22.1.

### Node.js and npm

To install Node.js and npm, follow these steps:

1. Download the Node.js installer from the [official Node.js website](https://nodejs.org/).
2. Follow the installation instructions for your operating system.
3. Verify the installation by running the following commands in your terminal:
   ```sh
   node --version
   ```
   ```sh
   npm --version
   ```
   If installed correctly, these commands should display the installed Node.js and npm versions respectively.
   
   **Note:** A working version for the project is Node.js v21.7.3 and npm 10.5.0.

## Setting up the environment

### Back-End Server 

Firstly navigate to the backend server, use the following command:
#### For Window
```sh
cd .\p2psearch-backend        
```
#### For Linux
```sh
cd p2psearch-backend      
```
To run the back-end, use the following command:

```sh
go run .
```

### Front-end Extension 

Firstly navigate to the frontend server, use the following command:
#### For Window
```sh
cd .\p2psearch-frontend-vreact       
```
#### For Linux
```sh
cd p2psearch-frontend-vreact     
```
To set up the enviroment for frontend development, use the following command:
```sh
npm install
```

To construct the front-end, use the following command:
```sh
npm run build
```
## Features

- **Peer Discovery:** Automatically discover and store a list of available hosts in the network.
- **Webpage Caching:** Download and cache static webpages.
- **Peer Communication:** Connect to and communicate with other peers to share cached pages.
- **Disconnect Functionality:** Ability to disconnect from peers gracefully.
- **Frontend Interface:** A user-friendly interface to manage the cache and connections.
- **Website Request:** Request specific webpages from peers.

## To Do

- [x] **Build functionality to download webpages:** Implement the ability to download and cache static webpages.
- [x] **Design frontend interface:** Create a user-friendly frontend to manage the P2P web cache.
- [x] **Disconnect Functionality:** Develop the ability to disconnect from other peers.
- [x] **Store List of available hosts:** Implemented the functionality to keep track of available peers in the network.
- [x] **Store List of websites downloaded on the available host:** Develop the functionality to store and share the list of cached websites on each peer.
- [x] **Request Website functionality:** Implement the ability to request specific webpages from other peers.

## Contribution

We welcome contributions to the P2P Web Cache project. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your changes to your forked repository.
5. Create a pull request to the main repository.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact

For any questions or suggestions, please contact us at:
2023-2024 Group
- Alex Kefer: [alex.kefer@example.com](mailto:alex.kefer@example.com)
- Keagan Edwards: [keaganmedwards@gmail.com](mailto:keaganmedwards@gmail.com)
- Ryan Martin: [ryan.business.work@gmail.com](mailto:ryan.business.work@gmail.com)
- Khang Tran: [khangnguyentran.it@gmail.com](mailto:khangnguyentran.it@gmail.com)

Thank you for checking out the P2P Web Cache project! We look forward to your feedback and contributions.


---