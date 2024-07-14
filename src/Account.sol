// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "suave-std/Context.sol";
import "./IAccount.sol";
import "./utils/PrivateKeyManager.sol";

contract Account is IAccount, PrivateKeyManager {

  function signup(string memory email) external override returns (string memory token) {
    bytes memory password = Context.confidentialInputs();

    // TODO: validate creadentials exist.
    
    string memory generatedToken = string(abi.encodePacked(email, password));
    super.createPrivateKey(generatedToken);

    // TODO: implement validity for N blocks.
    return generatedToken;
  }

  function login(string memory email) external override returns (string memory token) {
    // TODO: implement validate creadentials exist.
    // TODO: implement validity for N blocks.
    bytes memory password = Context.confidentialInputs();
    return string(abi.encodePacked(email, password));
  }

  function sendOperation(address target, bytes memory data, uint value, string memory chainId) external override returns (bytes memory result) {
    bytes memory token = Context.confidentialInputs();
    bytes memory signedData = super.signTx(data, string(token), chainId);

    // TODO: implement send operation and return result.
    return signedData;
  }

  function sendRecoverEmail(string memory email) external pure override returns (bool success) {
    // TODO: implement
    return true;
  }

  function recover(string memory email) external pure override returns (string memory token) {
    // TODO: implement
    return  "token";
  }

  function getPubKey(string memory email) external pure override returns (address pubKey) {
    // TODO: implement
    return address(0);
  }
}