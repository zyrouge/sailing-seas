import { defineAppOperator } from "../../operator";
import { textTemplate } from "../templates/render";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: "GET",
            path: "/static/styles.css",
            options: {
                auth: false,
            },
            handler: async (_, h) => {
                const content = await textTemplate("static/styles.css");
                return h.response(content).type("text/css");
            },
        });
    },
});
