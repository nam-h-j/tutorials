import express from "express";
import { getUsers, getUser, createUser, deleteUser, updateUser } from "../controllers/users.js";

const router = express.Router();

//이곳에 정의된 라우터는 /users로 시작함

//get all users
router.get("/", getUsers);

//get a user
router.get('/:id', getUser);

//create a user
router.post('/', createUser);

//delete a user
router.delete('/:id', deleteUser);

//update a user
router.patch('/:id', updateUser);

export default router;