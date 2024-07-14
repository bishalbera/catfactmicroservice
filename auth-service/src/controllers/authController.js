const jwt = require("jsonwebtoken");
const User = require("../models/userModel");

const generateToken = (id) => {
  return jwt.sign({id}, process.env.JWT_SECRET, { expiresIn: "30d" });
};

const registerUser = async (req, res) => {
    const {username, password } = req.body;
    
    try{
        const user = await User.create({username, password});
        
        res.status(201).json({
            id: user.id,
            username: user.username,
            token: generateToken(user.id),
        });
        
    } catch(error) {
        res.status(400).json({ message : error.message });
    }
};

const loginUser = async (req, res) => {
    const {username, password } = req.body;
    
    const user = await User.findOne({ where: {username} });

    if (user && (await user.matchPassword(password))) {
        res.json({
            id : user.id,
            username: user.username,
            token: generateToken(user.id),
        });
    } else {
        res.status(401).json({message: "Invalid credentials" });
    }

};

module.exports = {registerUser, loginUser};