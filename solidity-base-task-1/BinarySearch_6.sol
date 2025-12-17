// SPDX-License-Identifier: MIT

pragma solidity ^0.8.31;


/**
✅  二分查找 (Binary Search)
题目描述：在一个有序数组中查找目标值
**/
contract BinarySearch_6 {

    function binarySearch(uint[] calldata arr, uint target) external pure returns(bool,uint) {
        
         uint left = 0;
        uint right = arr.length - 1;

        while (left <= right) {
            uint mid = left + (right - left) / 2;
            if (arr[mid] == target) {
                return (true, mid); 
            } else if (arr[mid] < target) {
                left = mid + 1;      
            } else {
                right = mid - 1;     
            }
        }
        //未找到
        return (false, 0);  

    }

}