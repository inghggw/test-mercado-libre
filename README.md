**About API Coupon Test Backend Developer Mercado Libre**

- Desarrollada en Go v.1.14.4.
- Framework Gin para api rest.
- Clean Architecture.
- 5 Pruebas unitarias **(100% c/u)** y cobertura **(95.1%)**.
- Probado en sistema operativo Ubuntu.

**Nivel 1 - ejecuta el programa propuesto**

- Por consola se ejecuta el siguiente programa, y utiliza la data propuesta en el nivel 1:
  - **go run cmd/main.go**

**Nivel 2**

- Esta API con el servicio POST "/coupon", recibe unos ids de los item, donde se utilizó el **recurso de la api de mercado libre**. Con este recurso se válida cada uno de los items que existan para tomar el precio y así calcular los items que puede comprar con el monto del cupo enviado.
- En el directorio /capture, se dejan evidencias del consumo de servicios y ejeción de los test y el coverage. Adicional se adjunta las collections de postman utilizadas.
- Ejecutar el siguiente comando para iniciar el servicio web en **port:8800**:
  - **make dev**
  - En caso de no tener instalado en la máquina local make, se corre el servidor con el siguiente comando:
  - **go run httpd/main.go**

**Nivel 3**

- Url Api publicada: http://35.239.16.146:8800
- **Request POST: http://35.239.16.146:8800/coupon**

**Ejecutar tests y coverage**

- make test

**Autor**

Ing. Pedag. MEng. Henry Giovanny Gonzalez Waltero - © 2020.
