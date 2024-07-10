import express from "express";
import dotenv from "dontenv";

dotenv.config();

const app = express();

app.use(express.json());


app.use("/api/auth", authRoutes);

module.exports = app;