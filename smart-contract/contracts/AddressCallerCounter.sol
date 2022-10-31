// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract AddressCallerCounter {
    address public owner;

    mapping (address => uint) _addressCalled;

    struct User {
        address account;
        uint count;
    }

    event UserIncremented(
        address indexed account,
        uint indexed count
    );

    constructor () {
        owner = msg.sender;
    }

    modifier ownerOnly() {
        require(
            msg.sender == owner,
            "This function is restricted to the contract's owner"
        );
        _;
    }

    function incAddressCounter () public returns (uint) {
        uint counter = _addressCalled[msg.sender];

        counter++;

        _addressCalled[msg.sender] = counter;

        emit UserIncremented(msg.sender,_addressCalled[msg.sender]);

        return _addressCalled[msg.sender];
    }

    function getCountByAddress(address _address)
    public view
    ownerOnly
    returns(uint)
    {
        return _addressCalled[_address];
    }
}
