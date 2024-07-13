// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

interface IAccount {
  /**
	* creates a new private key and an associated token id in TEE.
  * private inputs:
  * - password
  */
	function signup(string memory email, string memory password) external returns (string memory token);
	
  /**
	* validates email and password in TEE and creates a new token 
	* id for using the private key
  * private inputs:
  * - password
  */
	function login(string memory email) external returns (string memory token);
	
  /**
	* sends any on chain operation using a token
	* uses the Gateway to send operations to the final chain
  * private inputs:
  * - token
  */
	function sendOperation(address target, bytes memory data, uint value) external returns (bytes memory result);
	
  /**
	* sends the email for recovering account
	* uses the Gateway to send emails off chain.
  */
	function sendRecoverEmail(string memory email) external returns (bool success);

  /**
  * recovers the account
	* private inputs:
  * - tempPass
	* - newPass
  */
	function recover(string memory email) external returns (string memory token);
	
	// Optional. we could also allow users to find the 
	// public key associated to some email, this could be 
	// useful for funding services
	function getPubKey(string memory email) external returns (address pubKey);
}