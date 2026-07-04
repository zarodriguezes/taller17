# Proyecto CuentaCuenta - Pruebas de Encapsulación en Go

## 📋 Descripción

Este proyecto demuestra el uso de **encapsulación** en Go, probando cómo las pruebas internas (`package banco`) pueden acceder a elementos privados, mientras que las pruebas externas (`package banco_test`) solo pueden acceder a la API pública.

---

## 📂 Estructura del Proyecto

![Estructura](docs/images/image.png)

## Código agregado en cuenta_externo_test.go:

![Pruebas externas 1](docs/images/image1.png)
![Pruebas externas 1](docs/images/image2.png)
![Pruebas externas 1](docs/images/image3.png)

## Ejecutar todas las pruebas en go 

go test ./banco_test -v

## Ejecutar solo pruebas externas (package banco_test):

1. Ejecuta las pruebas
2. Captura la salida completa de la ejecución (puedes copiar el texto de la terminal o tomar una captura de pantalla).
3. Compara esta nueva ejecución con la ejecución anterior
4. Analiza y responde en tu informe:
- ¿Qué nuevas pruebas aparecen en la ejecución?
- ¿Qué métodos o funcionalidades adicionales se están probando?
- ¿Cómo se refleja en los resultados el hecho de que los nuevos métodos son privados?
- ¿Qué diferencias observas en la cobertura de pruebas entre ambas ejecuciones?
- ¿Qué aprendizaje obtienes sobre la encapsulación en Go a partir de este ejercicio?
5. Entregar el proyecto
