// SPDX-License-Identifier: MIT

pragma solidity ^0.8.31;


/**
✅  合并两个有序数组 (Merge Sorted Array)
题目描述：将两个有序数组合并为一个有序数组。
**/
contract MergeSortedArray_5 {

    function mergeSortedArray( uint[] calldata arr1,uint[] calldata arr2 ) external pure returns( uint[] memory _mergeResult) {
        uint  x = 0;
        uint  y = 0;
        uint z = 0;
        _mergeResult = new uint[](arr1.length + arr2.length);
        while (x < arr1.length || y < arr2.length) {
            if (x < arr1.length && (y >= arr2.length || arr1[x] <= arr2[y])) {
                _mergeResult[z] = arr1[x];
                x++;
            } else {
                _mergeResult[z] = arr2[y];
                y++;
            }
            z++;
        }
    }

}