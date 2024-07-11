import express from "express";
import dotenv from "dontenv";

dotenv.config();

const app = express();

app.use(express.json());


sequelize.authenticate()
  .then(() => console.log('Database connected...'))
  .catch(err => console.log('Error: ' + err));


sequelize.sync();

app.use('/api/auth', authRoutes);


module.exports = app;