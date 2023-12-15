# Go Blockchain Implementation

This repository contains a simple yet functional implementation of a blockchain in Go. The project demonstrates the fundamental concepts of blockchain technology, including block creation, mining, and maintaining a chain of blocks with integrity.

## Features

- **Block Creation:** Ability to create new blocks with custom data.
- **Blockchain Integrity:** Ensures the integrity of the blockchain by linking blocks with their previous hash.
- **Mining Simulation:** Implements a basic proof-of-work algorithm for mining new blocks.
- **JSON Serialization:** Supports loading and saving the blockchain to and from a JSON file.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What you need to install the software:

- Go (version 1.15 or later recommended)
- Git (for cloning the repository)

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/araujo88/blockchain-golang.git
   cd blockchain-golang
   ```

2. **Run the program:**

   ```bash
   go run .
   ```

   This will execute the main program, which might include mining some blocks and displaying the blockchain.

3. **Run tests:**

   ```bash
   go test
   ```

   This will run the unit tests defined in the project, ensuring the core functionalities work as expected.

## Usage

Describe how to use the blockchain:

- **Creating a new blockchain:**

  ```go
  blockchain := NewBlockchain(difficulty)
  ```

- **Adding a block to the blockchain:**

  ```go
  blockchain.AddBlock("Your block data here")
  ```

- **Saving the blockchain to a file:**

  ```go
  err := blockchain.SaveToFile("blockchain.json")
  ```

- **Loading the blockchain from a file:**

  ```go
  blockchain, err := LoadBlockchainFromFile("blockchain.json")
  ```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
