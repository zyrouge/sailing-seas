import { defineAppOperator } from "../../operator";
import { tr } from "../templates/renderer";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: "GET",
            path: "/static/styles.css",
            options: {
                auth: false,
            },
            handler: async (_, h) => {
                const content = await tr.text("static/styles.css");
                return h.response(content).type("text/css");
            },
        });
    },
});
