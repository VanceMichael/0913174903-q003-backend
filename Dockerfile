FROM golang:1.23
WORKDIR /app
COPY . .
CMD ["sh", "-c", "echo 请先完成服务实现"]
