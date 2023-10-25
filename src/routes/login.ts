import { defineAppOperator } from "../operator";
import { ejsTemplate } from "./templates/render";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: ["GET", "POST"],
            path: "/login",
            options: {
                auth: false,
            },
            handler: async (req, h) => {
                const passcode =
                    typeof req.payload === "object" && "passcode" in req.payload
                        ? req.payload.passcode
                        : undefined;
                if (
                    typeof process.env.PASSCODE === "string" &&
                    process.env.PASSCODE === passcode
                ) {
                    const token = app.store.createToken();
                    req.cookieAuth.set({ token });
                    return h.redirect(req.query.next ?? "/");
                }
                const html = await ejsTemplate("login.ejs");
                return h.response(html).type("text/html");
            },
        });
    },
});
