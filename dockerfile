#imagem base de golang
FROM golang:1.22.2-alpine

#instalar dependencias adicionais
RUN apk add --no-cache git 

#definir diretorio de trabalho no container
WORKDIR /app

#copiar arquivos do projeto para o container
COPY . .

#baixar dependencias e copilar o app
RUN go mod tidy && go build -o app main.go

#expor a porta
EXPOSE 8080

#comando para rodar o app
CMD [ "./main" ]