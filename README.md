# GatePass 🚪🔑

# **#TODO**
Generate an API and GraphQL specification

# *English*
## Description
GatePass is a phone number-based authentication system developed in Go. It uses codes sent via SMS to verify a user's identity.

## Features
- Automatic sending of verification codes to phone numbers. 📱
- Code verification to authenticate users. 🔒
- RESTful interface for easy integration with other applications. ⚙️
- Support for GraphQL. <img src="https://graphql.org/img/logo.svg" alt="GraphQL Logo" width="15" height="14"/>
- Support for English and Spanish. 🌐

## Technologies Used
- Go
- Gin-Gonic
- Twilio API
- Docker / Docker Compose 🐳

## Getting Started

### Prerequisites
- Go 1.18 or higher
- Docker (optional)

### Instalation

In order to install the module, we must clone the repo using the command:

```bash
git clone https://github.com/usuario/gatepass.git
cd gatepass
```

After clonin the repo we must create the .env file using:

```bash
cp template.env .env
```

Then we need to change the xxxxx values for the correct values for our Twilio Account.


### Execute the module

#### Execute without Docker

In order to execute the module without Docker we must execute the following command:

```bash
go build
```
Then we execute the file from the build.

#### Execute with Docker

In order to execute the module using Docker we must follo the following commands:

```bash
git clone https://github.com/usuario/gatepass.git
cd gatepass
docker build -t gatepass . 
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='twilio account sid' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Phone number from which the messages will be sent' -e APP_NAME='App name' -e LANGUAGE='Language in which the message will be sent' gatepass
```

#### Execute with Docker Compose

In order to initiate the docker with Docker Compose you need to run the following command:

```bash
docker compose up -d --build
```

#### Execute from DockerHub

In order to initiate the docker from DockerHub you can pull the image with the command:

```bash
docker pull atralex/gatepass:v1.0.0
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='twilio account sid' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Phone number from which the messages will be sent' -e APP_NAME='App name' -e LANGUAGE='Language in which the message will be sent' gatepass
```

Or you can copy the service gatepass from the docker-compose.online.yml file into your docker-compose.yml file, adjust the environment and run:

```bash
docker compose up -d --build
```


## Feedback and Contributions 🚀
If you find any error or improvement in the code or in the Readme.md feel free to open a pull request.
Also if you know how to transalate the text message into other languages please feel free to add another case in the switch

# *Español*

## Descripción
GatePass es un sistema de autenticación basado en números de teléfono desarrollado en Go. Utiliza códigos enviados por SMS para verificar la identidad de un usuario.

## Características 
- Envío automático de códigos de verificación a números de teléfono. 📱
- Verificación de códigos para autenticar usuarios. 🔒
- Interfaz RESTful para integración fácil con otras aplicaciones. ⚙️
- Soporte para GraphQl. <img src="https://graphql.org/img/logo.svg" alt="GraphQL Logo" width="15" height="14"/>
- Soporte para inglés y español. 🌐

## Tecnologías Usadas
- Go
- Gin-Gonic
- Twilio API
- Docker / Docker Compose 🐳

## Cómo Empezar

#### Pre-requisitos
- Go 1.21
- Docker (opcional)

### Instalación

Debemos clonar el repositorio ejecutando los comandos:

```bash
git clone https://github.com/usuario/gatepass.git
cd gatepass
```

Después de clonar el repositorio, copiamos el archivo template.env con el comando:

```bash
cp template.env .env
```

**Y después sustituimos los valores xxxxx por los valores adecuados a tu cuenta de Twilio.**

### Ejecución

#### Ejecución sin Docker

Para ejecutar el módulo sin Docker ejecutamos el comando:

```bash
go build
```
Y luego ejecutamos el archivo resultante del build

#### Ejecución con Docker

Para ejecutar el módulo con Docker ejecutamos los siguientes comandos sustituyendo los valores de las variables de entorno necesarias

```bash
git clone https://github.com/usuario/gatepass.git
cd gatepass
docker build -t gatepass .
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='sid de twilio' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Telefono de twilio desde el que se mandaran los mensajes' -e APP_NAME='Nombre de la app' -e LANGUAGE='Idioma en el que se manda el mensaje' gatepass
```

#### Ejecución con Docker Compose

Para ejecutar el módulo con Docker Compose ejecutamos el comando:

```bash
docker compose up -d --build
```

#### Execute from DockerHub

Para iniciar el contenedor desde DockerHub puedes descargar la imagen con el comando:

```bash
docker pull atralex/gatepass:v1.0.0
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='sid de twilio' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Telefono de twilio desde el que se mandaran los mensajes' -e APP_NAME='Nombre de la app' -e LANGUAGE='Idioma en el que se manda el mensaje' atralex/gatepass:v1.0.0
```

O puedes copiar el servicio gatepass del archivo docker-compose.online.yml en tu archivo docker-compose.yml, ajustar las variables de entorno y ejecutar:

```bash
docker compose up -d --build
```


## Feedback y Contribuciones 🚀
Si encuentras algún error en el código o en el Readme.md por favor abre una pull request para corregirlo.
También si sabes traducir el mensaje de texto a otro idioma sientete libre de añadir otro case al switch