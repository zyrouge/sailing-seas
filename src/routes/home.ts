import { defineAppOperator } from "../operator";
import { tr } from "./templates/renderer";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: "GET",
            path: "/",
            handler: async (_, h) => {
                const html = await tr.ejs("home.ejs");
                return h.response(html).type("text/html");
            },
        });
    },
});
