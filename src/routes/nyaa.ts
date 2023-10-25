import { Nyaa, NyaaSearchResult } from "../helpers/nyaa";
import { defineAppOperator } from "../operator";
import { tr } from "./templates/renderer";

export default defineAppOperator({
    operate: async (app) => {
        app.server.route({
            method: ["GET", "POST"],
            path: "/nyaa",
            handler: async (req, h) => {
                const terms = req.query.terms;
                let result: NyaaSearchResult | undefined;
                if (typeof terms === "string" && terms.length > 0) {
                    result = await Nyaa.search(terms);
                }
                const html = await tr.ejs("nyaa.ejs", { result });
                return h.response(html).type("text/html");
            },
        });
    },
});
