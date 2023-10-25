import { Nyaa, NyaaSearchResult } from "../helpers/nyaa";
import { defineAppOperator } from "../operator";
import { ejsTemplate } from "./templates/render";

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
                const html = await ejsTemplate("nyaa.ejs", { result });
                return h.response(html).type("text/html");
            },
        });
    },
});
