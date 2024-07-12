import express from "express";
import dotenv from "dontenv";
const catRoutes = require("./routes/catRoutes");
const authRoutes = require("./routes/authRoutes");

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