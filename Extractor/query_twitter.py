import psycopg2



def conexion_pg():
    
    conexion = psycopg2.connect(database="customer",user="usuario",password="usuario",host="localhost",port="5432")
    
    cursor1 = conexion.cursor()

    cursor1.execute("SELECT * FROM query")

    datos = []

    for file in cursor1:

        datos.append(file)
        
        print(datos)
        
    
conexion_pg()