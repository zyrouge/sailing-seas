import { defineAppOperator } from "../operator";

export default defineAppOperator({
    operate: async (app) => {
        app.server.auth.strategy("session", "cookie", {
            cookie: {
                password: process.env.COOKIE_SECRET,
                isSecure: app.isProduction(),
            },
            redirectTo: "/login",
            appendNext: true,
            validate: (_: any, session: any) => {
                const token = session?.token;
                return {
                    isValid:
                        typeof token === "string" &&
                        app.store.isValidToken(token),
                };
            },
        });
        app.server.auth.default("session");
    },
});
