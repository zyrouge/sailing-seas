import { defineAppOperator } from "../../operator";
import { tr } from "../templates/renderer";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: "GET",
            path: "/static/icon.png",
            options: {
                auth: false,
            },
            handler: async (_, h) => {
                const content = await tr.stream("static/icon.png");
                return h.response(content).type("image/png");
            },
        });
    },
});
