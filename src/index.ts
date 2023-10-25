import dotenv from "dotenv";
dotenv.config();

import { App } from "./app";

const start = async () => {
    const { HOST, PORT } = process.env;
    if (!HOST) throw new Error("Missing 'process.env.HOST'");
    if (!PORT) throw new Error("Missing 'process.env.PORT'");

    const app = new App({
        server: {
            host: HOST,
            port: PORT,
        },
    });
    await app.init();
    await app.start();
};

start();
