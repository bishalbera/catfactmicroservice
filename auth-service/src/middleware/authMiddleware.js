import { verify } from 'jsonwebtoken';
import { findByPk } from '../models/userModel';


const protect = async (req, res, next) => {
    let token;

    if (req.headers.authorization && req.headers.authorization.startsWith("Bearer")) {
        try{
            token = req.headers.authorization.split(" ")[1];
            const decoded = verify(token, process.env.JWT_SECRET);

            req.user = await findByPk(decoded.id);

            next()
        } catch (error) {
            res.status(401).json({ message: " Not authorized"});
        }
    }
    if (!token) {
        res.status(401).json({message: "Not authorized "});
    }
};

export default {protect};