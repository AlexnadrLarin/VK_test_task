FROM node:alpine AS builder

WORKDIR /frontend_app_build

ARG REACT_APP_BACKEND_API_URL
ARG REACT_APP_TIME_INTERVAL

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:alpine

WORKDIR /frontend_app

COPY --from=builder /frontend_app_build/build /frontend_app/build

RUN npm install -g serve

CMD ["sh", "-c", "serve -s build -l ${PORT}"]

