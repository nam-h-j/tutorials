import mysql from 'mysql';
import { DBuserInfo } from './dbUserInfo.js';

export const connection = mysql.createConnection({
    host: DBuserInfo.host,
    user: DBuserInfo.user,
    password: DBuserInfo.password,
    database: DBuserInfo.database,
});

connection.connect();

connection.query('SELECT now() AS time', function (err, rows, fields) {
    if (err) throw err;
    console.log('The time is: ', rows[0].time);
});