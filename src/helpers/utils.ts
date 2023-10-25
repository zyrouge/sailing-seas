import pino from "pino";

export const logger = pino();

export const constants = {
    SITE_NAME: "Sailing Seas",
    SITE_BASE_URL: process.env.BASE_URL ?? "/",
};
