import psycopg2



def conexion_pg():
    
    conexion = psycopg2.connect(database="customer",user="usuario",password="usuario",host="localhost",port="5432")
    
    cursor1 = conexion.cursor()

    cursor1.execute("SELECT words FROM query")
    

    for elements  in cursor1:

        words_q = str(elements[0])

        words_all = words_q.split(",")


        
        lista =[]

        for term  in words_all:        
            
            word = str("#"+ term)

            lista.append(word)
      
        return lista


print(conexion_pg())