const express = require('express');
const dotenv = require('dotenv');

const catRoutes = require("./src/routes/catRoutes");
const authRoutes = require("./src/routes/authRoutes");
const { sequelize } = require('./src/models/userModel');


dotenv.config();

const app = express();

app.use(express.json());


sequelize.authenticate()
  .then(() => console.log('Database connected...'))
  .catch(err => console.log('Error: ' + err));


sequelize.sync();

app.use('/api/auth', authRoutes);
app.use('/api/cat', catRoutes);



module.exports = app;