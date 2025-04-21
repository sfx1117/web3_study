// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
/*
    TodoList: 是类似便签一样功能的东西，记录我们需要做的事情，以及完成状态。
    1.需要完成的功能
    ①、 创建任务
    ②、 修改任务名称
        - 任务名写错的时候
    ③、 修改完成状态：
        - 手动指定完成或者未完成
    ④、 自动切换
        - 如果未完成状态下，改为完成
        - 如果完成状态，改为未完成
    ⑤、获取任务
*/
contract ToDoList{

    //声明任务结构体
    struct ToDo{
        string name;
        bool isComplated;
    }

    //声明todolist
    ToDo[] public todoList;
    //1、创建任务
    function create(string memory name) external  {
        todoList.push(
            ToDo(name,false)
        );
    }

    //2、修改任务名称 
    //方法一，直接修改
    function updateName(uint index,string memory name) external {
        todoList[index].name=name;
    }
    //方法二，先取出对象，再修改
    function changeName(uint index,string memory name) external {
        ToDo storage todo=todoList[index];
        todo.name=name;
    }

    //3、修改完成状态
    function updateState(uint index,bool isComplated) external {
        todoList[index].isComplated=isComplated;
    }

    //4、自动切换完成状态
    function changeState(uint index) external {
        todoList[index].isComplated=!todoList[index].isComplated;
    }

    //5、获取任务
    //方式一，使用memory，需要将string的完整数据复制到内存，消耗gas高
    function getTodo(uint index) external view returns (string memory,bool){
        ToDo memory todo=todoList[index];
        return (todo.name,todo.isComplated);
    }
    //方式二，使用storage，只需要将string的引用地址复制到内存，消耗gas低
    function getTodo2(uint index) external view returns (string memory,bool){
        ToDo storage todo=todoList[index];
        return (todo.name,todo.isComplated);
    }


}
