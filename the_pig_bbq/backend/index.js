var express = require("express");
var bodyParser = require('body-parser');
var app = express();

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

// CORS
app.use(function (req, res, next) {
    res.setHeader("Access-Control-Allow-Origin", "*");
    res.setHeader('Access-Control-Allow-Methods', "*");
    res.setHeader("Access-Control-Allow-Headers", "*");
    next();
});

app.listen(4242, '0.0.0.0', () => {
    console.log("Server running on port 4242");
});

app.get("/menu", (req, res, next) => {
    if (req.get("X-Authenticate") == "EI46mGgilJPWOJhwfmEqsrdbsyibzDz1") {
        res.json([
            { "item": "DDOS", "price": "40" },
            { "item": "Single Hash Password Cracking", "price": "15" },
            { "item": "Instagram Login", "price": "150" },
            { "item": "Corporate Backdoor", "price": "5,000" },
            { "item": "STM32F042K6T6", "price": "25" },
            { "item": "0days", "price": "10,000" },
            { "item": "Chrome Sandbox Escape", "price": "30,000" },
            { "item": "Cracked Floppy Drive Reader", "price": "5" },
            { "item": "Cult of the Dead Pig T-Shirt", "price": "15" },
            { "item": "SolarWinds Intern Coffee Cup", "price": "25" },
            { "item": "iOS Rootkit", "price": "500,000" },
            { "item": "flag{AHHmt-uZyqy4Y3FNHXi8YBuRmPL9P1Sh}", "price": "0" }
        ]);
    } else {
        res.json([
            { "item": "BBQ Pork Plate", "price": "8" },
            { "item": "BBQ Pork Sandwich", "price": "8" },
            { "item": "Smoked Chicken Plate", "price": "7" },
            { "item": "Pork Ribs", "price": "17" },
            { "item": "Fried Catfish", "price": "12" },
            { "item": "BLT Sandwich", "price": "6" },
            { "item": "Collard Greens", "price": "3" },
            { "item": "French Fries", "price": "3" },
            { "item": "Hushpuppies", "price": "3" },
            { "item": "Baked Beans", "price": "3" }
        ]);
    }
});

app.post("/order", (req, res, next) => {
    if (req.body.order == "flag" || req.body.order == "null") {
        res.json({ flag: "flag{EI46mGgilJPWOJhwfmEqsrdbsyibzDz1}" });
    } else {
        res.json({ "orderNumber": Math.floor(Math.random() * (99 - 10) + 10) });
    }
});

app.get("/flag", (req, res, next) => {
    res.json({ flag: 'flag{uHSgglifeO9tlIJkGGxcMC1tKXCRwLN5}' });
});

app.post("/login", (req, res, next) => {
    if (req.body.username = "admin") {
        res.json({ "token": "EI46mGgilJPWOJhwfmEqsrdbsyibzDz1" });
    } else {
        res.json({ "token": "none" });
    }
});
