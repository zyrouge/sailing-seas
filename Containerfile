FROM docker.io/oven/bun:slim

WORKDIR /usr/app

COPY package*.json .
RUN bun install --frozen-lockfile

COPY . .

CMD [ "bun", "start" ]
