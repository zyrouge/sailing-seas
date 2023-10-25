import { defineAppOperator } from "../operator";
import { ejsTemplate } from "./templates/render";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: "GET",
            path: "/",
            handler: async (_, h) => {
                const html = await ejsTemplate("home.ejs");
                return h.response(html).type("text/html");
            },
        });
    },
});
