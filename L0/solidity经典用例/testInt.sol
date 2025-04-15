// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract testInt{

    function add(uint8 x,uint8 y) public pure returns (uint8){
        unchecked{

            return x+y;
        }

    }


}
