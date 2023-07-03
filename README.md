# Concurrente-TF

## **Integrantes:**

- Natalia Melissa Maury Castañeda (u201816996)

- Sergio Antonio Nuñez Lazo (u201910357)

- Oscar Daniel Flores Palermo (u201716498)

## Sobre el Trabajo Final (TF):
Para el Trabajo Final (TF) del curso Programación Concurrente y Distribuída, se tuvo que realizar el juego Hula Hoop Race como el realizado en la Tarea Académica 3 (TA3), pero la diferencia es que ahora los hula hulas son computadoras,  los resultados deben mostrarse en una interfaz gráfica usando una página web, y el juego termina cuando el primer jugador del otro equipo llega a la meta. Para entender bien los cambios se debe recordar la idea base del juego: se tienen dos equipos que curzarán un camino delimitado por hula hulas (Hula Hoop) y cuando dos jugadores se encuentran en la misma posición jugarán "piedra, papel o tijeras", el ganador del desafío continua y el perdedor regresar al inicio donde el siguiente compañero del equipo comienza a jugar. El proceso se repetirá hasta que el primer jugador de un equipo llegue a la meta, en vez de que todos los jugadores de un equipo lleguen a la meta como en la TA3. 

Para desarrollar el TF, se tuvo que primero convertir a las computadoras en los Hula Hulas/camino por recorrer. Para lograrlo se simuló que habían "múltiples computadoras" al ejecutar múltipoles cmd y haciendo que se conecten usando los puertos reservados desde el puerto 8000. Una vez que se ha establecido una conexión, se indica que ese nodo está en estado de escucha o "listening" para saber que puede recibir información. Una vez se tienen todos los nodos, el camino de hula hulas, se requieren de canales para enviar la información entre un nodo a otro. En este caso, la información son los jugadores que contienen su información relevante en una estructura, una vez reunida la información, se transforman los valores a formato json que el nodo se encargará de desencriptar e interpretra. Aquí los nodos son los procesos en vez de los jugadores, por lo que al recibir la información son los encargados de enviarla al siguiente nodo y llamar a las funciones correspondientes para hacer que se ejecuten los "saltos entre hula hulas", los enfrentamientos e informar el equipo ganador. De esta manera, al cada nodo ser un proceso que se ejecuta de forma independiente, se resolvió el problema utilizando canales y programación distribuída.

**Sobre el Juego:**
- **Equipos:** Como se mencionó previamente, los jugadores están divididos en dos equipos que en este caso serán "Cobras" o "Leones"
- **Enfrentamiento:** Para el enfrentamiento, ahora es como si se lanzara una moneda en vez de jugar yan ken po en sí, si el resultado es mayor a 50 gana el equipo A, si es de 49 o menos gana el equipo B.
- **Interfaz:** Debido a que en las indicaciones del TF brindadas en clase se indicó que la interfaz que se vea desde una página web, se realizó un backend y un frontend simple en go. El backend no tiene una  base de datos, pero la información que recibe es el jugador. Para ello, en el archivo "backend.go" se utiliza el mismo struct de Player que en el archivo principal, node.go, pero se le agrega el atributo id cuyo valor será dado con un contador. Es necesario hacer esta pequeña modificación para que el backend haga una lista de los jugadores que recibe y saber que es un nuevo juegador el que está entrando y evitar confusiones. Una vez recibida la información, realizará las consultas get y post para enviar los datos al frontend y que se vean bien. Para asegurar que el formato, los datos se verán de la forma deseada, también se implementó el archivo llamado "index.go" que se encarga de manejar el frontend. 

**Archivos**
- **node.go =** Archivo principal que se encarga de ejecutar el juego de "Hula Hoop Race: Cobras vs Leones"
- **backend.go =** Archivo que se encarga de gestionar todo lo relacionado con el backend
- **index.go =** Archivo que se encarga de gestionar todo lo relacionado con el frontend

**Puertos reservados:**
- **8080 =** Puerto que usa el backend
- **8081 =** Puerto que usa el frontend
- **8000 - 8006 =** Nodos/Hula Hulas (Puertos para la simulación). Se puede editar el archivo .bat para que el sistema tenga más nodos, pero el archivo subido tiene 7 nodos que son los listados aquí. 

## Cómo ejecutar el trabajo:
1. Ejecutar el backend en una cmd (Va desde el 8000 al 8006)
2. Ejecutar el .bat y esperar a que salgan todas en "listening"
3. Ejecutar el powershell el archivo node.go con el flag '-s'
4. Ejecutar el frontend
5. En un navegador, abrir el localhost:8081 para ver el resultado.

**IMPORANTE:** Sobre los archivos backend y frontend
- Se puede ver los resultados de ambos archivos en la cmd donde se están ejecutando. Esto es con el objetivo de verificar que se estaban mandando los datos correctamente y que el programa funcionaba bien.
- Para ver los resultados del backend en el navegador, se debe abrir una pestaña y escribir "localhost:8080/players"
