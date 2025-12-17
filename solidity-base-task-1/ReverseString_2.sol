// SPDX-License-Identifier: MIT

pragma solidity ^0.8.31;


/**
翻转字符串
题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
**/
contract ReverseString_2 {

    function reverseString( string memory  xxx) external pure returns(   string memory _result) {
        bytes memory xxxByte = bytes(xxx);
        bytes memory resultBytes = new bytes(xxxByte.length);

        for (uint i = 0; i < xxxByte.length; i++) {
            resultBytes[i] = xxxByte[xxxByte.length - i -1];
        }
        _result = string(resultBytes);
     
    }

}