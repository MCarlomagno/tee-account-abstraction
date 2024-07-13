// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./IAccount.sol";

contract Account is IAccount {
  function signup(string memory email, string memory password) external override returns (string memory token) {
    // TODO: implement
    return "token";
  }

  function login(string memory email) external override returns (string memory token) {
    // TODO: implement
    return "token";
  }

  function sendOperation(address target, bytes memory data, uint value) external override returns (bytes memory result) {
    // TODO: implement
    return data;
  }

  function sendRecoverEmail(string memory email) external override returns (bool success) {
    // TODO: implement
    return true;
  }

  function recover(string memory email) external override returns (string memory token) {
    // TODO: implement
    return  "token";
  }

  function getPubKey(string memory email) external override returns (address pubKey) {
    // TODO: implement
    return address(0);
  }
}