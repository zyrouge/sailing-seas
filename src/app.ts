import { randomUUID } from "crypto";
import Hapi from "@hapi/hapi";
import HapiCookie from "@hapi/cookie";
import HapiPino from "hapi-pino";
import { AppOperator } from "./operator";

import rAuth from "./routes/_auth";
import rIconPng from "./routes/static/icon.png";
import rStylesCss from "./routes/static/styles.css";
import rLogin from "./routes/login";
import rHome from "./routes/home";
import rNyaa from "./routes/nyaa";

export interface AppOptions {
    server: {
        host: string;
        port: string;
    };
}

export class App {
    store: AppStore;
    server: Hapi.Server;

    constructor(public options: AppOptions) {
        this.store = new AppStore();
        this.server = Hapi.server({
            ...this.options.server,
        });
    }

    async init() {
        await this.server.register(HapiCookie);
        await this.server.register(HapiPino);
        for (const x of App.operators) {
            await x.operate(this);
        }
    }

    async start() {
        await this.server.start();
    }

    isProduction() {
        return process.env.NODE_ENV === "production";
    }

    static operators: AppOperator[] = [
        rAuth,
        rIconPng,
        rStylesCss,
        rLogin,
        rHome,
        rNyaa,
    ];
}

export class AppStore {
    tokens: Map<string, number>;

    constructor() {
        this.tokens = new Map();
    }

    createToken() {
        const token = randomUUID();
        const expiresAt = Date.now() + 900000;
        this.tokens.set(token, expiresAt);
        return token;
    }

    isValidToken(token: string) {
        const now = Date.now();
        const expiresAt = this.tokens.get(token);
        if (expiresAt && now < expiresAt) {
            return true;
        }
        this.tokens.delete(token);
        return false;
    }
}
