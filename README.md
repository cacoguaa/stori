Stori APP

# Descripción
Este proyecto esta escrito en GO 1.22, contiene una funcionalidad básica, descrita a continuación
1. Se leeran configuraciones para cada servicio y repositorio necesarias para que el servicio arranque correctamente
2. Realiza la lectura de una archivo csv el cual tiene el siguiente formato

    | id | date | transaction |
    |----|------|-------------|
    | 0  | 7/15 | -45.05      |  
    
    El código ignorará el encabezado o header de este csv y empezará desde la primera row inmediata

3. Se leera el archivo linea a linea donde se iran construyendo objectos que cumplan con el formato establecido
4. Todos los objectos construidos se pasaran a una función donde se realizará 
   + inserción en la DB definida
   + Validación de mes de transacción
   + Validación de tipo de transacción realizada
   + Conteó y acumulación de valores a enviar
5. Luego de procesados estos datos se hará un envio a un correo electronico
   * El envio hace uso de smtp, bastante sencilo
   * Este código esta comentado por seguridad, si se desea probar se debe actualizar los datos del proveedor de correo, junto con datos de email emisor, receptor

# Ejecución
Dentro del código se anexa en la carpeta `downloads` un archivo de prueba llamado `transactions` de tipo csv que se ejecuta al momento de levantar el projecto.

Para ejecutar el projecto basta con ir a la terminar y ejecutar `docker-compose up` en la raiz del projecto, lo cual en primera instancia levantara una base de datos dummy y construira la tabla que permitirá almacenar los registros del archivo de prueba.

Cuando el servicio termine su ejecución la base de datos se mantendra activa por eso es necesario cerrarla manualmente en caso de necesitar terminar la ejecución del programa

El proyecto hace uso del archivo `run_env`, que contiene variables para conectarse a la base de datos

# Configuraciones

Para poder ejecutar el proyecto con una base de datos externa se debe
1. Iniciar el servicio con el comando
    ```
    docker run -e DB_USER={{db_user}} -e DB_PASS={{db_secret}} -e DB_PORT={{db_port}} -e DB_NAME={{db_name}} -e FILEPATH={{filepath}} -e DB_MAX_POOL_SIZE={{db_max_pool_size}}  -e DB_HOST={{db_host}} stori
    ```
    Donde cada variable debe ser reemplazada por el valor deseado, cabe recoradr que existen algún ajuste adicional para poder acceder desde docker a esta database externa