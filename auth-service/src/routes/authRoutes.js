import express from "express";

const router = express.Router();

router.post("/register", registerUser);
router.post("/login", loginUser);

module.exports = router;
