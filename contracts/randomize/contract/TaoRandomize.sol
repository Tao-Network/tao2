pragma solidity ^0.4.21;

import "./libs/SafeMath.sol";

contract TaoRandomize {
    using SafeMath for uint256;

    mapping (address=>bytes32[]) randomSecret;
    mapping (address=>bytes32) randomOpening;

    function TaoRandomize () public {
    }

    function setSecret(bytes32[] _secret) public {
        uint secretPoint =  block.number % 360;
        require(secretPoint >= 300);
        require(secretPoint < 330);
        randomSecret[msg.sender] = _secret;
    }

    function setOpening(bytes32 _opening) public {
        uint openingPoint =  block.number % 360;
        require(openingPoint >= 330);
        randomOpening[msg.sender] = _opening;
    }

    function getSecret(address _validator) public view returns(bytes32[]) {
        return randomSecret[_validator];
    }

    function getOpening(address _validator) public view returns(bytes32) {
        return randomOpening[_validator];
    }
}
