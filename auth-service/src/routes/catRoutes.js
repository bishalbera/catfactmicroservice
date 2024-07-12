import express from "express";


const router = express.Router();

router.get("/fact", protect, (req, res) => {
    getCatFact((error, response) => {
        if (error) {
            res.status(500).json({message: "Error fetching cat fact" });
        } else{
            res.json(response);
        }

    });
});

module.exports = router;