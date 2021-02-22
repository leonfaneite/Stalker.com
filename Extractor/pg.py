

import psycopg2

def conexion_pg():
    conexion = psycopg2.connect(database="customer",user="usuario",password="usuario",host="localhost",port="5432")
    cursor1=conexion.cursor()
    cursor1.execute("SELECT Id,Words FROM query")
    for fila in cursor1:
        print(fila)
    conexion.close()

    

print(conexion_pg())

