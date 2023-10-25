import p from "path";
import { readFile } from "fs/promises";
import ejs from "ejs";
import { constants } from "../../helpers/utils";

const ejsDefaultData = {
    site: {
        name: constants.SITE_NAME,
        baseUrl: constants.SITE_BASE_URL,
    },
};

export const ejsTemplate = async (
    basename: string,
    data: Record<any, any> = {}
) => {
    const path = p.join(__dirname, basename);
    const content = await ejs.renderFile(
        path,
        { ...ejsDefaultData, ...data },
        { root: p.dirname(path) }
    );
    return content;
};

export const textTemplate = async (basename: string) => {
    const path = p.join(__dirname, basename);
    const content = await readFile(path, "utf-8");
    return content;
};
