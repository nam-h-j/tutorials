import { v4 as uuidv4 } from "uuid";
let users = [];
//get all users
export const getUsers = (req, res) => {
    console.log(`get request users`);
    res.send(users);
}

//get a user
export const getUser = (req, res) => {
    const { id } = req.params;
    const foundUser = users.find((user) => user.id === id);
    res.send(foundUser);
}

//create a user
export const createUser = (req, res) => {
    const user = req.body;
    users.push({ ...user, id: uuidv4() });
    res.send(`User with the name ${user.name} added to the database`);
}

//delete a user
export const deleteUser = (req, res) => {
    const { id } = req.params;
    users = users.filter((user) => user.id !== id);
    res.send(`User with the id ${id} deleted from the database`);
}

//update a user
export const updateUser = (req, res) => {
    const { id } = req.params;
    const { name, age } = req.body;
    const user = users.find((user) => user.id === id);
    if (name) {
        user.name = name;
    }
    if (age) {
        user.age = age;
    }
    res.send(`User with the id ${id} has been updated`);
}