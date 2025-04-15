
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract StructExample{
    //结构体中是不能存在mapping的
    struct User{
        uint8 age;
        string name;
        // mapping (string =>string) fileMap;//后面写的函数会报错
        Hobbie[] hobbies;

    }

    struct Hobbie{
        string lev;
        string value;
    }

    User[] public users;
    function addUser(User calldata user) public {
        users.push(user);
    }

    function getUser(uint i) public view returns(User memory user){
        return users[i];
    }

    //将user对象 存到map中
    mapping(string name=> User user) public map;

    function setMap(string memory name,User calldata user) public {
        map[name]=user;
    }

    
}