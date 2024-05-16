# GatePass 🚪🔑

# **#TODO**


# *English*
## Description
GatePass is a phone number-based authentication system developed in Go. It uses codes sent via SMS to verify a user's identity.

## Features
- Automatic sending of verification codes to phone numbers. 📱
- Code verification to authenticate users. 🔒
- RESTful interface for easy integration with other applications. ⚙️
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
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='sid de twilio' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Telefono de twilio desde el que se mandaran los mensajes' gatepass
```

#### Execute with Docker Compose

Para ejecutar el módulo con Docker Compose ejecutamos el comando:

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
docker run -p 8080:8080 -e TWILIO_ACCOUNT_SID='sid de twilio' -e TWILIO_AUTH_TOKEN='auth token twilio' -e TWILIO_FROM_PHONE='Telefono de twilio desde el que se mandaran los mensajes' gatepass
```

#### Ejecución con Docker Compose

Para ejecutar el módulo con Docker Compose ejecutamos el comando:

```bash
docker compose up -d --build
```
## Feedback y Contribuciones 🚀
Si encuentras algún error en el código o en el Readme.md por favor abre una pull request para corregirlo.
También si sabes traducir el mensaje de texto a otro idioma sientete libre de añadir otro case al switch