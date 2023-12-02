import express from "express";
import { database } from "./index.js";
import { validateJWTadmin } from "./middlewares/validateJWTadmin.js";
import { itemsRouter } from "./routes/itemsToures.js";
import { loginRouter } from "./routes/loginRoutes.js";
import { registerRouter } from "./routes/registerRoutes.js";

export const router = express.Router();
router.use("/login", loginRouter);
router.use("/register", registerRouter);
router.use("/items", itemsRouter);

router.delete("/user", validateJWTadmin, async (req, res) => {
    console.log(`removing user`);
    const { _id } = req.body;
    console.log(req.body);
    if (!_id) return res.status(400).json({ status: "not enough data provided" });
    try {
        const dbuser = await database.getUserByID(_id, "users");
        if (dbuser == null) {
            console.error(`${_id} user does not exist`);
            return res.status(404).json({ status: "user does not exist" });
        }
        await database.releseUserItems(_id, "items");
        await database.deleteUser(_id, "users");
        return res.json({ status: "removed user properly", _id });
    } catch (err) {
        console.error(err);
        return res.status(500).json({ status: "internal server error" });
    }
});
