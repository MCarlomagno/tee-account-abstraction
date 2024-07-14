
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "suave-std/suavelib/Suave.sol";

contract PrivateKeyManager {
  Suave.DataId signingKeyBid;

  function createPrivateKey(string memory token) internal {
    string memory keyData = Suave.privateKeyGen(Suave.CryptoSignature.SECP256);

    address[] memory peekers = new address[](1);
    peekers[0] = address(this);

    Suave.DataRecord memory bid = Suave.newDataRecord(10, peekers, peekers, "private_key");
    Suave.confidentialStore(bid.id, token, abi.encodePacked(keyData));

    signingKeyBid = bid.id;
  }

  function getPrivateKey(string memory token) internal returns (string memory) {
    bytes memory signingKey = Suave.confidentialRetrieve(signingKeyBid, token);
    return string(signingKey);
  }

  function signTx(bytes memory data, string memory token, string memory chainId) internal returns (bytes memory) {
    string memory pk = getPrivateKey(token);
    return Suave.signEthTransaction(data, chainId, pk);
  }
}