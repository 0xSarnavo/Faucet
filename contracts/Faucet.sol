// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

// Faucet contract allows a designated signer to distribute a limited amount of tokens to users.
// It includes functionality for deposits, withdrawals by the owner, and token distribution
// to users who haven't previously received tokens.
contract Faucet {

    // The address of the contract owner who deployed the contract (immutable after deployment)
    address immutable owner;

    // The address of the signer allowed to authorize token distribution (immutable after deployment)
    address immutable signer;

    // Mapping to track if an address has already received tokens (prevents double withdrawal)
    mapping(address => bool) received;

    // Mapping to track the amount of tokens each address has received (can be expanded in future use cases)
    mapping(address => uint) amount;

    // The maximum token limit per distribution (set to 0.1 ETH in wei)
    uint tokenLimit = 100000000000000000;

    // Constructor sets the contract owner to the deployer and the signer to a specified address.
    // The contract is payable, allowing it to receive an initial balance during deployment.
    constructor(address _signer) payable  {
        owner = msg.sender;
        signer = _signer;
    }

    // Modifier to restrict certain functions to the owner of the contract only
    modifier onlyOwner {
        require(msg.sender == owner, "Only Owner can withdraw");
        _;
    }
    
    // Function to return the contract's current balance.
    // Can be called by anyone to check how much ether is stored in the contract.
    function getBalance() external view returns (uint) {
        return address(this).balance;
    }

    // External function allowing any address to deposit ether into the contract.
    // The function is marked as `payable` to enable ether transfers.
    function deposit() external payable {
    }

    // Function allowing the contract owner to withdraw all the ether stored in the contract.
    // The withdrawal transfers the entire contract balance to the owner.
    // The `onlyOwner` modifier ensures that only the owner can call this function.
    function withdraw() external payable onlyOwner returns (bytes memory) {
        // Transfer all contract balance to the owner and ensure success
        (bool sent, bytes memory data) = owner.call{value: address(this).balance}("Withdraw Successful");
        require(sent, "Failed to withdraw");
        return data; // Return transaction data
    }

    // Function allowing the signer to distribute tokens (ether) to a specified receiver.
    // Restrictions: 
    // 1. The signer must be the one initiating the transfer.
    // 2. The receiver must not be the signer.
    // 3. The receiver must not have previously received tokens.
    function sendToken(address _receiver) external payable returns (bytes memory) {
        // Check if the sender is the signer and the receiver isn't the signer
        require(signer == msg.sender && signer != _receiver, "Unauthorized action");

        // Ensure the receiver hasn't already received tokens
        require(received[_receiver] == false, "Receiver already received tokens");

        // Transfer the predefined token limit (0.1 ETH) to the receiver
        (bool sent, bytes memory data) = _receiver.call{value: tokenLimit}("Withdraw Successful");
        require(sent, "Failed to send tokens");

        // Mark the receiver as having received tokens to prevent future withdrawals
        received[_receiver] = true;

        return data; // Return transaction data
    }

}
