import p from "path";
import { readFile } from "fs/promises";
import ejs from "ejs";
import { constants } from "../../helpers/utils";
import { createReadStream } from "fs";

export class TemplateRenderer {
    defaultEjsData = {
        site: {
            name: constants.SITE_NAME,
            baseUrl: constants.SITE_BASE_URL,
        },
    };

    async ejs(basename: string, data: Record<any, any> = {}) {
        const path = p.join(__dirname, basename);
        const content = await ejs.renderFile(
            path,
            { ...this.defaultEjsData, ...data },
            { root: p.dirname(path) }
        );
        return content;
    }

    async text(basename: string) {
        const path = p.join(__dirname, basename);
        const content = await readFile(path, "utf-8");
        return content;
    }

    async stream(basename: string) {
        const path = p.join(__dirname, basename);
        const stream = createReadStream(path);
        return stream;
    }
}

export const tr = new TemplateRenderer();
